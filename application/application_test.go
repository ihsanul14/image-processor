package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	mocks "image-processor/framework/mocks/usecase"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	filePath      = "../framework/output/output.jpg"
	fileFormKey   = "file"
	fileFormValue = "file.png"
	errorMessage  = "error"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestNewApplication(t *testing.T) {
	NewApplication(nil)
}

func TestConvert(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUsecase := mocks.NewMockIUsecase(mockController)
	app := &Application{
		Usecase: mockUsecase,
	}
	assert.NotNil(t, app)

	actionHttpMethod := "POST"
	actionHttpUrl := "/api/image/convert"

	router := initRouter()
	router.POST(actionHttpUrl, app.Convert)

	t.Run("200", func(t *testing.T) {
		file, err := os.Open(filePath)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile(fileFormKey, fileFormValue)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		_, err = io.Copy(part, file)
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)
		mockUsecase.EXPECT().Convert(gomock.Any(), gomock.Any()).Return(nil)
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, body)
		res.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 200, response.Code)
	})

	t.Run("400", func(t *testing.T) {
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 400, response.Code)
	})

	t.Run("500", func(t *testing.T) {
		file, err := os.Open(filePath)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile(fileFormKey, fileFormValue)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		_, err = io.Copy(part, file)
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)
		mockUsecase.EXPECT().Convert(gomock.Any(), gomock.Any()).Return(fmt.Errorf(errorMessage))
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, body)
		res.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 500, response.Code)
	})
}

func TestResize(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUsecase := mocks.NewMockIUsecase(mockController)
	app := &Application{
		Usecase: mockUsecase,
	}
	assert.NotNil(t, app)

	actionHttpMethod := "POST"
	actionHttpUrl := "/api/image/resize"

	router := initRouter()
	router.POST(actionHttpUrl, app.Resize)

	t.Run("200", func(t *testing.T) {
		file, err := os.Open(filePath)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile(fileFormKey, fileFormValue)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		_, err = io.Copy(part, file)
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)
		mockUsecase.EXPECT().Resize(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, body)
		res.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 200, response.Code)
	})

	t.Run("400", func(t *testing.T) {
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 400, response.Code)
	})

	t.Run("500", func(t *testing.T) {
		file, err := os.Open(filePath)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile(fileFormKey, fileFormValue)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		_, err = io.Copy(part, file)
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)
		mockUsecase.EXPECT().Resize(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf(errorMessage))
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, body)
		res.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 500, response.Code)
	})
}

func TestCompress(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUsecase := mocks.NewMockIUsecase(mockController)
	app := &Application{
		Usecase: mockUsecase,
	}
	assert.NotNil(t, app)

	actionHttpMethod := "POST"
	actionHttpUrl := "/api/image/compress"

	router := initRouter()
	router.POST(actionHttpUrl, app.Compress)

	t.Run("200", func(t *testing.T) {
		file, err := os.Open(filePath)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile(fileFormKey, fileFormValue)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		_, err = io.Copy(part, file)
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)
		mockUsecase.EXPECT().Compress(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, body)
		res.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 200, response.Code)
	})

	t.Run("400", func(t *testing.T) {
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 400, response.Code)
	})

	t.Run("500", func(t *testing.T) {
		file, err := os.Open(filePath)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile(fileFormKey, fileFormValue)
		assert.Nil(t, err)
		assert.NotNil(t, file)
		_, err = io.Copy(part, file)
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)
		mockUsecase.EXPECT().Compress(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf(errorMessage))
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, body)
		res.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 500, response.Code)
	})
}
