// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy


type LbaasPolicyLoadBalancingMechanismPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#load_balancing_mechanism LbaasPolicy#load_balancing_mechanism}.
	LoadBalancingMechanism *string `field:"required" json:"loadBalancingMechanism" yaml:"loadBalancingMechanism"`
}

