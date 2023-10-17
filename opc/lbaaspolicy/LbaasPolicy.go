// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opc-go/opc/v7/lbaaspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy opc_lbaas_policy}.
type LbaasPolicy interface {
	cdktf.TerraformResource
	ApplicationCookieStickinessPolicy() LbaasPolicyApplicationCookieStickinessPolicyOutputReference
	ApplicationCookieStickinessPolicyInput() *LbaasPolicyApplicationCookieStickinessPolicy
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudgatePolicy() LbaasPolicyCloudgatePolicyOutputReference
	CloudgatePolicyInput() *LbaasPolicyCloudgatePolicy
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() *string
	SetLoadBalancer(val *string)
	LoadBalancerCookieStickinessPolicy() LbaasPolicyLoadBalancerCookieStickinessPolicyOutputReference
	LoadBalancerCookieStickinessPolicyInput() *LbaasPolicyLoadBalancerCookieStickinessPolicy
	LoadBalancerInput() *string
	LoadBalancingMechanismPolicy() LbaasPolicyLoadBalancingMechanismPolicyOutputReference
	LoadBalancingMechanismPolicyInput() *LbaasPolicyLoadBalancingMechanismPolicy
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RateLimitingRequestPolicy() LbaasPolicyRateLimitingRequestPolicyOutputReference
	RateLimitingRequestPolicyInput() *LbaasPolicyRateLimitingRequestPolicy
	// Experimental.
	RawOverrides() interface{}
	RedirectPolicy() LbaasPolicyRedirectPolicyOutputReference
	RedirectPolicyInput() *LbaasPolicyRedirectPolicy
	ResourceAccessControlPolicy() LbaasPolicyResourceAccessControlPolicyOutputReference
	ResourceAccessControlPolicyInput() *LbaasPolicyResourceAccessControlPolicy
	SetRequestHeaderPolicy() LbaasPolicySetRequestHeaderPolicyOutputReference
	SetRequestHeaderPolicyInput() *LbaasPolicySetRequestHeaderPolicy
	SslNegotiationPolicy() LbaasPolicySslNegotiationPolicyOutputReference
	SslNegotiationPolicyInput() *LbaasPolicySslNegotiationPolicy
	State() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrustedCertificatePolicy() LbaasPolicyTrustedCertificatePolicyOutputReference
	TrustedCertificatePolicyInput() *LbaasPolicyTrustedCertificatePolicy
	Type() *string
	Uri() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutApplicationCookieStickinessPolicy(value *LbaasPolicyApplicationCookieStickinessPolicy)
	PutCloudgatePolicy(value *LbaasPolicyCloudgatePolicy)
	PutLoadBalancerCookieStickinessPolicy(value *LbaasPolicyLoadBalancerCookieStickinessPolicy)
	PutLoadBalancingMechanismPolicy(value *LbaasPolicyLoadBalancingMechanismPolicy)
	PutRateLimitingRequestPolicy(value *LbaasPolicyRateLimitingRequestPolicy)
	PutRedirectPolicy(value *LbaasPolicyRedirectPolicy)
	PutResourceAccessControlPolicy(value *LbaasPolicyResourceAccessControlPolicy)
	PutSetRequestHeaderPolicy(value *LbaasPolicySetRequestHeaderPolicy)
	PutSslNegotiationPolicy(value *LbaasPolicySslNegotiationPolicy)
	PutTrustedCertificatePolicy(value *LbaasPolicyTrustedCertificatePolicy)
	ResetApplicationCookieStickinessPolicy()
	ResetCloudgatePolicy()
	ResetId()
	ResetLoadBalancerCookieStickinessPolicy()
	ResetLoadBalancingMechanismPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRateLimitingRequestPolicy()
	ResetRedirectPolicy()
	ResetResourceAccessControlPolicy()
	ResetSetRequestHeaderPolicy()
	ResetSslNegotiationPolicy()
	ResetTrustedCertificatePolicy()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LbaasPolicy
type jsiiProxy_LbaasPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LbaasPolicy) ApplicationCookieStickinessPolicy() LbaasPolicyApplicationCookieStickinessPolicyOutputReference {
	var returns LbaasPolicyApplicationCookieStickinessPolicyOutputReference
	_jsii_.Get(
		j,
		"applicationCookieStickinessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) ApplicationCookieStickinessPolicyInput() *LbaasPolicyApplicationCookieStickinessPolicy {
	var returns *LbaasPolicyApplicationCookieStickinessPolicy
	_jsii_.Get(
		j,
		"applicationCookieStickinessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) CloudgatePolicy() LbaasPolicyCloudgatePolicyOutputReference {
	var returns LbaasPolicyCloudgatePolicyOutputReference
	_jsii_.Get(
		j,
		"cloudgatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) CloudgatePolicyInput() *LbaasPolicyCloudgatePolicy {
	var returns *LbaasPolicyCloudgatePolicy
	_jsii_.Get(
		j,
		"cloudgatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) LoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) LoadBalancerCookieStickinessPolicy() LbaasPolicyLoadBalancerCookieStickinessPolicyOutputReference {
	var returns LbaasPolicyLoadBalancerCookieStickinessPolicyOutputReference
	_jsii_.Get(
		j,
		"loadBalancerCookieStickinessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) LoadBalancerCookieStickinessPolicyInput() *LbaasPolicyLoadBalancerCookieStickinessPolicy {
	var returns *LbaasPolicyLoadBalancerCookieStickinessPolicy
	_jsii_.Get(
		j,
		"loadBalancerCookieStickinessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) LoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) LoadBalancingMechanismPolicy() LbaasPolicyLoadBalancingMechanismPolicyOutputReference {
	var returns LbaasPolicyLoadBalancingMechanismPolicyOutputReference
	_jsii_.Get(
		j,
		"loadBalancingMechanismPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) LoadBalancingMechanismPolicyInput() *LbaasPolicyLoadBalancingMechanismPolicy {
	var returns *LbaasPolicyLoadBalancingMechanismPolicy
	_jsii_.Get(
		j,
		"loadBalancingMechanismPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) RateLimitingRequestPolicy() LbaasPolicyRateLimitingRequestPolicyOutputReference {
	var returns LbaasPolicyRateLimitingRequestPolicyOutputReference
	_jsii_.Get(
		j,
		"rateLimitingRequestPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) RateLimitingRequestPolicyInput() *LbaasPolicyRateLimitingRequestPolicy {
	var returns *LbaasPolicyRateLimitingRequestPolicy
	_jsii_.Get(
		j,
		"rateLimitingRequestPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) RedirectPolicy() LbaasPolicyRedirectPolicyOutputReference {
	var returns LbaasPolicyRedirectPolicyOutputReference
	_jsii_.Get(
		j,
		"redirectPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) RedirectPolicyInput() *LbaasPolicyRedirectPolicy {
	var returns *LbaasPolicyRedirectPolicy
	_jsii_.Get(
		j,
		"redirectPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) ResourceAccessControlPolicy() LbaasPolicyResourceAccessControlPolicyOutputReference {
	var returns LbaasPolicyResourceAccessControlPolicyOutputReference
	_jsii_.Get(
		j,
		"resourceAccessControlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) ResourceAccessControlPolicyInput() *LbaasPolicyResourceAccessControlPolicy {
	var returns *LbaasPolicyResourceAccessControlPolicy
	_jsii_.Get(
		j,
		"resourceAccessControlPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) SetRequestHeaderPolicy() LbaasPolicySetRequestHeaderPolicyOutputReference {
	var returns LbaasPolicySetRequestHeaderPolicyOutputReference
	_jsii_.Get(
		j,
		"setRequestHeaderPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) SetRequestHeaderPolicyInput() *LbaasPolicySetRequestHeaderPolicy {
	var returns *LbaasPolicySetRequestHeaderPolicy
	_jsii_.Get(
		j,
		"setRequestHeaderPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) SslNegotiationPolicy() LbaasPolicySslNegotiationPolicyOutputReference {
	var returns LbaasPolicySslNegotiationPolicyOutputReference
	_jsii_.Get(
		j,
		"sslNegotiationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) SslNegotiationPolicyInput() *LbaasPolicySslNegotiationPolicy {
	var returns *LbaasPolicySslNegotiationPolicy
	_jsii_.Get(
		j,
		"sslNegotiationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) State() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) TrustedCertificatePolicy() LbaasPolicyTrustedCertificatePolicyOutputReference {
	var returns LbaasPolicyTrustedCertificatePolicyOutputReference
	_jsii_.Get(
		j,
		"trustedCertificatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) TrustedCertificatePolicyInput() *LbaasPolicyTrustedCertificatePolicy {
	var returns *LbaasPolicyTrustedCertificatePolicy
	_jsii_.Get(
		j,
		"trustedCertificatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicy) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy opc_lbaas_policy} Resource.
func NewLbaasPolicy(scope constructs.Construct, id *string, config *LbaasPolicyConfig) LbaasPolicy {
	_init_.Initialize()

	if err := validateNewLbaasPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbaasPolicy{}

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/lbaas_policy opc_lbaas_policy} Resource.
func NewLbaasPolicy_Override(l LbaasPolicy, scope constructs.Construct, id *string, config *LbaasPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicy",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetLoadBalancer(val *string) {
	if err := j.validateSetLoadBalancerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancer",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a LbaasPolicy resource upon running "cdktf plan <stack-name>".
func LbaasPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLbaasPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicy",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func LbaasPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbaasPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbaasPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbaasPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LbaasPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLbaasPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LbaasPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LbaasPolicy) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LbaasPolicy) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LbaasPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LbaasPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LbaasPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutApplicationCookieStickinessPolicy(value *LbaasPolicyApplicationCookieStickinessPolicy) {
	if err := l.validatePutApplicationCookieStickinessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putApplicationCookieStickinessPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutCloudgatePolicy(value *LbaasPolicyCloudgatePolicy) {
	if err := l.validatePutCloudgatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCloudgatePolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutLoadBalancerCookieStickinessPolicy(value *LbaasPolicyLoadBalancerCookieStickinessPolicy) {
	if err := l.validatePutLoadBalancerCookieStickinessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLoadBalancerCookieStickinessPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutLoadBalancingMechanismPolicy(value *LbaasPolicyLoadBalancingMechanismPolicy) {
	if err := l.validatePutLoadBalancingMechanismPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLoadBalancingMechanismPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutRateLimitingRequestPolicy(value *LbaasPolicyRateLimitingRequestPolicy) {
	if err := l.validatePutRateLimitingRequestPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRateLimitingRequestPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutRedirectPolicy(value *LbaasPolicyRedirectPolicy) {
	if err := l.validatePutRedirectPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRedirectPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutResourceAccessControlPolicy(value *LbaasPolicyResourceAccessControlPolicy) {
	if err := l.validatePutResourceAccessControlPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putResourceAccessControlPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutSetRequestHeaderPolicy(value *LbaasPolicySetRequestHeaderPolicy) {
	if err := l.validatePutSetRequestHeaderPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSetRequestHeaderPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutSslNegotiationPolicy(value *LbaasPolicySslNegotiationPolicy) {
	if err := l.validatePutSslNegotiationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSslNegotiationPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) PutTrustedCertificatePolicy(value *LbaasPolicyTrustedCertificatePolicy) {
	if err := l.validatePutTrustedCertificatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTrustedCertificatePolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetApplicationCookieStickinessPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetApplicationCookieStickinessPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetCloudgatePolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetCloudgatePolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetLoadBalancerCookieStickinessPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadBalancerCookieStickinessPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetLoadBalancingMechanismPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetLoadBalancingMechanismPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetRateLimitingRequestPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetRateLimitingRequestPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetRedirectPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetRedirectPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetResourceAccessControlPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetResourceAccessControlPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetSetRequestHeaderPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetSetRequestHeaderPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetSslNegotiationPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetSslNegotiationPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) ResetTrustedCertificatePolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetTrustedCertificatePolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

