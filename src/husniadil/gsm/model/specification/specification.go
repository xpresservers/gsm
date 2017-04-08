package specification

type Specification struct {
	Overview SpecificationOverview `json:"overview,omitempty"`
	Detail   SpecificationDetail   `json:"detail,omitempty"`
}
