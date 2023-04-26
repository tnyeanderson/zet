# Kubernetes local-storage and hostPath mounts

Volumes of the `hostPath` and `local-storage`/`local` types are implemented in
interesting ways.

For either of these volume types, if the source directory (or any of its
parents) are mount points on the host system, two things happen:

1. The mount is replicated (with a separate connection) at
   `/var/lib/kubelet/pods/<id>/volumes/kubernetes.io~local-volume/<name>`,
   potentially with a subpath defined as needed (see `/proc/self/mountinfo`)
2. This path is bind mounted to the pod itself, resulting in a `/proc/mounts`
   entry on the pod which includes the connection information for the mount,
   even if the mount plugin (e.g. the NFS client) isn't available in the pod.

> NOTE: If a `hostPath` mount is defined (either with a backing PV or directly
defined in the pod spec) which is not part of a "mounted" path, the additional
mount point in `/var/lib/kubelet/pods` is not created. The directory is bind
mounted as-is.

This has some interesting side effects. Say you start with the following
manifest:

```yaml
---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: test-pv-nfs
spec:
  capacity:
    storage: 10Gi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteMany
  persistentVolumeReclaimPolicy: Retain
  storageClassName: local-storage
  local:
    path: /mnt/mymountpoint/subdir
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
          - callisto
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: test-pvc-nfs
spec:
  storageClassName: local-storage
  volumeName: test-pv-nfs
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 1Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: ubuntu-nfs
spec:
  containers:
    - name: ubuntu
      image: ubuntu
      command:
        - sleep
        - 3000s
      volumeMounts:
        - mountPath: /mount-target
          name: test-volume
  volumes:
    - name: test-volume
      persistentVolumeClaim:
        claimName: test-pvc-nfs

```

If you were to apply this, you may expect that running `umount
/mnt/mymountpoint` on the node where the PV/pod is scheduled would cause issues
when trying to write to that directory from the pod. However, this is not the
case! Since `/mnt/mymountpoint` gets evaluated at PV/pod creation, it ends up
resolving to (for example) `mynfs.host:/my/exported/path`, and a new connection
to this NFS server (or other mount provider) is created in `/var/lib/kubelet`
which gets bind mounted to the pod. Therefore, unmounting the
`/mnt/mymountpoint` on the host node has no effect on the pod's ability to
read/write to that filesystem, since it is now using an independent
mount/connection to the storage! This is the same for SeaweedFS or other
filesystems mounted on the host.

The best way to understand is to try it out. I spent days tracking down what
was happening in `pkg/volume/local/local.go` from the kubernetes repo. That
lead me to the kernel docs for shared subtrees and related information about
bind propagation, linked below. Realisitically, part of this nuance is at the
kernel level (cgroups and shared subtrees for mounts), and some is at the k8s
level (the creation of an independent connection/mount used for the pod).

## Related

- https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt
- https://man7.org/linux/man-pages/man7/mount_namespaces.7.html
- https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation

    #kubernetes #mount #volume #bind #kernel #cgroups
