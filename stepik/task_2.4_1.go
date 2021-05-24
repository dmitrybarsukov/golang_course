package main

import "fmt"

// Copy 2 stepik from here

type TestStruct struct {
	On bool
	Ammo int
	Power int
}

func (ts *TestStruct) Shoot() bool {
	if !ts.On || ts.Ammo <= 0 {
		return false
	} else {
		ts.Ammo -= 1
		return true
	}
}

func (ts *TestStruct) RideBike() bool {
	if !ts.On || ts.Power <= 0 {
		return false
	} else {
		ts.Power -= 1
		return true
	}
}

func main() {

    testStruct := new(TestStruct)

// Copy 2 stepik to here

	fmt.Println(testStruct)

}