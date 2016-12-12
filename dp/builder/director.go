package builder

type Director struct {
	b Builder
}

func (p Director) Create(name string, arms string) *Character {
	return p.b.SetName(name).SetArms(arms).Build()
}
