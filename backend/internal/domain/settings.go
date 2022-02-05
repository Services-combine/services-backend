package domain

type Settings struct {
	ID            string `json:"id" bson:"_id"`
	CountInviting int    `json:"countInviting" bson:"countInviting"`
	CountMailing  int    `json:"countMailing" bson:"countMailing"`
}
