package api

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// ListProjects returns a list of all Projects and their attributes.
func (a *api) ListProjects() (*[]models.Project, error) {
	params := url.Values{}
	var out []models.Project
	if err := a.Request("GET", "/api/v3/projects", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateProject is used to create a new Clubhouse Project.
func (a *api) CreateProject(project *models.CreateProject) error {
	params := url.Values{}
	var body *bytes.Buffer
	if project != nil {
		jsonbody, _ := toPayload(project, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/projects", params, body, &out); err != nil {
		return err
	}
	return nil
}

// DeleteProject can be used to delete a Project. Projects can only be deleted if all associated Stories are moved or deleted. In the case that the Project cannot be deleted, you will receive a 422 response.
func (a *api) DeleteProject(projectID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/projects/"+fmt.Sprint(projectID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// GetProject returns information about the selected Project.
func (a *api) GetProject(projectID int64) (*models.Project, error) {
	params := url.Values{}
	var out models.Project
	if err := a.Request("GET", "/api/v3/projects/"+fmt.Sprint(projectID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateProject can be used to change properties of a Project.
func (a *api) UpdateProject(projectID int64, project *models.UpdateProject) (*models.Project, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if project != nil {
		jsonbody, _ := toPayload(project, false)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.Project
	if err := a.Request("PUT", "/api/v3/projects/"+fmt.Sprint(projectID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
