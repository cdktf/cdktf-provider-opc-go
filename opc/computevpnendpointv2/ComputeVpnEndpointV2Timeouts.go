// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computevpnendpointv2


type ComputeVpnEndpointV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2#create ComputeVpnEndpointV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2#delete ComputeVpnEndpointV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2#update ComputeVpnEndpointV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

