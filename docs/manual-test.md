# manual-test

## curl endpoint

```bash
curl -v http://localhost:8080/healthcheck
```

```bash
 curl -v \
    --header "Content-Type: application/json" \
    --data '{"image_url": "dummy-image-url","user_id","dummy-user-id"}' \
    http://localhost:8080/proto.handler.v1.AnalyzeImageService/AnalyzeImage
```
