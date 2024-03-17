package usecase

import (
	"context"
	"image-processor/entity/processor"
	"image-processor/framework/utils"
	"mime/multipart"
)

type IUsecase interface {
	Convert(context.Context, *multipart.FileHeader) error
	Resize(context.Context, *multipart.FileHeader, *ResizeOpts) error
	Compress(context.Context, *multipart.FileHeader, string) error
}

type Usecase struct {
	Entity processor.IMemory
}

type ResizeOpts struct {
	Width  string
	Height string
	PointX string
	PointY string
}

func NewUsecase(memory processor.IMemory) IUsecase {
	return &Usecase{
		Entity: memory,
	}
}

func (u *Usecase) Convert(ctx context.Context, file *multipart.FileHeader) error {
	err := u.Entity.Convert(ctx, file, "")
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) Resize(ctx context.Context, file *multipart.FileHeader, param *ResizeOpts) error {
	width, err := utils.ParseStringToInt(param.Width)
	if err != nil {
		return err
	}
	height, err := utils.ParseStringToInt(param.Height)
	if err != nil {
		return err
	}
	pointX, err := utils.ParseStringToInt(param.PointX)
	if err != nil {
		return err
	}
	pointY, err := utils.ParseStringToInt(param.PointY)
	if err != nil {
		return err
	}
	params := &processor.ResizeOpts{
		Width:  width,
		Height: height,
		PointX: int(pointX),
		PointY: int(pointY),
	}
	err = u.Entity.Resize(ctx, file, params)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) Compress(ctx context.Context, file *multipart.FileHeader, q string) error {
	opts := &processor.CompressOpts{
		Quality: q,
	}
	err := u.Entity.Compress(ctx, file, opts)
	if err != nil {
		return err
	}
	return nil
}
