// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy


type LbaasPolicyTrustedCertificatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#trusted_certificate LbaasPolicy#trusted_certificate}.
	TrustedCertificate *string `field:"required" json:"trustedCertificate" yaml:"trustedCertificate"`
}

