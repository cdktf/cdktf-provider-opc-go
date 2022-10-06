package computeinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#name ComputeInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#shape ComputeInstance#shape}.
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#boot_order ComputeInstance#boot_order}.
	BootOrder *[]*float64 `field:"optional" json:"bootOrder" yaml:"bootOrder"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#desired_state ComputeInstance#desired_state}.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#hostname ComputeInstance#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#id ComputeInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#image_list ComputeInstance#image_list}.
	ImageList *string `field:"optional" json:"imageList" yaml:"imageList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#instance_attributes ComputeInstance#instance_attributes}.
	InstanceAttributes *string `field:"optional" json:"instanceAttributes" yaml:"instanceAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#label ComputeInstance#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// networking_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#networking_info ComputeInstance#networking_info}
	NetworkingInfo interface{} `field:"optional" json:"networkingInfo" yaml:"networkingInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#reverse_dns ComputeInstance#reverse_dns}.
	ReverseDns interface{} `field:"optional" json:"reverseDns" yaml:"reverseDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#ssh_keys ComputeInstance#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#storage ComputeInstance#storage}
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#tags ComputeInstance#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_instance#timeouts ComputeInstance#timeouts}
	Timeouts *ComputeInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

