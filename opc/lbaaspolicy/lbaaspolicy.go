package lbaaspolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicy",
		reflect.TypeOf((*LbaasPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationCookieStickinessPolicy", GoGetter: "ApplicationCookieStickinessPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "applicationCookieStickinessPolicyInput", GoGetter: "ApplicationCookieStickinessPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudgatePolicy", GoGetter: "CloudgatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cloudgatePolicyInput", GoGetter: "CloudgatePolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancer", GoGetter: "LoadBalancer"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerCookieStickinessPolicy", GoGetter: "LoadBalancerCookieStickinessPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerCookieStickinessPolicyInput", GoGetter: "LoadBalancerCookieStickinessPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerInput", GoGetter: "LoadBalancerInput"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancingMechanismPolicy", GoGetter: "LoadBalancingMechanismPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancingMechanismPolicyInput", GoGetter: "LoadBalancingMechanismPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putApplicationCookieStickinessPolicy", GoMethod: "PutApplicationCookieStickinessPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudgatePolicy", GoMethod: "PutCloudgatePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putLoadBalancerCookieStickinessPolicy", GoMethod: "PutLoadBalancerCookieStickinessPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putLoadBalancingMechanismPolicy", GoMethod: "PutLoadBalancingMechanismPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRateLimitingRequestPolicy", GoMethod: "PutRateLimitingRequestPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putRedirectPolicy", GoMethod: "PutRedirectPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putResourceAccessControlPolicy", GoMethod: "PutResourceAccessControlPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putSetRequestHeaderPolicy", GoMethod: "PutSetRequestHeaderPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putSslNegotiationPolicy", GoMethod: "PutSslNegotiationPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTrustedCertificatePolicy", GoMethod: "PutTrustedCertificatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "rateLimitingRequestPolicy", GoGetter: "RateLimitingRequestPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "rateLimitingRequestPolicyInput", GoGetter: "RateLimitingRequestPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redirectPolicy", GoGetter: "RedirectPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "redirectPolicyInput", GoGetter: "RedirectPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationCookieStickinessPolicy", GoMethod: "ResetApplicationCookieStickinessPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudgatePolicy", GoMethod: "ResetCloudgatePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoadBalancerCookieStickinessPolicy", GoMethod: "ResetLoadBalancerCookieStickinessPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoadBalancingMechanismPolicy", GoMethod: "ResetLoadBalancingMechanismPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRateLimitingRequestPolicy", GoMethod: "ResetRateLimitingRequestPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectPolicy", GoMethod: "ResetRedirectPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAccessControlPolicy", GoMethod: "ResetResourceAccessControlPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetRequestHeaderPolicy", GoMethod: "ResetSetRequestHeaderPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSslNegotiationPolicy", GoMethod: "ResetSslNegotiationPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedCertificatePolicy", GoMethod: "ResetTrustedCertificatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAccessControlPolicy", GoGetter: "ResourceAccessControlPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAccessControlPolicyInput", GoGetter: "ResourceAccessControlPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "setRequestHeaderPolicy", GoGetter: "SetRequestHeaderPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "setRequestHeaderPolicyInput", GoGetter: "SetRequestHeaderPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "sslNegotiationPolicy", GoGetter: "SslNegotiationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "sslNegotiationPolicyInput", GoGetter: "SslNegotiationPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "trustedCertificatePolicy", GoGetter: "TrustedCertificatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "trustedCertificatePolicyInput", GoGetter: "TrustedCertificatePolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "uri", GoGetter: "Uri"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyApplicationCookieStickinessPolicy",
		reflect.TypeOf((*LbaasPolicyApplicationCookieStickinessPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyApplicationCookieStickinessPolicyOutputReference",
		reflect.TypeOf((*LbaasPolicyApplicationCookieStickinessPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cookieName", GoGetter: "CookieName"},
			_jsii_.MemberProperty{JsiiProperty: "cookieNameInput", GoGetter: "CookieNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicyApplicationCookieStickinessPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyCloudgatePolicy",
		reflect.TypeOf((*LbaasPolicyCloudgatePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyCloudgatePolicyOutputReference",
		reflect.TypeOf((*LbaasPolicyCloudgatePolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudgateApplication", GoGetter: "CloudgateApplication"},
			_jsii_.MemberProperty{JsiiProperty: "cloudgateApplicationInput", GoGetter: "CloudgateApplicationInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudgatePolicyName", GoGetter: "CloudgatePolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "cloudgatePolicyNameInput", GoGetter: "CloudgatePolicyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "identityServiceInstanceGuid", GoGetter: "IdentityServiceInstanceGuid"},
			_jsii_.MemberProperty{JsiiProperty: "identityServiceInstanceGuidInput", GoGetter: "IdentityServiceInstanceGuidInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudgateApplication", GoMethod: "ResetCloudgateApplication"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudgatePolicyName", GoMethod: "ResetCloudgatePolicyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityServiceInstanceGuid", GoMethod: "ResetIdentityServiceInstanceGuid"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualHostnameForPolicyAttribution", GoGetter: "VirtualHostnameForPolicyAttribution"},
			_jsii_.MemberProperty{JsiiProperty: "virtualHostnameForPolicyAttributionInput", GoGetter: "VirtualHostnameForPolicyAttributionInput"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyConfig",
		reflect.TypeOf((*LbaasPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyLoadBalancerCookieStickinessPolicy",
		reflect.TypeOf((*LbaasPolicyLoadBalancerCookieStickinessPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyLoadBalancerCookieStickinessPolicyOutputReference",
		reflect.TypeOf((*LbaasPolicyLoadBalancerCookieStickinessPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cookieExpirationPeriod", GoGetter: "CookieExpirationPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "cookieExpirationPeriodInput", GoGetter: "CookieExpirationPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicyLoadBalancerCookieStickinessPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyLoadBalancingMechanismPolicy",
		reflect.TypeOf((*LbaasPolicyLoadBalancingMechanismPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyLoadBalancingMechanismPolicyOutputReference",
		reflect.TypeOf((*LbaasPolicyLoadBalancingMechanismPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancingMechanism", GoGetter: "LoadBalancingMechanism"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancingMechanismInput", GoGetter: "LoadBalancingMechanismInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicyLoadBalancingMechanismPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyRateLimitingRequestPolicy",
		reflect.TypeOf((*LbaasPolicyRateLimitingRequestPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyRateLimitingRequestPolicyOutputReference",
		reflect.TypeOf((*LbaasPolicyRateLimitingRequestPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "burstSize", GoGetter: "BurstSize"},
			_jsii_.MemberProperty{JsiiProperty: "burstSizeInput", GoGetter: "BurstSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delayExcessiveRequests", GoGetter: "DelayExcessiveRequests"},
			_jsii_.MemberProperty{JsiiProperty: "delayExcessiveRequestsInput", GoGetter: "DelayExcessiveRequestsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpErrorCode", GoGetter: "HttpErrorCode"},
			_jsii_.MemberProperty{JsiiProperty: "httpErrorCodeInput", GoGetter: "HttpErrorCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "loggingLevel", GoGetter: "LoggingLevel"},
			_jsii_.MemberProperty{JsiiProperty: "loggingLevelInput", GoGetter: "LoggingLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "rateLimitingCriteria", GoGetter: "RateLimitingCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "rateLimitingCriteriaInput", GoGetter: "RateLimitingCriteriaInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestsPerSecond", GoGetter: "RequestsPerSecond"},
			_jsii_.MemberProperty{JsiiProperty: "requestsPerSecondInput", GoGetter: "RequestsPerSecondInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpErrorCode", GoMethod: "ResetHttpErrorCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingLevel", GoMethod: "ResetLoggingLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetRateLimitingCriteria", GoMethod: "ResetRateLimitingCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "resetZoneMemorySize", GoMethod: "ResetZoneMemorySize"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
			_jsii_.MemberProperty{JsiiProperty: "zoneInput", GoGetter: "ZoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "zoneMemorySize", GoGetter: "ZoneMemorySize"},
			_jsii_.MemberProperty{JsiiProperty: "zoneMemorySizeInput", GoGetter: "ZoneMemorySizeInput"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyRedirectPolicy",
		reflect.TypeOf((*LbaasPolicyRedirectPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyRedirectPolicyOutputReference",
		reflect.TypeOf((*LbaasPolicyRedirectPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUri", GoGetter: "RedirectUri"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUriInput", GoGetter: "RedirectUriInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseCode", GoGetter: "ResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "responseCodeInput", GoGetter: "ResponseCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicyRedirectPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyResourceAccessControlPolicy",
		reflect.TypeOf((*LbaasPolicyResourceAccessControlPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyResourceAccessControlPolicyOutputReference",
		reflect.TypeOf((*LbaasPolicyResourceAccessControlPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deniedClients", GoGetter: "DeniedClients"},
			_jsii_.MemberProperty{JsiiProperty: "deniedClientsInput", GoGetter: "DeniedClientsInput"},
			_jsii_.MemberProperty{JsiiProperty: "disposition", GoGetter: "Disposition"},
			_jsii_.MemberProperty{JsiiProperty: "dispositionInput", GoGetter: "DispositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "permittedClients", GoGetter: "PermittedClients"},
			_jsii_.MemberProperty{JsiiProperty: "permittedClientsInput", GoGetter: "PermittedClientsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeniedClients", GoMethod: "ResetDeniedClients"},
			_jsii_.MemberMethod{JsiiMethod: "resetPermittedClients", GoMethod: "ResetPermittedClients"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicySetRequestHeaderPolicy",
		reflect.TypeOf((*LbaasPolicySetRequestHeaderPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicySetRequestHeaderPolicyOutputReference",
		reflect.TypeOf((*LbaasPolicySetRequestHeaderPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionWhenHeaderExists", GoGetter: "ActionWhenHeaderExists"},
			_jsii_.MemberProperty{JsiiProperty: "actionWhenHeaderExistsInput", GoGetter: "ActionWhenHeaderExistsInput"},
			_jsii_.MemberProperty{JsiiProperty: "actionWhenHeaderValueIs", GoGetter: "ActionWhenHeaderValueIs"},
			_jsii_.MemberProperty{JsiiProperty: "actionWhenHeaderValueIsInput", GoGetter: "ActionWhenHeaderValueIsInput"},
			_jsii_.MemberProperty{JsiiProperty: "actionWhenHeaderValueIsNot", GoGetter: "ActionWhenHeaderValueIsNot"},
			_jsii_.MemberProperty{JsiiProperty: "actionWhenHeaderValueIsNotInput", GoGetter: "ActionWhenHeaderValueIsNotInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headerName", GoGetter: "HeaderName"},
			_jsii_.MemberProperty{JsiiProperty: "headerNameInput", GoGetter: "HeaderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetActionWhenHeaderExists", GoMethod: "ResetActionWhenHeaderExists"},
			_jsii_.MemberMethod{JsiiMethod: "resetActionWhenHeaderValueIs", GoMethod: "ResetActionWhenHeaderValueIs"},
			_jsii_.MemberMethod{JsiiMethod: "resetActionWhenHeaderValueIsNot", GoMethod: "ResetActionWhenHeaderValueIsNot"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicySslNegotiationPolicy",
		reflect.TypeOf((*LbaasPolicySslNegotiationPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicySslNegotiationPolicyOutputReference",
		reflect.TypeOf((*LbaasPolicySslNegotiationPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerOrderPreference", GoMethod: "ResetServerOrderPreference"},
			_jsii_.MemberMethod{JsiiMethod: "resetSslCiphers", GoMethod: "ResetSslCiphers"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serverOrderPreference", GoGetter: "ServerOrderPreference"},
			_jsii_.MemberProperty{JsiiProperty: "serverOrderPreferenceInput", GoGetter: "ServerOrderPreferenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "sslCiphers", GoGetter: "SslCiphers"},
			_jsii_.MemberProperty{JsiiProperty: "sslCiphersInput", GoGetter: "SslCiphersInput"},
			_jsii_.MemberProperty{JsiiProperty: "sslProtocol", GoGetter: "SslProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "sslProtocolInput", GoGetter: "SslProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyTrustedCertificatePolicy",
		reflect.TypeOf((*LbaasPolicyTrustedCertificatePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyTrustedCertificatePolicyOutputReference",
		reflect.TypeOf((*LbaasPolicyTrustedCertificatePolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustedCertificate", GoGetter: "TrustedCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "trustedCertificateInput", GoGetter: "TrustedCertificateInput"},
		},
		func() interface{} {
			j := jsiiProxy_LbaasPolicyTrustedCertificatePolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
