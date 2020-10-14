package models

import (
	"encoding/json"
	"time"
)

// Team Group of Projects with the same Workflow.
type Team struct {
	// CreatedAt The time/date the Team was created.
	CreatedAt time.Time `json:"created_at"`
	// Description The description of the Team.
	Description string `json:"description"`
	// EntityType A string description of this resource.
	EntityType string `json:"entity_type"`
	// ID The unique identifier of the Team.
	ID int64 `json:"id"`
	// Name The name of the Team.
	Name string `json:"name"`
	// Position A number representing the position of the Team in relation to every other Team within the Organization.
	Position float64 `json:"position"`
	// ProjectIDs An array of IDs of projects within the Team.
	ProjectIDs []float64 `json:"project_ids"`
	// UpdatedAt The time/date the Team was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	Workflow  Workflow  `json:"workflow"`
}

func (m *Team) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Team) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
