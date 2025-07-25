---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_pnp_device Resource - terraform-provider-catalystcenter"
subcategory: "Plug and Play"
description: |-
  This resource can manage a PnP Device.
---

# catalystcenter_pnp_device (Resource)

This resource can manage a PnP Device.

## Example Usage

```terraform
resource "catalystcenter_pnp_device" "example" {
  serial_number = "FOC12345678"
  stack         = false
  pid           = "C9300-24P"
  hostname      = "switch1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `hostname` (String) Device hostname
- `pid` (String) Device product ID
- `serial_number` (String) Device serial number

### Optional

- `stack` (Boolean) Device is a stacked switch

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import catalystcenter_pnp_device.example "<serial_number>"
```
