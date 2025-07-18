package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	WindowSize Vector2
	update     = func() {
		if !isTabFocused() {
			return
		}
		UpdateWindowSize()
		PluginsUpdate()
		BeginDrawing()
		ClearBackground(Black)
		text := fmt.Sprintf("%v", WindowSize)
		DrawText(text, 0, 0, 20, White)
		PluginsDraw()
		EndDrawing()
	}
)

func UpdateWindowSize() {
	vw, vh := viewportSize()
	SetWindowSize(vw, vh)
	WindowSize = NewVector2(float32(vw), float32(vh))
}

func main() {
	InitWindow(0, 0, "")
	defer CloseWindow()
	UpdateWindowSize()
	plugins["particles"] = &Particles{}
	PluginsInit()
	SetTargetFPS(60)
	SetMainLoop(update)
	for !WindowShouldClose() {
		update()
	}
}
