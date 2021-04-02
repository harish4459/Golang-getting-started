package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Ford",
		LastName:  "Dent",
	}

	u2 := User{
		ID:        2,
		FirstName: "Ford",
		LastName:  "perfect",
	}
	if u1 == u2 {
		println("Same user")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user")
	} else {
		println("Different user")
	}

}
