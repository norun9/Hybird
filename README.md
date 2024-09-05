| Documents                                                |
|:---------------------------------------------------------|
| [Backend Architecture](backend/internal/api/doc/README.md) |


## AWS Serverless Architecture

PENDING

## Deploy

### Terraform

```bash
terraform plan -out=tfplan

terraform apply "tfplan"
```

### Frontend

Run GitHub Actions manually (ssg_deploy)