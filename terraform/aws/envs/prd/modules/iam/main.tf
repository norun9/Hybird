resource "aws_iam_role" "ec2_bastion_role" {
  name = "ec2_bastion_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Effect = "Allow",
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_role" "api_gateway_putlog_role" {
  name = "api_gateway_putlog_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Effect = "Allow",
      Principal = {
        Service = "apigateway.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_role_policy_attachment" "api_gateway_putlog" {
  role       = aws_iam_role.api_gateway_putlog_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonAPIGatewayPushToCloudWatchLogs"
}

resource "aws_iam_role_policy_attachment" "bastion_attach_policy" {
  role = aws_iam_role.ec2_bastion_role.name
  # EC2インスタンスへセッションマネージャーを使って接続するポリシー(AmazonSSMManagedInstanceCore)をアタッチする
  # https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AmazonSSMManagedInstanceCore.html
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
}

# Define the IAM role that the Lambda function will assume
resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Effect = "Allow",
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

# Attach policy to allow Lambda to decrypt environment variables using KMS
resource "aws_iam_policy" "lambda_kms_policy" {
  name        = "lambda_kms_policy"
  description = "Policy to allow Lambda function to decrypt environment variables using KMS"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        "Action" : ["kms:*"],
        Resource = "*"

      }
    ]
  })
}

# Attach the KMS policy to the Lambda execution role
resource "aws_iam_role_policy_attachment" "lambda_kms_policy_attachment" {
  role       = aws_iam_role.lambda_exec_role.name
  policy_arn = aws_iam_policy.lambda_kms_policy.arn
}

resource "aws_iam_policy" "lambda_invoke_policy" {
  name        = "lambda-invoke-policy"
  description = "Policy to allow invoking the Lambda function"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect   = "Allow",
        Action   = "lambda:InvokeFunction",
        Resource = "*"
      }
    ]
  })
}

# Attach the AWS managed policy for basic Lambda execution permissions
resource "aws_iam_policy_attachment" "lambda_exec_policy" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  roles      = [aws_iam_role.lambda_exec_role.name]
  name       = "LambdaExecRoleAttachment"
}

resource "aws_iam_role_policy_attachment" "lambda_invoke_policy_attachment" {
  role       = aws_iam_role.lambda_exec_role.name
  policy_arn = aws_iam_policy.lambda_invoke_policy.arn
}

# Attach the AWS managed policy for Lambda to access resources in a VPC
resource "aws_iam_policy_attachment" "lambda_vpc_access_policy" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
  roles      = [aws_iam_role.lambda_exec_role.name]
  name       = "LambdaVPCAccessRoleAttachment"
}