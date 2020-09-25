package api

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/pec1985/go-clubhouse.io/api/models"
)

func (a *api) ListGroups() (*[]models.Group, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out []models.Group
	if err := a.request("GET", "/api/v3/groups", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) CreateGroup(createGroup *models.CreateGroup) error {
	params := url.Values{}
	jsonbody, _ := json.Marshal(createGroup)
	body := bytes.NewBuffer(jsonbody)
	var out interface{}
	if err := a.request("POST", "/api/v3/groups", params, body, &out); err != nil {
		return err
	}
	return nil
}
func (a *api) GetGroup(groupPublicId string) (*models.Group, error) {
	params := url.Values{}
	body := bytes.Buffer{}
	var out models.Group
	if err := a.request("GET", "/api/v3/groups/"+groupPublicId+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) UpdateGroup(groupPublicId string, updateGroup *models.UpdateGroup) (*models.Group, error) {
	params := url.Values{}
	jsonbody, _ := json.Marshal(updateGroup)
	body := bytes.NewBuffer(jsonbody)
	var out models.Group
	if err := a.request("PUT", "/api/v3/groups/"+groupPublicId+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
