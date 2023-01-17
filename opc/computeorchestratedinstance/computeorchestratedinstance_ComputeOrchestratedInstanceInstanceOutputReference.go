package computeorchestratedinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v3/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/v3/computeorchestratedinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeOrchestratedInstanceInstanceOutputReference interface {
	cdktf.ComplexObject
	Attributes() *string
	AvailabilityDomain() *string
	BootOrder() *[]*float64
	SetBootOrder(val *[]*float64)
	BootOrderInput() *[]*float64
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
	Domain() *string
	Entry() *float64
	Fingerprint() *string
	Fqdn() *string
	// Experimental.
	Fqn() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	ImageFormat() *string
	ImageList() *string
	SetImageList(val *string)
	ImageListInput() *string
	InstanceAttributes() *string
	SetInstanceAttributes(val *string)
	InstanceAttributesInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkingInfo() ComputeOrchestratedInstanceInstanceNetworkingInfoList
	NetworkingInfoInput() interface{}
	Persistent() interface{}
	SetPersistent(val interface{})
	PersistentInput() interface{}
	PlacementRequirements() *[]*string
	Platform() *string
	Priority() *string
	QuotaReservation() *string
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
	Storage() ComputeOrchestratedInstanceInstanceStorageList
	StorageInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Vcable() *string
	Virtio() cdktf.IResolvable
	VncAddress() *string
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
	PutNetworkingInfo(value interface{})
	PutStorage(value interface{})
	ResetBootOrder()
	ResetHostname()
	ResetImageList()
	ResetInstanceAttributes()
	ResetLabel()
	ResetNetworkingInfo()
	ResetPersistent()
	ResetReverseDns()
	ResetSshKeys()
	ResetStorage()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeOrchestratedInstanceInstanceOutputReference
type jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Attributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) AvailabilityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) BootOrder() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"bootOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) BootOrderInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"bootOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Entry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ImageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ImageList() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ImageListInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) InstanceAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) InstanceAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) NetworkingInfo() ComputeOrchestratedInstanceInstanceNetworkingInfoList {
	var returns ComputeOrchestratedInstanceInstanceNetworkingInfoList
	_jsii_.Get(
		j,
		"networkingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) NetworkingInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkingInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Persistent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) PersistentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) PlacementRequirements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"placementRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) QuotaReservation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Relationships() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relationships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Resolvers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resolvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ReverseDns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reverseDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ReverseDnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reverseDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Site() *string {
	var returns *string
	_jsii_.Get(
		j,
		"site",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Storage() ComputeOrchestratedInstanceInstanceStorageList {
	var returns ComputeOrchestratedInstanceInstanceStorageList
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) StorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Vcable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Virtio() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"virtio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) VncAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vncAddress",
		&returns,
	)
	return returns
}


func NewComputeOrchestratedInstanceInstanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeOrchestratedInstanceInstanceOutputReference {
	_init_.Initialize()

	if err := validateNewComputeOrchestratedInstanceInstanceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.computeOrchestratedInstance.ComputeOrchestratedInstanceInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeOrchestratedInstanceInstanceOutputReference_Override(c ComputeOrchestratedInstanceInstanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.computeOrchestratedInstance.ComputeOrchestratedInstanceInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetBootOrder(val *[]*float64) {
	if err := j.validateSetBootOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootOrder",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetImageList(val *string) {
	if err := j.validateSetImageListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageList",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetInstanceAttributes(val *string) {
	if err := j.validateSetInstanceAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceAttributes",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetPersistent(val interface{}) {
	if err := j.validateSetPersistentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistent",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetReverseDns(val interface{}) {
	if err := j.validateSetReverseDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reverseDns",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetSshKeys(val *[]*string) {
	if err := j.validateSetSshKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) PutNetworkingInfo(value interface{}) {
	if err := c.validatePutNetworkingInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkingInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) PutStorage(value interface{}) {
	if err := c.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetBootOrder() {
	_jsii_.InvokeVoid(
		c,
		"resetBootOrder",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetImageList() {
	_jsii_.InvokeVoid(
		c,
		"resetImageList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetInstanceAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetNetworkingInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkingInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetPersistent() {
	_jsii_.InvokeVoid(
		c,
		"resetPersistent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetReverseDns() {
	_jsii_.InvokeVoid(
		c,
		"resetReverseDns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

