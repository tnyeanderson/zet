# Browse Kubernetes API with swagger-ui

It is extremely easy to browse the Kubernetes API with Swagger UI, and it can
be done without any custom images and only a single deployment (with associated
Service/Ingress).

Simply apply the manifest in this directory:

```bash
kubectl apply -f swagger-ui.yaml
```

This manifest uses an `initContainer` based on `bitnami/kubectl` to fetch the
current spec and save it to a file. The main container
(`swaggerapi/swagger-ui`) then reads the spec and provides a Swagger UI for it.

This keeps the design very simple, but it does mean that you need to
delete/recreate the pod to see any updates/changes to the API (for example, if
you recently added a CRD).

    #kubernetes #openapi #swagger #api
