package computeinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opc-go/opc/computeinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opc/r/compute_instance opc_compute_instance}.
type ComputeInstance interface {
	cdktf.TerraformResource
	Attributes() *string
	AvailabilityDomain() *string
	BootOrder() *[]*float64
	SetBootOrder(val *[]*float64)
	BootOrderInput() *[]*float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	Domain() *string
	Entry() *float64
	Fingerprint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageFormat() *string
	ImageList() *string
	SetImageList(val *string)
	ImageListInput() *string
	InstanceAttributes() *string
	SetInstanceAttributes(val *string)
	InstanceAttributesInput() *string
	IpAddress() *string
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkingInfo() ComputeInstanceNetworkingInfoList
	NetworkingInfoInput() interface{}
	// The tree node.
	Node() constructs.Node
	PlacementRequirements() *[]*string
	Platform() *string
	Priority() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QuotaReservation() *string
	// Experimental.
	RawOverrides() interface{}
	Relationships() *[]*string
	Resolvers() *[]*string
	ReverseDns() interface{}
	SetReverseDns(val interface{})
	ReverseDnsInput() interface{}
	Shape() *string
	SetShape(val *string)
	ShapeInput() *string
	Site() *string
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	StartTime() *string
	State() *string
	Storage() ComputeInstanceStorageList
	StorageInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Vcable() *string
	Virtio() cdktf.IResolvable
	VncAddress() *string
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
	PutNetworkingInfo(value interface{})
	PutStorage(value interface{})
	PutTimeouts(value *ComputeInstanceTimeouts)
	ResetBootOrder()
	ResetDesiredState()
	ResetHostname()
	ResetId()
	ResetImageList()
	ResetInstanceAttributes()
	ResetLabel()
	ResetNetworkingInfo()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReverseDns()
	ResetSshKeys()
	ResetStorage()
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

// The jsii proxy struct for ComputeInstance
type jsiiProxy_ComputeInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeInstance) Attributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) AvailabilityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) BootOrder() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"bootOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) BootOrderInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"bootOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Entry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ImageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ImageList() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ImageListInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) InstanceAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) InstanceAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) NetworkingInfo() ComputeInstanceNetworkingInfoList {
	var returns ComputeInstanceNetworkingInfoList
	_jsii_.Get(
		j,
		"networkingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) NetworkingInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkingInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) PlacementRequirements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"placementRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) QuotaReservation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Relationships() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relationships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Resolvers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resolvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ReverseDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reverseDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ReverseDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reverseDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Site() *string {
	var returns *string
	_jsii_.Get(
		j,
		"site",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Storage() ComputeInstanceStorageList {
	var returns ComputeInstanceStorageList
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) StorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Timeouts() ComputeInstanceTimeoutsOutputReference {
	var returns ComputeInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Vcable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) Virtio() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"virtio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstance) VncAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vncAddress",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opc/r/compute_instance opc_compute_instance} Resource.
func NewComputeInstance(scope constructs.Construct, id *string, config *ComputeInstanceConfig) ComputeInstance {
	_init_.Initialize()

	if err := validateNewComputeInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstance{}

	_jsii_.Create(
		"@cdktf/provider-opc.computeInstance.ComputeInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opc/r/compute_instance opc_compute_instance} Resource.
func NewComputeInstance_Override(c ComputeInstance, scope constructs.Construct, id *string, config *ComputeInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.computeInstance.ComputeInstance",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeInstance)SetBootOrder(val *[]*float64) {
	if err := j.validateSetBootOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootOrder",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetImageList(val *string) {
	if err := j.validateSetImageListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageList",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetInstanceAttributes(val *string) {
	if err := j.validateSetInstanceAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceAttributes",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetReverseDns(val interface{}) {
	if err := j.validateSetReverseDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reverseDns",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetSshKeys(val *[]*string) {
	if err := j.validateSetSshKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_ComputeInstance)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func ComputeInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.computeInstance.ComputeInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opc.computeInstance.ComputeInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeInstance) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstance) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeInstance) PutNetworkingInfo(value interface{}) {
	if err := c.validatePutNetworkingInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkingInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutStorage(value interface{}) {
	if err := c.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) PutTimeouts(value *ComputeInstanceTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstance) ResetBootOrder() {
	_jsii_.InvokeVoid(
		c,
		"resetBootOrder",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetDesiredState() {
	_jsii_.InvokeVoid(
		c,
		"resetDesiredState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetImageList() {
	_jsii_.InvokeVoid(
		c,
		"resetImageList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetInstanceAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetNetworkingInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkingInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetReverseDns() {
	_jsii_.InvokeVoid(
		c,
		"resetReverseDns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetSshKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

