package service

type UserProfileProperty struct {
	IsRequired  bool   `json:"isRequired"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	IsPublic    bool   `json:"isPublic"`
}
