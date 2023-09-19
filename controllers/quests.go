package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/krhone/go-quest/models"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

type QuestInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Reward      int    `json:"reward" validate:"required"`
}

// CreateQuest godoc
// @Summary Create a quest
// @Description Create a new quest item
// @Tags quests
// @Produce json
// @Param quest body QuestInput true "New Quest"
// @Success 200 {object} models.Quest
// @Failure 500 {object} object
// @Router /v1/quest [post]
func CreateQuest(context echo.Context) error {
	var input QuestInput
	if err := context.Bind(&input); err != nil {
		fmt.Println("esto es bind")
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Validation Error"})
	}

	fmt.Printf("Input: %+v\n", input)
	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		fmt.Println(err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Validation Error"})
	}

	quest, err := models.NewQuest(input.Title, input.Description, input.Reward)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating quest"})
	}

	return context.JSON(http.StatusOK, quest)

}

// GetQuests godoc
// @Summary Get all quests
// @Description Get all quests items
// @Tags quests
// @Produce json
// @Success 200 {array} models.Quest
// @Failure 500 {object} object
// @Router /v1/quests [get]
func GetAllQuests(context echo.Context) error {
	var quests []models.Quest

	models.DB.Find(&quests)

	return context.JSON(http.StatusOK, quests)
}

// GetQuest get Quest by Id
// @Summary Get one quest
// @Description Get one quest item
// @Tags quests
// @Produce json
// @Param id path string true "Quest ID"
// @Success 200 {object} models.Quest
// @Failure 400,404,500 {object} object
// @Router /v1/quest/{id} [get]
func GetQuest(context echo.Context) error {
	id := context.Param("id")

	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		return context.JSON(http.StatusNotFound, map[string]string{"error": "Quest not found"})
	}
	return context.JSON(200, quest)
}

// DeleteQuest godoc
// @Summary Delete one quest
// @Description Delete one quest item
// @Tags quests
// @Produce json
// @Param id path string true "Quest ID"
// @Success 200 {object} models.Quest
// @Failure 400,404,500 {object} object
// @Router /v1/quest/{id} [delete]
func DeleteQuest(context echo.Context) error {
	id := context.Param("id")

	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		return context.JSON(http.StatusNotFound, map[string]string{"error": "Quest not found"})
	}

	models.DB.Delete(&quest)

	return context.NoContent(http.StatusNoContent)
}

// UpdateQuest godoc
// @Summary Update quest
// @Description Update quest item
// @Tags quests
// @Produce json
// @Param id path string true "Quest ID"
// @Success 200 {object} models.Quest
// @Failure 400,404,500 {object} object
// @Router /v1/quest/{id} [put]
func UpdateQuest(context echo.Context) error {
	id := context.Param("id")
	var quest models.Quest

	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		return context.JSON(http.StatusNotFound, map[string]string{"error": "Quest not found"})
	}

	var input QuestInput

	if err := context.Bind(&input); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Validation Error"})
	}

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Validation Error"})
	}

	quest.Title = input.Title
	quest.Description = input.Description
	quest.Reward = input.Reward

	models.DB.Save(&quest)

	return context.JSON(http.StatusOK, quest)
}
