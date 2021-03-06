package models

type UpdateLabel struct {
	// Archived a true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived,omitempty"`
	// Color the hex color to be displayed with the Label (for example, "#ff0000").
	Color *string `json:"color,omitempty"`
	// Description the new description of the label.
	Description string `json:"description,omitempty"`
	// Name the new name of the label.
	Name string `json:"name,omitempty"`
}

func (m *UpdateLabel) Stringify() string {
	b, _ := toPayload(m, false)
	return string(b)
}
func (m *UpdateLabel) StringifyPretty() string {
	b, _ := toPayload(m, true)
	return string(b)
}
