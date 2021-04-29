package util

import (
	"reflect"
	"testing"

	"github.com/nats-io/stan.go"
)

func TestValidateNatsTopic(t *testing.T) {
	type args struct {
		tName string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateNatsTopic(tt.args.tName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateNatsTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateNatsTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConnection(t *testing.T) {
	type args struct {
		host    string
		cluster string
		client  string
		port    int
	}
	tests := []struct {
		name    string
		args    args
		want    stan.Conn
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				host:    "localhost",    // "localhost"
				cluster: "test-cluster", // nats-streaming-server clusterID
				client:  "client1",      //
				port:    4222,           // 4222
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConnection(tt.args.host, tt.args.cluster, tt.args.client, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNSConnection(t *testing.T) {
	type args struct {
		host    string
		cluster string
		client  string
		port    int
	}
	tests := []struct {
		name    string
		args    args
		want    stan.Conn
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				host:    "",        // empty: "localhost"
				cluster: "",        // empty: nats-streaming-server clusterID
				client:  "client1", //
				port:    0,         // 0 : 4222
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNSConnection(tt.args.host, tt.args.cluster, tt.args.client, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNSConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNSConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
