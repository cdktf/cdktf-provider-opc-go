package lbaaspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-opc-go/opc/v4/jsii"

	"github.com/cdktf/cdktf-provider-opc-go/opc/v4/lbaaspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbaasPolicyRateLimitingRequestPolicyOutputReference interface {
	cdktf.ComplexObject
	BurstSize() *float64
	SetBurstSize(val *float64)
	BurstSizeInput() *float64
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
	DelayExcessiveRequests() interface{}
	SetDelayExcessiveRequests(val interface{})
	DelayExcessiveRequestsInput() interface{}
	// Experimental.
	Fqn() *string
	HttpErrorCode() *float64
	SetHttpErrorCode(val *float64)
	HttpErrorCodeInput() *float64
	InternalValue() *LbaasPolicyRateLimitingRequestPolicy
	SetInternalValue(val *LbaasPolicyRateLimitingRequestPolicy)
	LoggingLevel() *string
	SetLoggingLevel(val *string)
	LoggingLevelInput() *string
	RateLimitingCriteria() *string
	SetRateLimitingCriteria(val *string)
	RateLimitingCriteriaInput() *string
	RequestsPerSecond() *float64
	SetRequestsPerSecond(val *float64)
	RequestsPerSecondInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	ZoneMemorySize() *float64
	SetZoneMemorySize(val *float64)
	ZoneMemorySizeInput() *float64
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
	ResetHttpErrorCode()
	ResetLoggingLevel()
	ResetRateLimitingCriteria()
	ResetZoneMemorySize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbaasPolicyRateLimitingRequestPolicyOutputReference
type jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) BurstSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"burstSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) BurstSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"burstSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) DelayExcessiveRequests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delayExcessiveRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) DelayExcessiveRequestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delayExcessiveRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) HttpErrorCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpErrorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) HttpErrorCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpErrorCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) InternalValue() *LbaasPolicyRateLimitingRequestPolicy {
	var returns *LbaasPolicyRateLimitingRequestPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) RateLimitingCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitingCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) RateLimitingCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateLimitingCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) RequestsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) RequestsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ZoneMemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"zoneMemorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ZoneMemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"zoneMemorySizeInput",
		&returns,
	)
	return returns
}


func NewLbaasPolicyRateLimitingRequestPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbaasPolicyRateLimitingRequestPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewLbaasPolicyRateLimitingRequestPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyRateLimitingRequestPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbaasPolicyRateLimitingRequestPolicyOutputReference_Override(l LbaasPolicyRateLimitingRequestPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opc.lbaasPolicy.LbaasPolicyRateLimitingRequestPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetBurstSize(val *float64) {
	if err := j.validateSetBurstSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burstSize",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetDelayExcessiveRequests(val interface{}) {
	if err := j.validateSetDelayExcessiveRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delayExcessiveRequests",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetHttpErrorCode(val *float64) {
	if err := j.validateSetHttpErrorCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpErrorCode",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetInternalValue(val *LbaasPolicyRateLimitingRequestPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetLoggingLevel(val *string) {
	if err := j.validateSetLoggingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetRateLimitingCriteria(val *string) {
	if err := j.validateSetRateLimitingCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateLimitingCriteria",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetRequestsPerSecond(val *float64) {
	if err := j.validateSetRequestsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestsPerSecond",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (j *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference)SetZoneMemorySize(val *float64) {
	if err := j.validateSetZoneMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneMemorySize",
		val,
	)
}

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ResetHttpErrorCode() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpErrorCode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		l,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ResetRateLimitingCriteria() {
	_jsii_.InvokeVoid(
		l,
		"resetRateLimitingCriteria",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ResetZoneMemorySize() {
	_jsii_.InvokeVoid(
		l,
		"resetZoneMemorySize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LbaasPolicyRateLimitingRequestPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

