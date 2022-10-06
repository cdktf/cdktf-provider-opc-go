package computestorageattachment


type ComputeStorageAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_attachment#create ComputeStorageAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opc/r/compute_storage_attachment#delete ComputeStorageAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

