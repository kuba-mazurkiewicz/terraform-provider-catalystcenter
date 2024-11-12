resource "catalystcenter_aaa_settings" "example" {
  site_id                         = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  network_aaa_server_type         = "AAA"
  network_aaa_protocol            = "RADIUS"
  network_aaa_primary_server_ip   = "1.2.3.4"
  network_aaa_secondary_server_ip = "1.2.3.5"
  network_aaa_shared_secret       = "Secret123"
  client_aaa_server_type          = "AAA"
  client_aaa_protocol             = "RADIUS"
  client_aaa_primary_server_ip    = "1.2.3.4"
  client_aaa_secondary_server_ip  = "1.2.3.5"
  client_aaa_shared_secret        = "Secret123"
}