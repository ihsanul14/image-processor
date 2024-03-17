package processor

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const pathTest = "../../framework/output/output.jpg"

func initFile(t *testing.T) *multipart.FileHeader {
	file, err := os.Open(pathTest)
	assert.Nil(t, err)
	defer file.Close()

	var buff bytes.Buffer
	buffWriter := io.Writer(&buff)

	formWriter := multipart.NewWriter(buffWriter)
	formPart, err := formWriter.CreateFormFile("file", "output.jpg")
	assert.Nil(t, err)

	_, err = io.Copy(formPart, file)
	assert.Nil(t, err)

	formWriter.Close()

	buffReader := bytes.NewReader(buff.Bytes())
	formReader := multipart.NewReader(buffReader, formWriter.Boundary())

	multipartForm, err := formReader.ReadForm(1 << 20)
	assert.Nil(t, err)

	files, exists := multipartForm.File["file"]
	assert.NotNil(t, exists)

	return files[0]
}

func TestConvert(t *testing.T) {
	memory := NewMemory()
	ctx := context.Background()
	file := initFile(t)
	err := memory.Convert(ctx, file, pathTest)
	assert.Nil(t, err)
}

func TestResize(t *testing.T) {
	memory := NewMemory()
	ctx := context.Background()
	file := initFile(t)
	opts := &ResizeOpts{
		Width:  1,
		Height: 1,
	}
	err := memory.Resize(ctx, file, opts)
	assert.Nil(t, err)
}

func TestCompress(t *testing.T) {
	memory := NewMemory()
	ctx := context.Background()
	file := initFile(t)
	opts := &CompressOpts{
		Quality: "4",
		Path:    pathTest,
	}
	err := memory.Compress(ctx, file, opts)
	assert.Nil(t, err)
}
