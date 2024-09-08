# Output the database endpoint (hostname)
output "db_host" {
  description = "The hostname of the RDS instance"
  value       = aws_db_instance.hybird_db.endpoint
}

# Output the database name
output "db_name" {
  description = "The name of the database"
  value       = aws_db_instance.hybird_db.db_name
}

# Output the database username
output "db_user" {
  description = "The username for the RDS instance"
  value       = aws_db_instance.hybird_db.username
}

# Output the database password (Sensitive)
output "db_pass" {
  description = "The password for the RDS instance"
  value       = aws_db_instance.hybird_db.password
  sensitive   = true
}