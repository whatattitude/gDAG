// monitor metric
package metric

import (
	"errors"
	"strconv"
)


type MetricItem struct{
	Idc string
	Labels map[string]string
	Value float64
}


type Metric []MetricItem

func (m *Metric)Len()int{
	return len(*m)
}


func (mList *Metric)GetAnalysisLabelValue(index int, name string) (value string, err error){
	m, err :=  mList.SafetyGetterItem(index)
	if err != nil{
		return "", err
	}
	switch name {
	case "idc":
		return m.Idc, nil
	case "value":
		return strconv.FormatFloat(m.Value, 'f', -1, 32), nil
	default: 
		if _, ok := m.Labels[name]; !ok {
   	 		// 不存在
			return "", errors.New(name + "is not in this Metric")
		}
		return m.Labels[name], nil
	}
}

//true代表正常，在阈值范围内，false代表异常不在阈值范围内
func (mList *Metric)ValueStatusChecker(index int, thresholdMax float64, thresholdMin float64 ) (status bool, err error){
	m, err :=  mList.SafetyGetterItem(index)
	if err != nil{
		return false, err
	}
	if m.Value > thresholdMin && m.Value < thresholdMax{
		return true, err
	}
	return false, err
}


func (mList *Metric)SafetyGetterItem(index int)(m MetricItem, err error){
	if mList != nil && len(*mList) > index{
		return (*mList)[index], nil
	}
	return m, errors.New("index out of range *Metric")
}