data "catalystcenter_anycast_gateway" "example" {
  id                   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  fabric_id            = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  virtual_network_name = "SDA_VN1"
  ip_pool_name         = "MyPool1"
}
