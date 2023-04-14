package statistics

import (
	"reflect"
	"testing"

	"gDAG/service/metric"
)

func TestLabelAnomalyAnalysis(t *testing.T) {
	type args struct {
		showDataType  string
		analysisLabel Name
		data          metric.Metric
	}
	tests := []struct {
		name                 string
		args                 args
		wantLabelAnalysisMap map[Value]Count
		wantErr              bool
	}{
		{
			name:"test",
			args: args{
				showDataType: "percentages",
				analysisLabel: "idc",
				data: metric.Metric{
					{Idc: "a", Value: 1},
					{Idc: "b", Value: 2},
					{Idc: "c", Value: 3},
					{Idc: "b", Value: 2},
					{Idc: "c", Value: 3},
				},
			},
			wantLabelAnalysisMap: make(map[string]float64),
			wantErr: false,

		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		
		gotLabelAnalysisMap, err := LabelAnomalyAnalysis(tt.args.showDataType, tt.args.analysisLabel, &tt.args.data)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. LabelAnomalyAnalysis() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(gotLabelAnalysisMap, tt.wantLabelAnalysisMap) {
			t.Errorf("%q. LabelAnomalyAnalysis() = %v, want %v", tt.name, gotLabelAnalysisMap, tt.wantLabelAnalysisMap)
		}
	}
}
