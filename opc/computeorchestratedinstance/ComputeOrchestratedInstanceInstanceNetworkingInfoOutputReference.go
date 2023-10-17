// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeorchestratedinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v7/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/v7/computeorchestratedinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference interface {
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

// The jsii proxy struct for ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference
type jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) Dns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) DnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) Index() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) IndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) IpNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) IpNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) IsDefaultGateway() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) IsDefaultGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) MacAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) MacAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) NameServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) NameServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nameServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) Nat() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) NatInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) SearchDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) SearchDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) SecLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) SecListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) SharedNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) SharedNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) Vnic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) VnicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) VnicSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vnicSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) VnicSetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vnicSetsInput",
		&returns,
	)
	return returns
}


func NewComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference {
	_init_.Initialize()

	if err := validateNewComputeOrchestratedInstanceInstanceNetworkingInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.computeOrchestratedInstance.ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference_Override(c ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.computeOrchestratedInstance.ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetDns(val *[]*string) {
	if err := j.validateSetDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetIndex(val *float64) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetIpNetwork(val *string) {
	if err := j.validateSetIpNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipNetwork",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetIsDefaultGateway(val interface{}) {
	if err := j.validateSetIsDefaultGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDefaultGateway",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetMacAddress(val *string) {
	if err := j.validateSetMacAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetNameServers(val *[]*string) {
	if err := j.validateSetNameServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameServers",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetNat(val *[]*string) {
	if err := j.validateSetNatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nat",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetSearchDomains(val *[]*string) {
	if err := j.validateSetSearchDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchDomains",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetSecLists(val *[]*string) {
	if err := j.validateSetSecListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secLists",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetSharedNetwork(val interface{}) {
	if err := j.validateSetSharedNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedNetwork",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetVnic(val *string) {
	if err := j.validateSetVnicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnic",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference)SetVnicSets(val *[]*string) {
	if err := j.validateSetVnicSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnicSets",
		val,
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetDns() {
	_jsii_.InvokeVoid(
		c,
		"resetDns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetIpNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetIpNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetIsDefaultGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetIsDefaultGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetMacAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetMacAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetNameServers() {
	_jsii_.InvokeVoid(
		c,
		"resetNameServers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetNat() {
	_jsii_.InvokeVoid(
		c,
		"resetNat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetSearchDomains() {
	_jsii_.InvokeVoid(
		c,
		"resetSearchDomains",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetSecLists() {
	_jsii_.InvokeVoid(
		c,
		"resetSecLists",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetSharedNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetSharedNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetVnic() {
	_jsii_.InvokeVoid(
		c,
		"resetVnic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ResetVnicSets() {
	_jsii_.InvokeVoid(
		c,
		"resetVnicSets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceNetworkingInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

