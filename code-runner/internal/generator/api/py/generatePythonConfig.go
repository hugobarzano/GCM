package py

func GenerateDependencies() []byte {
	return []byte(PythonRequirements)
}
