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
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TimeZoneSettings struct {
	Id         types.String `tfsdk:"id"`
	SiteId     types.String `tfsdk:"site_id"`
	Identifier types.String `tfsdk:"identifier"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TimeZoneSettings) getPath() string {
	return fmt.Sprintf("/dna/intent/api/v1/sites/%v/timeZoneSettings", url.QueryEscape(data.SiteId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TimeZoneSettings) toBody(ctx context.Context, state TimeZoneSettings) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.Identifier.IsNull() {
		body, _ = sjson.Set(body, "timeZone.identifier", data.Identifier.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TimeZoneSettings) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	data.Id = types.StringValue(fmt.Sprint(data.SiteId.ValueString()))
	if value := res.Get("response.timeZone.identifier"); value.Exists() {
		data.Identifier = types.StringValue(value.String())
	} else {
		data.Identifier = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TimeZoneSettings) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.timeZone.identifier"); value.Exists() && !data.Identifier.IsNull() {
		data.Identifier = types.StringValue(value.String())
	} else {
		data.Identifier = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TimeZoneSettings) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Identifier.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
