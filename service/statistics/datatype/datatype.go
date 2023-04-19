// Data format conversion for map type
package datatype

import (
	"github.com/whatattitude/gDAG/lib/log/logger"
)

type StatusCount struct {
	AbnormalCount         float64
	NormalCount           float64
	AbnormalPercentGlobal float64
	AbnormalPercentLocal  float64
	AbnormalData          []map[string]string
	NormalData            []map[string]string
}

type DataCount struct {
	ValueMap  map[string]*StatusCount
	LabelName string
}

func (d *DataCount) AddStatusCount(labelValue string, statusType string, data map[string]string) {
	if d.ValueMap == nil {
		d.ValueMap = make(map[string]*StatusCount)
	}
	_, ok := d.ValueMap[labelValue]
	if !ok {
		d.ValueMap[labelValue] = &StatusCount{AbnormalCount: 0, NormalCount: 0,
			AbnormalPercentLocal: 0, AbnormalPercentGlobal: 0}
	}
	switch statusType {
	case "normal":
		d.ValueMap[labelValue].NormalCount++
		d.ValueMap[labelValue].NormalData = append(d.ValueMap[labelValue].NormalData, data)
	case "abnormal":
		d.ValueMap[labelValue].AbnormalCount++
		d.ValueMap[labelValue].AbnormalData = append(d.ValueMap[labelValue].AbnormalData, data)
	default:
		logger.Logger.Error("statusType is  Do not support" + statusType)
	}

}

func (d *DataCount) DataTypeConvert(showDataType string) (err error) {
	switch showDataType {
	case "percentages":
		logger.Logger.Sugar().Info(d.LabelName, " convert PercentData ")
		d.convertPercent()
	case "count":
	default:
		logger.Logger.Error("DataCount showDataType Do not support" + showDataType)
	}

	return
}

func (d *DataCount) SetLabelName(labelName string) {
	d.LabelName = labelName
	d.ValueMap = nil
}

func (d *DataCount) convertPercent() {
	var allCount float64
	for _, v := range d.ValueMap {
		allCount += v.AbnormalCount + v.NormalCount
		//logger.Logger.Sugar().Debugf("convertData is AbnormalCount %f, NormalCount %f ", v.AbnormalCount, v.NormalCount)
		if (v.AbnormalCount + v.NormalCount) != 0 {
			v.AbnormalPercentLocal = v.AbnormalCount / (v.AbnormalCount + v.NormalCount)
		}
	}

	for k := range d.ValueMap {
		if allCount != 0 {
			d.ValueMap[k].AbnormalPercentGlobal = d.ValueMap[k].AbnormalCount / allCount
			break
		}
		logger.Logger.Debug("allCount is 0 ")
		d.ValueMap[k].AbnormalPercentGlobal = 0
	}

}
