package computevpnendpointv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v2/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/v2/computevpnendpointv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeVpnEndpointV2PhaseOneSettingsOutputReference interface {
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
	DhGroup() *string
	SetDhGroup(val *string)
	DhGroupInput() *string
	Encryption() *string
	SetEncryption(val *string)
	EncryptionInput() *string
	// Experimental.
	Fqn() *string
	Hash() *string
	SetHash(val *string)
	HashInput() *string
	InternalValue() *ComputeVpnEndpointV2PhaseOneSettings
	SetInternalValue(val *ComputeVpnEndpointV2PhaseOneSettings)
	Lifetime() *float64
	SetLifetime(val *float64)
	LifetimeInput() *float64
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
	ResetLifetime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeVpnEndpointV2PhaseOneSettingsOutputReference
type jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) DhGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) DhGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) Encryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) EncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) Hash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) HashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) InternalValue() *ComputeVpnEndpointV2PhaseOneSettings {
	var returns *ComputeVpnEndpointV2PhaseOneSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) Lifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) LifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeVpnEndpointV2PhaseOneSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeVpnEndpointV2PhaseOneSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewComputeVpnEndpointV2PhaseOneSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.computeVpnEndpointV2.ComputeVpnEndpointV2PhaseOneSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeVpnEndpointV2PhaseOneSettingsOutputReference_Override(c ComputeVpnEndpointV2PhaseOneSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.computeVpnEndpointV2.ComputeVpnEndpointV2PhaseOneSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetDhGroup(val *string) {
	if err := j.validateSetDhGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhGroup",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetEncryption(val *string) {
	if err := j.validateSetEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryption",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetHash(val *string) {
	if err := j.validateSetHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hash",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetInternalValue(val *ComputeVpnEndpointV2PhaseOneSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetLifetime(val *float64) {
	if err := j.validateSetLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetime",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) ResetLifetime() {
	_jsii_.InvokeVoid(
		c,
		"resetLifetime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2PhaseOneSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

