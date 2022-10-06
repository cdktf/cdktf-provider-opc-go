package computeinstance


type ComputeInstanceNetworkingInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#index ComputeInstance#index}.
	Index *float64 `field:"required" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#dns ComputeInstance#dns}.
	Dns *[]*string `field:"optional" json:"dns" yaml:"dns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#ip_address ComputeInstance#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#ip_network ComputeInstance#ip_network}.
	IpNetwork *string `field:"optional" json:"ipNetwork" yaml:"ipNetwork"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#is_default_gateway ComputeInstance#is_default_gateway}.
	IsDefaultGateway interface{} `field:"optional" json:"isDefaultGateway" yaml:"isDefaultGateway"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#mac_address ComputeInstance#mac_address}.
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#name_servers ComputeInstance#name_servers}.
	NameServers *[]*string `field:"optional" json:"nameServers" yaml:"nameServers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#nat ComputeInstance#nat}.
	Nat *[]*string `field:"optional" json:"nat" yaml:"nat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#search_domains ComputeInstance#search_domains}.
	SearchDomains *[]*string `field:"optional" json:"searchDomains" yaml:"searchDomains"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#sec_lists ComputeInstance#sec_lists}.
	SecLists *[]*string `field:"optional" json:"secLists" yaml:"secLists"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#shared_network ComputeInstance#shared_network}.
	SharedNetwork interface{} `field:"optional" json:"sharedNetwork" yaml:"sharedNetwork"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#vnic ComputeInstance#vnic}.
	Vnic *string `field:"optional" json:"vnic" yaml:"vnic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#vnic_sets ComputeInstance#vnic_sets}.
	VnicSets *[]*string `field:"optional" json:"vnicSets" yaml:"vnicSets"`
}

