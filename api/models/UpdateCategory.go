package models

import "encoding/json"

type UpdateCategory struct {
	// Archived A true/false boolean indicating if the Category has been archived.
	Archived bool `json:"archived"`
	// Color The hex color to be displayed with the Category (for example, "#ff0000").
	Color *string `json:"color"`
	// Name The new name of the Category.
	Name string `json:"name"`
}

func (m *UpdateCategory) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *UpdateCategory) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
