package huggo

import (
	"fmt"
	"time"
)

type Search struct {
	httpClient *HttpClient
}

// NewSearch creates a client to interact with the Search API.
func NewSearch(httpClient *HttpClient) *Search {
	return &Search{httpClient: httpClient}
}

// GetModels fetches paginated information from all models.
func (s *Search) GetModels() ([]Model, error) {
	var models []Model
	err := s.httpClient.Get("/models", &models)
	if err != nil {
		return nil, err
	}
	return models, nil
}

// GetModel fetches all the information for a specific model.
func (s *Search) GetModel(id string) (*Model, error) {
	var model Model
	path := fmt.Sprintf("/models/%s", id)
	err := s.httpClient.Get(path, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

// GetDatasets fetches information from all datasets.
func (s *Search) GetDatasets() ([]Dataset, error) {
	var datasets []Dataset
	err := s.httpClient.Get("/datasets", &datasets)
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

// GetDataset fetches all information for a specific dataset.
func (s *Search) GetDataset(id string) (*Dataset, error) {
	var dataset Dataset
	path := fmt.Sprintf("/datasets/%s", id)
	err := s.httpClient.Get(path, &dataset)
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

// GetMetrics fetches metrics
func (s *Search) GetDatasetsTags() (*DatasetTags, error) {
	var tags DatasetTags
	err := s.httpClient.Get("/datasets-tags-by-type", &tags)
	if err != nil {
		return nil, err
	}
	return &tags, nil
}

// GetSpaces fetches information from all spaces.
func (s *Search) GetSpaces() ([]Space, error) {
	var spaces []Space
	err := s.httpClient.Get("/spaces", &spaces)
	if err != nil {
		return nil, err
	}
	return spaces, nil
}

// GetSpaceByRepository fetches spaces associated to repository.
func (s *Search) GetSpacesByRepository(repositoryID string) (*Space, error) {
	var space Space
	path := fmt.Sprintf("/spaces/%s", repositoryID)
	err := s.httpClient.Get(path, &space)
	if err != nil {
		return nil, err
	}
	return &space, nil
}

// GetMetrics fetches metrics.
func (s *Search) GetMetrics() ([]Metric, error) {
	var metrics []Metric
	err := s.httpClient.Get("/metrics", &metrics)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

type Token struct {
	Type       string `json:"__type"`
	Content    string `json:"content"`
	Lstrip     bool   `json:"lstrip"`
	Normalized bool   `json:"normalized"`
	Rstrip     bool   `json:"rstrip"`
	SingleWord bool   `json:"single_word"`
}

type TokenizerConfig struct {
	BosToken     Token  `json:"bos_token,omitempty"`
	EosToken     Token  `json:"eos_token,omitempty"`
	PadToken     Token  `json:"pad_token,omitempty"`
	UnkToken     any    `json:"unk_token,omitempty"`
	ChatTemplate string `json:"chat_template,omitempty"`
}

type ModelConfig struct {
	Architectures      []string          `json:"architectures,omitempty"`
	AutoMap            map[string]string `json:"auto_map,omitempty"`
	ModelType          string            `json:"model_type,omitempty"`
	QuantizationConfig struct {
		QuantMethod string `json:"quant_method,omitempty"`
	} `json:"quantization_config,omitempty"`
	TokenizerConfig TokenizerConfig `json:"tokenizer_config,omitempty"`
}

type Model struct {
	ID_           string      `json:"_id"`
	ID            string      `json:"id"`
	Author        string      `json:"author"`
	Gated         bool        `json:"gated"`
	Inference     string      `json:"inference"`
	LastModified  time.Time   `json:"lastModified"`
	Likes         int         `json:"likes"`
	TrendingScore int         `json:"trendingScore"`
	Private       bool        `json:"private"`
	Sha           string      `json:"sha"`
	Config        ModelConfig `json:"config"`
	Downloads     int         `json:"downloads"`
	Tags          []string    `json:"tags"`
	PipelineTag   string      `json:"pipeline_tag"`
	LibraryName   string      `json:"library_name,omitempty"`
	CreatedAt     time.Time   `json:"createdAt"`
	ModelID       string      `json:"modelId"`
	Siblings      []Sibling   `json:"siblings"`
}

type Dataset struct {
	ID            string          `json:"_id"`
	Welcome8ID    string          `json:"id"`
	Author        string          `json:"author"`
	CardData      DatasetCardData `json:"cardData"`
	Disabled      bool            `json:"disabled"`
	Gated         bool            `json:"gated"`
	LastModified  string          `json:"lastModified"`
	Likes         int64           `json:"likes"`
	TrendingScore int64           `json:"trendingScore"`
	Private       bool            `json:"private"`
	SHA           string          `json:"sha"`
	Description   string          `json:"description"`
	Downloads     int64           `json:"downloads"`
	Tags          []string        `json:"tags"`
	CreatedAt     string          `json:"createdAt"`
	Key           string          `json:"key"`
}

type DatasetCardData struct {
	License        string   `json:"license"`
	Tags           []string `json:"tags,omitempty"`
	TaskCategories []string `json:"task_categories,omitempty"`
	SizeCategories []string `json:"size_categories,omitempty"`
	Language       []string `json:"language,omitempty"`
	Configs        []Config `json:"configs,omitempty"`
	PrettyName     *string  `json:"pretty_name,omitempty"`
}

type Config struct {
	ConfigName string     `json:"config_name"`
	DataFiles  []DataFile `json:"data_files"`
}

type DataFile struct {
	Split string `json:"split"`
	Path  string `json:"path"`
}

type Space struct {
	ID_           string        `json:"_id"`
	ID            string        `json:"id"`
	Author        string        `json:"author"`
	CardData      SpaceCardData `json:"cardData"`
	LastModified  string        `json:"lastModified"`
	Likes         int64         `json:"likes"`
	TrendingScore int64         `json:"trendingScore"`
	Private       bool          `json:"private"`
	SHA           string        `json:"sha"`
	Subdomain     string        `json:"subdomain"`
	SDK           string        `json:"sdk"`
	Tags          []string      `json:"tags"`
	CreatedAt     string        `json:"createdAt"`
	Siblings      []Sibling     `json:"siblings"`
}

type SpaceCardData struct {
	Title            string  `json:"title"`
	Emoji            string  `json:"emoji"`
	ColorFrom        string  `json:"colorFrom"`
	ColorTo          string  `json:"colorTo"`
	SDK              string  `json:"sdk"`
	SDKVersion       string  `json:"sdk_version"`
	AppFile          string  `json:"app_file"`
	Pinned           bool    `json:"pinned"`
	License          *string `json:"license,omitempty"`
	ShortDescription *string `json:"short_description,omitempty"`
}

type DatasetTags struct {
	Library []Library `json:"library"`
}

type Library struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

type Sibling struct {
	Rfilename string `json:"rfilename"`
}

type Metric struct {
	ID          string  `json:"id"`
	SpaceID     string  `json:"spaceId"`
	Description *string `json:"description,omitempty"`
}
