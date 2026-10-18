package main

import (
	"errors"
	"fmt"
)

const Pi = 3.14159

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Error %s : %s : %f", e.Message, e.Field, e.Value)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * Pi * c.Radius
}

func validateShape(s Shape) error {
	switch val := s.(type) {
	case Rectangle:
		if val.Width <= 0 {
			return &ValidationError{
				Message: "幅は正の値を入力してください。",
				Field:   "Rectangle.Width",
				Value:   val.Width,
			}
		} else if val.Height <= 0 {
			return &ValidationError{
				Message: "高さは正の値を入力してください。",
				Field:   "Rectangle.Height",
				Value:   val.Height,
			}
		}
	case Circle:
		if val.Radius <= 0 {
			return &ValidationError{
				Message: "半径は正の値を入力してください。",
				Field:   "Circle.Radius",
				Value:   val.Radius,
			}
		}
	default:
		return errors.New("その他の型が入力されました。")
	}

	return nil
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0},
		Circle{Radius: -1.0},
		Rectangle{Width: 1.0, Height: 0.0},
	}

	var calcShapes []Shape

	fmt.Println("--- 図形のバリデーション ---")
	for i, v := range shapes {

		err := validateShape(v)

		if err != nil {
			var numErr *ValidationError
			if errors.As(err, &numErr) {
				fmt.Printf("#%d：%T ⇒ NG：[バリデーションエラー] %s：%s（Value：%.2f）\n", i+1, v, numErr.Field, numErr.Message, numErr.Value)
				fmt.Printf("【詳細】Field=%s, value=%.2f\n", numErr.Field, numErr.Value)
			} else {
				fmt.Printf("バリデーションチェック：%s\n", err.Error())
			}
		} else {
			fmt.Printf("#%d：%T ⇒ バリデーションOK\n", i+1, v)
			calcShapes = append(calcShapes, v)
		}

	}

	fmt.Println("\n--- 検証済み図形の一括計算 ---")
	for _, v := range calcShapes {
		fmt.Printf("図形：%T | 面積：%.2f | 周長：%.2f\n", v, v.Area(), v.Perimeter())
	}
}
