## Yet another gRPC echo server

YAGES (yet another gRPC echo server) is an educational gRPC server implementation. The goal is to learn gRPC and communicate best practices around its deployment and usage in the context of Kubernetes.

### Changes from https://github.com/mhausenblas/yages:
- Adds `grpc_health_v1.HealthServer` implementation (see [documentation here](https://github.com/grpc/grpc/blob/master/doc/health-checking.md))
- Container image bundles [grpc-health-probe](https://github.com/grpc-ecosystem/grpc-health-probe) to ensure healthchecking in k8s is improved
- Releases with [goreleaser](https://goreleaser.com/)
