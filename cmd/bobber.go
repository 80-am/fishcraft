package cmd

import (
	"image"
	"github.com/go-vgo/robotgo"
)

// BobberFinder that detect
type BobberFinder interface {
	FindBobber(screen *robotgo.CBitmap, knownPosition *image.Point) *image.Point
}

type bobberFinderImpl struct {
}

// CreateBobberFinder that detects the bobber
func CreateBobberFinder() BobberFinder {
	return newbobberFinderImpl()
}

func (bf bobberFinderImpl) FindBobber(screen *robotgo.CBitmap, knownPosition *image.Point) *image.Point {
	var found *image.Point
	search := func(x, y int) bool {
		found = bf.check(screen, x, y)
		return nil != found
	}
	if nil != knownPosition {
		searchNearby(screen, knownPosition.X, knownPosition.Y, 50, 50, 50, 50, search)
	} else {
		w, h := imageSize(screen)
		searchNearby(screen, 0, 0, 0, w, 0, h, search)
	}
	return found
}

func (bf bobberFinderImpl) check(screen *robotgo.CBitmap, x int, y int) *image.Point {
	if isRed(screen, x, y) {
		redFeather := search(screen, x, y, 30, isRed)
		hook := search(screen, x, y, 15, isHook)
		if redFeather && hook {
			return &image.Point{X: x, Y: y}
		}
	}
	return nil
}

func color(screen *robotgo.CBitmap, x int, y int) (r int, g int, b int) {
	color := robotgo.CHex(robotgo.GetColor(robotgo.ToMMBitmapRef(*screen), x, y))
	return int((color >> 16) & 0xFF), int((color >> 8) & 0xFF), int((color) & 0xFF)
}

func isHook(screen *robotgo.CBitmap, a, c int) bool {
	minR := 70
	r, g, b := color(screen, a, c)
	return r >= minR && r <= g && r <= b
}

func isRed(screen *robotgo.CBitmap, x, y int) bool {
	r, g, b := color(screen, x, y)
	red := 1.9
	minRed := 50
	maxRed := 150
	gbMin := 0.50
	gbMax := 1.8
	gb := float64(g) / float64(b)
	return r >= minRed && r <= maxRed && r > int(float64(g)*red) && r > int(float64(b)*red) && gb >= gbMin && gb <= gbMax
}

func imageSize(bitmap *robotgo.CBitmap) (w, h int) {
	gbit := robotgo.ToBitmap(robotgo.ToMMBitmapRef(*bitmap))
	w = gbit.Width
	h = gbit.Height
	return w, h
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}

func newbobberFinderImpl() bobberFinderImpl {
	return bobberFinderImpl{}
}

func search(screen *robotgo.CBitmap, x, y, pixelNeeded int, check func(screen *robotgo.CBitmap, x, y int) bool) bool {
	pixelFound := 0
	return searchNearby(screen, x, y, 10, 10, 10, 10, func(x int, y int) bool {
		if check(screen, x, y) {
			pixelFound++
		}
		return pixelFound >= pixelNeeded
	})
}

func searchNearby(screen *robotgo.CBitmap, x, y, before, after, over, under int, call func(x, y int) bool) bool {
	w, h := imageSize(screen)
	searchWBegin := max(400, x - before)
	searchWEnd := min(w - 400, x + after)
	searchHBegin := max(400 , y - over)
	searchHEnd := min(h - 400, y + under)
	for searchY := searchHBegin; searchY < searchHEnd; searchY++ {
		for searchX := searchWBegin; searchX < searchWEnd; searchX++ {
			if call(searchX, searchY) {
				return true
			}
		}
	}
	return false
}
