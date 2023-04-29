package models

type Chapter struct {
	ChapterID        int    `json:"chapter_id"`
	ChapterProjectID int    `json:"chapter_project_id"`
	ChapterName      string `json:"chapter_name"`
}
