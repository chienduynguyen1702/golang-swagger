package controllers

var students = []Student{
	{
		ID:   1,
		Name: "John Doe",
		Age:  20,
	},
	{
		ID:   2,
		Name: "Jane Doe",
		Age:  21,
	},
	{
		ID:   3,
		Name: "John Smith",
		Age:  22,
	},
	{
		ID:   4,
		Name: "Emily Johnson",
		Age:  19,
	},
	{
		ID:   5,
		Name: "Michael Brown",
		Age:  20,
	},
	{
		ID:   6,
		Name: "Emma Wilson",
		Age:  21,
	},
}
var lectures = []Teacher{
	{
		ID:   1,
		Name: "Mr. Pep Guardiola",
	},

	{
		ID:   2,
		Name: "Mr. Jurgen Klopp",
	},
	{
		ID:   3,
		Name: "Mr. Thomas Tuchel",
	},
	{
		ID:   4,
		Name: "Mr. Ole Gunnar Solskjaer",
	},
	{
		ID:   5,
		Name: "Mr. Mikel Arteta",
	},
	{
		ID:   6,
		Name: "Mr. Frank Lampard",
	},
}

var courses = []Course{
	{
		ID:       1,
		Name:     "Computer Science",
		Lectures: lectures,
		Semester: "Spring 2021",
		Students: students,
	},
	{
		ID:       2,
		Name:     "Mathematics",
		Lectures: lectures,
		Semester: "Spring 2021",
		Students: students,
	},
	{
		ID:       3,
		Name:     "Physics",
		Lectures: lectures,
		Semester: "Spring 2021",
		Students: students,
	},
	{
		ID:       4,
		Name:     "Chemistry",
		Lectures: lectures,
		Semester: "Spring 2021",
		Students: students,
	},
	{
		ID:       5,
		Name:     "Biology",
		Lectures: lectures,
		Semester: "Spring 2021",
		Students: students,
	},
}
