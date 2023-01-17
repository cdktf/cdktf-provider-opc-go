package computestoragevolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opc-go/opc/v3/computestoragevolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume opc_compute_storage_volume}.
type ComputeStorageVolume interface {
	cdktf.TerraformResource
	Bootable() interface{}
	SetBootable(val interface{})
	BootableInput() interface{}
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hypervisor() *string
	SetHypervisor(val *string)
	HypervisorInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageList() *string
	SetImageList(val *string)
	ImageListEntry() *float64
	SetImageListEntry(val *float64)
	ImageListEntryInput() *float64
	ImageListInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MachineImage() *string
	SetMachineImage(val *string)
	MachineImageInput() *string
	Managed() interface{}
	SetManaged(val interface{})
	ManagedInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Platform() *string
	SetPlatform(val *string)
	PlatformInput() *string
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
	Readonly() interface{}
	SetReadonly(val interface{})
	ReadonlyInput() interface{}
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	Snapshot() *string
	SetSnapshot(val *string)
	SnapshotAccount() *string
	SetSnapshotAccount(val *string)
	SnapshotAccountInput() *string
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	SnapshotInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	StoragePool() *string
	SetStoragePool(val *string)
	StoragePoolInput() *string
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeStorageVolumeTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	PutTimeouts(value *ComputeStorageVolumeTimeouts)
	ResetBootable()
	ResetDescription()
	ResetHypervisor()
	ResetId()
	ResetImageList()
	ResetImageListEntry()
	ResetMachineImage()
	ResetManaged()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetReadonly()
	ResetSnapshot()
	ResetSnapshotAccount()
	ResetSnapshotId()
	ResetStatus()
	ResetStoragePool()
	ResetStorageType()
	ResetTags()
	ResetTimeouts()
	ResetUri()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeStorageVolume
type jsiiProxy_ComputeStorageVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeStorageVolume) Bootable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) BootableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Hypervisor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypervisor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) HypervisorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypervisorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) ImageList() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) ImageListEntry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageListEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) ImageListEntryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageListEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) ImageListInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) MachineImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) MachineImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Managed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) ManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Readonly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) ReadonlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Snapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) SnapshotAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) SnapshotAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) SnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) StoragePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) StoragePoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Timeouts() ComputeStorageVolumeTimeoutsOutputReference {
	var returns ComputeStorageVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeStorageVolume) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume opc_compute_storage_volume} Resource.
func NewComputeStorageVolume(scope constructs.Construct, id *string, config *ComputeStorageVolumeConfig) ComputeStorageVolume {
	_init_.Initialize()

	if err := validateNewComputeStorageVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeStorageVolume{}

	_jsii_.Create(
		"@cdktf/provider-opc.computeStorageVolume.ComputeStorageVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_volume opc_compute_storage_volume} Resource.
func NewComputeStorageVolume_Override(c ComputeStorageVolume, scope constructs.Construct, id *string, config *ComputeStorageVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.computeStorageVolume.ComputeStorageVolume",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetBootable(val interface{}) {
	if err := j.validateSetBootableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootable",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetHypervisor(val *string) {
	if err := j.validateSetHypervisorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hypervisor",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetImageList(val *string) {
	if err := j.validateSetImageListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageList",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetImageListEntry(val *float64) {
	if err := j.validateSetImageListEntryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageListEntry",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetMachineImage(val *string) {
	if err := j.validateSetMachineImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineImage",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetManaged(val interface{}) {
	if err := j.validateSetManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managed",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetPlatform(val *string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetReadonly(val interface{}) {
	if err := j.validateSetReadonlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetSnapshot(val *string) {
	if err := j.validateSetSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshot",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetSnapshotAccount(val *string) {
	if err := j.validateSetSnapshotAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotAccount",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetStoragePool(val *string) {
	if err := j.validateSetStoragePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePool",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ComputeStorageVolume)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
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
func ComputeStorageVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeStorageVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.computeStorageVolume.ComputeStorageVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeStorageVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeStorageVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.computeStorageVolume.ComputeStorageVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeStorageVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeStorageVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.computeStorageVolume.ComputeStorageVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeStorageVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opc.computeStorageVolume.ComputeStorageVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeStorageVolume) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeStorageVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeStorageVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeStorageVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeStorageVolume) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeStorageVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeStorageVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeStorageVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeStorageVolume) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeStorageVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeStorageVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeStorageVolume) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeStorageVolume) PutTimeouts(value *ComputeStorageVolumeTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetBootable() {
	_jsii_.InvokeVoid(
		c,
		"resetBootable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetHypervisor() {
	_jsii_.InvokeVoid(
		c,
		"resetHypervisor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetImageList() {
	_jsii_.InvokeVoid(
		c,
		"resetImageList",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetImageListEntry() {
	_jsii_.InvokeVoid(
		c,
		"resetImageListEntry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetMachineImage() {
	_jsii_.InvokeVoid(
		c,
		"resetMachineImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetManaged() {
	_jsii_.InvokeVoid(
		c,
		"resetManaged",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetPlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetPlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetReadonly() {
	_jsii_.InvokeVoid(
		c,
		"resetReadonly",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetSnapshot() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetSnapshotAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshotAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetStoragePool() {
	_jsii_.InvokeVoid(
		c,
		"resetStoragePool",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetStorageType() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) ResetUri() {
	_jsii_.InvokeVoid(
		c,
		"resetUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeStorageVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeStorageVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeStorageVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeStorageVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

