output "triggered_api_dkr_img" {
  value = null_resource.build_push_dkr_img.triggers
}

output "triggered_ws_connect_dkr_img" {
  value = null_resource.build_push_ws_connect_dkr_img.triggers
}

output "triggered_ws_disconnect_dkr_img" {
  value = null_resource.build_push_ws_disconnect_dkr_img.triggers
}

output "triggered_ws_default_dkr_img" {
  value = null_resource.build_push_ws_default_dkr_img.triggers
}