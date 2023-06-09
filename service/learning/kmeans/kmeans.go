package kmeans

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/whatattitude/gDAG/lib/log/logger"

	"math/rand"
	"sort"
	"time"
)

type DataInfo struct {
	Labels      map[string]string
	Value       float64
	PointIndex  int
	CenterIndex int
	Distance    float64
}

type OneDimensionalData interface {
	GetDataList() (dataList OneDimensionalDataInfo)
	RandomCenter(centerCount int) (CenterIndex OneDimensionalDataInfo, err error)
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
	logger.Logger.Sugar().Debugf("start OneDimensionalKmeansPlusPlus %+v", originalData)
	if len(originalData.GetDataList().DataSlice) == 0 {
		return errors.New("kmeans has no data , nothing to do")
	}
	centerSlice, err := originalData.RandomCenter(centerCount)
	if err != nil {
		return err
	}
	logger.Logger.Sugar().Debugf(" %+v", centerSlice)

	for kmeans.Runs < kmeans.MaxIterations {

		currCenterClusters, err := kmeans.CenterIteration(centerSlice, originalData)
		if err != nil {
			return err
		}

		output, _ := json.Marshal(currCenterClusters)
		logger.Logger.Sugar().Debug((string(output)))

		kmeans.ConvergenceCheck(currCenterClusters)
		kmeans.CenterClusters = currCenterClusters

		// kmeans.CenterClusters = make([]Cluster, len(currCenterClusters))
		// kmeans.CenterClusters = append(kmeans.CenterClusters, currCenterClusters...)
		output1, _ := json.Marshal(kmeans.CenterClusters)
		output2, _ := json.Marshal(currCenterClusters)
		logger.Logger.Sugar().Debugf(" %+v  ---  %+v ", string(output1), string(output2))
		logger.Logger.Sugar().Debugf(" %+v", kmeans)
		if kmeans.IsConvergence {
			logger.Logger.Info("kmeans.IsConvergence is true")
			break
		}

		err = kmeans.ClusterRandomCenter(centerSlice)
		if err != nil {
			logger.Logger.Sugar().Debugf(" %+v", kmeans)
			logger.Logger.Sugar().Debugf(err.Error())
			kmeans.IsConvergence = true
			return err
		}

	}

	if !kmeans.IsConvergence {
		return errors.New("kmeans failed cluster is not convergence, please add runs and retry")
	}
	return

}

func (kmeans *KmeansPlusPlus) ClusterRandomCenter(centerSlice OneDimensionalDataInfo) (err error) {
	rand.NewSource(time.Now().UnixNano())
	// centerIndex = append(centerIndex, (*o)[rand.Intn((*o).Len()-1)])
	for i := 0; i < len(kmeans.CenterClusters); i++ {
		if len(kmeans.CenterClusters[i].DataSlice) == 0 {
			return errors.New("one CenterClusters has no data  " + strconv.Itoa(kmeans.CenterClusters[i].CenterIndex))
		}
		centerPoint := kmeans.CenterClusters[i].DataSlice[rand.Intn(len(kmeans.CenterClusters[i].DataSlice))]
		centerSlice.DataSlice[i].CenterIndex = centerPoint.PointIndex
		centerSlice.DataSlice[i].PointIndex = centerPoint.PointIndex
		centerSlice.DataSlice[i].Value = centerPoint.Value
		centerSlice.DataSlice[i].Distance = -1
	}
	return
}

func (kmeans *KmeansPlusPlus) ConvergenceCheck(currCenterClusters []Cluster) {
	if currCenterClusters == nil || kmeans.CenterClusters == nil || len(kmeans.CenterClusters) == 0 {
		return
	}

	for index, curr := range currCenterClusters {

		original := kmeans.CenterClusters[index]

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

	}

	kmeans.IsConvergence = true

}

func (kmeans *KmeansPlusPlus) CenterIteration(centerIndexSlice OneDimensionalDataInfo,
	originalData OneDimensionalData) (currCenterClusters []Cluster, err error) {
	dataList := originalData.GetDataList()
	if len(dataList.DataSlice) == 0 {
		return nil, errors.New("originalData.GetDataList get no data, nothing need to do ")
	}
	currCenterClusters = make([]Cluster, 0)

	logger.Logger.Sugar().Debugf("%+v", dataList)
	for i := range centerIndexSlice.DataSlice {
		currCenterClustersItem := Cluster{
			CenterIndex: centerIndexSlice.DataSlice[i].CenterIndex,
		}
		currCenterClusters = append(currCenterClusters, currCenterClustersItem)
	}
	logger.Logger.Sugar().Debugf("%+v", currCenterClusters)
	for _, v := range dataList.DataSlice {
		DataSlice := OneDimensionalDataInfo{DataSlice: make([]DataInfo, 0), SortType: "distance"}
		for _, x1 := range centerIndexSlice.DataSlice {
			distance := originalData.GetDistance(x1, v)
			item := v.DeepCopy()
			item.Distance = distance
			item.CenterIndex = x1.PointIndex
			DataSlice.DataSlice = append(DataSlice.DataSlice, *item)
		}

		sort.Sort(DataSlice)
		logger.Logger.Sugar().Debugf("%+v", DataSlice)
		for i := range currCenterClusters {
			if currCenterClusters[i].CenterIndex == DataSlice.DataSlice[0].CenterIndex {
				currCenterClusters[i].DataSlice = append(currCenterClusters[i].DataSlice, DataSlice.DataSlice[0])
			}
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

	c2.DataSlice = make([]DataInfo, 0)
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

func (c Cluster) RemovalduplicateByLabel(labelkey string) (c2 Cluster) {
	c2.CenterIndex = c.CenterIndex
	c2.Labels = c.Labels
	c2.DataSlice = make([]DataInfo, 0)
	if len(c.DataSlice) == 0 {
		return c2
	}
	for i := range c.DataSlice {
		value, ok := c.DataSlice[i].Labels[labelkey]
		if !ok {
			logger.Logger.Debug("labelkey " + labelkey + " is not exsist")
			return c2
		}
		inC2 := false
		for j := range c2.DataSlice {
			value2, ok := c2.DataSlice[j].Labels[labelkey]
			if !ok {
				logger.Logger.Debug("labelkey " + labelkey + " is not exsist")
				return c2
			}
			if value2 == value {
				inC2 = true
			}

		}
		if !inC2 {
			c2.DataSlice = append(c2.DataSlice, *c.DataSlice[i].DeepCopy())
		}
	}
	return c2
}
