package lbaaspolicy


type LbaasPolicyLoadBalancingMechanismPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/lbaas_policy#load_balancing_mechanism LbaasPolicy#load_balancing_mechanism}.
	LoadBalancingMechanism *string `field:"required" json:"loadBalancingMechanism" yaml:"loadBalancingMechanism"`
}

