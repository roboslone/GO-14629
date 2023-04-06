// Commercial secret, LLC "RevTech". Refer to CONFIDENTIAL file in the root for details

package someProject

import (
	moduleLibrary "moduleLibrary"
)

var (
	_ = moduleLibrary.ModuleMap{
		"metrics": &moduleLibrary.ModMetrics[*Config, *State]{},
	}
)
