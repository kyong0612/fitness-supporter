FROM otel/opentelemetry-collector-contrib:0.100.0
COPY config.yaml /etc/otelcol-contrib/config.yaml
EXPOSE 4317
CMD ["--config", "/etc/otelcol-contrib/config.yaml"]
