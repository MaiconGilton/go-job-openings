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

func sendError(ctx *gin.Context, msg string, code int) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func GetOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get opening",
	})
}

func CreateOpening(ctx *gin.Context) {
	requestBody := struct {
		Role     string `json:"role" validate:"required"`
		Company  string `json:"company" validate:"required"`
		Location string `json:"location" validate:"required"`
		Salary   int32  `json:"salary" validate:"required,gt=0"`
		Link     string `json:"link" validate:"required,url"`
		Remote   *bool  `json:"remote"`
	}{}

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
		"success": true,
		"message": "Opening created with success!",
		"data":    opening,
	})
}

func UpdateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get opening",
	})
}

func DeleteOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get opening",
	})
}

func ListOpenings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get opening",
	})
}
