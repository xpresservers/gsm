package specification

type SpecificationDetail struct {
	Network  interface{} `json:"network,omitempty"`
	Launch   interface{} `json:"launch,omitempty"`
	Body     interface{} `json:"body,omitempty"`
	Display  interface{} `json:"display,omitempty"`
	Platform interface{} `json:"platform,omitempty"`
	Memory   interface{} `json:"memory,omitempty"`
	Camera   interface{} `json:"camera,omitempty"`
	Sound    interface{} `json:"sound,omitempty"`
	Comms    interface{} `json:"comms,omitempty"`
	Features interface{} `json:"features,omitempty"`
	Battery  interface{} `json:"battery,omitempty"`
	Misc     interface{} `json:"misc,omitempty"`
	Tests    interface{} `json:"tests,omitempty"`
}
