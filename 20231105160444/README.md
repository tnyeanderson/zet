# Remove an etcd member when the node is no longer available

Using `kubeadm reset` before taking a node down is obviously the correct way to
do things. However, sometimes a node fails suddenly and you are not able to run
`kubeadm reset` on the node.

For these examples, `badnode` will refer to this failed node, and `workingnode`
will refer to another etcd member node which is still functioning.

First, remove `badnode` from the kubernetes cluster:

```bash
kubectl delete node badnode
```

Once `badnode` is rebuilt, you can try to `kubeadm join` it back to the
cluster, but you'll probably get this error message:

```
error execution phase check-etcd: etcd cluster is not healthy: failed to dial endpoint https://badnode:2379 with maintenance client: context deadline exceeded
```

This is because the etcd cluster is not aware that `badnode` is no longer
available. Therefore, it must be removed manually. For in-cluster etcd, use the
following command to get the ID of the etcd member which needs to be removed:

```bash
kubectl exec etcd-workingnode -n kube-system -- etcdctl \
  --cacert /etc/kubernetes/pki/etcd/ca.crt \
  --cert /etc/kubernetes/pki/etcd/peer.crt \
  --key /etc/kubernetes/pki/etcd/peer.key \
  member list
```

Then remove it (for example, node `36693bb67a1292d7`):

```bash
kubectl exec etcd-workingnode -n kube-system -- etcdctl \
  --cacert /etc/kubernetes/pki/etcd/ca.crt \
  --cert /etc/kubernetes/pki/etcd/peer.crt \
  --key /etc/kubernetes/pki/etcd/peer.key \
  member remove 36693bb67a1292d7
```

Then run your `kubeadm join` as usual.

## Related

- <https://stackoverflow.com/a/64322313>

    #kubernetes #kubeadm #etcd #disaster
