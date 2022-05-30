package rainDrawer

import (
	"fmt"
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	dropChar = "'"
	minSpeed = 1
	maxSpeed = 3
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

func createDrop(posX int) *drop {
	speed := randomInt(minSpeed, maxSpeed)
	newDrop := drop{posX: posX, posY: 0, speed: speed}
	drops = append(drops, &newDrop)
	return &newDrop
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
		tview.Print(screen, dropChar, drop.posX, drop.posY, 100, tview.AlignLeft, tcell.ColorLime)
	}
}

func Draw(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	str := fmt.Sprintf("%d rain drops falling", len(drops))
	tview.Print(screen, str, x, height/2, width, tview.AlignCenter, tcell.ColorLime)

	incrementDrops()
	createDrop(randomInt(0, width))
	drawDrops(screen, drops)
	disposeDrops(height)

	return 0, 0, 0, 0
}
