---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_rf_profile Resource - terraform-provider-catalystcenter"
subcategory: "Wireless"
description: |-
  This resource manages a Wireless RF Profile.  Updating or re-creating this resource might lead to subsequent failures when modifying the resource due to a known API issue.
---

# catalystcenter_wireless_rf_profile (Resource)

This resource manages a Wireless RF Profile. <p/> Updating or re-creating this resource might lead to subsequent failures when modifying the resource due to a known API issue.

## Example Usage

```terraform
resource "catalystcenter_wireless_rf_profile" "example" {
  name                              = "RF_Profile_1"
  default_rf_profile                = false
  enable_radio_type_a               = true
  enable_radio_type_b               = true
  enable_radio_type_c               = false
  channel_width                     = "20"
  enable_custom                     = true
  enable_brown_field                = false
  radio_type_a_parent_profile       = "CUSTOM"
  radio_type_a_radio_channels       = "36,40,44,48,52,56,60,64,144,149,153,157,161,165,169,173"
  radio_type_a_data_rates           = "6,9,12,18,24,36,48,54"
  radio_type_a_mandatory_data_rates = "12,24"
  radio_type_a_power_threshold_v1   = -60
  radio_type_a_rx_sop_threshold     = "LOW"
  radio_type_a_min_power_level      = 8
  radio_type_a_max_power_level      = 20
  radio_type_b_parent_profile       = "CUSTOM"
  radio_type_b_radio_channels       = "1,6,11"
  radio_type_b_data_rates           = "9,11,12,18,24,36,48,54"
  radio_type_b_mandatory_data_rates = "12"
  radio_type_b_power_threshold_v1   = -60
  radio_type_b_rx_sop_threshold     = "LOW"
  radio_type_b_min_power_level      = 8
  radio_type_b_max_power_level      = 20
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `channel_width` (String) Channel Width
- `default_rf_profile` (Boolean) is Default Rf Profile
- `enable_brown_field` (Boolean) Enable Brown Field
- `enable_custom` (Boolean) Enable Custom
- `enable_radio_type_a` (Boolean) Enable Radio Type A
- `enable_radio_type_b` (Boolean) Enable Radio Type B
- `name` (String) RF Profile Name

### Optional

- `enable_radio_type_c` (Boolean) Enable Radio Type C (6GHz)
- `radio_type_a_data_rates` (String) Radio TypeA Properties - Data Rates
- `radio_type_a_mandatory_data_rates` (String) Radio TypeA Properties - Mandatory Data Rates
- `radio_type_a_max_power_level` (Number) Radio TypeA Properties - Max Power Level
- `radio_type_a_min_power_level` (Number) Radio TypeA Properties - Min Power Level
- `radio_type_a_parent_profile` (String) Radio TypeA Properties - Parent Profile
- `radio_type_a_power_threshold_v1` (Number) Radio TypeA Properties - Power Threshold V1
- `radio_type_a_radio_channels` (String) Radio TypeA Properties - Radio Channels
- `radio_type_a_rx_sop_threshold` (String) Radio TypeA Properties - Rx Sop Threshold
- `radio_type_b_data_rates` (String) Radio TypeB Properties - Data Rates
- `radio_type_b_mandatory_data_rates` (String) Radio TypeB Properties - Mandatory Data Rates
- `radio_type_b_max_power_level` (Number) Radio TypeB Properties - Max Power Level
- `radio_type_b_min_power_level` (Number) Radio TypeB Properties - Min Power Level
- `radio_type_b_parent_profile` (String) Radio TypeB Properties - Parent Profile
- `radio_type_b_power_threshold_v1` (Number) Radio TypeB Properties - Power Threshold V1
- `radio_type_b_radio_channels` (String) Radio TypeB Properties - Radio Channels
- `radio_type_b_rx_sop_threshold` (String) Radio TypeB Properties - Rx Sop Threshold
- `radio_type_c_data_rates` (String) Radio TypeC Properties - Data Rates
- `radio_type_c_mandatory_data_rates` (String) Radio TypeC Properties - Mandatory Data Rates
- `radio_type_c_max_power_level` (Number) Radio TypeC Properties - Max Power Level
- `radio_type_c_min_power_level` (Number) Radio TypeC Properties - Min Power Level
- `radio_type_c_parent_profile` (String) Radio TypeC Properties - Parent Profile
- `radio_type_c_power_threshold_v1` (Number) Radio TypeC Properties - Power Threshold V1
- `radio_type_c_radio_channels` (String) Radio TypeC Properties - Radio Channels
- `radio_type_c_rx_sop_threshold` (String) Radio TypeC Properties - Rx Sop Threshold

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_wireless_rf_profile.example "<name>"
```
