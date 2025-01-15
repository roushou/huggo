package huggo

type User struct {
	httpClient *HttpClient
}

// NewUser creates a client to interact with the User API.
func NewUser(httpClient *HttpClient) *User {
	return &User{httpClient: httpClient}
}

// WhoAmI fetches the user information.
func (u *User) WhoAmI() (*UserInfo, error) {
	var me UserInfo
	err := u.httpClient.Get("/whoami-v2", &me)
	if err != nil {
		return nil, err
	}
	return &me, nil
}

type UserInfo struct {
	Type      string        `json:"type"`
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Fullname  string        `json:"fullname"`
	IsPro     bool          `json:"isPro"`
	AvatarURL string        `json:"avatarUrl"`
	Orgs      []interface{} `json:"orgs"`
	Auth      Auth          `json:"auth"`
}

type Auth struct {
	Type        string      `json:"type"`
	AccessToken AccessToken `json:"accessToken"`
}

type AccessToken struct {
	DisplayName string      `json:"displayName"`
	Role        string      `json:"role"`
	CreatedAt   string      `json:"createdAt"`
	FineGrained FineGrained `json:"fineGrained"`
}

type FineGrained struct {
	CanReadGatedRepos bool     `json:"canReadGatedRepos"`
	Global            []string `json:"global"`
	Scoped            []Scoped `json:"scoped"`
}

type Scoped struct {
	Entity      Entity   `json:"entity"`
	Permissions []string `json:"permissions"`
}

type Entity struct {
	ID   string `json:"_id"`
	Type string `json:"type"`
	Name string `json:"name"`
}
