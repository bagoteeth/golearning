package main

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	people := map[string]Person{
		"Alice":   Person{Name: "Alice", Age: 25, Email: "alice@example.com"},
		"Bob":     Person{Name: "Bob", Age: 30, Email: "bob@example.com"},
		"Charlie": Person{Name: "Charlie", Age: 35, Email: "charlie@example.com"},
	}
	for _, person := range people {
		person.Age += 1
		people[person.Name] = person
	}

}
