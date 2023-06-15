package lbaaspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v5/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/v5/lbaaspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbaasPolicySetRequestHeaderPolicyOutputReference interface {
	cdktf.ComplexObject
	ActionWhenHeaderExists() *string
	SetActionWhenHeaderExists(val *string)
	ActionWhenHeaderExistsInput() *string
	ActionWhenHeaderValueIs() *[]*string
	SetActionWhenHeaderValueIs(val *[]*string)
	ActionWhenHeaderValueIsInput() *[]*string
	ActionWhenHeaderValueIsNot() *[]*string
	SetActionWhenHeaderValueIsNot(val *[]*string)
	ActionWhenHeaderValueIsNotInput() *[]*string
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
	HeaderName() *string
	SetHeaderName(val *string)
	HeaderNameInput() *string
	InternalValue() *LbaasPolicySetRequestHeaderPolicy
	SetInternalValue(val *LbaasPolicySetRequestHeaderPolicy)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetActionWhenHeaderExists()
	ResetActionWhenHeaderValueIs()
	ResetActionWhenHeaderValueIsNot()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbaasPolicySetRequestHeaderPolicyOutputReference
type jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ActionWhenHeaderExists() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionWhenHeaderExists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ActionWhenHeaderExistsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionWhenHeaderExistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ActionWhenHeaderValueIs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionWhenHeaderValueIs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ActionWhenHeaderValueIsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionWhenHeaderValueIsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ActionWhenHeaderValueIsNot() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionWhenHeaderValueIsNot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ActionWhenHeaderValueIsNotInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionWhenHeaderValueIsNotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) HeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) HeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) InternalValue() *LbaasPolicySetRequestHeaderPolicy {
	var returns *LbaasPolicySetRequestHeaderPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewLbaasPolicySetRequestHeaderPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbaasPolicySetRequestHeaderPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewLbaasPolicySetRequestHeaderPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicySetRequestHeaderPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbaasPolicySetRequestHeaderPolicyOutputReference_Override(l LbaasPolicySetRequestHeaderPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicySetRequestHeaderPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetActionWhenHeaderExists(val *string) {
	if err := j.validateSetActionWhenHeaderExistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionWhenHeaderExists",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetActionWhenHeaderValueIs(val *[]*string) {
	if err := j.validateSetActionWhenHeaderValueIsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionWhenHeaderValueIs",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetActionWhenHeaderValueIsNot(val *[]*string) {
	if err := j.validateSetActionWhenHeaderValueIsNotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionWhenHeaderValueIsNot",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetHeaderName(val *string) {
	if err := j.validateSetHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerName",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetInternalValue(val *LbaasPolicySetRequestHeaderPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ResetActionWhenHeaderExists() {
	_jsii_.InvokeVoid(
		l,
		"resetActionWhenHeaderExists",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ResetActionWhenHeaderValueIs() {
	_jsii_.InvokeVoid(
		l,
		"resetActionWhenHeaderValueIs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ResetActionWhenHeaderValueIsNot() {
	_jsii_.InvokeVoid(
		l,
		"resetActionWhenHeaderValueIsNot",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		l,
		"resetValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LbaasPolicySetRequestHeaderPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

