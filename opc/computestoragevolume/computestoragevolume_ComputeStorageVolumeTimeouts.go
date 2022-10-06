package computestoragevolume


type ComputeStorageVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume#create ComputeStorageVolume#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume#delete ComputeStorageVolume#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume#update ComputeStorageVolume#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

