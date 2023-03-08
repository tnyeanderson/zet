# Istio readiness/liveness probes

When a kubernetes pod includes both the app's container (`app`) and the istio
sidecar container (`istio-proxy`), the readiness and liveness probes for the
`app` container will be overwritten and directed through the proxy. This is
accomplished by setting an environment variable called `ISTIO_KUBE_APP_PROBERS`
to a JSON definition of the upstream probe endpoints in the `app` container.
Then the probes for the `app` container get overwritten to
`/app-health/api/livez` and `/app-health/api/readyz` on a port that is
listening from the `istio-proxy` container (`15020`).

This can be disabled by annotating the pod with
`sidecar.istio.io/rewriteAppHTTPProbers: "false"`.

Related:

- https://istio.io/latest/docs/ops/configuration/mesh/app-health-check/

    #istio #kubernetes #k8s
