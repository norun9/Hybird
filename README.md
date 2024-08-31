| Documents                                                |
|:---------------------------------------------------------|
| [Backend Architecture](backend/internal/api/doc/README.md) |


## AWS Serverless Architecture

PENDING

## Deploy

### Deprecated: Backend (Cloudformation)

```bash
DOCKER_HOST=unix:///Users/xxx/.docker/run/docker.sock sam build --use-container

DOCKER_HOST=unix:///Users/xxx/.docker/run/docker.sock sam deploy --profile xxx --guided
```

### Terraform

```bash
terraform plan -out=tfplan

terraform apply "tfplan"
```

### Frontend

Run GitHub Actions manually (ssg_deploy)