package request

type Scaling struct {
	Max int `json:"max,omitempty"`
	Min int `json:"min,omitempty"`
}
