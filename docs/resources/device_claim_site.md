---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_device_claim_site Resource - terraform-provider-catalystcenter"
subcategory: "Plug and Play"
description: |-
  This resource can manage a Device Claim Site.
---

# catalystcenter_device_claim_site (Resource)

This resource can manage a Device Claim Site.

## Example Usage

```terraform
resource "catalystcenter_device_claim_site" "example" {
  device_id  = "12345678-1234-1234-1234-123456789012"
  site_id    = "12345678-1234-1234-1234-123456789012"
  type       = "Default"
  image_id   = ""
  image_skip = true
  config_id  = "template1"
  config_parameters = [
    {
      name  = "HOSTNAME"
      value = "switch1"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (String) The device ID
- `site_id` (String) The site ID
- `type` (String) The device type
  - Choices: `Default`, `StackSwitch`, `AccessPoint`, `Sensor`, `CatalystWLC`, `MobilityExpress`

### Optional

- `config_id` (String) Config (temaplate) ID. Required if `type` is `Default` or `StackSwitch`.
- `config_parameters` (Attributes List) List of config (temaplate) parameters. (see [below for nested schema](#nestedatt--config_parameters))
- `gateway` (String) Gateway IP. Required if `type` is `CatalystWLC` or `MobilityExpress`.
- `image_id` (String) Image ID. Required if `type` is `Default` or `StackSwitch`.
- `image_skip` (Boolean) Skip image provisioning.
- `ip_interface_name` (String) IP interface name. Required for Catalyst 9800 WLC.
- `rf_profile` (String) RF profile. Required if `type` is `AccessPoint`.
- `sensor_profile` (String) Sensor profile. Required if `type` is `Sensor`.
- `static_ip` (String) Static IP address. Required if `type` is `CatalystWLC` or `MobilityExpress`.
- `subnet_mask` (String) Subnet mask. Required if `type` is `CatalystWLC` or `MobilityExpress`.
- `vlan_id` (String) Vlan ID. Required for Catalyst 9800 WLC.

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--config_parameters"></a>
### Nested Schema for `config_parameters`

Required:

- `name` (String) Name of config parameter.

Optional:

- `value` (String) Value of config parameter.