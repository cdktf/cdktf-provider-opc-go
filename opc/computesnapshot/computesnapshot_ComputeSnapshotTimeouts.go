package computesnapshot


type ComputeSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_snapshot#create ComputeSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_snapshot#delete ComputeSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

