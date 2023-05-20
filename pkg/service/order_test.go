package service

import (
	"context"
	"doan/pkg/model"
	"doan/pkg/repo"
	"reflect"
	"testing"
)

func TestOrderService_Create(t *testing.T) {

	type fields struct {
		repo repo.PGInterface
	}
	type args struct {
		ctx context.Context
		req model.OrderRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRs  *model.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &OrderService{
				repo: tt.fields.repo,
			}
			gotRs, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRs, tt.wantRs) {
				t.Errorf("Create() gotRs = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}
