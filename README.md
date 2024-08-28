| Documents                                                |
|:---------------------------------------------------------|
| [Backend Architecture](backend/internal/api/doc/README.md) |


## AWS Serverless Architecture

PENDING

## Deploy

### Backend

```bash
DOCKER_HOST=unix:///Users/xxx/.docker/run/docker.sock sam build --use-container

DOCKER_HOST=unix:///Users/xxx/.docker/run/docker.sock sam deploy --profile xxx --guided
```

### Frontend

Run GitHub Actions manually (ssg_deploy)