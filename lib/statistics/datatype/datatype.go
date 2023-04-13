package datatype

import (
	"errors"
	"sync"

	"github.com/younamebert/enum"
)

var (
	DataTypeEnum *DataEnum
	once         = &sync.Once{}
)

type DataEnum struct{
	Enum *enum.Enum
}

func GetLazySingletonInstance(args ...string) *DataEnum {
	if DataTypeEnum == nil {
		once.Do(func() {
			resp := enum.NewEnum(args... )
			DataTypeEnum  = &DataEnum{
				Enum:  &resp,
			}
		})
	}
	return DataTypeEnum
}

func MapValueConvertor(dataType string, dataMap map[string]float64)(err error){
	hasDataType := DataTypeEnum.Enum.IsEnum(dataType)
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
