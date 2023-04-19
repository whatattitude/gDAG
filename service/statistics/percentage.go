// Analysis library for composite structure data
package statistics

import (
	"github.com/whatattitude/gDAG/lib/log/logger"
)

var statisticsLogger = logger.Logger

type Name = string

type Value = string

type LabelSelector interface {
	Len() (len int)
	GetAnalysisLabelValue(index int, name string) (value string, err error)
	ValueStatusChecker(index int) (status string, err error)
	GetDataInfo(index int) (datainfo map[string]string, err error)
}

type DataInfo interface {
	GetDataInfo() map[string]string
}

type ValueStatusCount interface {
	SetLabelName(labelName string)
	AddStatusCount(labelValue string, statusType string, dataInfo map[string]string)
	DataTypeConvert(showDataType string) (err error)
}

func LabelAnomalyAnalysis(showDataType string, analysisLabel Name, data LabelSelector, labelAnalysis ValueStatusCount) (err error) {
	labelAnalysis.SetLabelName(analysisLabel)
	analysisCount := data.Len()

	//根据analysisLabel统计各取值数量
	for i := 0; i < analysisCount; i++ {
		value, err := data.GetAnalysisLabelValue(i, analysisLabel)
		if err != nil {
			return err
		}
		status, err := data.ValueStatusChecker(i)
		if err != nil {
			return err
		}
		dataInfo, err := data.GetDataInfo(i)
		if err != nil {
			return err
		}
		statisticsLogger.Sugar().Debugf("dataStatus=%s, data=%v ", status, dataInfo)
		labelAnalysisSliceLabelAdd(labelAnalysis, value, status, dataInfo)
	}

	labelAnalysis.DataTypeConvert(showDataType)

	return err
}

func labelAnalysisSliceLabelAdd(labelAnalysis ValueStatusCount, labelValue Value,
	status string, dataInfo map[string]string) {
	labelAnalysis.AddStatusCount(labelValue, status, dataInfo)

}
