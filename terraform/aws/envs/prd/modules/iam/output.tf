output "lambda_exec_role_arn" {
  description = "The ARN of the IAM role to be assumed by the Lambda function"
  value       = aws_iam_role.lambda_exec_role.arn
}

output "ec2_bastion_role_name" {
  value = aws_iam_role.ec2_bastion_role.name
}