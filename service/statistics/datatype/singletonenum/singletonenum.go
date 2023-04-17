// Data format conversion for map type
package singletonenum

import (
	"sync"

	"github.com/younamebert/enum"
)
type DataEnum struct{
	Enum *enum.Enum
}

var OneDataEnum = OneDataTypeEnum{
	Once: &sync.Once{},
	DataTypeEnum: nil,
}

type OneDataTypeEnum struct{
	DataTypeEnum *DataEnum
	Once         *sync.Once
}


func (OneDataEnum *OneDataTypeEnum) GetLazySingletonInstance(  args ...string) *OneDataTypeEnum {
	if OneDataEnum.DataTypeEnum == nil {
		OneDataEnum.Once.Do(func() {
			resp := enum.NewEnum(args... )
			OneDataEnum.DataTypeEnum  = &DataEnum{
				Enum:  &resp,
			}
		})
	}
	return OneDataEnum
}