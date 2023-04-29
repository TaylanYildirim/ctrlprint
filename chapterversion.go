package models

import "time"

type ChapterVersion struct {
	ChapterVersionID         int       `json:"chapter_version_id"`
	ChapterVersionChapterID  int       `json:"chapter_version_chapter_id"`
	ChapterVersionNumber     int       `json:"chapter_version_number"`
	ChapterVersionCreateDate time.Time `json:"chapter_version_create_date"`
	ChapterVersionPersonID   int       `json:"chapter_version_person_id"`
	ChapterVersionAppVersion string    `json:"chapter_version_appversion"`
}
