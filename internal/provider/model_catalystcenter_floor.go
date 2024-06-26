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
type Floor struct {
	Id          types.String  `tfsdk:"id"`
	Name        types.String  `tfsdk:"name"`
	ParentName  types.String  `tfsdk:"parent_name"`
	FloorNumber types.Int64   `tfsdk:"floor_number"`
	RfModel     types.String  `tfsdk:"rf_model"`
	Width       types.Float64 `tfsdk:"width"`
	Length      types.Float64 `tfsdk:"length"`
	Height      types.Float64 `tfsdk:"height"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Floor) getPath() string {
	return "/dna/intent/api/v1/site"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Floor) toBody(ctx context.Context, state Floor) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	body, _ = sjson.Set(body, "type", "floor")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "site.floor.name", data.Name.ValueString())
	}
	if !data.ParentName.IsNull() {
		body, _ = sjson.Set(body, "site.floor.parentName", data.ParentName.ValueString())
	}
	if !data.FloorNumber.IsNull() {
		body, _ = sjson.Set(body, "site.floor.floorNumber", data.FloorNumber.ValueInt64())
	}
	if !data.RfModel.IsNull() {
		body, _ = sjson.Set(body, "site.floor.rfModel", data.RfModel.ValueString())
	}
	if !data.Width.IsNull() {
		body, _ = sjson.Set(body, "site.floor.width", data.Width.ValueFloat64())
	}
	if !data.Length.IsNull() {
		body, _ = sjson.Set(body, "site.floor.length", data.Length.ValueFloat64())
	}
	if !data.Height.IsNull() {
		body, _ = sjson.Set(body, "site.floor.height", data.Height.ValueFloat64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Floor) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("floorIndex"); value.Exists() {
		data.FloorNumber = types.Int64Value(value.Int())
	} else {
		data.FloorNumber = types.Int64Null()
	}
	if value := res.Get("geometry.width"); value.Exists() {
		data.Width = types.Float64Value(value.Float())
	} else {
		data.Width = types.Float64Null()
	}
	if value := res.Get("geometry.length"); value.Exists() {
		data.Length = types.Float64Value(value.Float())
	} else {
		data.Length = types.Float64Null()
	}
	if value := res.Get("geometry.height"); value.Exists() {
		data.Height = types.Float64Value(value.Float())
	} else {
		data.Height = types.Float64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *Floor) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("floorIndex"); value.Exists() && !data.FloorNumber.IsNull() {
		data.FloorNumber = types.Int64Value(value.Int())
	} else {
		data.FloorNumber = types.Int64Null()
	}
	if value := res.Get("geometry.width"); value.Exists() && !data.Width.IsNull() {
		data.Width = types.Float64Value(value.Float())
	} else {
		data.Width = types.Float64Null()
	}
	if value := res.Get("geometry.length"); value.Exists() && !data.Length.IsNull() {
		data.Length = types.Float64Value(value.Float())
	} else {
		data.Length = types.Float64Null()
	}
	if value := res.Get("geometry.height"); value.Exists() && !data.Height.IsNull() {
		data.Height = types.Float64Value(value.Float())
	} else {
		data.Height = types.Float64Null()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *Floor) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.ParentName.IsNull() {
		return false
	}
	if !data.FloorNumber.IsNull() {
		return false
	}
	if !data.RfModel.IsNull() {
		return false
	}
	if !data.Width.IsNull() {
		return false
	}
	if !data.Length.IsNull() {
		return false
	}
	if !data.Height.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
