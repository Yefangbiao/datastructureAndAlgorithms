package main

// 接口
type fit interface {
	getRadius() float64
}

// 圆孔
type roundHole struct {
	radius float64
}

func (r *roundHole) Radius() float64 {
	return r.radius
}

func (r *roundHole) fits(f fit) bool {
	return r.radius >= f.getRadius()
}
