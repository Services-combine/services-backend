package domain

type Settings struct {
	CountInviting int `json:"countInviting" bson:"countInviting"`
	CountMailing  int `json:"countMailing" bson:"countMailing"`
}

type Marks struct {
	Marks []Mark `json:"marks" binding:"required"`
}

type Mark struct {
	Title string `json:"title" binding:"required"`
	Color string `json:"color" binding:"required"`
}
