package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	colors   = 6
	halfLife = 0.04
)

type Atom struct {
	position, velocity, force V2
}

type Particles struct {
	atoms   [colors][]Atom
	attract [colors][colors]float64
}

func (p *Particles) Init() error {
	for i := range colors {
		for j := range colors {
			p.attract[i][j] = float64(GetRandomValue(-100, 100)) / 100
		}
	}
	for g := range p.atoms {
		zero := V2{}
		n := GetRandomValue(100, 200)
		for range n {
			x := float64(GetRandomValue(0, int32(WindowSize.X)))
			y := float64(GetRandomValue(0, int32(WindowSize.Y)))
			position := V2{x, y}
			p.atoms[g] = append(p.atoms[g], Atom{position, zero, zero})
		}
	}
	return nil
}

func (p *Particles) rule(i, j int, g float64) {
	atoms1 := p.atoms[i]
	atoms2 := p.atoms[j]
	rmax := float64(120)
	for a := range atoms1 {
		force := V2{}
		for b := range atoms2 {
			if i == j && a == b {
				continue
			}
			var ab V2
			ab.X = toroidalDelta(atoms1[a].position.X, atoms2[b].position.X, WindowSize.X)
			ab.Y = toroidalDelta(atoms1[a].position.Y, atoms2[b].position.Y, WindowSize.Y)
			ab_len := V2Length(ab)
			if ab_len > rmax {
				continue
			}
			new := V2Scale(ab, compute_force(ab_len/rmax, g)/ab_len)
			force = V2Add(force, new)
		}
		scale := rmax * 8
		atoms1[a].force = V2Add(atoms1[a].force, V2Scale(force, scale))
	}
}

func mix(a, b, t float64) float64 {
	return a + (b-a)*t
}

func compute_force(d, g float64) float64 {
	rmin := float64(0.3)
	mid := (rmin + 1) / 2
	var scale float64
	if d < rmin {
		scale = mix(-1, 0, d/rmin)
	} else if d < mid {
		scale = mix(0, g, (d-rmin)/(mid-rmin))
	} else {
		scale = mix(g, 0, (d-mid)/(1-mid))
	}
	return scale
}

func (p *Particles) Update() {
	if !isTabFocused() {
		return
	}
	n := 8
	dt := float64(GetFrameTime()) / float64(n)
	friction := math.Pow(0.5, float64(dt)/halfLife)
	for range n {
		p.UpdatePart(dt, friction)
	}
}

func (p *Particles) UpdatePart(dt, friction float64) {
	for c := range colors {
		for i := range p.atoms[c] {
			p.atoms[c][i].force = V2{}
		}
	}
	for i := range colors {
		for j := range colors {
			p.rule(i, j, p.attract[i][j])
		}
	}
	for c := range colors {
		for i := range p.atoms[c] {
			atom := &p.atoms[c][i]
			v := V2Scale(atom.velocity, friction)
			atom.velocity = V2Lerp(v, atom.force, dt)
			change := V2Scale(atom.velocity, dt)
			atom.position = V2Add(atom.position, change)
		}
	}
}

func (p Particles) Draw() {
	colors := [colors]Color{Red, Green, Blue, Yellow, Orange, Purple}
	for c := range colors {
		for i := range p.atoms[c] {
			atom := &p.atoms[c][i]
			atom.position.X = modulo(atom.position.X, WindowSize.X)
			atom.position.Y = modulo(atom.position.Y, WindowSize.Y)
			DrawCircle(int32(atom.position.X), int32(atom.position.Y), 3, colors[c])
		}
	}
}

func toroidalDelta(a, b, size float64) float64 {
	d := b - a
	if d > size/2 {
		d -= size
	} else if d < -size/2 {
		d += size
	}
	return d
}

func modulo(a, size float64) float64 {
	for a < 0 {
		a += size
	}
	for a > size {
		a -= size
	}
	return a
}
