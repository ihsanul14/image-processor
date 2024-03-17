package usecase

import (
	"context"
	"fmt"
	mocks "image-processor/framework/mocks/entity"
	"mime/multipart"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	ctx := context.Background()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockMemory := mocks.NewMockIMemory(mockController)
	data := &multipart.FileHeader{}
	u := NewUsecase(mockMemory)

	t.Run("Success", func(t *testing.T) {
		mockMemory.EXPECT().Convert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		err := u.Convert(ctx, data)
		assert.Nil(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		mockMemory.EXPECT().Convert(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
		err := u.Convert(ctx, data)
		assert.NotNil(t, err)
	})
}

func TestResize(t *testing.T) {
	ctx := context.Background()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockMemory := mocks.NewMockIMemory(mockController)
	data := &multipart.FileHeader{}
	u := NewUsecase(mockMemory)
	opts := &ResizeOpts{
		Width:  "1",
		Height: "1",
		PointX: "500",
		PointY: "80",
	}

	t.Run("Success", func(t *testing.T) {
		mockMemory.EXPECT().Resize(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		err := u.Resize(ctx, data, opts)
		assert.Nil(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		mockMemory.EXPECT().Resize(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
		err := u.Resize(ctx, data, opts)
		assert.NotNil(t, err)
	})

	t.Run("Invalid Height Params", func(t *testing.T) {
		opts.Height = "a"
		err := u.Resize(ctx, data, opts)
		assert.NotNil(t, err)
	})

	t.Run("Invalid Width Params", func(t *testing.T) {
		opts.Width = "a"
		err := u.Resize(ctx, data, opts)
		assert.NotNil(t, err)
	})
	t.Run("Invalid PointX Params", func(t *testing.T) {
		opts.PointX = "a"
		opts.Height = "1"
		opts.Width = "1"
		err := u.Resize(ctx, data, opts)
		assert.NotNil(t, err)
	})

	t.Run("Invalid PointY Params", func(t *testing.T) {
		opts.PointX = "50"
		opts.PointY = "a"
		err := u.Resize(ctx, data, opts)
		assert.NotNil(t, err)
	})
}

func TestCompress(t *testing.T) {
	ctx := context.Background()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockMemory := mocks.NewMockIMemory(mockController)
	data := &multipart.FileHeader{}
	u := NewUsecase(mockMemory)

	t.Run("Success", func(t *testing.T) {
		mockMemory.EXPECT().Compress(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		err := u.Compress(ctx, data, "")
		assert.Nil(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		mockMemory.EXPECT().Compress(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
		err := u.Compress(ctx, data, "")
		assert.NotNil(t, err)
	})
}
