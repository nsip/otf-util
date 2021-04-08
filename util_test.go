package util

import (
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/nats-io/stan.go"
)

func Test_newNetClient(t *testing.T) {
	tests := []struct {
		name string
		want *http.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNetClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNetClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateName(t *testing.T) {
	type args struct {
		defaultName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				defaultName: "reader",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateName(tt.args.defaultName); got != tt.want {
				t.Errorf("GenerateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateID(); got != tt.want {
				t.Errorf("GenerateID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeTrack(t *testing.T) {
	type args struct {
		start time.Time
		name  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TimeTrack(tt.args.start, tt.args.name)
		})
	}
}

func TestAvailablePort(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			want: 12345,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AvailablePort()
			if (err != nil) != tt.wantErr {
				t.Errorf("AvailablePort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AvailablePort() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestFetch(t *testing.T) {
	type args struct {
		method string
		url    string
		header map[string]string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fetch(tt.args.method, tt.args.url, tt.args.header, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}
