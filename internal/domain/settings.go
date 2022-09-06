package domain

type Settings struct {
	CountInviting int `json:"countInviting" bson:"countInviting"`
	CountMailing  int `json:"countMailing" bson:"countMailing"`
}
