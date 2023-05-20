package repo

import (
	"context"
	"doan/conf"
	"testing"
)

func TestRepoPG_GetOneUserByEmail(t *testing.T) {
	config := conf.Config{
		Driver: "mysql",
		Host:   "localhost",
		Port:   "3306",
		User:   "doan",
		Pass:   "doan",
		Name:   "doan",
		Schema: "public",
	}
	db, err := conf.Open(&config)
	if err != nil {
		t.Error(err)
	}

	repo := NewRepo(db)
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetOneUserByEmail",
			args: args{
				ctx:   context.Background(),
				email: "notfound@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "GetOneUserByEmail",
			args: args{
				ctx:   context.Background(),
				email: "duchieu.ctk41@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.GetOneUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOneUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
