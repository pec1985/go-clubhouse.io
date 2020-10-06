package api

import (
	"net/url"
)

// Disables the Story Template feature for the given Organization.
func (a *api) DisableStoryTemplates() error {
	params := url.Values{}
	var out interface{}
	if err := a.request("PUT", "/api/v3/entity-templates/disable", params, nil, &out); err != nil {
		return err
	}
	return nil
}
