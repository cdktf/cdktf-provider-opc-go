// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecurityprotocol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeSecurityProtocolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_protocol#name ComputeSecurityProtocol#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_protocol#description ComputeSecurityProtocol#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_protocol#dst_ports ComputeSecurityProtocol#dst_ports}.
	DstPorts *[]*string `field:"optional" json:"dstPorts" yaml:"dstPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_protocol#id ComputeSecurityProtocol#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_protocol#ip_protocol ComputeSecurityProtocol#ip_protocol}.
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_protocol#src_ports ComputeSecurityProtocol#src_ports}.
	SrcPorts *[]*string `field:"optional" json:"srcPorts" yaml:"srcPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_protocol#tags ComputeSecurityProtocol#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

