package day3

import (
	"math"
)

// PlaneObject is an interface for an object that
// is plotted on a coordinate plane.
type PlaneObject interface {
	Steps() []Step
}

// Step represents a single step of an object that is plotted on a plane.
// A step is like an (x,y) point, but it has additional information about the
// the object it belongs to
type Step struct {
	lineDirection lineType
	point         Point
	Value         int
	objectID      int
}

func (s Step) Intersects(s2 Step) bool {
	if s.objectID == s2.objectID {
		return false
	}
	if s.lineDirection == s2.lineDirection {
		return false
	}
	if isCentralPoint(s) || isCentralPoint(s2) {
		return false
	}
	return s.point.X == s2.point.X && s.point.Y == s2.point.Y
}

func isCentralPoint(step Step) bool {
	if step.point.X == 0 && step.point.Y == 0 {
		return true
	}
	return false
}

type Point struct {
	X, Y int
}

// Plane represents an (x,y) coordinate plane.
// A plane can plot many different objects and assigns each new object
// an objectID to tell them apart.
// The (x,y) coordinate plane is represented by a map where the key in
// the map is an x,y point and the value in the map is a list of steps,
// where a step is one step of an object.
type Plane struct {
	centralPoint        Point
	closestIntersection int
	points              map[Point][]Step
	objectCount         int
}

func NewPlane(centralPoint Point) *Plane {
	return &Plane{
		centralPoint:        centralPoint,
		closestIntersection: math.MaxInt32,
		points:              map[Point][]Step{},
	}
}

func (p *Plane) Plot(obj PlaneObject) {
	p.objectCount++
	for _, step := range obj.Steps() {
		step.objectID = p.objectCount
		p.add(step)
	}
}

func (p *Plane) add(newStep Step) {
	existingSteps, ok := p.points[newStep.point]
	if ok {
		for _, existingStep := range existingSteps {
			if existingStep.Intersects(newStep) {
				// if the new step intersects with an existing step then check if the distance
				// of this intersection is closer than any other intersection to the central point.
				distance := manhattenDistance(p.centralPoint, newStep.point)
				if distance < p.closestIntersection {
					p.closestIntersection = distance
				}
			}
		}
	}
	p.points[newStep.point] = append(p.points[newStep.point], newStep)
}

func manhattenDistance(point1, point2 Point) int {
	x1, y1 := point1.X, point1.Y
	x2, y2 := point2.X, point2.Y
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
