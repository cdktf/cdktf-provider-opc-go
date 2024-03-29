// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy


type LbaasPolicyRedirectPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#redirect_uri LbaasPolicy#redirect_uri}.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#response_code LbaasPolicy#response_code}.
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
}

