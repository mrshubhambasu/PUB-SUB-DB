package datastore

import (
	"assignment/model"
	"testing"

	"gorm.io/gorm"
)

func TestPush(t *testing.T) {
	type args struct {
		dataModel model.Data
		db        *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Push(tt.args.dataModel, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInit(t *testing.T) {
	_, err := Init("root@tcp(127.0.0.1:3306)/hotel?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Error("Expected error is nil")
	}
}
