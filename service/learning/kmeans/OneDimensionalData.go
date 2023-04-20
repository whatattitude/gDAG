package kmeans

import (
	"errors"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/whatattitude/gDAG/lib/log/logger"
)

// type OneDimensionalDataInfo interface {
// 	GetDataList() (dataList []DataInfo)
// 	RandomCenter(centerCount int) (CenterIndex []int)
// 	GetDistance(x1, x2 int) (distance float64)
// }

// type DataInfo struct {
// 	Labels      map[string]string
// 	Value       float64
// 	PointIndex  int
// 	CenterIndex int
// 	Distance    float64
// }

// type OneDimensionalDataInfo []DataInfo
type OneDimensionalDataInfo struct {
	DataSlice []DataInfo
	SortType  string
}

// type DataSort struct {
// 	DataSlice []DataInfo
// 	SortType  string
// }

func (o *OneDimensionalDataInfo) GetDataList() (dataList OneDimensionalDataInfo) {

	for i := range o.DataSlice {
		o.DataSlice[i].CenterIndex = o.DataSlice[i].PointIndex
		o.DataSlice[i].Distance = 0
		o.DataSlice[i].CenterIndex = -1
	}
	return *o
}

func (o *OneDimensionalDataInfo) RandomCenter(centerCount int) (dataInfo OneDimensionalDataInfo, err error) {
	if o.Len() == 0 {
		return dataInfo, errors.New("OneDimensionalDataInfo has no data , can not generate center point")
	}
	o.SortType = "value"
	logger.Logger.Sugar().Debugf("---  %+v", *o)
	sort.Sort(o)
	logger.Logger.Sugar().Debugf("-=-=-  %+v", *o)
	centerIndex := dataInfo.DataSlice
	if centerCount == 2 {

		centerIndex = append(centerIndex, (*o).DataSlice[0])
		centerIndex = append(centerIndex, (*o).DataSlice[o.Len()-1])
	} else {
		for i := 0; i < centerCount; i++ {
			rand.NewSource(time.Now().UnixNano())
			centerIndex = append(centerIndex, (*o).DataSlice[rand.Intn((*o).Len()-1)])
			centerIndex = append(centerIndex, (*o).DataSlice[rand.Intn((*o).Len()-1)])
		}

	}
	for i := range centerIndex {
		centerIndex[i].CenterIndex = centerIndex[i].PointIndex
	}
	dataInfo.DataSlice = centerIndex
	return
}

func (o *OneDimensionalDataInfo) GetDistance(x1, x2 DataInfo) (distance float64) {
	return math.Abs(x1.Value - x2.Value)
}

func (o OneDimensionalDataInfo) Len() int {
	return len(o.DataSlice)
}

func (o OneDimensionalDataInfo) Less(i, j int) bool {
	logger.Logger.Sugar().Debugf(o.SortType)
	switch o.SortType {
	case "value":
		return (o.DataSlice)[i].Value < (o.DataSlice)[j].Value
	case "distance":
		return (o.DataSlice)[i].Distance < (o.DataSlice)[j].Distance
	}
	return (o.DataSlice)[i].Distance < (o.DataSlice)[j].Distance
}

func (o OneDimensionalDataInfo) Swap(i, j int) {
	tmp := *o.DataSlice[i].DeepCopy()
	o.DataSlice[i] = *o.DataSlice[j].DeepCopy()
	o.DataSlice[j] = *tmp.DeepCopy()
}
