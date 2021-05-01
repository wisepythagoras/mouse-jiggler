package main

import (
	"math"
	"time"

	"github.com/go-vgo/robotgo"
)

var jigglingOn = true
var mousePosQueue [][]int

func jiggle(x int, y int) {
	for {
		robotgo.MoveMouseSmooth(x, y, 0.1, 20.0)
		robotgo.MoveMouseSmooth(x+5, y, 0.1, 20.0)

		if !jigglingOn {
			return
		}
	}
}

func mousePosLoop(queue *Queue) {
	for {
		x, y := robotgo.GetMousePos()
		queue.Add([]int{x, y})
		time.Sleep(10 * time.Millisecond)
	}
}

func stdDeviation(items []int) float64 {
	var sum, mean, sd float64

	for _, v := range items {
		sum += float64(v)
	}

	mean = sum / float64(len(items))

	for _, v := range items {
		sd += math.Pow(float64(v)-mean, 2)
	}

	return math.Sqrt(sd / float64(len(items)))
}

func main() {
	queue := &Queue{
		MaxSize: 300,
	}

	go mousePosLoop(queue)

	lastMoveAt := time.Now()
	jigglingOn = false

	for {
		rawQueue := queue.Get()

		if len(rawQueue) < 300 {
			continue
		}

		var xArr, yArr []int

		// Add the positions into individual arrays.
		for _, v := range rawQueue {
			if len(v) == 0 {
				continue
			}

			xArr = append(xArr, v[0])
			yArr = append(yArr, v[1])
		}

		xStdDeviation := stdDeviation(xArr)
		yStdDeviation := stdDeviation(yArr)

		isStill := xStdDeviation == 0 && yStdDeviation == 0
		isJiggling := xStdDeviation < 2 && yStdDeviation == 0
		timeSinceLastMove := time.Now().Sub(lastMoveAt)

		if isStill && timeSinceLastMove > 30*time.Second && !jigglingOn {
			go jiggle(xArr[len(xArr)-1], yArr[len(yArr)-1])
			jigglingOn = true
		} else if !isStill && !isJiggling && jigglingOn {
			jigglingOn = false
		} else if !isStill && !isJiggling && !jigglingOn {
			lastMoveAt = time.Now()
		}

		time.Sleep(50 * time.Millisecond)
	}
}
