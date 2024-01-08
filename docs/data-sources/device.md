---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_device Data Source - terraform-provider-catalystcenter"
subcategory: "Plug and Play"
description: |-
  This data source can read the Device.
---

# catalystcenter_device (Data Source)

This data source can read the Device.

## Example Usage

```terraform
data "catalystcenter_device" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the object
- `serial_number` (String) Device serial number

### Read-Only

- `hostname` (String) Device hostname
- `pid` (String) Device product ID
- `stack` (Boolean) Device is a stacked switch