package lbaaspolicy


type LbaasPolicyApplicationCookieStickinessPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#cookie_name LbaasPolicy#cookie_name}.
	CookieName *string `field:"required" json:"cookieName" yaml:"cookieName"`
}

