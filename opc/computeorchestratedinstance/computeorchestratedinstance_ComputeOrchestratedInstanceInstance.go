package computeorchestratedinstance


type ComputeOrchestratedInstanceInstance struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#name ComputeOrchestratedInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#shape ComputeOrchestratedInstance#shape}.
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#boot_order ComputeOrchestratedInstance#boot_order}.
	BootOrder *[]*float64 `field:"optional" json:"bootOrder" yaml:"bootOrder"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#hostname ComputeOrchestratedInstance#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#image_list ComputeOrchestratedInstance#image_list}.
	ImageList *string `field:"optional" json:"imageList" yaml:"imageList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#instance_attributes ComputeOrchestratedInstance#instance_attributes}.
	InstanceAttributes *string `field:"optional" json:"instanceAttributes" yaml:"instanceAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#label ComputeOrchestratedInstance#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// networking_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#networking_info ComputeOrchestratedInstance#networking_info}
	NetworkingInfo interface{} `field:"optional" json:"networkingInfo" yaml:"networkingInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#persistent ComputeOrchestratedInstance#persistent}.
	Persistent interface{} `field:"optional" json:"persistent" yaml:"persistent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#reverse_dns ComputeOrchestratedInstance#reverse_dns}.
	ReverseDns interface{} `field:"optional" json:"reverseDns" yaml:"reverseDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#ssh_keys ComputeOrchestratedInstance#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#storage ComputeOrchestratedInstance#storage}
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#tags ComputeOrchestratedInstance#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

