---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_fabric_l2_virtual_network Resource - terraform-provider-catalystcenter"
subcategory: "SDA"
description: |-
  This resource can manage a Fabric L2 Virtual Network.
---

# catalystcenter_fabric_l2_virtual_network (Resource)

This resource can manage a Fabric L2 Virtual Network.

## Example Usage

```terraform
resource "catalystcenter_fabric_l2_virtual_network" "example" {
  fabric_id               = "5e6f7b3a-2b0b-4a7d-8b1c-0d4b1cd5e1b1"
  vlan_name               = "VLAN401"
  vlan_id                 = 401
  traffic_type            = "DATA"
  fabric_enabled_wireless = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `fabric_id` (String) ID of the fabric this layer 2 virtual network is to be assigned to
- `traffic_type` (String) The type of traffic that is served
  - Choices: `DATA`, `VOICE`
- `vlan_name` (String) Name of the VLAN of the layer 2 virtual network. Must contain only alphanumeric characters, underscores, and hyphens

### Optional

- `associated_l3_virtual_network_name` (String) Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring. The layer 3 virtual network must have already been added to the fabric before association. This field must either be present in all payload elements or none
- `fabric_enabled_wireless` (Boolean) Set to true to enable wireless. Default is false
- `vlan_id` (Number) ID of the VLAN of the layer 2 virtual network. Allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, and 2046. If deploying on a fabric zone, this vlanId must match the vlanId of the corresponding layer 2 virtual network on the fabric site

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_fabric_l2_virtual_network.example "<fabric_id>,<vlan_name>"
```