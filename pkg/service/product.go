package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"doan/pkg/valid"
	"github.com/praslar/lib/common"
)

type ProductService struct {
	repo repo.PGInterface
}

type ProductInterface interface {
	Create(ctx context.Context, ob model.ProductRequest) (rs *model.ProductResponse, err error)
	Update(ctx context.Context, ob model.ProductRequest) (rs *model.Product, err error)
	Delete(ctx context.Context, id string) (err error)
	GetOne(ctx context.Context, id string) (rs *model.Product, err error)
	GetList(ctx context.Context, req model.ProductParams) (rs *model.ListProductResponse, err error)
}

func NewProductService(repo repo.PGInterface) ProductInterface {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(ctx context.Context, req model.ProductRequest) (rs *model.ProductResponse, err error) {

	respone := model.ProductResponse{}
	// create product
	ob := &model.Product{}
	common.Sync(req, ob)
	if err := s.repo.CreateProduct(ctx, ob); err != nil {
		return nil, err
	}
	respone.Product = ob

	// create multi sku
	for _, v := range req.Sku {
		sku := &model.Sku{}
		common.Sync(v, sku)
		sku.ProductId = ob.ID
		if err := s.repo.CreateSku(ctx, sku); err != nil {
			return nil, err
		}
		respone.Sku = append(respone.Sku, sku)
	}

	return &respone, nil
}

func (s *ProductService) Update(ctx context.Context, req model.ProductRequest) (rs *model.Product, err error) {

	ob, err := s.repo.GetOneProduct(ctx, valid.String(req.ID))
	if err != nil {
		return nil, err
	}

	common.Sync(req, ob)

	if err := s.repo.UpdateProduct(ctx, ob); err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ProductService) Delete(ctx context.Context, id string) (err error) {
	return s.repo.DeleteProduct(ctx, id)
}

func (s *ProductService) GetOne(ctx context.Context, id string) (rs *model.Product, err error) {

	ob, err := s.repo.GetOneProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return ob, nil
}

func (s *ProductService) GetList(ctx context.Context, req model.ProductParams) (rs *model.ListProductResponse, err error) {

	ob, err := s.repo.GetListProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return ob, nil
}
