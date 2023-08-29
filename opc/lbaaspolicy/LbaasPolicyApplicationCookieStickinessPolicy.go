// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy


type LbaasPolicyApplicationCookieStickinessPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#cookie_name LbaasPolicy#cookie_name}.
	CookieName *string `field:"required" json:"cookieName" yaml:"cookieName"`
}

