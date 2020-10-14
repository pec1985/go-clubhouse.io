package models

import (
	"encoding/json"
	"time"
)

// Iteration An Iteration is a defined, time-boxed period of development for a collection of Stories. See https://help.clubhouse.io/hc/en-us/articles/360028953452-Iterations-Overview for more information.
type Iteration struct {
	// AppURL The Clubhouse application url for the Iteration.
	AppURL string `json:"app_url"`
	// CreatedAt The instant when this iteration was created.
	CreatedAt time.Time `json:"created_at"`
	// Description The description of the iteration.
	Description string `json:"description"`
	// EndDate The date this iteration begins.
	EndDate string `json:"end_date"`
	// EntityType A string description of this resource
	EntityType string `json:"entity_type"`
	// FollowerIDs An array of UUIDs for any Members listed as Followers.
	FollowerIDs []string `json:"follower_ids"`
	// GroupIDs An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIDs []string `json:"group_ids"`
	// GroupMentionIDs An array of Group IDs that have been mentioned in the Story description.
	GroupMentionIDs []string `json:"group_mention_ids"`
	// ID The ID of the iteration.
	ID int64 `json:"id"`
	// Labels An array of labels attached to the iteration.
	Labels []Label `json:"labels"`
	// MemberMentionIDs An array of Member IDs that have been mentioned in the Story description.
	MemberMentionIDs []string `json:"member_mention_ids"`
	// MentionIDs Deprecated: use member_mention_ids.
	MentionIDs []string `json:"mention_ids"`
	// Name The name of the iteration.
	Name string `json:"name"`
	// StartDate The date this iteration begins.
	StartDate string         `json:"start_date"`
	Stats     IterationStats `json:"stats"`
	// Status The status of the iteration. Values are either "unstarted", "started", or "done".
	Status string `json:"status"`
	// UpdatedAt The instant when this iteration was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Iteration) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Iteration) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
