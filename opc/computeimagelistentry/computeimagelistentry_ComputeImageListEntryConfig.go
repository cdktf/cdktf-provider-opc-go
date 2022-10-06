package computeimagelistentry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeImageListEntryConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_image_list_entry#machine_images ComputeImageListEntry#machine_images}.
	MachineImages *[]*string `field:"required" json:"machineImages" yaml:"machineImages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_image_list_entry#name ComputeImageListEntry#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_image_list_entry#version ComputeImageListEntry#version}.
	Version *float64 `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_image_list_entry#attributes ComputeImageListEntry#attributes}.
	Attributes *string `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_image_list_entry#id ComputeImageListEntry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

