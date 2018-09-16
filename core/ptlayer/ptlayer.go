package ptlayer

import (
	"math"
	"math/rand"
	"time"
)

type Pts struct {
	X int
	Y int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate scattered points on page
func RandPts(count int, width int, height int) (ptlist []Pts) {
	// divide the whole area to boxes, we are essentially picking `count` points, each from different boxes
	// `5.0` is a constant used to increase available box count.
	var mindist int = int(math.Sqrt(float64((width * height) / (5.0 * count)))) // magic number 5.0 -- see above comments

randpt:
	for {
		xpos := (rand.Int() % width)
		ypos := (rand.Int() % height)
		if xpos < mindist/2 || (width-xpos) < mindist/2 || ypos < mindist/2 || (height-ypos) < mindist/2 {
			continue randpt
		}
		for _, pt := range ptlist {
			var absdist int = 0
			if pt.X < xpos {
				absdist += (xpos - pt.X)
			} else {
				absdist += (pt.X - xpos)
			}
			if pt.Y < ypos {
				absdist += (ypos - pt.Y)
			} else {
				absdist += (pt.Y - ypos)
			}
			if absdist < mindist {
				continue randpt
			}
		}
		ptlist = append(ptlist, Pts{
			X: xpos,
			Y: ypos,
		})
		if len(ptlist) >= count {
			break randpt
		}
	}
	return
}
