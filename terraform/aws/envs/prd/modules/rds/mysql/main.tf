resource "aws_security_group" "db_sg" {
  name        = "hybird-db-sg"
  description = "Security group for the RDS instance"
  vpc_id      = var.vpc_id  # Assuming you have the VPC ID defined or passed in as a variable

  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr_block]  # Adjust this to allow access from your application or Lambda
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_subnet_group" "db_subnet_group" {
  name       = "hybird-db-subnet-group"
  subnet_ids = var.subnet_ids  # Assuming you're using a VPC module that outputs private subnet IDs
}

resource "random_password" "db_password" {
  length  = 16
  special = true
}

resource "aws_db_instance" "hybird_db" {
  identifier             = "hybird-db-instance"
  allocated_storage      = 20
  storage_type           = "gp2"
  engine                 = "mysql"
  engine_version         = "8.0"
  instance_class         = "db.t3.micro"
  db_name                = "mydb"
  username               = "admin"
  password               = random_password.db_password.result
  parameter_group_name   = "default.mysql8.0"
  db_subnet_group_name   = aws_db_subnet_group.db_subnet_group.name
  vpc_security_group_ids = [aws_security_group.db_sg.id]
  availability_zone = "ap-northeast-1a"

  backup_retention_period = 7
  skip_final_snapshot     = true

  # Specify single AZ deployment
  multi_az = false

  tags = {
    Name = "hybird-db-instance"
  }
}