package cmd

import (
	"fmt"
	"image"
	"math/rand"
	"time"
	"github.com/go-vgo/robotgo"
)

// Fish inside World of Warcraft
func Fish(finder BobberFinder) {
	cast()
	pos := find(finder, nil)
	if nil != pos {
		bait := CreateBaitDetector()
		Wait: for i := 0; i < 200; i++ {
			if bait.CheckBait(find(finder, pos)) {
				loot(pos)
				break Wait
			}
			time.Sleep(time.Duration(rand.Intn(1100 - 950 + 1) + 1100) * time.Millisecond)
		}
	}
	randomSleep()
}

func cast() {
	robotgo.KeyTap(Key)
	randomSleep()
}

func find(finder BobberFinder, knownPosition *image.Point) *image.Point {
	screen := robotgo.CBitmap(robotgo.CaptureScreen())
	defer robotgo.FreeBitmap(robotgo.ToMMBitmapRef(screen))
	return finder.FindBobber(&screen, nil)
}

func loot(pos *image.Point) {
	robotgo.Move(pos.X, pos.Y)
	robotgo.Click("right")
}

func randomSleep() {
	r := rand.Intn(3000 - 1337 + 1) + 1337
	t := time.Duration(r) * time.Millisecond
	fmt.Printf("Sleeping for %v\n", t)
	time.Sleep(t)
}