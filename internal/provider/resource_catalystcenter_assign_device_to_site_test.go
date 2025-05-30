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
func TestAccCcAssignDeviceToSite(t *testing.T) {
	if os.Getenv("INVENTORY") == "" {
		t.Skip("skipping test, set environment variable INVENTORY")
	}
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcAssignDeviceToSitePrerequisitesConfig + testAccCcAssignDeviceToSiteConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcAssignDeviceToSitePrerequisitesConfig + testAccCcAssignDeviceToSiteConfig_all(),
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
const testAccCcAssignDeviceToSitePrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_building" "test" {
  name        = "Building1"
  parent_name = "${catalystcenter_area.test.parent_name}/${catalystcenter_area.test.name}"
  country     = "United States"
  address     = "150 W Tasman Dr, San Jose"
  latitude    = 37.338
  longitude   = -121.832
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcAssignDeviceToSiteConfig_minimum() string {
	config := `resource "catalystcenter_assign_device_to_site" "test" {` + "\n"
	config += `	device_ids = ["024f383c-14a5-421c-b21d-b80910cde422"]` + "\n"
	config += `	site_id = catalystcenter_building.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAssignDeviceToSiteConfig_all() string {
	config := `resource "catalystcenter_assign_device_to_site" "test" {` + "\n"
	config += `	device_ids = ["024f383c-14a5-421c-b21d-b80910cde422"]` + "\n"
	config += `	site_id = catalystcenter_building.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
