package service

import (
	"context"
	"fmt"

	"github.com/otiai10/gosseract"
	pb "github.com/pet/ocr/protos"
)

type OCRService struct {
	client  *gosseract.Client
	storage Storage
}

func (s *OCRService) Init(storage Storage) {
	s.client = gosseract.NewClient()
	s.storage = storage
}

func (s *OCRService) Close() {
	s.client.Close()
}

func (s *OCRService) DoOCR(ctx context.Context, req *pb.OCRRequest, resp *pb.OCRResponse) {
	path := req.RawImgPath
	if path == "" {
		resp.Status = pb.OCRResponse_UNDEFINED
		return
	}

	s.client.SetImage(path)
	t, err := s.client.Text()
	fmt.Println(t, err)

	if err != nil {
		resp.Status = pb.OCRResponse_FAILED
	}

	id, _ := s.storage.InsertParsed(t)
	fmt.Println(id)
	resp.Status = pb.OCRResponse_SUCCESS
}

type Storage interface {
	InsertParsed(img string) (string, error)
}
