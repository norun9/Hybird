output "lambda_invoke_arn" {
  value       = aws_lambda_function.hybird_lambda.invoke_arn
  description = "The name of the Lambda function."
}

output "lambda_function_name" {
  value       = aws_lambda_function.hybird_lambda.function_name
  description = "The name of the Lambda function."
}