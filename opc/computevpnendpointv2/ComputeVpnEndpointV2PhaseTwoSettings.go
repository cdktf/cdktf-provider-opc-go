package computevpnendpointv2


type ComputeVpnEndpointV2PhaseTwoSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2#encryption ComputeVpnEndpointV2#encryption}.
	Encryption *string `field:"required" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2#hash ComputeVpnEndpointV2#hash}.
	Hash *string `field:"required" json:"hash" yaml:"hash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2#lifetime ComputeVpnEndpointV2#lifetime}.
	Lifetime *float64 `field:"optional" json:"lifetime" yaml:"lifetime"`
}

