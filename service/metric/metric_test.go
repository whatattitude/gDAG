// monitor metric
package metric

import "testing"

func TestMetric_GetAnalysisLabelValue(t *testing.T) {
	type args struct {
		index int
		name  string
	}
	tests := []struct {
		name      string
		mSlice    Metric
		args      args
		wantValue string
		wantErr   bool
	}{
		{
			name: "test",
			mSlice: []MetricItem{
				{
					Idc: "a",
					Labels: map[string]string{
						"a": "a",
						"b": "b",
						"c": "c",
					},
					Value: 1,
				},
				{
					Idc: "b",
					Labels: map[string]string{
						"a": "a",
						"b": "b",
						"c": "c",
					},
					Value: 1,
				},
			},
			args: args{
				index: 1,
				name:  "Labels",
			},
			wantValue: "b",
			wantErr:   false,
		},
	}
	// TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := tt.mSlice.GetAnalysisLabelValue(tt.args.index, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Metric.GetAnalysisLabelValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("Metric.GetAnalysisLabelValue() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
