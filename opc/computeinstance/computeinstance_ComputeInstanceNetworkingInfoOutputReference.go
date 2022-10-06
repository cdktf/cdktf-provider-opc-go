package computeinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/computeinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceNetworkingInfoOutputReference interface {
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
	Dns() *[]*string
	SetDns(val *[]*string)
	DnsInput() *[]*string
	// Experimental.
	Fqn() *string
	Index() *float64
	SetIndex(val *float64)
	IndexInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	IpNetwork() *string
	SetIpNetwork(val *string)
	IpNetworkInput() *string
	IsDefaultGateway() interface{}
	SetIsDefaultGateway(val interface{})
	IsDefaultGatewayInput() interface{}
	MacAddress() *string
	SetMacAddress(val *string)
	MacAddressInput() *string
	NameServers() *[]*string
	SetNameServers(val *[]*string)
	NameServersInput() *[]*string
	Nat() *[]*string
	SetNat(val *[]*string)
	NatInput() *[]*string
	SearchDomains() *[]*string
	SetSearchDomains(val *[]*string)
	SearchDomainsInput() *[]*string
	SecLists() *[]*string
	SetSecLists(val *[]*string)
	SecListsInput() *[]*string
	SharedNetwork() interface{}
	SetSharedNetwork(val interface{})
	SharedNetworkInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Vnic() *string
	SetVnic(val *string)
	VnicInput() *string
	VnicSets() *[]*string
	SetVnicSets(val *[]*string)
	VnicSetsInput() *[]*string
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
	ResetDns()
	ResetIpAddress()
	ResetIpNetwork()
	ResetIsDefaultGateway()
	ResetMacAddress()
	ResetNameServers()
	ResetNat()
	ResetSearchDomains()
	ResetSecLists()
	ResetSharedNetwork()
	ResetVnic()
	ResetVnicSets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceNetworkingInfoOutputReference
type jsiiProxy_ComputeInstanceNetworkingInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) Dns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) DnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) Index() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) IndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) IpNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) IpNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) IsDefaultGateway() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) IsDefaultGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) MacAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) NameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) NameServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) Nat() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) NatInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) SearchDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) SearchDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) SecLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) SecListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) SharedNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) SharedNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) Vnic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) VnicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) VnicSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vnicSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) VnicSetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vnicSetsInput",
		&returns,
	)
	return returns
}


func NewComputeInstanceNetworkingInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeInstanceNetworkingInfoOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceNetworkingInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceNetworkingInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.computeInstance.ComputeInstanceNetworkingInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeInstanceNetworkingInfoOutputReference_Override(c ComputeInstanceNetworkingInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.computeInstance.ComputeInstanceNetworkingInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetDns(val *[]*string) {
	if err := j.validateSetDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetIndex(val *float64) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetIpNetwork(val *string) {
	if err := j.validateSetIpNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipNetwork",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetIsDefaultGateway(val interface{}) {
	if err := j.validateSetIsDefaultGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDefaultGateway",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetMacAddress(val *string) {
	if err := j.validateSetMacAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetNameServers(val *[]*string) {
	if err := j.validateSetNameServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameServers",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetNat(val *[]*string) {
	if err := j.validateSetNatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nat",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetSearchDomains(val *[]*string) {
	if err := j.validateSetSearchDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchDomains",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetSecLists(val *[]*string) {
	if err := j.validateSetSecListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secLists",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetSharedNetwork(val interface{}) {
	if err := j.validateSetSharedNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedNetwork",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetVnic(val *string) {
	if err := j.validateSetVnicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnic",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference)SetVnicSets(val *[]*string) {
	if err := j.validateSetVnicSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnicSets",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetDns() {
	_jsii_.InvokeVoid(
		c,
		"resetDns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetIpNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetIpNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetIsDefaultGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetIsDefaultGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetMacAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetMacAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetNameServers() {
	_jsii_.InvokeVoid(
		c,
		"resetNameServers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetNat() {
	_jsii_.InvokeVoid(
		c,
		"resetNat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetSearchDomains() {
	_jsii_.InvokeVoid(
		c,
		"resetSearchDomains",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetSecLists() {
	_jsii_.InvokeVoid(
		c,
		"resetSecLists",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetSharedNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetSharedNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetVnic() {
	_jsii_.InvokeVoid(
		c,
		"resetVnic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ResetVnicSets() {
	_jsii_.InvokeVoid(
		c,
		"resetVnicSets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeInstanceNetworkingInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

