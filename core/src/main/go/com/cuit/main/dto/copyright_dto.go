package dto

type CreateCopyrightDTO struct {
	CopyrightName string `json:"copyright_name" binding:"required"`
	Label         string `json:"label"`
	DigitalMark   string `json:"digital_mark"`
	Hash          string `json:"hash"`
}
