package builder

type Character struct {
	name string
	arms string
}

func (p *Character) SetName(name string) {
	p.name = name
}

func (p *Character) SetArms(arms string) {
	p.arms = arms
}

func (p Character) GetName() string {
	return p.name
}

func (p Character) GetArms() string {
	return p.arms
}
