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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AssignDeviceToSite struct {
	Id        types.String `tfsdk:"id"`
	DeviceIds types.Set    `tfsdk:"device_ids"`
	SiteId    types.String `tfsdk:"site_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AssignDeviceToSite) getPath() string {
	return "/dna/intent/api/v1/networkDevices/assignToSite/apply"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data AssignDeviceToSite) getPathDelete() string {
	return "/dna/intent/api/v1/networkDevices/unassignFromSite/apply"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AssignDeviceToSite) toBody(ctx context.Context, state AssignDeviceToSite) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.DeviceIds.IsNull() {
		var values []string
		data.DeviceIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "deviceIds", values)
	}
	if !data.SiteId.IsNull() {
		body, _ = sjson.Set(body, "siteId", data.SiteId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AssignDeviceToSite) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.siteId"); value.Exists() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AssignDeviceToSite) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.siteId"); value.Exists() && !data.SiteId.IsNull() {
		data.SiteId = types.StringValue(value.String())
	} else {
		data.SiteId = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AssignDeviceToSite) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.DeviceIds.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
