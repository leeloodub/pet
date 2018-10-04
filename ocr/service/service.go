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

// func (s *OCRService) Init(storage Storage) {
// 	s.client = gosseract.NewClient()
// 	s.storage = storage
// }
//
// func (s *OCRService) Close() {
// 	s.client.Close()
// }

func (s *OCRService) DoOCR(ctx context.Context, req *pb.OCRRequest) (*pb.OCRResponse, error) {
	path := req.RawImgPath
	if path == "" {
		return &pb.OCRResponse{Status: pb.OCRResponse_UNDEFINED}, nil
	}

	s.client.SetImage(path)
	t, err := s.client.Text()
	fmt.Println(t, err)

	if err != nil {
		return &pb.OCRResponse{Status: pb.OCRResponse_FAILED}, nil
	}

	id, _ := s.storage.InsertParsed(t)
	fmt.Println(id)
	return &pb.OCRResponse{Status: pb.OCRResponse_SUCCESS}, nil
}

type Storage interface {
	InsertParsed(img string) (string, error)
}
