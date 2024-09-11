# WebSocket API の作成
resource "aws_apigatewayv2_api" "websocket_api" {
  name                      = "websocket-api"
  protocol_type             = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

# WebSocketの $connect ルート
resource "aws_apigatewayv2_route" "connect_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$connect"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_connect_integration.id}"
}

# WebSocketの $disconnect ルート
resource "aws_apigatewayv2_route" "disconnect_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$disconnect"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_disconnect_integration.id}"
}

# WebSocketの $default ルート
resource "aws_apigatewayv2_route" "default_route" {
  api_id    = aws_apigatewayv2_api.websocket_api.id
  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_default_integration.id}"
}

data "aws_caller_identity" "current" {}

# Lambda ロールの作成
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

  # Lambda に必要な権限を追加（DynamoDB, CloudWatch Logs, API Gatewayなど）
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

# Lambda 関数 (connect)
resource "aws_lambda_function" "connect_lambda" {
  function_name = "hybird-websocket-connect"
  image_uri     = data.aws_ecr_image.ws_connect.image_uri # ECR イメージのURIを指定
  role          = aws_iam_role.lambda_execution_role.arn
  timeout       = 15
  package_type = "Image"

  environment {
    variables = {
      TABLE_NAME = aws_dynamodb_table.connections_table.name
    }
  }
}

# Lambda 関数 (disconnect)
resource "aws_lambda_function" "disconnect_lambda" {
  function_name = "hybird-websocket-disconnect"
  image_uri     = data.aws_ecr_image.ws_disconnect.image_uri # ECR イメージのURIを指定
  role          = aws_iam_role.lambda_execution_role.arn
  timeout       = 15
  package_type = "Image"

  environment {
    variables = {
      TABLE_NAME = aws_dynamodb_table.connections_table.name
    }
  }
}

# Lambda 関数 (default)
resource "aws_lambda_function" "default_lambda" {
  function_name = "hybird-websocket-default"
  image_uri     = data.aws_ecr_image.ws_default.image_uri # ECR イメージのURIを指定
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

# Lambda インテグレーション (connect)
resource "aws_apigatewayv2_integration" "lambda_connect_integration" {
  api_id           = aws_apigatewayv2_api.websocket_api.id
  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.connect_lambda.invoke_arn
}

# Lambda インテグレーション (disconnect)
resource "aws_apigatewayv2_integration" "lambda_disconnect_integration" {
  api_id           = aws_apigatewayv2_api.websocket_api.id
  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.disconnect_lambda.invoke_arn
}

# Lambda インテグレーション (default)
resource "aws_apigatewayv2_integration" "lambda_default_integration" {
  api_id           = aws_apigatewayv2_api.websocket_api.id
  integration_type = "AWS_PROXY"
  integration_uri  = aws_lambda_function.default_lambda.invoke_arn
}

# WebSocket のデプロイ
resource "aws_apigatewayv2_stage" "websocket_stage" {
  api_id      = aws_apigatewayv2_api.websocket_api.id
  name        = "prd"
  auto_deploy = true
}
