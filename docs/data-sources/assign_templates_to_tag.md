---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_assign_templates_to_tag Data Source - terraform-provider-catalystcenter"
subcategory: "Tags"
description: |-
  This data source can read the Assign Templates to Tag.
---

# catalystcenter_assign_templates_to_tag (Data Source)

This data source can read the Assign Templates to Tag.

## Example Usage

```terraform
data "catalystcenter_assign_templates_to_tag" "example" {
  tag_id = "ea505070-6bb8-493f-bff0-8058e8e03ee5"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `tag_id` (String) Tag Id to be associated with the template

### Read-Only

- `id` (String) The id of the object
- `template_ids` (Set of String) Template Ids List