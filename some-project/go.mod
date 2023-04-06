module someProject

go 1.20

require (
	moduleLibrary v0.0.0
)

replace (
	moduleLibrary => ../module-library
)
