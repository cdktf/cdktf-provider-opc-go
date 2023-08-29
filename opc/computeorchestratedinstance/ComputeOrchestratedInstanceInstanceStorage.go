// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeorchestratedinstance


type ComputeOrchestratedInstanceInstanceStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#index ComputeOrchestratedInstance#index}.
	Index *float64 `field:"required" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/opc/1.4.1/docs/resources/compute_orchestrated_instance#volume ComputeOrchestratedInstance#volume}.
	Volume *string `field:"required" json:"volume" yaml:"volume"`
}

