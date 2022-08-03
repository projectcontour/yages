FROM alpine:latest as build
ADD https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.11/grpc_health_probe-linux-amd64 /grpc-health-probe
RUN chmod +x /grpc-health-probe
FROM scratch as final
COPY yages /
COPY --from=build /grpc-health-probe /grpc-health-probe
ENTRYPOINT ["/yages"]
