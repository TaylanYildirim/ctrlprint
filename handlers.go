package handlers

import (
	"chapter-history-api/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var ErrChapterNotFound = fmt.Errorf("chapter not found")

func GetChapterVersionsHandler(service *service.ChapterVersionService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		chapterID, err := strconv.Atoi(c.Params("chapterID"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid chapterID",
			})
		}

		chapterData, err := service.GetChapterData(chapterID)
		if err != nil {
			if err == ErrChapterNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Chapter not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		return c.JSON(chapterData)
	}
}
