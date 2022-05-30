package rainDrawer

import (
	"fmt"
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	DROP_CHAR      = "'"
	RAIN_INTENSITY = 2
	MIN_DROP_SPEED = 1
	MAX_DROP_SPEED = 3
)

var (
	drops []*drop
)

type drop struct {
	posX  int
	posY  int
	speed int
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func createDrops(width int) {
	for i := 0; i < RAIN_INTENSITY; i++ {
		posX := randomInt(0, width)
		speed := randomInt(MIN_DROP_SPEED, MAX_DROP_SPEED)
		newDrop := drop{posX: posX, posY: 0, speed: speed}
		drops = append(drops, &newDrop)
	}
}

func incrementDrops() {
	for _, drop := range drops {
		drop.posY += drop.speed
	}
}

func disposeDrops(maxHeight int) {
	for i := 0; i < len(drops); i++ {
		if drops[i].posY > maxHeight {
			drops[i] = nil
		}
	}

	newDrops := make([]*drop, 0, len(drops))
	for _, drop := range drops {
		if drop != nil {
			newDrops = append(newDrops, drop)
		}
	}

	drops = newDrops
}

func drawDrops(screen tcell.Screen, drops []*drop) {
	for _, drop := range drops {
		tview.Print(screen, DROP_CHAR, drop.posX, drop.posY, 1, tview.AlignLeft, tcell.ColorLime)
	}
}

func Draw(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	str := fmt.Sprintf("%d rain drops falling", len(drops))
	tview.Print(screen, str, x, height/2, width, tview.AlignCenter, tcell.ColorLime)

	incrementDrops()
	createDrops(width)
	drawDrops(screen, drops)
	disposeDrops(height)

	return 0, 0, 0, 0
}
