package lbaasloadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbaasLoadBalancerConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#name LbaasLoadBalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#region LbaasLoadBalancer#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#scheme LbaasLoadBalancer#scheme}.
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#description LbaasLoadBalancer#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#enabled LbaasLoadBalancer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#id LbaasLoadBalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#ip_network LbaasLoadBalancer#ip_network}.
	IpNetwork *string `field:"optional" json:"ipNetwork" yaml:"ipNetwork"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#parent_load_balancer LbaasLoadBalancer#parent_load_balancer}.
	ParentLoadBalancer *string `field:"optional" json:"parentLoadBalancer" yaml:"parentLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#permitted_clients LbaasLoadBalancer#permitted_clients}.
	PermittedClients *[]*string `field:"optional" json:"permittedClients" yaml:"permittedClients"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#permitted_methods LbaasLoadBalancer#permitted_methods}.
	PermittedMethods *[]*string `field:"optional" json:"permittedMethods" yaml:"permittedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#policies LbaasLoadBalancer#policies}.
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#server_pool LbaasLoadBalancer#server_pool}.
	ServerPool *string `field:"optional" json:"serverPool" yaml:"serverPool"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_load_balancer#tags LbaasLoadBalancer#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

