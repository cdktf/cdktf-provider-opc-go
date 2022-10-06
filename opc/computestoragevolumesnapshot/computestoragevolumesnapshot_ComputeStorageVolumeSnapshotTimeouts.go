package computestoragevolumesnapshot


type ComputeStorageVolumeSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#create ComputeStorageVolumeSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#delete ComputeStorageVolumeSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

