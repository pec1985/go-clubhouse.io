package models

import "encoding/json"

// ProjectStats a group of calculated values for this Project.
type ProjectStats struct {
	// NumPoints the total number of points in this Project.
	NumPoints int64 `json:"num_points"`
	// NumRelatedDocuments the total number of documents related to this Project
	NumRelatedDocuments int64 `json:"num_related_documents"`
	// NumStories the total number of stories in this Project.
	NumStories int64 `json:"num_stories"`
}

func (m *ProjectStats) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *ProjectStats) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
