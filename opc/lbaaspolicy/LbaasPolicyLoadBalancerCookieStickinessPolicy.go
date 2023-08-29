// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy


type LbaasPolicyLoadBalancerCookieStickinessPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#cookie_expiration_period LbaasPolicy#cookie_expiration_period}.
	CookieExpirationPeriod *float64 `field:"required" json:"cookieExpirationPeriod" yaml:"cookieExpirationPeriod"`
}

