// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbaaspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v8/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/v8/lbaaspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbaasPolicyResourceAccessControlPolicyOutputReference interface {
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
	DeniedClients() *[]*string
	SetDeniedClients(val *[]*string)
	DeniedClientsInput() *[]*string
	Disposition() *string
	SetDisposition(val *string)
	DispositionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *LbaasPolicyResourceAccessControlPolicy
	SetInternalValue(val *LbaasPolicyResourceAccessControlPolicy)
	PermittedClients() *[]*string
	SetPermittedClients(val *[]*string)
	PermittedClientsInput() *[]*string
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
	ResetDeniedClients()
	ResetPermittedClients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbaasPolicyResourceAccessControlPolicyOutputReference
type jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) DeniedClients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deniedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) DeniedClientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deniedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) Disposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) DispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) InternalValue() *LbaasPolicyResourceAccessControlPolicy {
	var returns *LbaasPolicyResourceAccessControlPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) PermittedClients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) PermittedClientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLbaasPolicyResourceAccessControlPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbaasPolicyResourceAccessControlPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewLbaasPolicyResourceAccessControlPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyResourceAccessControlPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbaasPolicyResourceAccessControlPolicyOutputReference_Override(l LbaasPolicyResourceAccessControlPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyResourceAccessControlPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference)SetDeniedClients(val *[]*string) {
	if err := j.validateSetDeniedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deniedClients",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference)SetDisposition(val *string) {
	if err := j.validateSetDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disposition",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference)SetInternalValue(val *LbaasPolicyResourceAccessControlPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference)SetPermittedClients(val *[]*string) {
	if err := j.validateSetPermittedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedClients",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) ResetDeniedClients() {
	_jsii_.InvokeVoid(
		l,
		"resetDeniedClients",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) ResetPermittedClients() {
	_jsii_.InvokeVoid(
		l,
		"resetPermittedClients",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LbaasPolicyResourceAccessControlPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

