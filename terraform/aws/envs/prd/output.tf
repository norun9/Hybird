# output "triggered_api_dkr_img" {
#   value       = null_resource.build_push_dkr_img.triggers
#   description = "Triggers for building and pushing the API Docker image."
# }
#
# output "triggered_ws_connect_dkr_img" {
#   value       = null_resource.build_push_ws_connect_dkr_img.triggers
#   description = "Triggers for building and pushing the WebSocket connect Docker image."
# }
#
# output "triggered_ws_disconnect_dkr_img" {
#   value       = null_resource.build_push_ws_disconnect_dkr_img.triggers
#   description = "Triggers for building and pushing the WebSocket disconnect Docker image."
# }
#
# output "triggered_ws_default_dkr_img" {
#   value       = null_resource.build_push_ws_default_dkr_img.triggers
#   description = "Triggers for building and pushing the WebSocket default Docker image."
# }