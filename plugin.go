package main

type Plugin interface {
	Init() error
	Update()
	Draw()
}

var (
	plugins = make(map[string]Plugin)
)

func PluginsInit() error {
	for _, p := range plugins {
		err := p.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func PluginsUpdate() {
	for _, p := range plugins {
		p.Update()
	}
}

func PluginsDraw() {
	for _, p := range plugins {
		p.Draw()
	}
}
