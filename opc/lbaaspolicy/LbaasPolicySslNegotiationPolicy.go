// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy


type LbaasPolicySslNegotiationPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#port LbaasPolicy#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#ssl_protocol LbaasPolicy#ssl_protocol}.
	SslProtocol *[]*string `field:"required" json:"sslProtocol" yaml:"sslProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#server_order_preference LbaasPolicy#server_order_preference}.
	ServerOrderPreference *string `field:"optional" json:"serverOrderPreference" yaml:"serverOrderPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#ssl_ciphers LbaasPolicy#ssl_ciphers}.
	SslCiphers *[]*string `field:"optional" json:"sslCiphers" yaml:"sslCiphers"`
}

