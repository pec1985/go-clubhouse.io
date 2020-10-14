package models

import "encoding/json"

type CreateOrDeleteReaction struct {
	// Emoji The emoji short-code to add / remove. E.g. `:thumbsup::skin-tone-4:`.
	Emoji string `json:"emoji"`
}

func (m *CreateOrDeleteReaction) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *CreateOrDeleteReaction) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
