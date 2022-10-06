package computemachineimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeMachineImageConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_machine_image#account ComputeMachineImage#account}.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_machine_image#file ComputeMachineImage#file}.
	File *string `field:"required" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_machine_image#name ComputeMachineImage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_machine_image#attributes ComputeMachineImage#attributes}.
	Attributes *string `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_machine_image#description ComputeMachineImage#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_machine_image#id ComputeMachineImage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

