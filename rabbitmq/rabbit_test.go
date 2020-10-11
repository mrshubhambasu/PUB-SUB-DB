package rabbitmq

import (
	"reflect"
	"testing"

	"github.com/streadway/amqp"
)

func TestNewConnection(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *AMQP
	}{
		// TBD: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConnection(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAMQP_CreateChannel(t *testing.T) {
	type fields struct {
		conn  *amqp.Connection
		Ch    *amqp.Channel
		Queue *amqp.Queue
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TBD: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AMQP{
				conn:  tt.fields.conn,
				Ch:    tt.fields.Ch,
				Queue: tt.fields.Queue,
			}
			c.CreateChannel()
		})
	}
}

func TestAMQP_DeclareQueue(t *testing.T) {
	type fields struct {
		conn  *amqp.Connection
		Ch    *amqp.Channel
		Queue *amqp.Queue
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TBD: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AMQP{
				conn:  tt.fields.conn,
				Ch:    tt.fields.Ch,
				Queue: tt.fields.Queue,
			}
			c.DeclareQueue(tt.args.name)
		})
	}
}

func TestAMQP_SendData(t *testing.T) {
	type fields struct {
		conn  *amqp.Connection
		Ch    *amqp.Channel
		Queue *amqp.Queue
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TBD: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AMQP{
				conn:  tt.fields.conn,
				Ch:    tt.fields.Ch,
				Queue: tt.fields.Queue,
			}
			c.SendData(tt.args.data)
		})
	}
}

func TestAMQP_RecieveData(t *testing.T) {
	type fields struct {
		conn  *amqp.Connection
		Ch    *amqp.Channel
		Queue *amqp.Queue
	}
	tests := []struct {
		name   string
		fields fields
		want   <-chan amqp.Delivery
	}{
		// TBD: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AMQP{
				conn:  tt.fields.conn,
				Ch:    tt.fields.Ch,
				Queue: tt.fields.Queue,
			}
			if got := c.RecieveData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AMQP.RecieveData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_failOnError(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TBD: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			failOnError(tt.args.err, tt.args.msg)
		})
	}
}
