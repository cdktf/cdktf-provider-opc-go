package lbaaspolicy


type LbaasPolicyLoadBalancerCookieStickinessPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#cookie_expiration_period LbaasPolicy#cookie_expiration_period}.
	CookieExpirationPeriod *float64 `field:"required" json:"cookieExpirationPeriod" yaml:"cookieExpirationPeriod"`
}

