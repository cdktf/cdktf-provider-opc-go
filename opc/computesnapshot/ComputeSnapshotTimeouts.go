package computesnapshot


type ComputeSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_snapshot#create ComputeSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_snapshot#delete ComputeSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
