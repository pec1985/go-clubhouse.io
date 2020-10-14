package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// List Files returns a list of all Files and related attributes in your Clubhouse.
func (a *api) ListFiles() (*[]models.File, error) {
	params := url.Values{}
	var out []models.File
	if err := a.Request("GET", "/api/v3/files", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
func (a *api) CreateFiles(files *models.CreateFiles) error {
	params := url.Values{}
	var body *bytes.Buffer
	if files != nil {
		jsonbody, _ := json.Marshal(files)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.Request("POST", "/api/v3/files", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Delete File can be used to delete any previously attached File.
func (a *api) DeleteFile(fileID int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.Request("DELETE", "/api/v3/files/"+fmt.Sprint(fileID)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Get File returns information about the selected File.
func (a *api) GetFile(fileID int64) (*models.File, error) {
	params := url.Values{}
	var out models.File
	if err := a.Request("GET", "/api/v3/files/"+fmt.Sprint(fileID)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update File can used to update the properties of a file uploaded to Clubhouse.
func (a *api) UpdateFile(fileID int64, file *models.UpdateFile) (*models.File, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if file != nil {
		jsonbody, _ := json.Marshal(file)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.File
	if err := a.Request("PUT", "/api/v3/files/"+fmt.Sprint(fileID)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
