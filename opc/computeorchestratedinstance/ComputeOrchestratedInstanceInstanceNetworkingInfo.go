// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeorchestratedinstance


type ComputeOrchestratedInstanceInstanceNetworkingInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#index ComputeOrchestratedInstance#index}.
	Index *float64 `field:"required" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#dns ComputeOrchestratedInstance#dns}.
	Dns *[]*string `field:"optional" json:"dns" yaml:"dns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#ip_address ComputeOrchestratedInstance#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#ip_network ComputeOrchestratedInstance#ip_network}.
	IpNetwork *string `field:"optional" json:"ipNetwork" yaml:"ipNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#is_default_gateway ComputeOrchestratedInstance#is_default_gateway}.
	IsDefaultGateway interface{} `field:"optional" json:"isDefaultGateway" yaml:"isDefaultGateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#mac_address ComputeOrchestratedInstance#mac_address}.
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#name_servers ComputeOrchestratedInstance#name_servers}.
	NameServers *[]*string `field:"optional" json:"nameServers" yaml:"nameServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#nat ComputeOrchestratedInstance#nat}.
	Nat *[]*string `field:"optional" json:"nat" yaml:"nat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#search_domains ComputeOrchestratedInstance#search_domains}.
	SearchDomains *[]*string `field:"optional" json:"searchDomains" yaml:"searchDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#sec_lists ComputeOrchestratedInstance#sec_lists}.
	SecLists *[]*string `field:"optional" json:"secLists" yaml:"secLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#shared_network ComputeOrchestratedInstance#shared_network}.
	SharedNetwork interface{} `field:"optional" json:"sharedNetwork" yaml:"sharedNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#vnic ComputeOrchestratedInstance#vnic}.
	Vnic *string `field:"optional" json:"vnic" yaml:"vnic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#vnic_sets ComputeOrchestratedInstance#vnic_sets}.
	VnicSets *[]*string `field:"optional" json:"vnicSets" yaml:"vnicSets"`
}

