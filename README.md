## Backend Architecture

[See our "Clean Architecture" for more details](backend/api/doc/README.md)


## AWS Architecture

### REST API

<div align="center">
  <img src="https://github.com/user-attachments/assets/830a0e75-4cd3-437f-be9c-c5726c3081fe" alt="backend">
</div>

### WebSocket API

<div align="center">
  <img src="https://github.com/user-attachments/assets/cb3d706a-1aa0-46ee-ac21-6be7873ecf99" alt="websocket">
</div>

### Frontend

<div align="center">
  <img src="https://github.com/user-attachments/assets/c0b308f3-79a8-4251-87eb-d23375ad274f" alt="frontend">
</div>


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

### Goose Migration to RDS via Bastion Host

> [!NOTE]
> Integrate the migration process into the Continuous Deployment (CD) pipeline to automate database updates during deployments.

Follow these steps to perform Goose migration to RDS instance via Bastion host.

#### Step 1: SSH into the Bastion Host
```bash
ssh -i "path/to/your-keypair.pem" ec2-user@<EC2_Public_IP>
```

#### Step 2: Connect to the RDS Instance
Once logged into the Bastion, use the following command to connect to your RDS instance in the shell.
```bash
mysql -h <RDS_Endpoint> -u <DB_User> -p
```

#### Step 3: Install Goose

Use the following command to install Goose in the shell.
```bash
# Download the Goose installation script
curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh -o goose | sh

# Make the Goose binary executable
chmod +x goose

# Move Goose to /usr/local/bin to make it accessible globally
sudo mv goose /usr/local/bin/

# Verify that Goose is installed correctly
goose --version
```

#### Step 4: Transfer Migration Directory to the Home Directory on Bastion

To transfer the migration directory from "your local machine" to the Bastion instance, use the following command.
```bash
# /home/ec2-user/ is the Bastion home directory
scp -r -i "your_keypair.pem" your_migration_files_dir ec2-user@<DNS>:/home/ec2-user/
```


#### Step 5: Run Goose Migration

After transferring the migration files, run the Goose migration using the following command in the shell.
```bash
goose -dir /path/to/your_migrations_dir mysql '<DB_User>:<DB_Password>@tcp(<RDS_Endpoint>:3306)/<DB_Name>?parseTime=true' up
```





