package provider


type OpcProviderConfig struct {
	// The OPC identity domain for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#identity_domain OpcProvider#identity_domain}
	IdentityDomain *string `field:"required" json:"identityDomain" yaml:"identityDomain"`
	// The user password for OPC API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#password OpcProvider#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name for OPC API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#user OpcProvider#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#alias OpcProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The HTTP endpoint for OPC API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#endpoint OpcProvider#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Skip TLS Verification for self-signed certificates. Should only be used if absolutely required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#insecure OpcProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The HTTP endpoint for the Load Balancer Classic service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#lbaas_endpoint OpcProvider#lbaas_endpoint}
	LbaasEndpoint *string `field:"optional" json:"lbaasEndpoint" yaml:"lbaasEndpoint"`
	// Maximum number retries to wait for a successful response when operating on resources within OPC (defaults to 1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#max_retries OpcProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The HTTP endpoint for Oracle Storage operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#storage_endpoint OpcProvider#storage_endpoint}
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
	// The Storage Service ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs#storage_service_id OpcProvider#storage_service_id}
	StorageServiceId *string `field:"optional" json:"storageServiceId" yaml:"storageServiceId"`
}

