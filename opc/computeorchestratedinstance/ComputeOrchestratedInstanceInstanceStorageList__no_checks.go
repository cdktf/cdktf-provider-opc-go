// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package computeorchestratedinstance

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceStorageList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeOrchestratedInstanceInstanceStorageList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceStorageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceStorageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceStorageList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeOrchestratedInstanceInstanceStorageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeOrchestratedInstanceInstanceStorageListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

