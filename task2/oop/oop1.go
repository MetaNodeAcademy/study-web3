package main

import (
	"fmt"
	"math"
)

func main() {
	// 创建矩形实例并计算
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形面积: %.2f\n", rect.Area())
	fmt.Printf("矩形周长: %.2f\n", rect.Perimeter())

	// 创建圆形实例并计算
	circle := Circle{Radius: 4}
	fmt.Printf("圆形面积: %.2f\n", circle.Area())
	fmt.Printf("圆形周长: %.2f\n", circle.Perimeter())

	// 通过接口调用（可选）
	var s Shape = &rect
	fmt.Printf("接口调用 - 矩形面积: %.2f\n", s.Area())

	s = &circle
	fmt.Printf("接口调用 - 圆形面积: %.2f\n", s.Area())
}

/*
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/
type Shape interface {
	Area() float64      //面积
	Perimeter() float64 //周长
}

// 矩形
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r *Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}