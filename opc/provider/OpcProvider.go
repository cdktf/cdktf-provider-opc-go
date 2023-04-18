package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-opc-go/opc/v4/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs opc}.
type OpcProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IdentityDomain() *string
	SetIdentityDomain(val *string)
	IdentityDomainInput() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	LbaasEndpoint() *string
	SetLbaasEndpoint(val *string)
	LbaasEndpointInput() *string
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	RawOverrides() interface{}
	StorageEndpoint() *string
	SetStorageEndpoint(val *string)
	StorageEndpointInput() *string
	StorageServiceId() *string
	SetStorageServiceId(val *string)
	StorageServiceIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	User() *string
	SetUser(val *string)
	UserInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetEndpoint()
	ResetInsecure()
	ResetLbaasEndpoint()
	ResetMaxRetries()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStorageEndpoint()
	ResetStorageServiceId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OpcProvider
type jsiiProxy_OpcProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_OpcProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) IdentityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) IdentityDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) LbaasEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbaasEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) LbaasEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbaasEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) StorageEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) StorageEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) StorageServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) StorageServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpcProvider) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs opc} Resource.
func NewOpcProvider(scope constructs.Construct, id *string, config *OpcProviderConfig) OpcProvider {
	_init_.Initialize()

	if err := validateNewOpcProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpcProvider{}

	_jsii_.Create(
		"@cdktf/provider-opc.provider.OpcProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs opc} Resource.
func NewOpcProvider_Override(o OpcProvider, scope constructs.Construct, id *string, config *OpcProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.provider.OpcProvider",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpcProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetIdentityDomain(val *string) {
	_jsii_.Set(
		j,
		"identityDomain",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetLbaasEndpoint(val *string) {
	_jsii_.Set(
		j,
		"lbaasEndpoint",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetStorageEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageEndpoint",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetStorageServiceId(val *string) {
	_jsii_.Set(
		j,
		"storageServiceId",
		val,
	)
}

func (j *jsiiProxy_OpcProvider)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
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
func OpcProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpcProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.provider.OpcProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpcProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpcProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.provider.OpcProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpcProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpcProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opc.provider.OpcProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpcProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opc.provider.OpcProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpcProvider) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpcProvider) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpcProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		o,
		"resetAlias",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpcProvider) ResetEndpoint() {
	_jsii_.InvokeVoid(
		o,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpcProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		o,
		"resetInsecure",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpcProvider) ResetLbaasEndpoint() {
	_jsii_.InvokeVoid(
		o,
		"resetLbaasEndpoint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpcProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpcProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpcProvider) ResetStorageEndpoint() {
	_jsii_.InvokeVoid(
		o,
		"resetStorageEndpoint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpcProvider) ResetStorageServiceId() {
	_jsii_.InvokeVoid(
		o,
		"resetStorageServiceId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpcProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpcProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpcProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpcProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

