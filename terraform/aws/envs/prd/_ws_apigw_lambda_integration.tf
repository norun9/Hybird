# Create WebSocket API
resource "aws_apigatewayv2_api" "websocket_api" {
  name                      = "websocket-api"
  protocol_type             = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

# Route $connect of WebSocket
resource "aws_apigatewayv2_route" "connect_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$connect"
  # The "integrations/" prefix is used to indicate that the route is linked to an AWS Lambda or HTTP integration,
  # followed by the unique integration ID. This ensures that the WebSocket API knows which integration to invoke.
  target    = "integrations/${aws_apigatewayv2_integration.lambda_connect_integration.id}"
}

# Route $disconnect of WebSocket
resource "aws_apigatewayv2_route" "disconnect_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$disconnect"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_disconnect_integration.id}"
}

# Route $default of WebSocket
resource "aws_apigatewayv2_route" "default_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_default_integration.id}"
}

# This data source retrieves information about the IAM identity currently making
# the AWS API calls, such as the AWS account ID, user ARN (Amazon Resource Name),
# and user ID. It's useful when you need to reference details about the caller's
# identity, such as determining the account ID for resource tagging or creating
# specific permissions.
data "aws_caller_identity" "current" {}

# Create Lambda Role
resource "aws_iam_role" "lambda_execution_role" {
  name = "lambda_execution_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action    = "sts:AssumeRole",
      Effect    = "Allow",
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })

  inline_policy {
    name = "lambda-dynamodb-policy"
    policy = jsonencode({
      Version = "2012-10-17",
      Statement = [
        {
          Effect = "Allow",
          Action = [
            "dynamodb:PutItem",
            "dynamodb:GetItem",
            "dynamodb:DeleteItem",
            "dynamodb:Scan"
          ],
          # Add required policy to manipulate Connections table of DynamoDB
          Resource = "arn:aws:dynamodb:${var.aws_region}:${data.aws_caller_identity.current.account_id}:table/Connections"
        },
        {
          Effect = "Allow",
          Action = [
            "logs:CreateLogGroup",
            "logs:CreateLogStream",
            "logs:PutLogEvents"
          ],
          Resource = "*"
        }
      ]
    })
  }
}

# Create Lambda function required for connect
resource "aws_lambda_function" "connect_lambda" {
  function_name = "hybird-websocket-connect"
  image_uri     = data.aws_ecr_image.ws_connect.image_uri # specified URI of ECR image
  role          = aws_iam_role.lambda_execution_role.arn
  timeout       = 15
  package_type = "Image"

  environment {
    variables = {
      TABLE_NAME = aws_dynamodb_table.connections_table.name
    }
  }
}

# Create Lambda function required for disconnect
resource "aws_lambda_function" "disconnect_lambda" {
  function_name = "hybird-websocket-disconnect"
  image_uri     = data.aws_ecr_image.ws_disconnect.image_uri # specified URI of ECR image
  role          = aws_iam_role.lambda_execution_role.arn
  timeout       = 15
  package_type = "Image"

  environment {
    variables = {
      TABLE_NAME = aws_dynamodb_table.connections_table.name
    }
  }
}

# Create Lambda function required for default
resource "aws_lambda_function" "default_lambda" {
  function_name = "hybird-websocket-default"
  image_uri     = data.aws_ecr_image.ws_default.image_uri # specified URI of ECR image
  role          = aws_iam_role.lambda_execution_role.arn
  timeout       = 15
  package_type = "Image"

  environment {
    variables = {
      TABLE_NAME = aws_dynamodb_table.connections_table.name
      API_GATEWAY_ENDPOINT = aws_apigatewayv2_api.websocket_api.api_endpoint
    }
  }
}

# Create Lambda integration required for connect
resource "aws_apigatewayv2_integration" "lambda_connect_integration" {
  api_id           = aws_apigatewayv2_api.websocket_api.id
  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.connect_lambda.invoke_arn
}

# Create Lambda integration required for disconnect
resource "aws_apigatewayv2_integration" "lambda_disconnect_integration" {
  api_id           = aws_apigatewayv2_api.websocket_api.id
  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.disconnect_lambda.invoke_arn
}

# Create Lambda integration required for default
resource "aws_apigatewayv2_integration" "lambda_default_integration" {
  api_id           = aws_apigatewayv2_api.websocket_api.id
  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.default_lambda.invoke_arn
}

# Deploy WebSocket API Gateway to the stage as prd
resource "aws_apigatewayv2_stage" "websocket_stage" {
  api_id      = aws_apigatewayv2_api.websocket_api.id
  name        = "prd"
  auto_deploy = true
}