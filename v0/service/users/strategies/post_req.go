package strategies

type CopyRequest struct {
	SourceStrategySignature string `json:"sourceStrategySignature"`
}

type CreateRequest struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	SavedName   *string `json:"savedName,omitempty"`
	IsSaved     bool    `json:"isSaved"`
	IsPublic    bool    `json:"isPublic"`
	StepTree    Step    `json:"stepTree"`
}

func (c *CreateRequest) HasDescription() bool {
	return c.Description != nil
}

func (c *CreateRequest) GetDescription() string {
	return *c.Description
}

func (c *CreateRequest) SetDescription(desc string) *CreateRequest {
	c.Description = &desc
	return c
}

func (c *CreateRequest) HasSavedName() bool {
	return c.SavedName != nil
}

func (c *CreateRequest) GetSavedName() string {
	return *c.SavedName
}

func (c *CreateRequest) SetSavedName(name string) *CreateRequest {
	c.SavedName = &name
	return c
}

