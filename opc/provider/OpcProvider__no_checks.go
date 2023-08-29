// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpcProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (o *jsiiProxy_OpcProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateOpcProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOpcProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateOpcProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_OpcProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func validateNewOpcProviderParameters(scope constructs.Construct, id *string, config *OpcProviderConfig) error {
	return nil
}

