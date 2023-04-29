package service

import (
	"chapter-history-api/models"
	"database/sql"
	"fmt"
	"time"
)

type ChapterVersionService struct {
	db *sql.DB
}

func NewChapterVersionService(db *sql.DB) *ChapterVersionService {
	return &ChapterVersionService{db: db}
}

type ChapterData struct {
	Company  string                  `json:"company"`
	Project  string                  `json:"project"`
	Chapter  string                  `json:"chapter"`
	Versions []models.ChapterVersion `json:"versions"`
}
type ChapterVersionData struct {
	ChapterVersionID         int       `json:"chapter_version_id"`
	ChapterVersionNumber     int       `json:"chapter_version_number"`
	ChapterVersionCreateDate time.Time `json:"chapter_version_create_date"`
	CreatedBy                string    `json:"created_by"`
	AppVersionName           string    `json:"appversion_name"`
}

var ErrChapterNotFound = fmt.Errorf("chapter not found")

func (c *ChapterVersionService) GetChapterData(chapterID int) (*ChapterData, error) {
	sqlQuery := `
		SELECT 
			chapter.chapter_name,
			project.project_name,
			company.company_name,
			chapter_version.chapter_version_id,
			chapter_version.chapter_version_number,
			chapter_version.chapter_version_create_date,
			person.person_username
		FROM 
			chapter
			JOIN project ON chapter.chapter_project_id = project.project_id
			JOIN company ON project.project_company_id = company.company_id
			JOIN chapter_version ON chapter.chapter_id = chapter_version.chapter_version_chapter_id
			JOIN person ON chapter_version.chapter_version_person_id = person.person_id
		WHERE 
			chapter.chapter_id = $1
		ORDER BY 
			chapter_version.chapter_version_number ASC
	`

	rows, err := c.db.Query(sqlQuery, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chapterData := &ChapterData{}

	for rows.Next() {
		cv := models.ChapterVersion{}
		person := models.Person{}

		err := rows.Scan(
			&chapterData.Chapter,
			&chapterData.Project,
			&chapterData.Company,
			&cv.ChapterVersionID,
			&cv.ChapterVersionNumber,
			&cv.ChapterVersionCreateDate,
			&person.PersonUsername,
		)
		if err != nil {
			return nil, err
		}

		// Convert app version to marketing name
		/*appVersionName, err := convertAppVersionToMarketingName(cv.ChapterVersionAppVersion)
		if err != nil {
			return nil, err
		}*/

		cvData := models.ChapterVersion{
			ChapterVersionID:         cv.ChapterVersionID,
			ChapterVersionNumber:     cv.ChapterVersionNumber,
			ChapterVersionCreateDate: cv.ChapterVersionCreateDate,
		}

		chapterData.Versions = append(chapterData.Versions, cvData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(chapterData.Versions) == 0 {
		return nil, ErrChapterNotFound
	}

	return chapterData, nil
}
