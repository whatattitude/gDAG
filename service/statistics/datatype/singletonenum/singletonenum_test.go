package singletonenum

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/younamebert/enum"
)

func TestGetLazySingletonInstance(t *testing.T) {
	tests := []struct {
		name string
		want *enum.Enum
	}{
		{
			name: "Enum",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := OneDataEnum.GetLazySingletonInstance(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetLazySingletonInstance() = %v, want %v", tt.name, got, tt.want)
		}
		fmt.Println(OneDataEnum.DataTypeEnum)
		
	}
}
