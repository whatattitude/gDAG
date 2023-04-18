package kmeans

import (
	"encoding/json"
	"errors"
	"fmt"
	"gDAG/lib/log/logger"
	"sort"
)

type DataInfo struct {
	Labels      map[string]string
	Value       float64
	PointIndex  int
	CenterIndex int
	Distance    float64
}

type OneDimensionalData interface {
	GetDataList() (dataList []DataInfo)
	RandomCenter(centerCount int) (CenterIndex []DataInfo)
	GetDistance(x1, x2 DataInfo) (distance float64)
}

type KmeansPlusPlus struct {
	CenterClusters []Cluster
	MaxIterations  int
	IsConvergence  bool
	Runs           int
	CenterCount    int
}

type Cluster struct {
	DataSlice   []DataInfo
	CenterIndex int
	Labels      map[string]string
}

type SegmentDistance struct {
	CenterIndex   int
	EndPointIndex int
	Distance      float64
}

func (kmeans *KmeansPlusPlus) OneDimensionalKmeansPlusPlus(originalData OneDimensionalData,
	centerCount int) (err error) {

	kmeans.CenterCount = centerCount
	logger.Logger.Sugar().Debugf(" %+v", originalData)
	centerIndexSlice := originalData.RandomCenter(centerCount)
	logger.Logger.Sugar().Debugf(" %+v", centerIndexSlice)

	for kmeans.Runs < kmeans.MaxIterations {

		currCenterClusters, err := kmeans.CenterIteration(centerIndexSlice, originalData)
		if err != nil {
			return err
		}

		output, _ := json.Marshal(currCenterClusters)
		logger.Logger.Sugar().Debug((json.RawMessage(output)))
		fmt.Println(string(output))

		kmeans.ConvergenceCheck(currCenterClusters)
		kmeans.CenterClusters = currCenterClusters
		// kmeans.CenterClusters = make([]Cluster, len(currCenterClusters))
		// kmeans.CenterClusters = append(kmeans.CenterClusters, currCenterClusters...)
		logger.Logger.Sugar().Debugf(" %+v", kmeans)
		if kmeans.IsConvergence {
			logger.Logger.Info("kmeans.IsConvergence is true")
			break
		}
	}

	if !kmeans.IsConvergence {
		return errors.New("kmeans failed cluster is not convergence, please add runs and retry")
	}
	return

}

func (kmeans *KmeansPlusPlus) ConvergenceCheck(currCenterClusters []Cluster) {
	if currCenterClusters == nil && kmeans.CenterClusters == nil {
		return
	}

	for _, curr := range currCenterClusters {
		var inOriginal = false
		for _, original := range kmeans.CenterClusters {
			if curr.CenterIndex == original.CenterIndex {
				inOriginal = true

				for _, currDataSlice := range curr.DataSlice {
					var inOriginalDataSlice = false
					for _, originalDataSlice := range original.DataSlice {
						if currDataSlice.PointIndex == originalDataSlice.PointIndex {
							inOriginalDataSlice = true
							break
						}
					}
					if !inOriginalDataSlice {
						kmeans.IsConvergence = false
						return
					}
				}
				break
			}
		}
		if !inOriginal {
			kmeans.IsConvergence = false
			return
		}
	}
	kmeans.IsConvergence = true
	return
}

func (kmeans *KmeansPlusPlus) CenterIteration(centerIndexSlice []DataInfo,
	originalData OneDimensionalData) (currCenterClusters []Cluster, err error) {
	dataList := originalData.GetDataList()
	if len(dataList) == 0 {
		return nil, errors.New("originalData.GetDataList get no data, nothing need to do ")
	}
	currCenterClusters = make([]Cluster, len(centerIndexSlice))

	for _, v := range dataList {
		DataSlice := make([]DataInfo, 0)
		for _, x1 := range centerIndexSlice {
			distance := originalData.GetDistance(x1, v)
			item := v.DeepCopy()
			item.Distance = distance
			item.CenterIndex = x1.CenterIndex
			DataSlice = append(DataSlice, *item)
		}

		DataSliceOneDimensional := OneDimensionalDataInfo(DataSlice)
		sort.Sort(DataSliceOneDimensional)

		var inCurr = false
		for i := range currCenterClusters {
			if currCenterClusters[i].CenterIndex == DataSlice[0].CenterIndex {
				currCenterClusters[i].DataSlice = append(currCenterClusters[i].DataSlice, DataSlice[0])
				inCurr = true
			}
		}
		if !inCurr {
			currCenterClusters = append(currCenterClusters)
		}

	}
	kmeans.Runs++
	return
}

func (c *Cluster) DeepCopy() (c2 *Cluster) {
	if c == nil {
		return nil
	}
	c2 = &Cluster{}
	*c2 = *c
	if c.Labels != nil {
		c2.Labels = make(map[string]string)
		for k, v := range c.Labels {
			c2.Labels[k] = v
		}
	}
	if len(c.DataSlice) < 0 {
		return
	}
	c2.DataSlice = make([]DataInfo, len(c.DataSlice))
	for _, v := range c.DataSlice {
		c2.DataSlice = append(c2.DataSlice, *v.DeepCopy())
	}
	return
}

func (d *DataInfo) DeepCopy() (d2 *DataInfo) {
	if d == nil {
		return nil
	}
	d2 = &DataInfo{}
	*d2 = *d
	if d.Labels != nil {
		d2.Labels = make(map[string]string)
		for k, v := range d.Labels {
			d2.Labels[k] = v
		}
	}
	return
}

// func (d DataInfo) DeepCopy() (d2 DataInfo) {

// 	d2 = d
// 	if d.Labels != nil {
// 		d2.Labels = make(map[string]string)
// 		for k, v := range d.Labels {
// 			d2.Labels[k] = v
// 		}
// 	}
// 	return
// }
