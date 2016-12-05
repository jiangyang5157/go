package set

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	age  int
}

func Test_ThreadUnsafeSet(t *testing.T) {
	i1, i2, i3, i4, i33 := 1, 2, 3, 4, 3
	intSet := NewThreadUnsafeSet(i1, i3, i2, i4, i33)
	fmt.Println(intSet.String())
	intAddressSet := NewThreadUnsafeSet(&i1, &i3, &i2, &i4, &i33)
	fmt.Println(intAddressSet.String())

	int2Set := intSet.Clone()
	fmt.Println(int2Set.Add(4))
	fmt.Println(int2Set.Add(5))
	fmt.Println(int2Set.Add(6))
	fmt.Println("intSet: ", intSet.String())
	fmt.Println("int2Set: ", int2Set.String())
	fmt.Println("intSet.IsSubset(int2Set): ", intSet.IsSubset(int2Set))
	fmt.Println("int2Set.IsSubset(intSet): ", int2Set.IsSubset(intSet))
	fmt.Println("intSet.IsSuperset(int2Set): ", intSet.IsSuperset(int2Set))
	fmt.Println("int2Set.IsSuperset(intSet): ", int2Set.IsSuperset(intSet))
	fmt.Println("intSet.Difference(int2Set).String(): ", intSet.Difference(int2Set).String())
	fmt.Println("int2Set.Difference(intSet).String(): ", int2Set.Difference(intSet).String())
	fmt.Println("intSet.SymmetricDifference(int2Set).String(): ", intSet.SymmetricDifference(int2Set).String())
	fmt.Println("int2Set.SymmetricDifference(intSet).String(): ", int2Set.SymmetricDifference(intSet).String())

	sa, sb, sc, sd, scc := "a", "b", "c", "d", "c"
	stringSet := NewThreadUnsafeSet(sa, sc, sb, sd, scc)
	fmt.Println(stringSet.String())
	stringAddressSet := NewThreadUnsafeSet(&sa, &sc, &sb, &sd, &scc)
	fmt.Println(stringAddressSet.String())

	personSet := NewThreadUnsafeSet(
		Person{name: "Alise", age: 11},
		Person{name: "Bob", age: 33},
		Person{name: "John", age: 22},
		Person{name: "Nick", age: 44},
		Person{name: "Bob", age: 33},
	)
	fmt.Println(personSet.String())
	personAddressSet := NewThreadUnsafeSet(
		&Person{name: "Alise", age: 11},
		&Person{name: "Bob", age: 33},
		&Person{name: "John", age: 22},
		&Person{name: "Nick", age: 44},
		&Person{name: "Bob", age: 33},
	)
	fmt.Println(personAddressSet.String())

	personSet.Remove(Person{name: "John", age: 22})
	fmt.Println(personSet.String())
	personAddressSet.Remove(&Person{name: "John", age: 22})
	fmt.Println(personAddressSet.String())
}

func Test_ThreadSafeSet(t *testing.T) {
	i1, i2, i3, i4, i33 := 1, 2, 3, 4, 3
	intSet := NewThreadSafeSet(i1, i3, i2, i4, i33)
	fmt.Println(intSet.String())
	intAddressSet := NewThreadSafeSet(&i1, &i3, &i2, &i4, &i33)
	fmt.Println(intAddressSet.String())

	int2Set := intSet.Clone()
	fmt.Println(int2Set.Add(4))
	fmt.Println(int2Set.Add(5))
	fmt.Println(int2Set.Add(6))
	fmt.Println("intSet: ", intSet.String())
	fmt.Println("int2Set: ", int2Set.String())
	fmt.Println("intSet.IsSubset(int2Set): ", intSet.IsSubset(int2Set))
	fmt.Println("int2Set.IsSubset(intSet): ", int2Set.IsSubset(intSet))
	fmt.Println("intSet.IsSuperset(int2Set): ", intSet.IsSuperset(int2Set))
	fmt.Println("int2Set.IsSuperset(intSet): ", int2Set.IsSuperset(intSet))
	fmt.Println("intSet.Difference(int2Set).String(): ", intSet.Difference(int2Set).String())
	fmt.Println("int2Set.Difference(intSet).String(): ", int2Set.Difference(intSet).String())
	fmt.Println("intSet.SymmetricDifference(int2Set).String(): ", intSet.SymmetricDifference(int2Set).String())
	fmt.Println("int2Set.SymmetricDifference(intSet).String(): ", int2Set.SymmetricDifference(intSet).String())

	sa, sb, sc, sd, scc := "a", "b", "c", "d", "c"
	stringSet := NewThreadSafeSet(sa, sc, sb, sd, scc)
	fmt.Println(stringSet.String())
	stringAddressSet := NewThreadSafeSet(&sa, &sc, &sb, &sd, &scc)
	fmt.Println(stringAddressSet.String())

	personSet := NewThreadSafeSet(
		Person{name: "Alise", age: 11},
		Person{name: "Bob", age: 33},
		Person{name: "John", age: 22},
		Person{name: "Nick", age: 44},
		Person{name: "Bob", age: 33},
	)
	fmt.Println(personSet.String())
	personAddressSet := NewThreadSafeSet(
		&Person{name: "Alise", age: 11},
		&Person{name: "Bob", age: 33},
		&Person{name: "John", age: 22},
		&Person{name: "Nick", age: 44},
		&Person{name: "Bob", age: 33},
	)
	fmt.Println(personAddressSet.String())

	personSet.Remove(Person{name: "John", age: 22})
	fmt.Println(personSet.String())
	personAddressSet.Remove(&Person{name: "John", age: 22})
	fmt.Println(personAddressSet.String())
}

func Test_Iterator(t *testing.T) {
	personSet := NewThreadSafeSet(
		Person{name: "Alise", age: 11},
		Person{name: "Bob", age: 33},
		Person{name: "John", age: 22},
		Person{name: "Nick", age: 44},
		Person{name: "Bob", age: 33},
	)
	var personSetFound Person
	personSetIter := personSet.Iterator()
	for e := range personSetIter.ch {
		if e.(Person).name == "John" {
			personSetFound = e.(Person)
			personSetIter.Stop()
		}
	}
	fmt.Printf("personSetFound %+v in %v\n", personSetFound, personSet.String())
	personSet.Remove(personSetFound)
	fmt.Printf("After remove the personSetFound, %v\n", personSet.String())

	personAddressSet := NewThreadSafeSet(
		&Person{name: "Alise", age: 11},
		&Person{name: "Bob", age: 33},
		&Person{name: "John", age: 22},
		&Person{name: "Nick", age: 44},
		&Person{name: "Bob", age: 33},
	)
	var personAddressSetFound *Person = nil
	personAddressSetIter := personAddressSet.Iterator()
	for e := range personAddressSetIter.ch {
		if e.(*Person).name == "John" {
			personAddressSetFound = e.(*Person)
			personAddressSetIter.Stop()
		}
	}
	fmt.Printf("personAddressSetFound %+v in %v\n", personAddressSetFound, personAddressSet.String())
	personAddressSet.Remove(personAddressSetFound)
	fmt.Printf("After remove the personAddressSetFound, %v\n", personAddressSet.String())
}
