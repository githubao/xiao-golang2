// 附近的人
// author: baoqiang
// time: 2019/2/28 下午10:25
package impl

import (
	"github.com/gansidui/geohash"
	"fmt"
	"github.com/gansidui/nearest"
)

func RunGeo() {
	latitude := 39.92324
	longitude := 116.3906
	precision := 5

	hash, box := geohash.Encode(latitude, latitude, precision)

	fmt.Println(hash)
	fmt.Println(box.MinLat, box.MaxLat, box.MinLng, box.MaxLng)

	neighbors := geohash.GetNeighbors(latitude, longitude, precision)
	for _, hash := range neighbors {
		fmt.Println(hash)
	}

}

func RunNearest() {
	near := nearest.NewNearest()
	near.SetPrecision(5)
	near.AddCoord("A", 40.92424, 116.3906)
	near.AddCoord("B", 39.93224, 116.3927)
	near.AddCoord("C", 39.92484, 116.3916)
	near.AddCoord("D", 39.92494, 116.3923)
	near.AddCoord("E", 39.92220, 116.3915)
	near.AddCoord("F", 39.92424, 117.3906)

	keys := near.QueryNearestSquareFromKey("C")
	coordNode1, ok := near.GetCoordNode("C")
	if !ok {
		return
	}
	for _, key := range keys {
		coordNode2, _ := near.GetCoordNode(key)
		fmt.Println(key, nearest.DistanceCoordNode(coordNode1, coordNode2))
	}

}
