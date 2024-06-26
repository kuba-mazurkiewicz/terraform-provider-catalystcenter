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
type CredentialsSNMPv3 struct {
	Id              types.String `tfsdk:"id"`
	Description     types.String `tfsdk:"description"`
	Username        types.String `tfsdk:"username"`
	PrivacyType     types.String `tfsdk:"privacy_type"`
	PrivacyPassword types.String `tfsdk:"privacy_password"`
	AuthType        types.String `tfsdk:"auth_type"`
	AuthPassword    types.String `tfsdk:"auth_password"`
	SnmpMode        types.String `tfsdk:"snmp_mode"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data CredentialsSNMPv3) getPath() string {
	return "/dna/intent/api/v2/global-credential"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CredentialsSNMPv3) toBody(ctx context.Context, state CredentialsSNMPv3) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
		body, _ = sjson.Set(body, "snmpV3.0.id", state.Id.ValueString())
	}
	_ = put
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "snmpV3.0.description", data.Description.ValueString())
	}
	if !data.Username.IsNull() {
		body, _ = sjson.Set(body, "snmpV3.0.username", data.Username.ValueString())
	}
	if !data.PrivacyType.IsNull() {
		body, _ = sjson.Set(body, "snmpV3.0.privacyType", data.PrivacyType.ValueString())
	}
	if !data.PrivacyPassword.IsNull() {
		body, _ = sjson.Set(body, "snmpV3.0.privacyPassword", data.PrivacyPassword.ValueString())
	}
	if !data.AuthType.IsNull() {
		body, _ = sjson.Set(body, "snmpV3.0.authType", data.AuthType.ValueString())
	}
	if !data.AuthPassword.IsNull() {
		body, _ = sjson.Set(body, "snmpV3.0.authPassword", data.AuthPassword.ValueString())
	}
	if !data.SnmpMode.IsNull() {
		body, _ = sjson.Set(body, "snmpV3.0.snmpMode", data.SnmpMode.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CredentialsSNMPv3) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("username"); value.Exists() {
		data.Username = types.StringValue(value.String())
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get("privacyType"); value.Exists() {
		data.PrivacyType = types.StringValue(value.String())
	} else {
		data.PrivacyType = types.StringNull()
	}
	if value := res.Get("authType"); value.Exists() {
		data.AuthType = types.StringValue(value.String())
	} else {
		data.AuthType = types.StringNull()
	}
	if value := res.Get("snmpMode"); value.Exists() {
		data.SnmpMode = types.StringValue(value.String())
	} else {
		data.SnmpMode = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *CredentialsSNMPv3) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("username"); value.Exists() && !data.Username.IsNull() {
		data.Username = types.StringValue(value.String())
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get("privacyType"); value.Exists() && !data.PrivacyType.IsNull() {
		data.PrivacyType = types.StringValue(value.String())
	} else {
		data.PrivacyType = types.StringNull()
	}
	if value := res.Get("authType"); value.Exists() && !data.AuthType.IsNull() {
		data.AuthType = types.StringValue(value.String())
	} else {
		data.AuthType = types.StringNull()
	}
	if value := res.Get("snmpMode"); value.Exists() && !data.SnmpMode.IsNull() {
		data.SnmpMode = types.StringValue(value.String())
	} else {
		data.SnmpMode = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *CredentialsSNMPv3) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Description.IsNull() {
		return false
	}
	if !data.Username.IsNull() {
		return false
	}
	if !data.PrivacyType.IsNull() {
		return false
	}
	if !data.PrivacyPassword.IsNull() {
		return false
	}
	if !data.AuthType.IsNull() {
		return false
	}
	if !data.AuthPassword.IsNull() {
		return false
	}
	if !data.SnmpMode.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
