// monitor metric
package metric

import (
	"errors"
	"reflect"
)

type MetricItem struct {
	Idc    string
	Labels map[string]string
	Value  float64
}

type Metric []MetricItem

func (mSlice *Metric) Len() int {
	return len(*mSlice)
}

// 从指定的结构体获取类型为string的成员变量的值
func (mSlice *Metric) GetAnalysisLabelValue(index int, name string) (value string, err error) {
	m, err := mSlice.SafetyGetterItem(index)
	if err != nil {
		return "", err
	}

	//从结构体中的string类型的成员变量获取结果
	refType := reflect.TypeOf(m)
	refValue := reflect.ValueOf(m)
	nameType, status := refType.FieldByName(name)
	if status && nameType.Type.Kind() == reflect.String{
		index := nameType.Index
		return refValue.FieldByIndex(index).String(), nil
	}

	//从结构体中的map类型的变量获取结果
	if _, ok := m.Labels[name]; !ok {
		// 不存在
		return "", errors.New(name + "is not in this struct or its kind is not string")
	}
	return m.Labels[name], nil

}

// true代表正常，在阈值范围内，false代表异常不在阈值范围内
func (mSlice *Metric) ValueStatusChecker(index int, thresholdMax float64, thresholdMin float64) (status bool, err error) {
	m, err := mSlice.SafetyGetterItem(index)
	if err != nil {
		return false, err
	}
	if m.Value > thresholdMin && m.Value < thresholdMax {
		return true, err
	}
	return false, err
}

func (mSlice *Metric) SafetyGetterItem(index int) (m MetricItem, err error) {
	if mSlice != nil && len(*mSlice) > index {
		return (*mSlice)[index], nil
	}
	return m, errors.New("index out of range *Metric")
}
