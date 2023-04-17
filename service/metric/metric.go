// monitor metric
package metric

import (
	"errors"
	"reflect"
	"strconv"
)

type MetricItem struct {
	Idc    string
	Labels map[string]string
	Value  float64
	ThresholdMax float64
	ThresholdMin float64
	TimeStamp string
}

type Metric []MetricItem

func (mSlice *Metric) Len() int {
	return len(*mSlice)
}

func (mSlice *Metric)	GetDataInfo(index int) (datainfo map[string]string, err error){
	m, err := mSlice.SafetyGetterItem(index)
	if err != nil {
		return datainfo, err
	}
	datainfo = make(map[string]string)
	for k, v := range m.Labels{
		datainfo[k] = v
	}
	datainfo["Idc"] = m.Idc
	datainfo["TimeStamp"] = m.TimeStamp
	datainfo["ThresholdMax"] = strconv.FormatFloat(m.ThresholdMax, 'f', 1, 64) 
	datainfo["ThresholdMin"] = strconv.FormatFloat(m.ThresholdMin, 'f', 1, 64) 
	datainfo["Value"] = strconv.FormatFloat(m.Value, 'f', 3, 64) 
	return datainfo, err
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
func (mSlice *Metric) ValueStatusChecker(index int) (status string, err error) {
	m, err := mSlice.SafetyGetterItem(index)
	if err != nil {
		return "abnormal", err
	}
	if m.Value > m.ThresholdMin && m.Value < m.ThresholdMax {
		return "normal", err
	}
	return "abnormal", err
}

func (mSlice *Metric) SafetyGetterItem(index int) (m MetricItem, err error) {
	if mSlice != nil && len(*mSlice) > index {
		return (*mSlice)[index], nil
	}
	return m, errors.New("index out of range *Metric")
}

func (mSlice *Metric)DeepCopy()( m2 *Metric){
	*m2 = make([]MetricItem, mSlice.Len())
	for i, v := range *mSlice {
		(*m2)[i] = *(v.DeepCopy())
	}
		
	return
}

func (mItem *MetricItem)DeepCopy()( m2 *MetricItem){
	m2 = mItem
	m2.Labels = make(map[string]string, len(mItem.Labels))
	for key, value := range mItem.Labels {
    	m2.Labels[key] = value
	}
	return
}
