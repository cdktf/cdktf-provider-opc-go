package computeorchestratedinstance


type ComputeOrchestratedInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#create ComputeOrchestratedInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#delete ComputeOrchestratedInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_orchestrated_instance#update ComputeOrchestratedInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

