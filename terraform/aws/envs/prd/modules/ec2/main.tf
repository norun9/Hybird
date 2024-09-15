resource "aws_security_group" "bastion_sg" {
  name        = "bastion_sg"
  description = "Security group for EC2 instance in private subnet"
  vpc_id      = var.vpc_id

  # Inbound Rules:
  # Permit SSH connection from Bastion host
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Permit the connection to RDS(MySQL)
  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr_block]
  }

  # Outbound Rules:
  # https
  egress {
    protocol    = "tcp"
    from_port   = 443
    to_port     = 443
    cidr_blocks = ["0.0.0.0/0"]
  }

  # http
  egress {
    protocol    = "tcp"
    from_port   = 80
    to_port     = 80
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Permit the entire outbound traffic
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Fetch Amazon Linux 2 AMI
data "aws_ssm_parameter" "amzn2_ami" {
  name = "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
}

resource "aws_instance" "bastion" {
  ami                         = data.aws_ssm_parameter.amzn2_ami.value
  instance_type               = "t2.micro"
  vpc_security_group_ids      = [aws_security_group.bastion_sg.id]
  subnet_id                   = var.public_subnet_id # Public subnet id
  associate_public_ip_address = true
  iam_instance_profile        = "ec2_bastion_profile"
  key_name                    = "hybird-keypair"
  capacity_reservation_specification {
    capacity_reservation_preference = "none"
  }
  tags = {
    Name = "hybird-bastion"
  }
  # Settings of EBS Route Volume
  root_block_device {
    volume_size           = 8
    volume_type           = "gp3"
    iops                  = 3000
    throughput            = 125
    delete_on_termination = true
    encrypted             = true
    kms_key_id            = "alias/aws/ebs"
    tags = {
      Name = "gp3-ec2"
    }
  }
  user_data = <<EOF
#!/bin/bash
# Update and install necessary packages
sudo yum update -y
sudo yum install -y mysql
EOF
}

resource "aws_iam_instance_profile" "ec2_bastion_profile" {
  name = "ec2_bastion_profile"
  role = var.ec2_bastion_role_name
}