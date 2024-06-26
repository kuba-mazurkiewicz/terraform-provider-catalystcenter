---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_profile Data Source - terraform-provider-catalystcenter"
subcategory: "Network Settings"
description: |-
  This data source can read the Network Profile.
---

# catalystcenter_network_profile (Data Source)

This data source can read the Network Profile.

## Example Usage

```terraform
data "catalystcenter_network_profile" "example" {
  id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the object
- `name` (String) The name of the network profile

### Read-Only

- `templates` (Attributes List) (see [below for nested schema](#nestedatt--templates))
- `type` (String) Profile type

<a id="nestedatt--templates"></a>
### Nested Schema for `templates`

Read-Only:

- `attributes` (Attributes List) (see [below for nested schema](#nestedatt--templates--attributes))
- `type` (String) Template type

<a id="nestedatt--templates--attributes"></a>
### Nested Schema for `templates.attributes`

Read-Only:

- `template_id` (String) Template ID
