package lbaaspolicy


type LbaasPolicyRateLimitingRequestPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#burst_size LbaasPolicy#burst_size}.
	BurstSize *float64 `field:"required" json:"burstSize" yaml:"burstSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#delay_excessive_requests LbaasPolicy#delay_excessive_requests}.
	DelayExcessiveRequests interface{} `field:"required" json:"delayExcessiveRequests" yaml:"delayExcessiveRequests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#requests_per_second LbaasPolicy#requests_per_second}.
	RequestsPerSecond *float64 `field:"required" json:"requestsPerSecond" yaml:"requestsPerSecond"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#zone LbaasPolicy#zone}.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#http_error_code LbaasPolicy#http_error_code}.
	HttpErrorCode *float64 `field:"optional" json:"httpErrorCode" yaml:"httpErrorCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#logging_level LbaasPolicy#logging_level}.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#rate_limiting_criteria LbaasPolicy#rate_limiting_criteria}.
	RateLimitingCriteria *string `field:"optional" json:"rateLimitingCriteria" yaml:"rateLimitingCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy#zone_memory_size LbaasPolicy#zone_memory_size}.
	ZoneMemorySize *float64 `field:"optional" json:"zoneMemorySize" yaml:"zoneMemorySize"`
}

