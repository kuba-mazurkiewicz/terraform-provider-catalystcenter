---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_interface Resource - terraform-provider-catalystcenter"
subcategory: "Wireless"
description: |-
  This resource can manage a Wireless Interface.
---

# catalystcenter_wireless_interface (Resource)

This resource can manage a Wireless Interface.

## Example Usage

```terraform
resource "catalystcenter_wireless_interface" "example" {
  interface_name = "GigabitEthernet1/0/1"
  vlan_id        = 100
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_name` (String) Interface Name
- `vlan_id` (Number) VLAN ID
  - Range: `1`-`4094`

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_wireless_interface.example "<interface_name>"
```
