package settings

type SaveSettingsInput struct {
	CountInviting int `json:"countInviting"`
	CountMailing  int `json:"countMailing"`
}

func NewSaveSettingsInput(countInviting, countMailing int) SaveSettingsInput {
	return SaveSettingsInput{
		CountInviting: countInviting,
		CountMailing:  countMailing,
	}
}
