package huggo

type Hub struct {
	Collection *Collection
	Search     *Search
	User       *User
}

// NewHub creates a Hub client
func NewHub(apiKey string) (*Hub, error) {
	httpClient, err := NewHttpClient(apiKey)
	if err != nil {
		return nil, err
	}
	hub := &Hub{
		Collection: NewCollection(httpClient),
		Search:     NewSearch(httpClient),
		User:       NewUser(httpClient),
	}
	return hub, nil
}
