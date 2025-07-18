---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_tag Resource - terraform-provider-catalystcenter"
subcategory: "Tags"
description: |-
  This resource can manage a Tag.
---

# catalystcenter_tag (Resource)

This resource can manage a Tag.

## Example Usage

```terraform
resource "catalystcenter_tag" "example" {
  name        = "Tag1"
  description = "Tag1 Description"
  system_tag  = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `description` (String) Description of the tag
- `dynamic_rules` (Attributes List) Dynamic rules details (see [below for nested schema](#nestedatt--dynamic_rules))
- `system_tag` (Boolean) true for system created tags, false for user defined tag

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--dynamic_rules"></a>
### Nested Schema for `dynamic_rules`

Optional:

- `items` (Attributes List) items details, multiple rules can be defined by items (see [below for nested schema](#nestedatt--dynamic_rules--items))
- `member_type` (String) memberType of the tag
  - Choices: `networkdevice`, `interface`
- `name` (String) Name of the parameter (e.g. for interface:portName,adminStatus,speed,status,description. for networkdevice:family,series,hostname,managementIpAddress,groupNameHierarchy,softwareVersion)
- `operation` (String) Operation of the rule
  - Choices: `OR`, `IN`, `EQ`, `LIKE`, `ILIKE`, `AND`
- `value` (String) Value of the parameter (e.g. for portName:1/0/1,for adminStatus,status:up/down, for speed: any integer value, for description: any valid string, for family:switches, for series:C3650, for managementIpAddress:10.197.124.90, groupNameHierarchy:Global, softwareVersion: 16.9.1)
- `values` (List of String) values of the parameter,Only one of the value or values can be used for the given parameter. (for managementIpAddress e.g. ["10.197.124.90","10.197.124.91"])

<a id="nestedatt--dynamic_rules--items"></a>
### Nested Schema for `dynamic_rules.items`

Optional:

- `name` (String) Name of the parameter (e.g. managementIpAddress,hostname)
- `operation` (String) Operation of the rule
  - Choices: `OR`, `IN`, `EQ`, `LIKE`, `ILIKE`, `AND`
- `value` (String) Value of the parameter (e.g. %10%,%NA%)

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import catalystcenter_tag.example "<name>"
```
