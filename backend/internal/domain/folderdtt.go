package domain

type MoveRename struct {
	Path string `json:"name" binding:"required"`
}

type FolderRename struct {
	Name string `json:"name" binding:"required"`
}
