package cmd

import (
	"image"
	"math"
)

// BaitDetector to find the bait
type BaitDetector interface {
	CheckBait(position *image.Point) bool
}

// BaitDetectorImpl Implementor
type BaitDetectorImpl struct {
	top, right, bottom, left int
	once                     bool
}

// CreateBaitDetector create the detector
func CreateBaitDetector() BaitDetector {
	return newBaitDetectorImpl()
}

// CheckBait if movement
func (bait *BaitDetectorImpl) CheckBait(position *image.Point) bool {
	if nil == position {
		if Debug {
			DebugLogger.Println("Bobber not deteted.")
		}
		NotFound++
		return false
	}
	if bait.once && (math.Abs(float64(bait.left - position.X)) > 50 || math.Abs(float64(bait.top - position.Y)) > 50) {
		if Debug {
			DebugLogger.Printf("Found something outside outlier at %v.\n", position)
		}
		NotFound++
		return false
	}


	if !bait.once || position.X < bait.left {
		bait.left = position.X
	}
	if !bait.once || position.X > bait.right {
		bait.right = position.X
	}

	if !bait.once || position.Y < bait.top {
		bait.top = position.Y
	}
	if !bait.once || position.Y > bait.bottom {
		bait.bottom = position.Y
	}
	bait.once = true
	return bait.detect()
}

func (bait *BaitDetectorImpl) detect() bool {
	x := bait.left
	y := bait.top
	w := bait.right - bait.left
	h := bait.bottom - bait.top
	if Debug {
		DebugLogger.Printf("Bobber positioned at [%v, %v - %v, %v].\n", x, y, w, h)
	}
	return w >= Threshold || h >= Threshold
}

func newBaitDetectorImpl() BaitDetector {
	return &BaitDetectorImpl{}
}
