package form

// UnstructuredRequest doc
type UnstructuredRequest struct {
	Namespace     string `json:"namespace" validate:"required"`
	TemplateName  string `json:"templateName" validate:"required,min=2,max=50"`
	CPURequest    string `json:"cpuRequest"`
	MemoryRequest string `json:"memoryRequest"`
	CPULimit      string `json:"cpuLimit"`
	MemoryLimit   string `json:"memoryLimit"`
	CommandParams string `json:"commandParams" validate:"required"`
	TestLabel     string `json:"testLabel" validate:"required,min=2,max=50"`
}
