---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_devices Data Source - terraform-provider-catalystcenter"
subcategory: "Inventory"
description: |-
  This data source fetches all network devices defined on the Catalyst Center.
  This data source does not report physical location (site) information. Obtain it instead from the data source
  data.catalystcenter_device_detail. To determine physical locations of multiple devices use the same data
  source with for_each Terraform meta-argument.
---

# catalystcenter_network_devices (Data Source)

This data source fetches all network devices defined on the Catalyst Center.

This data source does not report physical location (site) information. Obtain it instead from the data source
`data.catalystcenter_device_detail`. To determine physical locations of multiple devices use the same data
source with `for_each` Terraform meta-argument.

## Example Usage

```terraform
data "catalystcenter_network_devices" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `devices` (Attributes List) (see [below for nested schema](#nestedatt--devices))

<a id="nestedatt--devices"></a>
### Nested Schema for `devices`

Read-Only:

- `hostname` (String) Hostname of the network device
- `id` (String) UUID of the network device
- `management_ip_address` (String) Management IP address
- `management_state` (String) Management state of a network device. If it is not "Managed" for a device, then the `catalystcenter_network_device_detail` may fail for that device.
- `platform_id` (String) Platform identifier
- `role` (String) Role of the network device, such as `ACCESS` or `DISTRIBUTION`.
- `software_type` (String) Type of software