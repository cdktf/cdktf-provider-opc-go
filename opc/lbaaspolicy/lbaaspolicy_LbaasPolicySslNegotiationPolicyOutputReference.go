package lbaaspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/lbaaspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbaasPolicySslNegotiationPolicyOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *LbaasPolicySslNegotiationPolicy
	SetInternalValue(val *LbaasPolicySslNegotiationPolicy)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ServerOrderPreference() *string
	SetServerOrderPreference(val *string)
	ServerOrderPreferenceInput() *string
	SslCiphers() *[]*string
	SetSslCiphers(val *[]*string)
	SslCiphersInput() *[]*string
	SslProtocol() *[]*string
	SetSslProtocol(val *[]*string)
	SslProtocolInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetServerOrderPreference()
	ResetSslCiphers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbaasPolicySslNegotiationPolicyOutputReference
type jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) InternalValue() *LbaasPolicySslNegotiationPolicy {
	var returns *LbaasPolicySslNegotiationPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) ServerOrderPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverOrderPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) ServerOrderPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverOrderPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) SslCiphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sslCiphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) SslCiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sslCiphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) SslProtocol() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sslProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) SslProtocolInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sslProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLbaasPolicySslNegotiationPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbaasPolicySslNegotiationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewLbaasPolicySslNegotiationPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicySslNegotiationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbaasPolicySslNegotiationPolicyOutputReference_Override(l LbaasPolicySslNegotiationPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicySslNegotiationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetInternalValue(val *LbaasPolicySslNegotiationPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetServerOrderPreference(val *string) {
	if err := j.validateSetServerOrderPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverOrderPreference",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetSslCiphers(val *[]*string) {
	if err := j.validateSetSslCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCiphers",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetSslProtocol(val *[]*string) {
	if err := j.validateSetSslProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslProtocol",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) ResetServerOrderPreference() {
	_jsii_.InvokeVoid(
		l,
		"resetServerOrderPreference",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) ResetSslCiphers() {
	_jsii_.InvokeVoid(
		l,
		"resetSslCiphers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LbaasPolicySslNegotiationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

