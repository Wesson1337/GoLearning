package geometry

import "math"

type Point struct{ X, Y float64 }

// Традиционная функция
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.X)
}

// То же, но как метод типа Point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.X)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor // equivalent to (*p).X *= factor
	p.Y *= factor // equivalent to (*p).Y *= factor
}

