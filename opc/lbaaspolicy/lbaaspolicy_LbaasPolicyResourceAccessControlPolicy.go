package lbaaspolicy


type LbaasPolicyResourceAccessControlPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#disposition LbaasPolicy#disposition}.
	Disposition *string `field:"required" json:"disposition" yaml:"disposition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#denied_clients LbaasPolicy#denied_clients}.
	DeniedClients *[]*string `field:"optional" json:"deniedClients" yaml:"deniedClients"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#permitted_clients LbaasPolicy#permitted_clients}.
	PermittedClients *[]*string `field:"optional" json:"permittedClients" yaml:"permittedClients"`
}

