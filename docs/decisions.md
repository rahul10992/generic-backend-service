# Decisions


HTTP on :8080

gRPC on :9090

/healthz, /readyz, /metrics

metrics: Prometheus, OTel traces later

TLS: edge termination at ingress, no mTLS (for now)