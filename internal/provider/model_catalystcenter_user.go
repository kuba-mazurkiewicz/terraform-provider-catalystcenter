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

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type User struct {
	Id        types.String `tfsdk:"id"`
	FirstName types.String `tfsdk:"first_name"`
	LastName  types.String `tfsdk:"last_name"`
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
	Email     types.String `tfsdk:"email"`
	RoleIds   types.Set    `tfsdk:"role_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data User) getPath() string {
	return "/dna/system/api/v1/user"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data User) toBody(ctx context.Context, state User) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "userId", state.Id.ValueString())
	}
	_ = put
	if !data.FirstName.IsNull() {
		body, _ = sjson.Set(body, "firstName", data.FirstName.ValueString())
	}
	if !data.LastName.IsNull() {
		body, _ = sjson.Set(body, "lastName", data.LastName.ValueString())
	}
	if !data.Username.IsNull() {
		body, _ = sjson.Set(body, "username", data.Username.ValueString())
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, "password", data.Password.ValueString())
	}
	if !data.Email.IsNull() {
		body, _ = sjson.Set(body, "email", data.Email.ValueString())
	}
	if !data.RoleIds.IsNull() {
		var values []string
		data.RoleIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "roleList", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *User) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("firstName"); value.Exists() {
		data.FirstName = types.StringValue(value.String())
	} else {
		data.FirstName = types.StringNull()
	}
	if value := res.Get("lastName"); value.Exists() {
		data.LastName = types.StringValue(value.String())
	} else {
		data.LastName = types.StringNull()
	}
	if value := res.Get("username"); value.Exists() {
		data.Username = types.StringValue(value.String())
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get("email"); value.Exists() {
		data.Email = types.StringValue(value.String())
	} else {
		data.Email = types.StringNull()
	}
	if value := res.Get("roleList"); value.Exists() && len(value.Array()) > 0 {
		data.RoleIds = helpers.GetStringSet(value.Array())
	} else {
		data.RoleIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *User) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("firstName"); value.Exists() && !data.FirstName.IsNull() {
		data.FirstName = types.StringValue(value.String())
	} else {
		data.FirstName = types.StringNull()
	}
	if value := res.Get("lastName"); value.Exists() && !data.LastName.IsNull() {
		data.LastName = types.StringValue(value.String())
	} else {
		data.LastName = types.StringNull()
	}
	if value := res.Get("username"); value.Exists() && !data.Username.IsNull() {
		data.Username = types.StringValue(value.String())
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get("email"); value.Exists() && !data.Email.IsNull() {
		data.Email = types.StringValue(value.String())
	} else {
		data.Email = types.StringNull()
	}
	if value := res.Get("roleList"); value.Exists() && !data.RoleIds.IsNull() {
		data.RoleIds = helpers.GetStringSet(value.Array())
	} else {
		data.RoleIds = types.SetNull(types.StringType)
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *User) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FirstName.IsNull() {
		return false
	}
	if !data.LastName.IsNull() {
		return false
	}
	if !data.Username.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	if !data.Email.IsNull() {
		return false
	}
	if !data.RoleIds.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
