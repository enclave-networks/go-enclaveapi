package trustrequirement

type TrustRequirementSettings struct {
	Configuration map[string]string   `json:"Configuration,omitempty"`
	Conditions    []map[string]string `json:"Conditions,omitempty"`
}
