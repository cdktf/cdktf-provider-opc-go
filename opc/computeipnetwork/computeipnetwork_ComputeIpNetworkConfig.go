package computeipnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeIpNetworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_ip_network#ip_address_prefix ComputeIpNetwork#ip_address_prefix}.
	IpAddressPrefix *string `field:"required" json:"ipAddressPrefix" yaml:"ipAddressPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_ip_network#name ComputeIpNetwork#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_ip_network#description ComputeIpNetwork#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_ip_network#id ComputeIpNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_ip_network#ip_network_exchange ComputeIpNetwork#ip_network_exchange}.
	IpNetworkExchange *string `field:"optional" json:"ipNetworkExchange" yaml:"ipNetworkExchange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_ip_network#public_napt_enabled ComputeIpNetwork#public_napt_enabled}.
	PublicNaptEnabled interface{} `field:"optional" json:"publicNaptEnabled" yaml:"publicNaptEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_ip_network#tags ComputeIpNetwork#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

