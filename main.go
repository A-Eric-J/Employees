package main

import "fmt"

// minStartTime / with this function we find the earliest
// time start working of employees based on the visited array
func minStartTime(employees []Employee, visited []bool) int {
	minValue := 2050
	minIndex := -1

	for i := 0; i < len(employees); i++ {
		if !visited[i] && employees[i].startTime < minValue {
			minValue = employees[i].startTime
			minIndex = i
		}
	}

	return minIndex
}

// sortIndexes / we want to know the order index of sending
// employees based on start working time
func sortIndexes(employees []Employee, orderIndexes []int) {
	// this array is for knowing that we have sent the employee or not,
	// and you see that, at first the initialized values are false
	// it means that we have not sent the employees yet
	visited := []bool{false, false, false, false, false}

	for i := 0; i < len(employees); i++ {
		minIndex := minStartTime(employees, visited)
		orderIndexes[i] = minIndex
		visited[minIndex] = true
	}
}

// whichCityIsNear / this function works based on the table of distances in README.md
func whichCityIsNear(cityName string, departCityCapacity []int) string {
	switch cityName {
	case "Munich":
		if departCityCapacity[0] != 0 {
			departCityCapacity[0]--
			return "a"
		} else if departCityCapacity[1] != 0 {
			departCityCapacity[1]--
			return "b"
		} else if departCityCapacity[2] != 0 {
			departCityCapacity[2]--
			return "c"
		}
	case "Hamburg":
		if departCityCapacity[1] != 0 {
			departCityCapacity[1]--
			return "b"
		} else if departCityCapacity[2] != 0 {
			departCityCapacity[2]--
			return "c"
		} else if departCityCapacity[0] != 0 {
			departCityCapacity[0]--
			return "a"
		}
	case "Berlin":
		if departCityCapacity[0] != 0 {
			departCityCapacity[0]--
			return "a"
		} else if departCityCapacity[1] != 0 {
			departCityCapacity[1]--
			return "b"
		} else if departCityCapacity[2] != 0 {
			departCityCapacity[2]--
			return "c"
		}

	}
	return ""
}

// departedCityIndexing / this function works based on whichCityIsNear function and indexing departed cities
func departedCityIndexing(orderIndexes []int, employees []Employee, departCityCapacity []int) {

	for i := 0; i < len(employees); i++ {
		departedIndexing := whichCityIsNear(employees[orderIndexes[i]].city, departCityCapacity)
		employees[orderIndexes[i]].departCityName = departedIndexing
	}

}

func printResult(employees []Employee) {
	for i := 0; i < len(employees); i++ {
		fmt.Println(i+1, ")", "Name: ", employees[i].name, " City: ", employees[i].city, " StartTime: ", employees[i].startTime, " DepartCity: ", employees[i].departCityName)
	}
}

// Employee / we have some employees for sending on a mission
//
//	employees have some data like name, city of working, start working time in this company
//
// and the departed city for the mission
type Employee struct {
	name           string
	city           string
	startTime      int
	departCityName string
}

func main() {
	///  The information of employees
	employees := []Employee{{name: "Eric", city: "Berlin", startTime: 2012, departCityName: ""}, {name: "John", city: "Munich", startTime: 2010, departCityName: ""}, {name: "Nami", city: "Hamburg", startTime: 2009, departCityName: ""}, {name: "Anthony", city: "Hamburg", startTime: 2021, departCityName: ""}, {name: "Jeniffer", city: "Munich", startTime: 2009, departCityName: ""}}
	// We need to know the order of the sending the employees
	// to the departed city based on the start working time for the company
	// and the capacity of the departed city, so first the order of indexes
	// is 0 to 4 based on these 5 employees
	orderIndexes := []int{0, 1, 2, 3, 4} /// we want to have this order for sending them : 2,4,1,0,3 based  on the start working time for the company
	/// we have 3 cities (a , b, c) that are for departed mission and
	/// these cities have some capacities => a : 2 , b : 1 , c : 2
	departCityCapacity := []int{2, 1, 2}

	sortIndexes(employees, orderIndexes)

	departedCityIndexing(orderIndexes, employees, departCityCapacity)
	printResult(employees)

}
