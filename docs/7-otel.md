# OpenTelemetry

## References

- <https://zenn.dev/google_cloud_jp/articles/20230516-cloud-run-otel>
- <https://github.com/GoogleCloudPlatform/devrel-demos/tree/main/devops/otel-col-cloud-run-multicontainer>
- <https://cloud.google.com/run/docs/tutorials/custom-metrics-opentelemetry-sidecar?hl=ja>

## build sidecar container

- ref:<https://cloud.google.com/run/docs/deploying?hl=ja#multicontainer-yaml>

```bash
docker buildx build .otelcollector/ --platform linux/amd64 --no-cache --tag asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/sidecar/otel:latest
docker push asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/sidecar/otel:latest
# docker image rm asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/sidecar/otel:latest
```

## replace config

- cloud runのyaml fileを編集する
<https://share.cleanshot.com/ncrCNj1T>
