# Subresources in kubernetes

Some resources in kubernetes (e.g. `services`) can have subcomponents called
subresources. One example of this is `services/proxy`, which may be used in
RBAC rules. For example, a user must have `create` permissions on
`services/proxy` to be able to access the kubernetes dashboard using `kubectl
proxy`.

You can use `kubectl get --raw /` to list the different `apiGroups`, then call
the same command on one of those groups to get the information about its
resources and subresources. For example:

```bash
kubectl get --raw /api/v1 | jq '.resources[] | select(.name == "services/proxy")'
```

This returns:

```json
{
  "name": "services/proxy",
  "singularName": "",
  "namespaced": true,
  "kind": "ServiceProxyOptions",
  "verbs": [
    "create",
    "delete",
    "get",
    "patch",
    "update"
  ]
}
```

Searching the kind `ServiceProxyOptions` will bring up some
[documentation](https://github.com/godaddy/kubernetes-client/blob/master/docs/1.12/ServiceProxyOptions.md)
about the subresource, which can be very hard to locate otherwise.

    #kubernetes #tips
