package computestoragevolumesnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeStorageVolumeSnapshotConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#volume_name ComputeStorageVolumeSnapshot#volume_name}.
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#collocated ComputeStorageVolumeSnapshot#collocated}.
	Collocated interface{} `field:"optional" json:"collocated" yaml:"collocated"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#description ComputeStorageVolumeSnapshot#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#id ComputeStorageVolumeSnapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#name ComputeStorageVolumeSnapshot#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#parent_volume_bootable ComputeStorageVolumeSnapshot#parent_volume_bootable}.
	ParentVolumeBootable interface{} `field:"optional" json:"parentVolumeBootable" yaml:"parentVolumeBootable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#tags ComputeStorageVolumeSnapshot#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume_snapshot#timeouts ComputeStorageVolumeSnapshot#timeouts}
	Timeouts *ComputeStorageVolumeSnapshotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

