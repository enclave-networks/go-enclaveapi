package policy

type PolicyCreate struct {
	Description  string      `json:"Description,omitempty"`
	IsEnabled    bool        `json:"IsEnabled,omitempty"`
	SenderTags   []string    `json:"SenderTags,omitempty"`
	RecieverTags []string    `json:"RecieverTags,omitempty"`
	Acls         []PolicyAcl `json:"Acls,omitempty"`
	Notes        string      `json:"Notes,omitempty"`
}
