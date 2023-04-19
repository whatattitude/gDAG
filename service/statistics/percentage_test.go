package statistics

import (
	"encoding/json"
	"testing"

	"github.com/whatattitude/gDAG/service/metric"
	"github.com/whatattitude/gDAG/service/statistics/datatype"
)

func TestLabelAnomalyAnalysis(t *testing.T) {
	type args struct {
		showDataType  string
		analysisLabel Name
		data          LabelSelector
		labelAnalysis datatype.DataCount
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				showDataType:  "percentages",
				analysisLabel: "Idc",
				data: &metric.Metric{
					{Idc: "a", Value: 1, ThresholdMax: 1, ThresholdMin: 0, TimeStamp: "1"},
					{Idc: "b", Value: 0.33, ThresholdMax: 1, ThresholdMin: 0, TimeStamp: "1"},
					{Idc: "c", Value: 3, ThresholdMax: 1, ThresholdMin: 0, TimeStamp: "1"},
					{Idc: "b", Value: 2, ThresholdMax: 1, ThresholdMin: 0, TimeStamp: "1"},
					{Idc: "c", Value: 3, ThresholdMax: 1, ThresholdMin: 0, TimeStamp: "1"},
				},
				labelAnalysis: datatype.DataCount{},
			},

			wantErr: false,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LabelAnomalyAnalysis(tt.args.showDataType, tt.args.analysisLabel, tt.args.data, &tt.args.labelAnalysis); (err != nil) != tt.wantErr {
				t.Errorf("LabelAnomalyAnalysis() error = %v, wantErr %v", err, tt.wantErr)
			}
			dataType, _ := json.Marshal(tt.args.labelAnalysis.ValueMap)
			dataString := string(dataType)
			t.Logf("labelAnalysis is: %s", dataString)
			if err := LabelAnomalyAnalysis(tt.args.showDataType, "TimeStamp", tt.args.data, &tt.args.labelAnalysis); (err != nil) != tt.wantErr {
				t.Errorf("LabelAnomalyAnalysis() error = %v, wantErr %v", err, tt.wantErr)
			}
			dataType, _ = json.Marshal(tt.args.labelAnalysis)
			dataString = string(dataType)
			t.Logf("labelAnalysis is: %s", dataString)
		})
	}
}
