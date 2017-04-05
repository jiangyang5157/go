package builder

type CharacterBuilder struct {
	c *Character
}

func (p *CharacterBuilder) SetName(name string) Builder {
	if p.c == nil {
		p.c = &Character{}
	}
	p.c.name = name
	return p
}

func (p *CharacterBuilder) SetArms(arms string) Builder {
	if p.c == nil {
		p.c = &Character{}
	}
	p.c.arms = arms
	return p
}

func (p *CharacterBuilder) Build() *Character {
	return p.c
}
