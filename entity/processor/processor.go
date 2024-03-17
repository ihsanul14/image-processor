package processor

import (
	"context"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"os/exec"

	"gocv.io/x/gocv"
)

const filePath = "framework/output/output.jpg"
const filePathResize = "framework/output/output_resized.jpg"
const filePathCompress = "framework/output/output_compressed.jpg"

type Processor struct{}

type ResizeOpts struct {
	Width  float64
	Height float64
	PointX int
	PointY int
}

type CompressOpts struct {
	Quality string
	Path    string
}

type IMemory interface {
	Convert(context.Context, *multipart.FileHeader, string) error
	Resize(context.Context, *multipart.FileHeader, *ResizeOpts) error
	Compress(context.Context, *multipart.FileHeader, *CompressOpts) error
}

func NewMemory() IMemory {
	return &Processor{}
}

func (m *Processor) Convert(ctx context.Context, file *multipart.FileHeader, path string) error {
	if path == "" {
		path = filePath
	}
	fileContent, err := m.retrieveFile(file)
	if err != nil {
		return err
	}
	img, err := gocv.IMDecode(fileContent, gocv.IMReadColor)
	if err != nil {
		return fmt.Errorf("empty image")
	}
	defer img.Close()
	gocv.IMWrite(path, img)
	return nil
}

func (m *Processor) Resize(ctx context.Context, file *multipart.FileHeader, param *ResizeOpts) error {
	fileContent, err := m.retrieveFile(file)
	if err != nil {
		return err
	}
	img, err := gocv.IMDecode(fileContent, gocv.IMReadColor)
	if err != nil {
		return fmt.Errorf("empty image")
	}
	defer img.Close()

	resized := gocv.NewMat()
	gocv.Resize(img, &resized, image.Point{X: param.PointX, Y: param.PointY}, param.Width, param.Height, gocv.InterpolationLanczos4)
	gocv.IMWrite(filePathResize, resized)
	return nil
}

func (m *Processor) Compress(ctx context.Context, file *multipart.FileHeader, opts *CompressOpts) error {
	if opts.Path == "" {
		opts.Path = filePathCompress
	}
	err := m.Convert(ctx, file, filePathCompress)
	if err != nil {
		return err
	}
	cmd := exec.Command("ffmpeg", "-i", opts.Path, "-q:v", opts.Quality, "-y", opts.Path)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (m *Processor) retrieveFile(data *multipart.FileHeader) ([]byte, error) {
	file, err := data.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}
