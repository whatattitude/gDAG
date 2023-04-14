package env

import (
	"reflect"
	"testing"
)

func TestGetEnvDefault(t *testing.T) {
	tests := []struct {
		name      string
		wantGoEnv GOENV
		wantErr   bool
	}{
		{
			name: "test",
			wantGoEnv: GOENV{},
			wantErr: false,

		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGoEnv, err := GetEnvDefault()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEnvDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGoEnv, tt.wantGoEnv) {
				t.Errorf("GetEnvDefault() = %v, want %v", gotGoEnv, tt.wantGoEnv)
			}
		})
	}
}
