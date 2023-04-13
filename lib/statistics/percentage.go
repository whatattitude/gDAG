package statistics

import (
	"fmt"
	"gDAG/lib/log/logger"
	"gDAG/lib/statistics/datatype"
)

var statisticsLogger = logger.InitLogger("./log/statistics/")

type Name = string
type Value = string
type Count = float64

type LabelSelector interface {
	Len()int
	GetAnalysisLabelValue(index int , name  string) (value string, err error)
	ValueStatusChecker(index int ,thresholdMax float64, thresholdMin float64 )(status bool, err error)
}


func LabelAnomalyAnalysis(showDataType string, analysisLabel Name, data LabelSelector) ( labelAnalysisMap map[Value]Count, err error) {
	labelAnalysisMap = make(map[string]float64)
	analysisCount := data.Len()

	//根据analysisLabel统计各取值数量
	for i := 0; i < analysisCount; i++ {
		value, err := data.GetAnalysisLabelValue(i, analysisLabel)
		if err != nil {
			return labelAnalysisMap,  err
		}
		statisticsLogger.Sugar().Debugln(value, " ",  analysisLabel, data)
		labelAnalysisSliceLabelAdd(labelAnalysisMap, value)
	}

	//根据showDataType转换展示数据格式
	datatype.GetLazySingletonInstance("count","percentages")
	err = datatype.MapValueConvertor(showDataType, labelAnalysisMap) 
	if err != nil{
		return labelAnalysisMap,  err
	}
	return labelAnalysisMap,  err
}

func labelAnalysisSliceLabelAdd(labelAnalysisMap map[Value]Count, labelValue Value) {
	fmt.Println(labelValue)
	_, ok := labelAnalysisMap[labelValue]
	if !ok {
		labelAnalysisMap[labelValue] = 0
	}
	labelAnalysisMap[labelValue]++
}
