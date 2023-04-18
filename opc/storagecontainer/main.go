package storagecontainer

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opc.storageContainer.StorageContainer",
		reflect.TypeOf((*StorageContainer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOrigins", GoGetter: "AllowedOrigins"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOriginsInput", GoGetter: "AllowedOriginsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "exposedHeaders", GoGetter: "ExposedHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "exposedHeadersInput", GoGetter: "ExposedHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxAge", GoGetter: "MaxAge"},
			_jsii_.MemberProperty{JsiiProperty: "maxAgeInput", GoGetter: "MaxAgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyInput", GoGetter: "PrimaryKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "quotaBytes", GoGetter: "QuotaBytes"},
			_jsii_.MemberProperty{JsiiProperty: "quotaBytesInput", GoGetter: "QuotaBytesInput"},
			_jsii_.MemberProperty{JsiiProperty: "quotaCount", GoGetter: "QuotaCount"},
			_jsii_.MemberProperty{JsiiProperty: "quotaCountInput", GoGetter: "QuotaCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readAcls", GoGetter: "ReadAcls"},
			_jsii_.MemberProperty{JsiiProperty: "readAclsInput", GoGetter: "ReadAclsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOrigins", GoMethod: "ResetAllowedOrigins"},
			_jsii_.MemberMethod{JsiiMethod: "resetExposedHeaders", GoMethod: "ResetExposedHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxAge", GoMethod: "ResetMaxAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKey", GoMethod: "ResetPrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetQuotaBytes", GoMethod: "ResetQuotaBytes"},
			_jsii_.MemberMethod{JsiiMethod: "resetQuotaCount", GoMethod: "ResetQuotaCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadAcls", GoMethod: "ResetReadAcls"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecondaryKey", GoMethod: "ResetSecondaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteAcls", GoMethod: "ResetWriteAcls"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryKey", GoGetter: "SecondaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryKeyInput", GoGetter: "SecondaryKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "writeAcls", GoGetter: "WriteAcls"},
			_jsii_.MemberProperty{JsiiProperty: "writeAclsInput", GoGetter: "WriteAclsInput"},
		},
		func() interface{} {
			j := jsiiProxy_StorageContainer{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opc.storageContainer.StorageContainerConfig",
		reflect.TypeOf((*StorageContainerConfig)(nil)).Elem(),
	)
}
