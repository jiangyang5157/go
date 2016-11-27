package set

import (
	"testing"
	"fmt"
)

type Person struct {
	name string
	age  int
}

func Test_ThreadUnsafeSet(t *testing.T) {
	i1, i2, i3, i4, i33 := 1, 2, 3, 4, 3
	intSet := NewThreadUnsafeSet(i1, i3, i2, i4, i33)
	fmt.Println(intSet.ToString())
	intAddressSet := NewThreadUnsafeSet(&i1, &i3, &i2, &i4, &i33)
	fmt.Println(intAddressSet.ToString())

	int2Set := intSet.Clone()
	fmt.Println(int2Set.Add(4))
	fmt.Println(int2Set.Add(5))
	fmt.Println(int2Set.Add(6))
	fmt.Println("intSet: ", intSet.ToString())
	fmt.Println("int2Set: ", int2Set.ToString())
	fmt.Println("intSet.IsSubset(int2Set): ", intSet.IsSubset(int2Set))
	fmt.Println("int2Set.IsSubset(intSet): ", int2Set.IsSubset(intSet))
	fmt.Println("intSet.IsSuperset(int2Set): ", intSet.IsSuperset(int2Set))
	fmt.Println("int2Set.IsSuperset(intSet): ", int2Set.IsSuperset(intSet))
	fmt.Println("intSet.Difference(int2Set).ToString(): ", intSet.Difference(int2Set).ToString())
	fmt.Println("int2Set.Difference(intSet).ToString(): ", int2Set.Difference(intSet).ToString())
	fmt.Println("intSet.SymmetricDifference(int2Set).ToString(): ", intSet.SymmetricDifference(int2Set).ToString())
	fmt.Println("int2Set.SymmetricDifference(intSet).ToString(): ", int2Set.SymmetricDifference(intSet).ToString())

	sa, sb, sc, sd, scc := "a", "b", "c", "d", "c"
	stringSet := NewThreadUnsafeSet(sa, sc, sb, sd, scc)
	fmt.Println(stringSet.ToString())
	stringAddressSet := NewThreadUnsafeSet(&sa, &sc, &sb, &sd, &scc)
	fmt.Println(stringAddressSet.ToString())

	personSet := NewThreadUnsafeSet(
		Person{name: "Alise", age: 11},
		Person{name: "Bob", age: 33},
		Person{name: "John", age: 22},
		Person{name: "Nick", age: 44},
		Person{name: "Bob", age: 33},
	)
	fmt.Println(personSet.ToString())
	personAddressSet := NewThreadUnsafeSet(
		&Person{name: "Alise", age: 11},
		&Person{name: "Bob", age: 33},
		&Person{name: "John", age: 22},
		&Person{name: "Nick", age: 44},
		&Person{name: "Bob", age: 33},
	)
	fmt.Println(personAddressSet.ToString())

	personSet.Remove(Person{name: "John", age: 22})
	fmt.Println(personSet.ToString())
	personAddressSet.Remove(&Person{name: "John", age: 22})
	fmt.Println(personAddressSet.ToString())
}

func Test_ThreadSafeSet(t *testing.T) {
	i1, i2, i3, i4, i33 := 1, 2, 3, 4, 3
	intSet := NewThreadSafeSet(i1, i3, i2, i4, i33)
	fmt.Println(intSet.ToString())
	intAddressSet := NewThreadSafeSet(&i1, &i3, &i2, &i4, &i33)
	fmt.Println(intAddressSet.ToString())

	int2Set := intSet.Clone()
	fmt.Println(int2Set.Add(4))
	fmt.Println(int2Set.Add(5))
	fmt.Println(int2Set.Add(6))
	fmt.Println("intSet: ", intSet.ToString())
	fmt.Println("int2Set: ", int2Set.ToString())
	fmt.Println("intSet.IsSubset(int2Set): ", intSet.IsSubset(int2Set))
	fmt.Println("int2Set.IsSubset(intSet): ", int2Set.IsSubset(intSet))
	fmt.Println("intSet.IsSuperset(int2Set): ", intSet.IsSuperset(int2Set))
	fmt.Println("int2Set.IsSuperset(intSet): ", int2Set.IsSuperset(intSet))
	fmt.Println("intSet.Difference(int2Set).ToString(): ", intSet.Difference(int2Set).ToString())
	fmt.Println("int2Set.Difference(intSet).ToString(): ", int2Set.Difference(intSet).ToString())
	fmt.Println("intSet.SymmetricDifference(int2Set).ToString(): ", intSet.SymmetricDifference(int2Set).ToString())
	fmt.Println("int2Set.SymmetricDifference(intSet).ToString(): ", int2Set.SymmetricDifference(intSet).ToString())

	sa, sb, sc, sd, scc := "a", "b", "c", "d", "c"
	stringSet := NewThreadSafeSet(sa, sc, sb, sd, scc)
	fmt.Println(stringSet.ToString())
	stringAddressSet := NewThreadSafeSet(&sa, &sc, &sb, &sd, &scc)
	fmt.Println(stringAddressSet.ToString())

	personSet := NewThreadSafeSet(
		Person{name: "Alise", age: 11},
		Person{name: "Bob", age: 33},
		Person{name: "John", age: 22},
		Person{name: "Nick", age: 44},
		Person{name: "Bob", age: 33},
	)
	fmt.Println(personSet.ToString())
	personAddressSet := NewThreadSafeSet(
		&Person{name: "Alise", age: 11},
		&Person{name: "Bob", age: 33},
		&Person{name: "John", age: 22},
		&Person{name: "Nick", age: 44},
		&Person{name: "Bob", age: 33},
	)
	fmt.Println(personAddressSet.ToString())

	personSet.Remove(Person{name: "John", age: 22})
	fmt.Println(personSet.ToString())
	personAddressSet.Remove(&Person{name: "John", age: 22})
	fmt.Println(personAddressSet.ToString())
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
	fmt.Printf("personSetFound %+v in %v\n", personSetFound, personSet.ToString())
	personSet.Remove(personSetFound)
	fmt.Printf("After remove the personSetFound, %v\n", personSet.ToString())

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
	fmt.Printf("personAddressSetFound %+v in %v\n", personAddressSetFound, personAddressSet.ToString())
	personAddressSet.Remove(personAddressSetFound)
	fmt.Printf("After remove the personAddressSetFound, %v\n", personAddressSet.ToString())
}