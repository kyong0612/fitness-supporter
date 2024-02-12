FROM ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-contrib:0.94.0-amd64
COPY config.yaml /etc/otel/config.yaml
EXPOSE 4317
CMD ["--config", "/etc/otel/config.yaml"]
