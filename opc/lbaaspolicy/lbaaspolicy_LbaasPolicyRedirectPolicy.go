package lbaaspolicy


type LbaasPolicyRedirectPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#redirect_uri LbaasPolicy#redirect_uri}.
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#response_code LbaasPolicy#response_code}.
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
}

