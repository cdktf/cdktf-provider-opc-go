// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computevpnendpointv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opc-go/opc/v6/computevpnendpointv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2 opc_compute_vpn_endpoint_v2}.
type ComputeVpnEndpointV2 interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CustomerVpnGateway() *string
	SetCustomerVpnGateway(val *string)
	CustomerVpnGatewayInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	IkeIdentifier() *string
	SetIkeIdentifier(val *string)
	IkeIdentifierInput() *string
	IpNetwork() *string
	SetIpNetwork(val *string)
	IpNetworkInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalGatewayIpAddress() *string
	LocalGatewayPrivateIpAddress() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PhaseOneSettings() ComputeVpnEndpointV2PhaseOneSettingsOutputReference
	PhaseOneSettingsInput() *ComputeVpnEndpointV2PhaseOneSettings
	PhaseTwoSettings() ComputeVpnEndpointV2PhaseTwoSettingsOutputReference
	PhaseTwoSettingsInput() *ComputeVpnEndpointV2PhaseTwoSettings
	PreSharedKey() *string
	SetPreSharedKey(val *string)
	PreSharedKeyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReachableRoutes() *[]*string
	SetReachableRoutes(val *[]*string)
	ReachableRoutesInput() *[]*string
	RequirePerfectForwardSecrecy() interface{}
	SetRequirePerfectForwardSecrecy(val interface{})
	RequirePerfectForwardSecrecyInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeVpnEndpointV2TimeoutsOutputReference
	TimeoutsInput() interface{}
	TunnelStatus() *string
	Uri() *string
	VnicSets() *[]*string
	SetVnicSets(val *[]*string)
	VnicSetsInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutPhaseOneSettings(value *ComputeVpnEndpointV2PhaseOneSettings)
	PutPhaseTwoSettings(value *ComputeVpnEndpointV2PhaseTwoSettings)
	PutTimeouts(value *ComputeVpnEndpointV2Timeouts)
	ResetDescription()
	ResetEnabled()
	ResetId()
	ResetIkeIdentifier()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhaseOneSettings()
	ResetPhaseTwoSettings()
	ResetRequirePerfectForwardSecrecy()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeVpnEndpointV2
type jsiiProxy_ComputeVpnEndpointV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeVpnEndpointV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) CustomerVpnGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerVpnGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) CustomerVpnGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerVpnGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) IkeIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) IkeIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) IpNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) IpNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) LocalGatewayIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) LocalGatewayPrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayPrivateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) PhaseOneSettings() ComputeVpnEndpointV2PhaseOneSettingsOutputReference {
	var returns ComputeVpnEndpointV2PhaseOneSettingsOutputReference
	_jsii_.Get(
		j,
		"phaseOneSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) PhaseOneSettingsInput() *ComputeVpnEndpointV2PhaseOneSettings {
	var returns *ComputeVpnEndpointV2PhaseOneSettings
	_jsii_.Get(
		j,
		"phaseOneSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) PhaseTwoSettings() ComputeVpnEndpointV2PhaseTwoSettingsOutputReference {
	var returns ComputeVpnEndpointV2PhaseTwoSettingsOutputReference
	_jsii_.Get(
		j,
		"phaseTwoSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) PhaseTwoSettingsInput() *ComputeVpnEndpointV2PhaseTwoSettings {
	var returns *ComputeVpnEndpointV2PhaseTwoSettings
	_jsii_.Get(
		j,
		"phaseTwoSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) PreSharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) PreSharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) ReachableRoutes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reachableRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) ReachableRoutesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reachableRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) RequirePerfectForwardSecrecy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePerfectForwardSecrecy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) RequirePerfectForwardSecrecyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePerfectForwardSecrecyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Timeouts() ComputeVpnEndpointV2TimeoutsOutputReference {
	var returns ComputeVpnEndpointV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) TunnelStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) VnicSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vnicSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeVpnEndpointV2) VnicSetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vnicSetsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2 opc_compute_vpn_endpoint_v2} Resource.
func NewComputeVpnEndpointV2(scope constructs.Construct, id *string, config *ComputeVpnEndpointV2Config) ComputeVpnEndpointV2 {
	_init_.Initialize()

	if err := validateNewComputeVpnEndpointV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeVpnEndpointV2{}

	_jsii_.Create(
		"@cdktf/provider-opc.computeVpnEndpointV2.ComputeVpnEndpointV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_vpn_endpoint_v2 opc_compute_vpn_endpoint_v2} Resource.
func NewComputeVpnEndpointV2_Override(c ComputeVpnEndpointV2, scope constructs.Construct, id *string, config *ComputeVpnEndpointV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.computeVpnEndpointV2.ComputeVpnEndpointV2",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetCustomerVpnGateway(val *string) {
	if err := j.validateSetCustomerVpnGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerVpnGateway",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetIkeIdentifier(val *string) {
	if err := j.validateSetIkeIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeIdentifier",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetIpNetwork(val *string) {
	if err := j.validateSetIpNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipNetwork",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetPreSharedKey(val *string) {
	if err := j.validateSetPreSharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSharedKey",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetReachableRoutes(val *[]*string) {
	if err := j.validateSetReachableRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reachableRoutes",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetRequirePerfectForwardSecrecy(val interface{}) {
	if err := j.validateSetRequirePerfectForwardSecrecyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirePerfectForwardSecrecy",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ComputeVpnEndpointV2)SetVnicSets(val *[]*string) {
	if err := j.validateSetVnicSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnicSets",
		val,
	)
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
func ComputeVpnEndpointV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeVpnEndpointV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.computeVpnEndpointV2.ComputeVpnEndpointV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeVpnEndpointV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeVpnEndpointV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.computeVpnEndpointV2.ComputeVpnEndpointV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeVpnEndpointV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeVpnEndpointV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.computeVpnEndpointV2.ComputeVpnEndpointV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeVpnEndpointV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opc.computeVpnEndpointV2.ComputeVpnEndpointV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeVpnEndpointV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) PutPhaseOneSettings(value *ComputeVpnEndpointV2PhaseOneSettings) {
	if err := c.validatePutPhaseOneSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhaseOneSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) PutPhaseTwoSettings(value *ComputeVpnEndpointV2PhaseTwoSettings) {
	if err := c.validatePutPhaseTwoSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhaseTwoSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) PutTimeouts(value *ComputeVpnEndpointV2Timeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetIkeIdentifier() {
	_jsii_.InvokeVoid(
		c,
		"resetIkeIdentifier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetPhaseOneSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetPhaseOneSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetPhaseTwoSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetPhaseTwoSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetRequirePerfectForwardSecrecy() {
	_jsii_.InvokeVoid(
		c,
		"resetRequirePerfectForwardSecrecy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeVpnEndpointV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeVpnEndpointV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

