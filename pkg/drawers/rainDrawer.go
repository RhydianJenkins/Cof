package drawers

import (
	"fmt"
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	DROP_CHARS      = "'`|-"
	RAIN_INTENSITY  = 5
	MIN_DROP_SPEED  = 0.5
	MAX_DROP_SPEED  = 3.0
	RAIN_DROP_DRIFT = 0.25
)

type drop struct {
	posX  float64
	posY  float64
	speed float64
	drift float64
	char  rune
}

type RainDrawer struct {
	drops []*drop
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (rainDrawer *RainDrawer) createDrops(width int) {
	drift := RAIN_DROP_DRIFT
	for i := 0; i < RAIN_INTENSITY; i++ {
		posX := randomFloat(0, float64(width))
		speed := randomFloat(MIN_DROP_SPEED, MAX_DROP_SPEED)
		char := rune(DROP_CHARS[len(rainDrawer.drops)%len(DROP_CHARS)])
		newDrop := drop{posX: posX, posY: -1, speed: speed, drift: drift, char: char}
		rainDrawer.drops = append(rainDrawer.drops, &newDrop)
	}
}

func (rainDrawer *RainDrawer) incrementDrops() {
	for _, drop := range rainDrawer.drops {
		drop.posY += drop.speed
		drop.posX += drop.drift
	}
}

func (rainDrawer *RainDrawer) disposeDrops(maxHeight int) {
	newDrops := make([]*drop, 0, len(rainDrawer.drops))
	for i := 0; i < len(rainDrawer.drops); i++ {
		if rainDrawer.drops[i].posY <= float64(maxHeight) {
			newDrops = append(newDrops, rainDrawer.drops[i])
		}
	}

	rainDrawer.drops = newDrops
}

func (rainDrawer *RainDrawer) drawDrops(screen tcell.Screen) {
	for _, drop := range rainDrawer.drops {
		tview.Print(screen, string(drop.char), int(drop.posX), int(drop.posY), 1, tview.AlignLeft, tcell.ColorBlue)
	}
}

func (rainDrawer *RainDrawer) Draw(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	rainDrawer.incrementDrops()
	rainDrawer.createDrops(width)
	rainDrawer.drawDrops(screen)
	rainDrawer.disposeDrops(height)

	str := fmt.Sprintf("%d rain drops falling", len(rainDrawer.drops))
	tview.Print(screen, str, x, height/2, width, tview.AlignCenter, tcell.ColorLime)

	return 0, 0, 0, 0
}

func CreateRainDrawer() *RainDrawer {
	return &RainDrawer{}
}
