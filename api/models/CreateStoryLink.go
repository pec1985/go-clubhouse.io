package models

import "encoding/json"

type CreateStoryLink struct {
	// ObjectID the ID of the object Story.
	ObjectID int64 `json:"object_id"`
	// SubjectID the ID of the subject Story.
	SubjectID int64 `json:"subject_id"`
	// Verb the type of link.
	Verb string `json:"verb"`
}

func (m *CreateStoryLink) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateStoryLink) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
