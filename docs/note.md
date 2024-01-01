# note

## manual build and deploy

```bash
-docker image rm asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/prd

docker buildx build . --platform linux/amd64 --no-cache --tag asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/prd:latest

docker push asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/prd:latest

gcloud run deploy fitness-supporter \
    --image asia-northeast1-docker.pkg.dev/kyong0612-lab/fitness-supporter/prd:latest \
    --project kyong0612-lab \
    --region asia-northeast1 \
    --allow-unauthenticated \
    --set-env-vars=ENV=production,LINE_CHANNEL_TOKEN=dummy
```
