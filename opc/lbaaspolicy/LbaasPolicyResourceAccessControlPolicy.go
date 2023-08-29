// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy


type LbaasPolicyResourceAccessControlPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#disposition LbaasPolicy#disposition}.
	Disposition *string `field:"required" json:"disposition" yaml:"disposition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#denied_clients LbaasPolicy#denied_clients}.
	DeniedClients *[]*string `field:"optional" json:"deniedClients" yaml:"deniedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#permitted_clients LbaasPolicy#permitted_clients}.
	PermittedClients *[]*string `field:"optional" json:"permittedClients" yaml:"permittedClients"`
}

