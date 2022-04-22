package format

type Properties struct {
	prop map[string]string
}

func (p *Properties) Get(key string) string {
	return p.prop[key]
}

func (p *Properties) Set(key string, val string) {
	p.prop[key] = val
}

type Ini struct {
	Sections map[string]*Properties
}

func (i *Ini) Get(section string, key string) string {
	return i.Sections[section].Get(key)
}
