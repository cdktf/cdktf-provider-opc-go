package computevpnendpointv2


type ComputeVpnEndpointV2PhaseOneSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_vpn_endpoint_v2#dh_group ComputeVpnEndpointV2#dh_group}.
	DhGroup *string `field:"required" json:"dhGroup" yaml:"dhGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_vpn_endpoint_v2#encryption ComputeVpnEndpointV2#encryption}.
	Encryption *string `field:"required" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_vpn_endpoint_v2#hash ComputeVpnEndpointV2#hash}.
	Hash *string `field:"required" json:"hash" yaml:"hash"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_vpn_endpoint_v2#lifetime ComputeVpnEndpointV2#lifetime}.
	Lifetime *float64 `field:"optional" json:"lifetime" yaml:"lifetime"`
}

