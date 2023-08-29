// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeSnapshotConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_snapshot#instance ComputeSnapshot#instance}.
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_snapshot#account ComputeSnapshot#account}.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_snapshot#id ComputeSnapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_snapshot#machine_image ComputeSnapshot#machine_image}.
	MachineImage *string `field:"optional" json:"machineImage" yaml:"machineImage"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_snapshot#timeouts ComputeSnapshot#timeouts}
	Timeouts *ComputeSnapshotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

