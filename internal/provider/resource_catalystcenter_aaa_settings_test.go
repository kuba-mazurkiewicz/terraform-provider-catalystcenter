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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcAAASettings(t *testing.T) {
	if os.Getenv("AAA") == "" {
		t.Skip("skipping test, set environment variable AAA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_aaa_settings.test", "network_aaa_server_type", "AAA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_aaa_settings.test", "network_aaa_protocol", "RADIUS"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_aaa_settings.test", "network_aaa_primary_server_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_aaa_settings.test", "network_aaa_secondary_server_ip", "1.2.3.5"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_aaa_settings.test", "client_aaa_server_type", "AAA"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_aaa_settings.test", "client_aaa_protocol", "RADIUS"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_aaa_settings.test", "client_aaa_primary_server_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_aaa_settings.test", "client_aaa_secondary_server_ip", "1.2.3.5"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcAAASettingsPrerequisitesConfig + testAccCcAAASettingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcAAASettingsPrerequisitesConfig + testAccCcAAASettingsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccCcAAASettingsPrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcAAASettingsConfig_minimum() string {
	config := `resource "catalystcenter_aaa_settings" "test" {` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	network_aaa_server_type = "AAA"` + "\n"
	config += `	network_aaa_protocol = "RADIUS"` + "\n"
	config += `	network_aaa_primary_server_ip = "1.2.3.4"` + "\n"
	config += `	client_aaa_server_type = "AAA"` + "\n"
	config += `	client_aaa_protocol = "RADIUS"` + "\n"
	config += `	client_aaa_primary_server_ip = "1.2.3.4"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAAASettingsConfig_all() string {
	config := `resource "catalystcenter_aaa_settings" "test" {` + "\n"
	config += `	site_id = catalystcenter_area.test.id` + "\n"
	config += `	network_aaa_server_type = "AAA"` + "\n"
	config += `	network_aaa_protocol = "RADIUS"` + "\n"
	config += `	network_aaa_primary_server_ip = "1.2.3.4"` + "\n"
	config += `	network_aaa_secondary_server_ip = "1.2.3.5"` + "\n"
	config += `	network_aaa_shared_secret = "Secret123"` + "\n"
	config += `	client_aaa_server_type = "AAA"` + "\n"
	config += `	client_aaa_protocol = "RADIUS"` + "\n"
	config += `	client_aaa_primary_server_ip = "1.2.3.4"` + "\n"
	config += `	client_aaa_secondary_server_ip = "1.2.3.5"` + "\n"
	config += `	client_aaa_shared_secret = "Secret123"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll