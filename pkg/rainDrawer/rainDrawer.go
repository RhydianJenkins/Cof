package rainDrawer

import (
	"fmt"
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	DROP_CHARS      = "`'|"
	RAIN_INTENSITY  = 5
	MIN_DROP_SPEED  = 0.5
	RAIN_DROP_DRIFT = 0.2
)

var (
	drops        []*drop
	numDropChars float64
)

func init() {
	numDropChars = float64(len(DROP_CHARS))
}

type drop struct {
	posX  float64
	posY  float64
	speed float64
	drift float64
	char  rune
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func createDrops(width int) {
	for i := 0; i < RAIN_INTENSITY; i++ {
		posX := randomFloat(0, float64(width))
		speed := randomFloat(MIN_DROP_SPEED, numDropChars)
		char := rune(DROP_CHARS[int(speed)+1])
		newDrop := drop{posX: posX, posY: -1, speed: speed, drift: RAIN_DROP_DRIFT, char: char}
		drops = append(drops, &newDrop)
	}
}

func incrementDrops() {
	for _, drop := range drops {
		drop.posY += drop.speed
		drop.posX += drop.drift
	}
}

func disposeDrops(maxHeight int) {
	newDrops := make([]*drop, 0, len(drops))
	for i := 0; i < len(drops); i++ {
		if drops[i].posY <= float64(maxHeight) {
			newDrops = append(newDrops, drops[i])
		}
	}

	drops = newDrops
}

func drawDrops(screen tcell.Screen, drops []*drop) {
	for _, drop := range drops {
		tview.Print(screen, string(drop.char), int(drop.posX), int(drop.posY), 1, tview.AlignLeft, tcell.ColorBlue)
	}
}

func Draw(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	incrementDrops()
	createDrops(width)
	drawDrops(screen, drops)
	disposeDrops(height)

	str := fmt.Sprintf("%d rain drops falling", len(drops))
	tview.Print(screen, str, x, height/2, width, tview.AlignCenter, tcell.ColorLime)

	return 0, 0, 0, 0
}
