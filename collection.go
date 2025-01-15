package huggo

import (
	"encoding/json"
	"fmt"
)

type Collection struct {
	httpClient *HttpClient
}

// NewCollection creates a client to interact with the Collection API.
func NewCollection(httpClient *HttpClient) *Collection {
	return &Collection{httpClient: httpClient}
}

func (c *Collection) GetCollections() ([]CollectionInfo, error) {
	var colInfo []CollectionInfo
	err := c.httpClient.Get("/collections", &colInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get collections: %w", err)
	}
	return colInfo, nil
}

type CollectionInfo struct {
	Slug            string           `json:"slug"`
	Title           string           `json:"title"`
	Description     string           `json:"description,omitempty"`
	Gating          bool             `json:"gating"`
	LastUpdated     string           `json:"lastUpdated"`
	Owner           CollectionOwner  `json:"owner"`
	Items           []CollectionItem `json:"items"`
	Theme           string           `json:"theme"`
	Private         bool             `json:"private"`
	Upvotes         int64            `json:"upvotes"`
	IsUpvotedByUser bool             `json:"isUpvotedByUser"`
}

type CollectionItem struct {
	ID               string          `json:"_id"`
	Position         int64           `json:"position"`
	Type             string          `json:"type"`
	Author           string          `json:"author,omitempty"`
	AuthorData       CollectionOwner `json:"authorData,omitempty"`
	Downloads        int64           `json:"downloads,omitempty"`
	Gated            Gated           `json:"gated"`
	ItemID           string          `json:"id"`
	Inference        string          `json:"inference,omitempty"`
	LastModified     string          `json:"lastModified,omitempty"`
	Likes            int64           `json:"likes,omitempty"`
	Private          bool            `json:"private,omitempty"`
	RepoType         string          `json:"repoType,omitempty"`
	IsLikedByUser    bool            `json:"isLikedByUser,omitempty"`
	PipelineTag      string          `json:"pipeline_tag,omitempty"`
	WidgetOutputUrls []interface{}   `json:"widgetOutputUrls,omitempty"`
	Gallery          []string        `json:"gallery,omitempty"`
	Note             struct {
		HTML string `json:"html"`
		Text string `json:"text"`
	} `json:"note,omitempty"`
	Title           string `json:"title,omitempty"`
	ThumbnailURL    string `json:"thumbnailUrl,omitempty"`
	Upvotes         int64  `json:"upvotes,omitempty"`
	PublishedAt     string `json:"publishedAt,omitempty"`
	IsUpvotedByUser bool   `json:"isUpvotedByUser,omitempty"`
}

type CollectionOwner struct {
	AvatarURL     string `json:"avatarUrl"`
	Fullname      string `json:"fullname"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	IsHF          bool   `json:"isHf"`
	IsMod         bool   `json:"isMod"`
	IsEnterprise  bool   `json:"isEnterprise,omitempty"`
	FollowerCount int64  `json:"followerCount"`
	ID            string `json:"_id,omitempty"`
	IsPro         bool   `json:"isPro,omitempty"`
}

type Gated struct {
	val interface{}
}

func (g *Gated) UnmarshalJSON(data []byte) error {
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		g.val = b
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		g.val = s
		return nil
	}
	return fmt.Errorf("unsupported type for Gated: %s", string(data))
}

func (g Gated) Value() interface{} {
	return g.val
}
