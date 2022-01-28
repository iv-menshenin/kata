package pointsSymmetry

type (
	coords struct {
		x int
		y int
	}
	info struct {
		shifted  bool
		center   int
		accurate bool
	}
)

func getPointCloudInfo(arr []coords) info {
	var minX, maxX int
	for i, c := range arr {
		switch true {
		case i == 0:
			minX, maxX = c.x, c.x
		case c.x < minX:
			minX = c.x
		case c.x > maxX:
			maxX = c.x
		}
	}
	center := (maxX + minX) / 2
	return info{
		shifted:  center-minX > maxX-center,
		center:   center,
		accurate: (maxX+minX)%2 == 0,
	}
}

func normalizePoint(point coords, i info) coords {
	if i.accurate {
		point.x -= i.center
	} else {
		switch true {
		case point.x >= i.center && i.shifted:
			point.x -= i.center - 1
		case point.x <= i.center && !i.shifted:
			point.x -= i.center + 1
		default:
			point.x -= i.center
		}
	}
	return point
}

// isSymmetric will determine if the point cloud is symmetric along the X-axis
func isSymmetric(pointCloud []coords) bool {
	i := getPointCloudInfo(pointCloud)
	var pointMap = make(map[coords]struct{})
	for _, point := range pointCloud {
		point = normalizePoint(point, i)
		if point.x == 0 {
			continue
		}
		if point.x < 0 {
			point.x *= -1
		}
		if _, ok := pointMap[point]; ok {
			delete(pointMap, point)
		} else {
			pointMap[point] = struct{}{}
		}
	}
	return len(pointMap) == 0
}
