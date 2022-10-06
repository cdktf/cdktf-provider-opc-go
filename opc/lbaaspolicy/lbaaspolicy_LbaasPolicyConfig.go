package lbaaspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbaasPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#load_balancer LbaasPolicy#load_balancer}.
	LoadBalancer *string `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#name LbaasPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// application_cookie_stickiness_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#application_cookie_stickiness_policy LbaasPolicy#application_cookie_stickiness_policy}
	ApplicationCookieStickinessPolicy *LbaasPolicyApplicationCookieStickinessPolicy `field:"optional" json:"applicationCookieStickinessPolicy" yaml:"applicationCookieStickinessPolicy"`
	// cloudgate_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#cloudgate_policy LbaasPolicy#cloudgate_policy}
	CloudgatePolicy *LbaasPolicyCloudgatePolicy `field:"optional" json:"cloudgatePolicy" yaml:"cloudgatePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#id LbaasPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// load_balancer_cookie_stickiness_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#load_balancer_cookie_stickiness_policy LbaasPolicy#load_balancer_cookie_stickiness_policy}
	LoadBalancerCookieStickinessPolicy *LbaasPolicyLoadBalancerCookieStickinessPolicy `field:"optional" json:"loadBalancerCookieStickinessPolicy" yaml:"loadBalancerCookieStickinessPolicy"`
	// load_balancing_mechanism_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#load_balancing_mechanism_policy LbaasPolicy#load_balancing_mechanism_policy}
	LoadBalancingMechanismPolicy *LbaasPolicyLoadBalancingMechanismPolicy `field:"optional" json:"loadBalancingMechanismPolicy" yaml:"loadBalancingMechanismPolicy"`
	// rate_limiting_request_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#rate_limiting_request_policy LbaasPolicy#rate_limiting_request_policy}
	RateLimitingRequestPolicy *LbaasPolicyRateLimitingRequestPolicy `field:"optional" json:"rateLimitingRequestPolicy" yaml:"rateLimitingRequestPolicy"`
	// redirect_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#redirect_policy LbaasPolicy#redirect_policy}
	RedirectPolicy *LbaasPolicyRedirectPolicy `field:"optional" json:"redirectPolicy" yaml:"redirectPolicy"`
	// resource_access_control_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#resource_access_control_policy LbaasPolicy#resource_access_control_policy}
	ResourceAccessControlPolicy *LbaasPolicyResourceAccessControlPolicy `field:"optional" json:"resourceAccessControlPolicy" yaml:"resourceAccessControlPolicy"`
	// set_request_header_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#set_request_header_policy LbaasPolicy#set_request_header_policy}
	SetRequestHeaderPolicy *LbaasPolicySetRequestHeaderPolicy `field:"optional" json:"setRequestHeaderPolicy" yaml:"setRequestHeaderPolicy"`
	// ssl_negotiation_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#ssl_negotiation_policy LbaasPolicy#ssl_negotiation_policy}
	SslNegotiationPolicy *LbaasPolicySslNegotiationPolicy `field:"optional" json:"sslNegotiationPolicy" yaml:"sslNegotiationPolicy"`
	// trusted_certificate_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#trusted_certificate_policy LbaasPolicy#trusted_certificate_policy}
	TrustedCertificatePolicy *LbaasPolicyTrustedCertificatePolicy `field:"optional" json:"trustedCertificatePolicy" yaml:"trustedCertificatePolicy"`
}

