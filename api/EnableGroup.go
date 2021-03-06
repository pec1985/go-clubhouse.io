package api

import (
	"net/url"
)

// EnableGroups enables Groups for the current workspace2
func (a *api) EnableGroups() error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("PUT", "/api/v3/groups/enable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
