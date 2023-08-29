// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy


type LbaasPolicySetRequestHeaderPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#header_name LbaasPolicy#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#action_when_header_exists LbaasPolicy#action_when_header_exists}.
	ActionWhenHeaderExists *string `field:"optional" json:"actionWhenHeaderExists" yaml:"actionWhenHeaderExists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#action_when_header_value_is LbaasPolicy#action_when_header_value_is}.
	ActionWhenHeaderValueIs *[]*string `field:"optional" json:"actionWhenHeaderValueIs" yaml:"actionWhenHeaderValueIs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#action_when_header_value_is_not LbaasPolicy#action_when_header_value_is_not}.
	ActionWhenHeaderValueIsNot *[]*string `field:"optional" json:"actionWhenHeaderValueIsNot" yaml:"actionWhenHeaderValueIsNot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#value LbaasPolicy#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

