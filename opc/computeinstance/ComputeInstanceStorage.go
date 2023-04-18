package computeinstance


type ComputeInstanceStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_instance#index ComputeInstance#index}.
	Index *float64 `field:"required" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_instance#volume ComputeInstance#volume}.
	Volume *string `field:"required" json:"volume" yaml:"volume"`
}

