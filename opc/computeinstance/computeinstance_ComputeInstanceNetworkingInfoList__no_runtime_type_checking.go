//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package computeinstance

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeInstanceNetworkingInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeInstanceNetworkingInfoList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeInstanceNetworkingInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeInstanceNetworkingInfoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

