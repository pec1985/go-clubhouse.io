package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListEpicStories get a list of all Stories in an Epic.
func (a *api) ListEpicStories(epicID int64, getEpicStories *models.GetEpicStories) (*[]models.StorySlim, error) {
	params := url.Values{}
	if getEpicStories != nil {
		kv := map[string]interface{}{}
		b, _ := json.Marshal(getEpicStories)
		json.Unmarshal(b, &kv)
		for k, v := range kv {
			params.Set(k, fmt.Sprint(v))
		}
	}
	var out []models.StorySlim
	if err := a.Request("GET", "/api/v3/epics/"+fmt.Sprint(epicID)+"/stories", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
