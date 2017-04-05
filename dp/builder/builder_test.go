package builder

import (
	"fmt"
	"testing"
)

func Test_builder(t *testing.T) {
	var d *Director = &Director{b: &CharacterBuilder{}}
	var c *Character = d.Create("Loader", "AK47")
	fmt.Println(c.name + "," + c.arms)
}
