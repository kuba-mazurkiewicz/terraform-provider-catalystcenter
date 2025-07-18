---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_lan_automation Resource - terraform-provider-catalystcenter"
subcategory: "LAN Automation"
description: |-
  This resource can Start LAN Automation on resource creation and Stop LAN Automation on resource deletion
---

# catalystcenter_lan_automation (Resource)

This resource can Start LAN Automation on resource creation and Stop LAN Automation on resource deletion

## Example Usage

```terraform
resource "catalystcenter_lan_automation" "example" {
  discovered_device_site_name_hierarchy = "Global/Area1/Area2"
  primary_device_management_ip_address  = "1.2.3.4"
  primary_device_interface_names        = ["HundredGigE1/0/1"]
  ip_pools = [
    {
      ip_pool_name = "POOL1"
      ip_pool_role = "MAIN_POOL"
    }
  ]
  multicast_enabled        = true
  host_name_prefix         = "TEST"
  isis_domain_password     = "cisco123"
  redistribute_isis_to_bgp = true
  discovery_level          = 2
  discovery_timeout        = 20
  discovery_devices = [
    {
      device_serial_number         = "FOC12345678"
      device_host_name             = "EDGE01"
      device_site_name_hierarchy   = "Global/Area1/Area2"
      device_management_ip_address = "10.0.0.1"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `discovered_device_site_name_hierarchy` (String) Discovered device site name.
- `ip_pools` (Attributes List) The list of IP pools with its name and role. (see [below for nested schema](#nestedatt--ip_pools))
- `primary_device_interface_names` (Set of String) The list of interfaces on primary seed via which the discovered devices are connected.
- `primary_device_management_ip_address` (String) Primary seed management IP address.

### Optional

- `discovery_devices` (Attributes Set) List of specific devices that will be LAN Automated in this session (see [below for nested schema](#nestedatt--discovery_devices))
- `discovery_level` (Number) Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier
  - Range: `1`-`5`
  - Default value: `2`
- `discovery_timeout` (Number) Discovery timeout in minutes. Until this time, the stop processing will not be triggered.
  - Range: `20`-`10080`
- `host_name_file_id` (String) File ID of the CSV file containing the host name list.
- `host_name_prefix` (String) Host name prefix which shall be assigned to the discovered device.
- `isis_domain_password` (String) ISIS domain password.
- `multicast_enabled` (Boolean) To enable underlay native multicast.
- `peer_device_management_ip_address` (String) Secondary seed management IP address.
- `redistribute_isis_to_bgp` (Boolean) Advertise LAN Automation summary route into BGP.

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--ip_pools"></a>
### Nested Schema for `ip_pools`

Required:

- `ip_pool_name` (String) Name of the IP pool.
- `ip_pool_role` (String) Role of the IP pool.
  - Choices: `MAIN_POOL`, `PHYSICAL_LINK_POOL`


<a id="nestedatt--discovery_devices"></a>
### Nested Schema for `discovery_devices`

Required:

- `device_serial_number` (String) Serial number of the device

Optional:

- `device_host_name` (String) Hostname of the device
- `device_management_ip_address` (String) Management IP Address of the device
- `device_site_name_hierarchy` (String) Site name hierarchy for the device, must be a child site of the discoveredDeviceSiteNameHierarchy or same if it’s not area type

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import catalystcenter_lan_automation.example "<id>"
```
