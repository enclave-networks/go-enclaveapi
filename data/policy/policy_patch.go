package policy

// Defines the modifiable properties of a policy.
type PolicyPatch struct {
	Description  string      `json:"Description,omitempty"`
	IsEnabled    bool        `json:"IsEnabled,omitempty"`
	SenderTags   []string    `json:"SenderTags,omitempty"`
	RecieverTags []string    `json:"RecieverTags,omitempty"`
	Acls         []PolicyAcl `json:"Acls,omitempty"`
	Notes        string      `json:"Notes,omitempty"`
}
