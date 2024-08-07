package models

type Image struct {
	ID         string `json:"id"`
	Containers int64  `json:"containers"`
	Created    int64  `json:"created"`
	Size       int64  `json:"size"`
}
