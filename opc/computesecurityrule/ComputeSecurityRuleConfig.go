package computesecurityrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeSecurityRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#flow_direction ComputeSecurityRule#flow_direction}.
	FlowDirection *string `field:"required" json:"flowDirection" yaml:"flowDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#name ComputeSecurityRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#acl ComputeSecurityRule#acl}.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#description ComputeSecurityRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#dst_ip_address_prefixes ComputeSecurityRule#dst_ip_address_prefixes}.
	DstIpAddressPrefixes *[]*string `field:"optional" json:"dstIpAddressPrefixes" yaml:"dstIpAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#dst_vnic_set ComputeSecurityRule#dst_vnic_set}.
	DstVnicSet *string `field:"optional" json:"dstVnicSet" yaml:"dstVnicSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#enabled ComputeSecurityRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#id ComputeSecurityRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#security_protocols ComputeSecurityRule#security_protocols}.
	SecurityProtocols *[]*string `field:"optional" json:"securityProtocols" yaml:"securityProtocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#src_ip_address_prefixes ComputeSecurityRule#src_ip_address_prefixes}.
	SrcIpAddressPrefixes *[]*string `field:"optional" json:"srcIpAddressPrefixes" yaml:"srcIpAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#src_vnic_set ComputeSecurityRule#src_vnic_set}.
	SrcVnicSet *string `field:"optional" json:"srcVnicSet" yaml:"srcVnicSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_security_rule#tags ComputeSecurityRule#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

