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

## google cloud workflow identity set up

ref: <https://github.com/google-github-actions/auth?tab=readme-ov-file#setup>

```bash
gcloud iam workload-identity-pools create "github" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --display-name="GitHub Actions Pool"
```

```bash
gcloud iam workload-identity-pools describe "github" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --format="value(name)"
```

```bash
gcloud iam workload-identity-pools providers create-oidc "fitness-supporter" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --workload-identity-pool="github" \
  --display-name="My GitHub repo Provider" \
  --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository" \
  --issuer-uri="https://token.actions.githubusercontent.com"
```

- see provider id

```bash
gcloud iam workload-identity-pools providers describe "fitness-supporter" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --workload-identity-pool="github" \
  --format="value(name)"
```

- set CD parameter

```bash
- uses: 'google-github-actions/auth@v2'
  with:
    project_id: 'my-project'
    workload_identity_provider: '...' # "projects/123456789/locations/global/workloadIdentityPools/github/providers/my-repo"
```
