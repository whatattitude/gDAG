package datatype

import (
	"errors"
	"sync"

	"github.com/younamebert/enum"
)



var OneDataEnum = OneDataTypeEnum{
	Once: &sync.Once{},
	DataTypeEnum: nil,
}

type OneDataTypeEnum struct{
	DataTypeEnum *DataEnum
	Once         *sync.Once
}

type DataEnum struct{
	Enum *enum.Enum
}

func GetLazySingletonInstance( OneDataEnum *OneDataTypeEnum, args ...string) *OneDataTypeEnum {
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

func MapValueConvertor( OneDataEnum *OneDataTypeEnum, dataType string, dataMap map[string]float64)(err error){
	hasDataType := OneDataEnum.DataTypeEnum.Enum.IsEnum(dataType)
	if !hasDataType{
		return errors.New(dataType + "is not in datatype.DataTypeEnum")
	}
	switch dataType{
		case "percentages":
			percentagesMapConvertor( dataMap)
		case "count":
			return
	}
	return
	
}

func percentagesMapConvertor(dataMap map[string]float64){
	valueSum := 0.0
	for _, v := range dataMap {
		valueSum += v
	}
	for k := range dataMap {
		dataMap[k] = dataMap[k]/ valueSum
	}
}
