package lbaaspolicy


type LbaasPolicyTrustedCertificatePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#trusted_certificate LbaasPolicy#trusted_certificate}.
	TrustedCertificate *string `field:"required" json:"trustedCertificate" yaml:"trustedCertificate"`
}

