---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_lan_automation Data Source - terraform-provider-catalystcenter"
subcategory: "LAN Automation"
description: |-
  This data source can read the LAN Automation.
---

# catalystcenter_lan_automation (Data Source)

This data source can read the LAN Automation.

## Example Usage

```terraform
data "catalystcenter_lan_automation" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `discovered_device_site_name_hierarchy` (String) Discovered device site name.
- `discovery_devices` (Attributes Set) List of specific devices that will be LAN Automated in this session (see [below for nested schema](#nestedatt--discovery_devices))
- `discovery_level` (Number) Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier
- `discovery_timeout` (Number) Discovery timeout in minutes. Until this time, the stop processing will not be triggered.
- `host_name_file_id` (String) File ID of the CSV file containing the host name list.
- `host_name_prefix` (String) Host name prefix which shall be assigned to the discovered device.
- `ip_pools` (Attributes List) The list of IP pools with its name and role. (see [below for nested schema](#nestedatt--ip_pools))
- `isis_domain_password` (String) ISIS domain password.
- `multicast_enabled` (Boolean) To enable underlay native multicast.
- `peer_device_management_ip_address` (String) Secondary seed management IP address.
- `primary_device_interface_names` (Set of String) The list of interfaces on primary seed via which the discovered devices are connected.
- `primary_device_management_ip_address` (String) Primary seed management IP address.
- `redistribute_isis_to_bgp` (Boolean) Advertise LAN Automation summary route into BGP.

<a id="nestedatt--discovery_devices"></a>
### Nested Schema for `discovery_devices`

Read-Only:

- `device_host_name` (String) Hostname of the device
- `device_management_ip_address` (String) Management IP Address of the device
- `device_serial_number` (String) Serial number of the device
- `device_site_name_hierarchy` (String) Site name hierarchy for the device, must be a child site of the discoveredDeviceSiteNameHierarchy or same if it’s not area type


<a id="nestedatt--ip_pools"></a>
### Nested Schema for `ip_pools`

Read-Only:

- `ip_pool_name` (String) Name of the IP pool.
- `ip_pool_role` (String) Role of the IP pool.
