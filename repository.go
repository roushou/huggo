package huggo

import "fmt"

type Repository struct {
	httpClient *HttpClient
}

// NewRepository creates a client to interact with the Repository API.
func NewRepository(httpClient *HttpClient) *Repository {
	return &Repository{httpClient: httpClient}
}

type CreateRepositoryPayload struct {
	// Type specifies the type of repository. Valid values are "dataset", "space", or "model". Defaults to "model".
	Type string `json:"type"`
	// Name is the name of the repository.
	Name string `json:"name"`
	// Organization is the organization ID in which to create the repository.
	Organization string `json:"organization"`
	// Visibility determines if the repository is "public" or "private".
	Visibility string `json:"private"`
	// SDK specifies the SDK to use for the repository when its type is "space". Valid values are "streamlit", "gradio", "docker", or "static".
	SDK string `json:"sdk,omitempty"`
}

// CreateRepository creates a new repository.
func (r *Repository) CreateRepository(payload CreateRepositoryPayload) error {
	err := r.httpClient.Post("/repos/create", payload, nil)
	if err != nil {
		return fmt.Errorf("failed to create repository: %v", err)
	}
	return nil
}

type DeleteRepositoryPayload struct {
	// Type specifies the type of repository. Valid values are "dataset", "space", or "model". Defaults to "model".
	Type string `json:"type"`
	// Name is the name of the repository.
	Name string `json:"name"`
	// Organization is the organization id in which to create the repository.
	Organization string `json:"organizatoin,omitempty"`
}

// DeleteRepository deletes a repository.
func (r *Repository) DeleteRepository(payload DeleteRepositoryPayload) error {
	err := r.httpClient.Delete("/repos/delete", payload, nil)
	if err != nil {
		return fmt.Errorf("failed to delete repository: %v", err)
	}
	return nil
}

type MoveRepositoryPayload struct {
	// Type specifies the type of repository. Valid values are "dataset", "space", or "model". Defaults to "model".
	Type string `json:"type"`
	// From is the current name of the repository.
	From string `json:"fromRepo"`
	// To is the new name for the repository.
	To string `json:"toRepo"`
}

// MoveRepository moves a repository within the same namespace or transfer from a user to an organization.
func (r *Repository) MoveRepository(payload MoveRepositoryPayload) error {
	err := r.httpClient.Post("/repos/move", payload, nil)
	if err != nil {
		return fmt.Errorf("failed to move repository: %v", err)
	}
	return nil
}

type UpdateVisibilityPayload struct {
	// Visibility determines if the repository should be "public" or "private".
	Visibility string `json:"private"`
}

// UpdateRepositoryVisibility updates the repository's visibility.
func (r *Repository) UpdateRepositoryVisibility(repositoryType string, repositoryID string, payload UpdateVisibilityPayload) error {
	path := fmt.Sprintf("/repos/%s/%s", repositoryType, repositoryID)
	err := r.httpClient.Put(path, payload, nil)
	if err != nil {
		return fmt.Errorf("failed to create repository: %v", err)
	}
	return nil
}
