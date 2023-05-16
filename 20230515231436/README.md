# Kubernetes host networking and CNI plugins

If `hostNetwork: true` is set on a pod, *none of the CNI plugins in the chain
will be called*. This can be demonstrated by adding the
[debug](https://github.com/containernetworking/cni/tree/main/plugins/debug)
plugin to the chain. First, compile the debug plugin and add the resulting
binary to `/opt/cni/bin` (or `/var/lib/rancher/k3s/data/current/bin` for k3s).
Then edit the CNI configuration as follows:

```yaml
{
  ...
  "plugins":[
    ...
    {
      "type":"debug",
      "cniOutput":"/tmp/cni-output.txt",
    }
    ...
  ]
}

```

If a pod is scheduled with `hostNetwork: true`, *the CNI plugin chain is not
called at all, so there will be no output*. Otherwise, when a pod is scheduled,
the debug plugin will write all the information that was available to it at
that point in the chain to the provided file.

    #kubernetes #k3s #cni
