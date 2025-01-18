package handler

import (
	"job-openings/config"
	"job-openings/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	logger   *config.Logger
	db       *gorm.DB
	validate *validator.Validate
)

func InitHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetDatabase()
	validate = validator.New()
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func sendError(ctx *gin.Context, msg string, code int) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

// @BasePath /api/v1

// @Summary      Get opening
// @Description  Get opening by id
// @Tags         openings
// @Accept       json
// @Produce      json
// @Param        id   query  int  true  "opening id"
// @Success      200  {object}  schemas.OpeningResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /opening [get]
func GetOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, "Missing required id param!", http.StatusBadRequest)
		return
	}

	var opening = schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, err.Error(), http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": opening,
	})
}

type CreateOpeningRequest struct {
	Role     string `json:"role" validate:"required"`
	Company  string `json:"company" validate:"required"`
	Location string `json:"location" validate:"required"`
	Salary   int32  `json:"salary" validate:"required,gt=0"`
	Link     string `json:"link" validate:"required,url"`
	Remote   *bool  `json:"remote"`
}

// @Summary      Create opening
// @Description  Create opening
// @Tags         openings
// @Accept       json
// @Produce      json
// @Param        body  body  CreateOpeningRequest  true  "body"
// @Success      200  {object}  schemas.OpeningResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /opening [post]
func CreateOpening(ctx *gin.Context) {
	requestBody := CreateOpeningRequest{}

	if err := ctx.BindJSON(&requestBody); err != nil {
		sendError(ctx, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(&requestBody); err != nil {
		sendError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	opening := schemas.Opening{
		Role:     requestBody.Role,
		Company:  requestBody.Company,
		Location: requestBody.Location,
		Salary:   float32(requestBody.Salary),
		Link:     requestBody.Link,
		Remote:   *requestBody.Remote,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("Error at create opening %v", err)
		sendError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Opening created with success!",
		"data":    opening,
	})
}

type UpdateOpeningRequest struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Salary   int32  `json:"salary"`
	Link     string `json:"link"`
	Remote   *bool  `json:"remote"`
}

// @Summary      Update opening
// @Description  Update opening
// @Tags         openings
// @Accept       json
// @Produce      json
// @Param        body  body  UpdateOpeningRequest  true  "Request body"
// @Success      200  {object}  schemas.OpeningResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /opening [put]
func UpdateOpening(ctx *gin.Context) {
	requestBody := UpdateOpeningRequest{}

	if err := ctx.BindJSON(&requestBody); err != nil {
		sendError(ctx, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestBody.ID == "" {
		sendError(ctx, "Missing required id param!", http.StatusBadRequest)
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, requestBody.ID).Error; err != nil {
		sendError(ctx, err.Error(), http.StatusNotFound)
		return
	}

	if requestBody.Role != opening.Role {
		opening.Role = requestBody.Role
	}
	if requestBody.Company != opening.Company {
		opening.Company = requestBody.Company
	}
	if requestBody.Location != opening.Location {
		opening.Location = requestBody.Location
	}
	if requestBody.Salary != int32(opening.Salary) {
		opening.Salary = float32(requestBody.Salary)
	}
	if requestBody.Link != opening.Link {
		opening.Link = requestBody.Link
	}
	if *requestBody.Remote != opening.Remote {
		opening.Remote = *requestBody.Remote
	}

	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, "Error at updating opening: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Opening updated with success!",
		"data":    opening,
	})
}

type DeleteOpeningResponse struct {
	Message string `json:"message"`
}

// @Summary      Delete opening
// @Description  Delete opening by id
// @Tags         openings
// @Accept       json
// @Produce      json
// @Param        id   query  int  true  "opening id"
// @Success      200  {object}  DeleteOpeningResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /opening [delete]
func DeleteOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, "Missing required id param!", http.StatusBadRequest)
		return
	}

	var opening = schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, err.Error(), http.StatusNotFound)
		return
	}

	if err := db.Delete(&opening, id).Error; err != nil {
		sendError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Opening deleted with success!",
	})
}

// @Summary      List openings
// @Description  List openings
// @Tags         openings
// @Accept       json
// @Produce      json
// @Success      200  {object}  []schemas.OpeningResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /openings [get]
func ListOpenings(ctx *gin.Context) {
	var openings = []schemas.Opening{}
	if err := db.First(&openings).Error; err != nil {
		sendError(ctx, err.Error(), http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": openings,
	})
}
