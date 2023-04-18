package kmeans

import (
	"math"
	"math/rand"
	"sort"
	"time"
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

type OneDimensionalDataInfo []DataInfo

func (o *OneDimensionalDataInfo) GetDataList() (dataList []DataInfo) {
	for i := range *o {
		(*o)[i].CenterIndex = (*o)[i].PointIndex
		(*o)[i].Distance = 0
	}
	return *o
}

func (o *OneDimensionalDataInfo) RandomCenter(centerCount int) (centerIndex []DataInfo) {
	if centerCount == 2 {
		sort.Sort(o)
		centerIndex = append(centerIndex, (*o)[0])
		centerIndex = append(centerIndex, (*o)[o.Len()-1])
	} else {
		for i := 0; i < centerCount; i++ {
			rand.Seed(time.Now().UnixNano())
			centerIndex = append(centerIndex, (*o)[rand.Intn((*o).Len()-1)])
			centerIndex = append(centerIndex, (*o)[rand.Intn((*o).Len()-1)])
		}

	}
	return
}

func (o *OneDimensionalDataInfo) GetDistance(x1, x2 DataInfo) (distance float64) {
	return math.Abs(x1.Distance - x2.Distance)
}

func (n OneDimensionalDataInfo) Len() int {
	return len(n)
}

func (n OneDimensionalDataInfo) Less(i, j int) bool {
	return (n)[i].Value < (n)[j].Value
}

func (n OneDimensionalDataInfo) Swap(i, j int) {
	tmp := (n)[i].DeepCopy()
	(n)[i] = *n[j].DeepCopy()
	(n)[j] = *tmp.DeepCopy()
}
