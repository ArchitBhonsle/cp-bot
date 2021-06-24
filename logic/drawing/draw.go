package drawing

import "strings"

func HorizontalTop(num int) string {
	return TopLeft + strings.Repeat(Horizontal, num-2) + TopRight
}

func HorizontalMiddle(num int) string {
	return LeftHorizontalIntersection + strings.Repeat(Horizontal, num-2) + RightHorizontalIntersection
}

func HorizontalBottom(num int) string {
	return BottomLeft + strings.Repeat(Horizontal, num-2) + BottomRight
}
