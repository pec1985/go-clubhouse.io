package api

import (
	"net/url"
)

// Enables the Story Template feature for the given Organization.
func (a *api) EnableStoryTemplates() error {
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/entity-templates/enable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
