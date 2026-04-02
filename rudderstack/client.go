package rudderstack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rudderlabs/rudder-api-go/client"
)

type Client struct {
	Sources         SourcesService
	Destinations    DestinationsService
	Connections     ConnectionsService
	Transformations TransformationsService
}

type Transformation struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	IsPublished bool   `json:"isPublished"`
	TestJSON    string `json:"testJson,omitempty"`
}

type TransformationsService interface {
	Create(ctx context.Context, transformation *Transformation) (*Transformation, error)
	Get(ctx context.Context, id string) (*Transformation, error)
	Update(ctx context.Context, transformation *Transformation) (*Transformation, error)
	Delete(ctx context.Context, id string) error
}

func NewAPIClient(accessToken string, options ...client.Option) (*Client, error) {
	api, err := client.New(accessToken, options...)
	if err != nil {
		return nil, err
	}

	return &Client{
		Sources:         api.Sources,
		Destinations:    api.Destinations,
		Connections:     api.Connections,
		Transformations: &transformationsService{accessToken: accessToken, baseURL: apiUrl},
	}, nil
}

type transformationsService struct {
	accessToken string
	baseURL     string
}

func (s *transformationsService) Create(ctx context.Context, transformation *Transformation) (*Transformation, error) {
	data, err := json.Marshal(transformation)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/transformations", s.baseURL), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create transformation: %s", string(body))
	}

	var result struct {
		Transformation *Transformation `json:"transformation"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Transformation, nil
}

func (s *transformationsService) Get(ctx context.Context, id string) (*Transformation, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/transformations/%s", s.baseURL, id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("transformation not found")
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get transformation: %s", string(body))
	}

	var result struct {
		Transformation *Transformation `json:"transformation"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Transformation, nil
}

func (s *transformationsService) Update(ctx context.Context, transformation *Transformation) (*Transformation, error) {
	data, err := json.Marshal(transformation)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf("%s/transformations/%s", s.baseURL, transformation.ID), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to update transformation: %s", string(body))
	}

	var result struct {
		Transformation *Transformation `json:"transformation"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Transformation, nil
}

func (s *transformationsService) Delete(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("%s/transformations/%s", s.baseURL, id), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete transformation: %s", string(body))
	}

	return nil
}
