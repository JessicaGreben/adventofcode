package day3

import (
	"math"
)

var defaultCentralPoint = Point{0, 0}

type Point struct {
	X, Y int
}

func (p *Point) x() int {
	return p.X
}

func (p *Point) y() int {
	return p.Y
}

// planePoint is an x,y point that is plotted on a plane.
// A collection of points gets plotted on a plane starting at the
// plane's central point. Points are added one at a time and are assigned
// an order when they are plotted. The order starts at 1 and increments one
// at a time.
// When a collection of points gets plotted on a plane as a single
// object, the points for that object all have the same objectID.
type planePoint struct {
	Point
	objectID int
	order    int
}

func (p *planePoint) getPoint() Point {
	return p.Point
}

func (p *planePoint) getObjectID() int {
	return p.objectID
}

func (p *planePoint) setObjectID(id int) {
	p.objectID = id
}

func (p *planePoint) getOrder() int {
	return p.order
}

// Plane represents an x,y coordinate plane.
// A plane can plot many different objects and it assigns each new object
// an objectID to tell them apart.
//
// The x,y coordinate plane is represented by a map defined as:
//   - key: an x,y point on the plane
//   - value: list of object's orderedPoints at that position on the plane
//
// The centralPoint is the start location of all objects when they
// are plotted on the plane.
type Plane struct {
	centralPoint        Point
	closestIntersection int
	points              map[Point][]planePoint
	objectIDCounter     int
}

func NewPlane(centralPoint Point) *Plane {
	return &Plane{
		centralPoint:        centralPoint,
		closestIntersection: math.MaxInt32,
		points:              map[Point][]planePoint{},
	}
}

// Plot plots a new object on the plane.
// An object is represented by a collection of planePoints.
// When a new object is plotted all the planePoints get the same objectID.
func (p *Plane) Plot(points []planePoint) {
	p.objectIDCounter++
	for _, point := range points {
		point.setObjectID(p.objectIDCounter)
		p.add(point)
	}
}

// add adds a new x,y point to the plane. When a new point is added to
// the plane, it checks if it intersects with any other objects. If the new
// point does intersect with other objects, then it calculates if its the
// closest intersection to the central point.
func (p *Plane) add(newPoint planePoint) {
	existingPoints, ok := p.points[newPoint.getPoint()]
	if ok {
		for _, existingPoint := range existingPoints {
			if p.pointsIntersect(existingPoint, newPoint) {
				// if the newPoint intersects with an existingPoint then check if the distance
				// of this intersection is closer than any other intersection to the central point.
				distance := manhattenDistance(p.centralPoint, newPoint.getPoint())
				if distance < p.closestIntersection {
					p.closestIntersection = distance
				}
			}
		}
	}
	p.points[newPoint.getPoint()] = append(p.points[newPoint.getPoint()], newPoint)
}

func (p *Plane) pointsIntersect(point1, point2 planePoint) bool {
	if point1.getObjectID() == point2.getObjectID() {
		return false
	}
	if p.isCentralPoint(point1) || p.isCentralPoint(point2) {
		return false
	}
	return point1.x() == point2.x() && point1.y() == point2.y()
}

func (p *Plane) isCentralPoint(point planePoint) bool {
	return point.x() == p.centralPoint.x() && point.y() == p.centralPoint.y()
}

func manhattenDistance(point1, point2 Point) int {
	x1, y1 := point1.x(), point1.y()
	x2, y2 := point2.x(), point2.y()
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
