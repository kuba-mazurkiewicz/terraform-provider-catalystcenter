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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcFabricControlPlane(t *testing.T) {
	if os.Getenv("SDA") == "" {
		t.Skip("skipping test, set environment variable SDA")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_control_plane.test", "device_management_ip_address", "192.168.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_control_plane.test", "site_name_hierarchy", "Global/Area1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_fabric_control_plane.test", "route_distribution_protocol", "LISP_BGP"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcFabricControlPlaneConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcFabricControlPlaneConfig() string {
	config := `resource "catalystcenter_fabric_control_plane" "test" {` + "\n"
	config += `	device_management_ip_address = "192.168.10.1"` + "\n"
	config += `	site_name_hierarchy = "Global/Area1"` + "\n"
	config += `	route_distribution_protocol = "LISP_BGP"` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_fabric_control_plane" "test" {
			device_management_ip_address = "192.168.10.1"
			depends_on = [catalystcenter_fabric_control_plane.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig