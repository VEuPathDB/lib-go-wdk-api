package read

type Service struct {
	ChangePasswordUrl string `json:"changePasswordUrl"`

	ReleaseDate string `json:"releaseDate"`

	DisplayName string `json:"displayName"`

	UserDatasetStoreStatus struct {
		IsAvailable bool `json:"isAvailable"`
	} `json:"userDatasetStoreStatus"`

	CategoriesOntologyName string `json:"categoriesOntologyName"`

	Description string `json:"description"`

	StartupTime uint64 `json:"startupTime"`

	ProjectId string `json:"projectId"`

	BuildNumber string `json:"buildNumber"`

	UserProfileProperties []UserProfileProperty `json:"userProfileProperties"`

	Authentication struct {
		OAuthUrl       string `json:"oauthUrl"`
		Method         string `json:"method"`
		OAuthClientId  string `json:"oauthClientId"`
		OAuthClientUrl string `json:"oauthClientUrl"`
	} `json:"authentication"`
}

type UserProfileProperty struct {
	IsRequired  bool   `json:"isRequired"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	IsPublic    bool   `json:"isPublic"`
}
