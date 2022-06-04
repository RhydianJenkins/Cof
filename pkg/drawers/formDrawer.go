package drawers

import (
	"github.com/gdamore/tcell/v2"
)

type FormDrawer struct {
}

func (formDrawer *FormDrawer) Draw(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return 0, 0, 0, 0
}

func CreateFormDrawer() *FormDrawer {
	return &FormDrawer{}
}
