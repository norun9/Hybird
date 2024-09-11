# DynamoDB テーブル (接続情報)
resource "aws_dynamodb_table" "connections_table" {
  name         = "Connections"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ConnectionId"

  attribute {
    name = "ConnectionId"
    type = "S"
  }
}