package drawing

import "strings"

func HorizontalTop(num int) string {
	return TLCorner + strings.Repeat(Horizontal, num-2) + TRCorner
}

func HorizontalMiddle(num int) string {
	return LMIntersection + strings.Repeat(Horizontal, num-2) + RMIntersection
}

func HorizontalBottom(num int) string {
	return BLCorner + strings.Repeat(Horizontal, num-2) + BRCorner
}
