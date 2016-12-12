package builder

type Builder interface {
	Build() *Character
	SetName(name string) Builder
	SetArms(arms string) Builder
}
