package data

//Student adalah berisi ID, Nama dan Grade
type Student struct {
	ID    string
	Name  string
	Grade int
}

//DData merupakan data yang berisi strcut
var DData = []Student{
	Student{"E001", "ethan", 21},
	Student{"W001", "wick", 22},
	Student{"B001", "bourne", 23},
	Student{"B002", "bond", 23},
}
