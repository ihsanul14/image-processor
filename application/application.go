package application

import (
	"image-processor/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type IApplication interface {
	Convert(*gin.Context)
	Resize(*gin.Context)
	Compress(*gin.Context)
}

type Application struct {
	Usecase usecase.IUsecase
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewApplication(uc usecase.IUsecase) IApplication {
	return &Application{
		Usecase: uc,
	}
}

// @Summary Upload a file
// @Description Uploads a file using form-data
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} Response
// @Router /api/image/convert [post]
func (a *Application) Convert(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if !strings.HasSuffix(strings.ToLower(file.Filename), ".png") {
		c.JSON(http.StatusBadRequest, &Response{
			Code:    400,
			Message: "input file should be in .png",
		})
		return
	}

	err = a.Usecase.Convert(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Code:    200,
		Message: "success",
	})
}

func (a *Application) Resize(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	param := &usecase.ResizeOpts{
		Width:  c.PostForm("width"),
		Height: c.PostForm("height"),
		PointX: c.PostForm("pointX"),
		PointY: c.PostForm("pointY"),
	}

	err = a.Usecase.Resize(c, file, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Code:    200,
		Message: "success",
	})
}

func (a *Application) Compress(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	q := c.PostForm("quality")
	err = a.Usecase.Compress(c, file, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Code:    200,
		Message: "success",
	})
}
