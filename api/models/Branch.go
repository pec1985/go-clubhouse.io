package models

import (
	"encoding/json"
	"time"
)

type Branch struct {
	// An array of PullRequests attached to the Branch (there is usually only one).
	PullRequests []PullRequest `json:"pull_requests"`
	// A true/false boolean indicating if the Branch has been deleted.
	Deleted bool `json:"deleted"`
	// The unique ID of the Branch.
	Id *int64 `json:"id"`
	// The name of the Branch.
	Name string `json:"name"`
	// A true/false boolean indicating if the Branch is persistent; e.g. master.
	Persistent bool `json:"persistent"`
	// The time/date the Branch was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// The URL of the Branch.
	Url string `json:"url"`
	// The time/date the Branch was created.
	CreatedAt *time.Time `json:"created_at"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The IDs of the Branches the Branch has been merged into.
	MergedBranchIds []int64 `json:"merged_branch_ids"`
	// The ID of the Repository that contains the Branch.
	RepositoryId *int64 `json:"repository_id"`
}

func (m *Branch) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Branch) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
