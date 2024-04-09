package user

type Preferences struct {
	Global  map[string]string `json:"global"`
	Project map[string]string `json:"project"`
}

type Profile struct {
	Id      uint `json:"id"`
	IsGuest bool `json:"isGuest"`

	// Email will only be set if the user being requested is
	// also the user making the request.
	Email *string `json:"email"`

	// Properties will only be set if the user being requested
	// is also the user making the request.
	Properties map[Property]string `json:"properties"`

	// Preferences will only be set if the user being
	// requested is also the user making the request and
	// additionally the query param
	// "includedPreferences=true" is set on the HTTP request.
	Preferences *Preferences `json:"preferences"`
}

const (
	PropertyFirstName    Property = "firstName"
	PropertyLastName     Property = "lastName"
	PropertyMiddleName   Property = "middleName"
	PropertyOrganization Property = "organization"
)

type Property string
