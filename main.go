package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

var (
	WindowSize V2
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
	WindowSize = V2{float64(vw), float64(vh)}
}

func main() {
	InitWindow(0, 0, "")
	defer CloseWindow()
	UpdateWindowSize()
	plugins["particles"] = &Particles{}
	PluginsInit()
	SetTargetFPS(600)
	SetMainLoop(update)
	for !WindowShouldClose() {
		update()
	}
}
