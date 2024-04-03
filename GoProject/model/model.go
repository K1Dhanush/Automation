package model

type Report struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Namespace struct {
	Region   string `json:"region,omitempty"`
	Instance string `json:"instance"`
	Id       string `json:"id"`
}
