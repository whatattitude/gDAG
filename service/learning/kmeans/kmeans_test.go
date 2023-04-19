package kmeans

import "testing"

func TestKmeansPlusPlus_OneDimensionalKmeansPlusPlus(t *testing.T) {
	type fields struct {
		CenterClusters []Cluster
		MaxIterations  int
		IsConvergence  bool
		Runs           int
		CenterCount    int
	}
	type args struct {
		originalData OneDimensionalData
		centerCount  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				MaxIterations: 10,
				CenterCount:   2,
			},
			args: args{
				originalData: &OneDimensionalDataInfo{
					{PointIndex: 1, Value: 1},
					{PointIndex: 2, Value: 1},
					{PointIndex: 3, Value: 1},
					{PointIndex: 4, Value: 1},
					{PointIndex: 5, Value: 1},
					{PointIndex: 6, Value: 1},
					{PointIndex: 7, Value: 1},
					{PointIndex: 8, Value: 1},
					{PointIndex: 9, Value: 1},
					{PointIndex: 10, Value: 1},
				},
				centerCount: 2,
			},
			wantErr: false,
		},
		{
			name: "test",
			fields: fields{
				MaxIterations: 10,
				CenterCount:   2,
			},
			args: args{
				originalData: &OneDimensionalDataInfo{
					{PointIndex: 1, Value: 1},
					{PointIndex: 2, Value: 0.1},
					{PointIndex: 3, Value: 1},
					{PointIndex: 4, Value: 1},
					{PointIndex: 5, Value: 1},
					{PointIndex: 6, Value: 1},
					{PointIndex: 7, Value: 1},
					{PointIndex: 8, Value: 100},
					{PointIndex: 9, Value: 90},
					{PointIndex: 10, Value: 80},
				},
				centerCount: 2,
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kmeans := &KmeansPlusPlus{
				CenterClusters: tt.fields.CenterClusters,
				MaxIterations:  tt.fields.MaxIterations,
				IsConvergence:  tt.fields.IsConvergence,
				Runs:           tt.fields.Runs,
				CenterCount:    tt.fields.CenterCount,
			}
			if err := kmeans.OneDimensionalKmeansPlusPlus(tt.args.originalData, tt.args.centerCount); (err != nil) != tt.wantErr {
				t.Errorf("KmeansPlusPlus.OneDimensionalKmeansPlusPlus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
