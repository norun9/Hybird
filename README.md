## Backend Architecture

| Documents                                                |
|:---------------------------------------------------------|
| [Clean Architecture](backend/api/doc/README.md) |


## AWS Serverless Architecture

<div align="center">
  <img src="https://github.com/user-attachments/assets/dcaf8479-a440-471b-a7dc-4b8c2b6d36ec" alt="serverless">
</div>


![hybird_serverless drawio (1)]()

### WebSocket API
PENDING

## Deployment Process

> The following are personal notes

### Terraform

```bash
export AWS_PROFILE=your_iam_user_name

terraform plan -out=tfplan

terraform apply "tfplan"
```

### Frontend

Run GitHub Actions manually ([ssg_deploy.yml](.github/workflows/ssg_deploy.yml))

### Goose Migration to RDS via Bastion

> [!NOTE]
> Integrate the migration process into the Continuous Deployment (CD) pipeline to automate database updates during deployments.

Follow these steps to perform a Goose migration to an RDS instance via a Bastion host.

#### Step 1: SSH into the Bastion Host
```bash
ssh -i "path/to/your-keypair.pem" ec2-user@<EC2_Public_IP>
```

#### Step 2: Connect to the RDS Instance
Once logged into the Bastion, use the following command to connect to your RDS instance.
```bash
mysql -h <RDS_Endpoint> -u <DB_User> -p
```

#### Step 3: Install Goose
```bash
# Download the Goose installation script:
curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh -o goose | sh

# Make the Goose binary executable:
chmod +x goose

# Move Goose to /usr/local/bin to make it accessible globally:
sudo mv goose /usr/local/bin/

# Verify that Goose is installed correctly:
goose --version
```

#### Step 4: Transfer Migration Directory to the Home Directory on Bastion

To transfer the migration directory from your local machine to the Bastion instance, use the following command:
```bash
# /home/ec2-user/ is the Bastion home directory
scp -r -i "your_keypair.pem" your_migration_files_dir ec2-user@<DNS>:/home/ec2-user/
```


#### Step 5: Run Goose Migration

After transferring the migration files, run the Goose migration using the following command:
```bash
goose -dir /path/to/your_migrations_dir mysql '<DB_User>:<DB_Password>@tcp(<RDS_Endpoint>:3306)/<DB_Name>?parseTime=true' up
```





