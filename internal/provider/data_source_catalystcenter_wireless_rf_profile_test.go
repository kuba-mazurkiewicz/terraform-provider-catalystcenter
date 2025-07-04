// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcWirelessRFProfile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "rf_profile_name", "RF_Profile_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "default_rf_profile", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "enable_radio_type_a", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "enable_radio_type_b", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "enable_radio_type6_g_hz", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_parent_profile", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_radio_channels", "36,40,44,48,124,128,157,161"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_data_rates", "6,9,12,18,24,36,48,54"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_mandatory_data_rates", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_power_threshold_v1", "-60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_rx_sop_threshold", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_min_power_level", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_max_power_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_channel_width", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_a_custom_rx_sop_threshold", "-70"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_parent_profile", "HIGH"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_radio_channels", "1,6,11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_data_rates", "1,2,5.5,6,9,11,12,18,24,36,48,54"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_mandatory_data_rates", "1,2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_power_threshold_v1", "-60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_rx_sop_threshold", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_min_power_level", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_max_power_level", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_wireless_rf_profile.test", "radio_type_b_custom_rx_sop_threshold", "-70"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcWirelessRFProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcWirelessRFProfileConfig() string {
	config := `resource "catalystcenter_wireless_rf_profile" "test" {` + "\n"
	config += `	rf_profile_name = "RF_Profile_1"` + "\n"
	config += `	default_rf_profile = false` + "\n"
	config += `	enable_radio_type_a = true` + "\n"
	config += `	enable_radio_type_b = true` + "\n"
	config += `	enable_radio_type6_g_hz = false` + "\n"
	config += `	radio_type_a_parent_profile = "CUSTOM"` + "\n"
	config += `	radio_type_a_radio_channels = "36,40,44,48,124,128,157,161"` + "\n"
	config += `	radio_type_a_data_rates = "6,9,12,18,24,36,48,54"` + "\n"
	config += `	radio_type_a_mandatory_data_rates = "6"` + "\n"
	config += `	radio_type_a_power_threshold_v1 = -60` + "\n"
	config += `	radio_type_a_rx_sop_threshold = "CUSTOM"` + "\n"
	config += `	radio_type_a_min_power_level = 8` + "\n"
	config += `	radio_type_a_max_power_level = 20` + "\n"
	config += `	radio_type_a_channel_width = "20"` + "\n"
	config += `	radio_type_a_custom_rx_sop_threshold = -70` + "\n"
	config += `	radio_type_b_parent_profile = "HIGH"` + "\n"
	config += `	radio_type_b_radio_channels = "1,6,11"` + "\n"
	config += `	radio_type_b_data_rates = "1,2,5.5,6,9,11,12,18,24,36,48,54"` + "\n"
	config += `	radio_type_b_mandatory_data_rates = "1,2"` + "\n"
	config += `	radio_type_b_power_threshold_v1 = -60` + "\n"
	config += `	radio_type_b_rx_sop_threshold = "CUSTOM"` + "\n"
	config += `	radio_type_b_min_power_level = 8` + "\n"
	config += `	radio_type_b_max_power_level = 20` + "\n"
	config += `	radio_type_b_custom_rx_sop_threshold = -70` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_wireless_rf_profile" "test" {
			rf_profile_name = "RF_Profile_1"
			depends_on = [catalystcenter_wireless_rf_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
