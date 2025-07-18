package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Atom struct {
	x, y, vx, vy float64
}

type Particles struct {
	atoms   [4][]Atom
	attract [4][4]float32
}

func (p *Particles) Init() error {
	p.attract = [4][4]float32{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, -0.1},
	}
	for g := range p.atoms {
		for range 200 {
			x := float64(GetRandomValue(0, int32(WindowSize.X)))
			y := float64(GetRandomValue(0, int32(WindowSize.Y)))
			p.atoms[g] = append(p.atoms[g], Atom{x, y, 0, 0})
		}
	}
	return nil
}

func (p *Particles) rule(i, j int, g float64) {
	atoms1 := p.atoms[i]
	atoms2 := p.atoms[j]
	g *= 100
	for a := range atoms1 {
		fx := float64(0)
		fy := float64(0)
		for b := range atoms2 {
			if i == j && a == b {
				continue
			}
			dx := atoms1[a].x - atoms2[b].x
			dy := atoms1[a].y - atoms2[b].y
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < 10 || dist > 80 {
				continue
			}
			F := g / dist
			fx += dx * F
			fy += dy * F
		}
		atoms1[a].vx = (atoms1[a].vx + fx) * 0.5
		atoms1[a].vy = (atoms1[a].vy + fy) * 0.5
		atoms1[a].x += atoms1[a].vx
		atoms1[a].y += atoms1[a].vy
		strength := float64(2)
		r := float64(10)
		if atoms1[a].x < r {
			// atoms1[a].vx *= -1
			atoms1[a].vx += (r - atoms1[a].x) * strength
			atoms1[a].x = r
		}
		if atoms1[a].x > float64(WindowSize.X)-r {
			// atoms1[a].vx *= -1
			atoms1[a].vx += (float64(WindowSize.X) - r - atoms1[a].x) * strength
			atoms1[a].x = float64(WindowSize.X) - r
		}
		if atoms1[a].y < r {
			// atoms1[a].vy *= -1
			atoms1[a].vy += (r - atoms1[a].y) * strength
			atoms1[a].y = r
		}
		if atoms1[a].y > float64(WindowSize.Y)-r {
			// atoms1[a].vy *= -1
			atoms1[a].vy += (float64(WindowSize.Y) - r - atoms1[a].y) * strength
			atoms1[a].y = float64(WindowSize.Y) - r
		}
	}
}

func (p *Particles) Update() {
	const (
		red    = 0
		green  = 1
		blue   = 2
		yellow = 3
	)
	p.rule(yellow, yellow, -0.1)
	// p.rule(yellow, yellow, 0.01)
	p.rule(blue, yellow, -0.0124)
	p.rule(blue, blue, -0.1)
	p.rule(yellow, blue, 0.004)
	// p.rule(green, blue, 0.002)
	p.rule(yellow, green, -0.097)
	p.rule(green, yellow, -0.097)
	p.rule(green, yellow, 0.0095)
	p.rule(red, red, 0.0095)
	p.rule(red, yellow, -0.0095)
	p.rule(red, green, 0.002)
	p.rule(green, blue, -0.04)
	p.rule(green, green, 0.1)
}

func (p Particles) Draw() {
	colors := [4]Color{Red, Green, Blue, Yellow}
	for g, group := range p.atoms {
		for _, atom := range group {
			// screen := Vector2Scale(pos, WindowSize.Y/2)
			// screen.X += WindowSize.X / 2
			// screen.Y += WindowSize.Y / 2
			DrawCircle(int32(atom.x), int32(atom.y), 3, colors[g])
		}
	}
}
