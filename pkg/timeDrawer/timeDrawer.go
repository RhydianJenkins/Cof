package timeDrawer

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Draw(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	timeStr := time.Now().Format("Current time is 15:04:05")
	tview.Print(screen, timeStr, x, height/2, width, tview.AlignCenter, tcell.ColorLime)
	return 0, 0, 0, 0
}
