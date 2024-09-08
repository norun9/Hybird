resource "aws_security_group" "bastion_sg" {
  name        = "bastion_sg"
  description = "Security group for EC2 instance in private subnet"
  vpc_id      = var.vpc_id

  # インバウンドルール: BastionホストからのSSH接続を許可
  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # 踏み台サーバから最新のパッケージのダウンロードができるようにするため
  egress {
    protocol    = "tcp"
    from_port   = 443
    to_port     = 443
    cidr_blocks = ["0.0.0.0/0"]
  }

  # 踏み台サーバから最新のパッケージのダウンロードができるようにするため
  egress {
    protocol    = "tcp"
    from_port   = 80
    to_port     = 80
    cidr_blocks = ["0.0.0.0/0"]
  }

  # インバウンドルール: RDSへの接続を許可 (例えば、DBがMySQLなら3306)
  ingress {
    from_port = 3306
    to_port   = 3306
    protocol  = "tcp"
    cidr_blocks = [var.vpc_cidr_block]
  }

  # アウトバウンドルール: 全てのアウトバウンドトラフィックを許可
  egress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Amazon Linux 2 AMIを取得
data "aws_ssm_parameter" "amzn2_ami" {
  name = "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
}

# TODO: SSHトンネルのためパブリックサブネットに配置する
resource "aws_instance" "bastion" {
  ami           =  data.aws_ssm_parameter.amzn2_ami.value
  instance_type = "t2.micro"  # 無料枠対象のインスタンスタイプ
  vpc_security_group_ids = [aws_security_group.bastion_sg.id]
  subnet_id              = var.subnet_id
  associate_public_ip_address = true
  iam_instance_profile   = "ec2_bastion_profile"
  key_name = "hybird-keypair"
  capacity_reservation_specification {
    capacity_reservation_preference = "none"
  }
  tags = {
    Name = "hybird-bastion"
  }
  # EBSのルートボリューム設定
  root_block_device {
    # ボリュームサイズ(GiB)
    volume_size = 8
    # ボリュームタイプ
    volume_type = "gp3"
    # GP3のIOPS
    iops = 3000
    # GP3のスループット
    throughput = 125
    # EC2終了時に削除
    delete_on_termination = true

    #暗号化KMS設定
    encrypted  = true
    kms_key_id = "alias/aws/ebs"

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