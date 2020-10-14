package models

import "encoding/json"

type CreateFiles struct {
	// StoryID The story ID that this file will be associated with.
	StoryID int64 `json:"story_id"`
}

func (m *CreateFiles) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateFiles) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
