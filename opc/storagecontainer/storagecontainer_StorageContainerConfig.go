package storagecontainer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageContainerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#name StorageContainer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#allowed_origins StorageContainer#allowed_origins}.
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#exposed_headers StorageContainer#exposed_headers}.
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#id StorageContainer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#max_age StorageContainer#max_age}.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#metadata StorageContainer#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#primary_key StorageContainer#primary_key}.
	PrimaryKey *string `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#quota_bytes StorageContainer#quota_bytes}.
	QuotaBytes *float64 `field:"optional" json:"quotaBytes" yaml:"quotaBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#quota_count StorageContainer#quota_count}.
	QuotaCount *float64 `field:"optional" json:"quotaCount" yaml:"quotaCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#read_acls StorageContainer#read_acls}.
	ReadAcls *[]*string `field:"optional" json:"readAcls" yaml:"readAcls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#secondary_key StorageContainer#secondary_key}.
	SecondaryKey *string `field:"optional" json:"secondaryKey" yaml:"secondaryKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/storage_container#write_acls StorageContainer#write_acls}.
	WriteAcls *[]*string `field:"optional" json:"writeAcls" yaml:"writeAcls"`
}

