---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_profile_for_sites_assignments Data Source - terraform-provider-catalystcenter"
subcategory: "Network Profiles"
description: |-
  This data source can read the Network Profile for Sites Assignments.
---

# catalystcenter_network_profile_for_sites_assignments (Data Source)

This data source can read the Network Profile for Sites Assignments.

## Example Usage

```terraform
data "catalystcenter_network_profile_for_sites_assignments" "example" {
  network_profile_id = "ff19242d-9989-434c-beee-92b16b6b30a3"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_profile_id` (String) Network-Profile Id to be associated

### Read-Only

- `id` (String) The id of the object
- `items` (Attributes Set) List of sites where the network profile is assigned (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String) Site Id
