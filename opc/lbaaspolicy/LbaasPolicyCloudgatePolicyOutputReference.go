package lbaaspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v4/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/v4/lbaaspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbaasPolicyCloudgatePolicyOutputReference interface {
	cdktf.ComplexObject
	CloudgateApplication() *string
	SetCloudgateApplication(val *string)
	CloudgateApplicationInput() *string
	CloudgatePolicyName() *string
	SetCloudgatePolicyName(val *string)
	CloudgatePolicyNameInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IdentityServiceInstanceGuid() *string
	SetIdentityServiceInstanceGuid(val *string)
	IdentityServiceInstanceGuidInput() *string
	InternalValue() *LbaasPolicyCloudgatePolicy
	SetInternalValue(val *LbaasPolicyCloudgatePolicy)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualHostnameForPolicyAttribution() *string
	SetVirtualHostnameForPolicyAttribution(val *string)
	VirtualHostnameForPolicyAttributionInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCloudgateApplication()
	ResetCloudgatePolicyName()
	ResetIdentityServiceInstanceGuid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbaasPolicyCloudgatePolicyOutputReference
type jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) CloudgateApplication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudgateApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) CloudgateApplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudgateApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) CloudgatePolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudgatePolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) CloudgatePolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudgatePolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) IdentityServiceInstanceGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityServiceInstanceGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) IdentityServiceInstanceGuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityServiceInstanceGuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) InternalValue() *LbaasPolicyCloudgatePolicy {
	var returns *LbaasPolicyCloudgatePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) VirtualHostnameForPolicyAttribution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHostnameForPolicyAttribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) VirtualHostnameForPolicyAttributionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHostnameForPolicyAttributionInput",
		&returns,
	)
	return returns
}


func NewLbaasPolicyCloudgatePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbaasPolicyCloudgatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewLbaasPolicyCloudgatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyCloudgatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbaasPolicyCloudgatePolicyOutputReference_Override(l LbaasPolicyCloudgatePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyCloudgatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetCloudgateApplication(val *string) {
	if err := j.validateSetCloudgateApplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudgateApplication",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetCloudgatePolicyName(val *string) {
	if err := j.validateSetCloudgatePolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudgatePolicyName",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetIdentityServiceInstanceGuid(val *string) {
	if err := j.validateSetIdentityServiceInstanceGuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityServiceInstanceGuid",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetInternalValue(val *LbaasPolicyCloudgatePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference)SetVirtualHostnameForPolicyAttribution(val *string) {
	if err := j.validateSetVirtualHostnameForPolicyAttributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualHostnameForPolicyAttribution",
		val,
	)
}

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) ResetCloudgateApplication() {
	_jsii_.InvokeVoid(
		l,
		"resetCloudgateApplication",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) ResetCloudgatePolicyName() {
	_jsii_.InvokeVoid(
		l,
		"resetCloudgatePolicyName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) ResetIdentityServiceInstanceGuid() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentityServiceInstanceGuid",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicyCloudgatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

