package main

import "fmt"

func main(){
	fmt.Println("Welcome to the Map!")

	student_grades := make(map[string]int)

	student_grades["Sid"] = 90
	student_grades["Ram"] = 20
	student_grades["Shyam"] = 30

	// fmt.Println(student_grades["Sid"])
	// fmt.Println(student_grades)

	// fmt.Println("Name Marks")
	// for key,value := range student_grades{
	// 	fmt.Printf("%s %d\n", key, value)
	// }

	// delete(student_grades, "Ram")

	// fmt.Println("Name Marks")
	// for key,value := range student_grades{
	// 	fmt.Printf("%s %d\n", key, value)
	// }

	// Marks, Exist := student_grades["Sid"]

	// fmt.Println("Grades of Sid", Marks)
	// fmt.Println("Sid exists: ", Exist)

	employee_data := map[string]int{
		"Alice": 10000,
		"Ram": 20000,
		"Shyam": 30000,
	}

	for name, salary := range employee_data{
		fmt.Printf("%s : %d\n", name, salary)
	}

	hostel_detail := map[string]int{
		"Sid": 308,
		"Raj": 410,
		"Kareen": 305,
	}

	for name, room := range hostel_detail{
		fmt.Printf("%s: %d\n", name, room)
	}
}