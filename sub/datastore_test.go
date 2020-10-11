package main

import (
	"testing"

	"gorm.io/gorm"
)

func Test_byteToStruct(t *testing.T) {
	// Test for simple argument & no variadic arguments
	data := Data{}
	resp, err := byteToStruct([]byte("some value"), data)

	if err != nil {
		t.Log("Recieved as expected: ", resp, err)
	} else {
		t.Error("Received unexpected: ", resp, err.Error())
	}
}

func Test_pushToDatabase(t *testing.T) {
	type args struct {
		bytes []byte
		db    *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TBD: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pushToDatabase(tt.args.bytes, tt.args.db)
		})
	}
}
