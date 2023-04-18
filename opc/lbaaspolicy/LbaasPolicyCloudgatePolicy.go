package lbaaspolicy


type LbaasPolicyCloudgatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#virtual_hostname_for_policy_attribution LbaasPolicy#virtual_hostname_for_policy_attribution}.
	VirtualHostnameForPolicyAttribution *string `field:"required" json:"virtualHostnameForPolicyAttribution" yaml:"virtualHostnameForPolicyAttribution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#cloudgate_application LbaasPolicy#cloudgate_application}.
	CloudgateApplication *string `field:"optional" json:"cloudgateApplication" yaml:"cloudgateApplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#cloudgate_policy_name LbaasPolicy#cloudgate_policy_name}.
	CloudgatePolicyName *string `field:"optional" json:"cloudgatePolicyName" yaml:"cloudgatePolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#identity_service_instance_guid LbaasPolicy#identity_service_instance_guid}.
	IdentityServiceInstanceGuid *string `field:"optional" json:"identityServiceInstanceGuid" yaml:"identityServiceInstanceGuid"`
}

