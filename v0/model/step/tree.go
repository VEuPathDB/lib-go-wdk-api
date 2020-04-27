package step

type TreeStep struct {
	Id             uint64    `json:"stepId"`
	PrimaryInput   *TreeStep `json:"primaryInput"`
	SecondaryInput *TreeStep `json:"secondaryInput"`
}
