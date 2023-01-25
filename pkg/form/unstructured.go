package form

// UnstructuredRequest doc
type UnstructuredRequest struct {
	Namespace     string `json:"namespace" validate:"required"`
	TemplateName  string `json:"template_name" validate:"required,min=2,max=40"`
	CPURequest    string `json:"cpu_request"`
	MemoryRequest string `json:"memory_request"`
	CPULimit      string `json:"cpu_limit"`
	MemoryLimit   string `json:"memory_limit"`
	CommandParams string `json:"commandParams" validate:"required"`
}
