package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type ProductUsecase interface {
	Create(ctx context.Context, name string) (string, error)
	Read(ctx context.Context) (map[string]string, error)
	Update(ctx context.Context, id string, name string) (string, string, error)
	Delete(ctx context.Context, id string) error
}

type productUsecase struct {
	productData map[string]string
}

func NewProductUsecase(productData map[string]string) ProductUsecase {
	return &productUsecase{
		productData: productData,
	}
}

func (uc *productUsecase) Create(ctx context.Context, name string) (string, error) {
	id := uuid.NewString()
	uc.productData[id] = name
	return id, nil
}

func (uc *productUsecase) Read(ctx context.Context) (map[string]string, error) {
	return uc.productData, nil
}

func (uc *productUsecase) Update(ctx context.Context, id string, name string) (string, string, error) {
	_, ok := uc.productData[id]
	if !ok {
		return "", "", errors.New("not found")
	}

	uc.productData[id] = name
	return id, name, nil
}

func (uc *productUsecase) Delete(ctx context.Context, id string) error {
	_, ok := uc.productData[id]
	if !ok {
		return errors.New("not found")
	}

	delete(uc.productData, id)
	return nil
}
