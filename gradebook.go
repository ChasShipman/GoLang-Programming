/*  CS424 Summer 2019
Programming Assignment 1
Charles Shipman
Jul 8th, 2019

This program opens a file from user input
reads the file and adds student data to a set
This program sorts students by last name and will calulate
each student's test average, homework average and grade average

This program will print to the console the statistics for the class
including how many students/the weight of test and homework/overall class average
and then the statistics for each student
including last, first names/ tests average (number of tests),
homework average (number of homeworks)
grade average.

If a student has fewer than the max number of homework it will
be indicated
*/

package main

import ( //import libraries
	"bufio"   //used for io
	"fmt"     //used for printing
	"os"      //used for os function commands
	"sort"    //used to interface sort
	"strconv" //used to convert strings to int and float64
	"strings" //used to split strings
)

type Student struct { //struct student type
	firstName      string  //first name field of student
	lastName       string  //last name field of student
	testGrades     []int   //array of int test values
	homeworkGrades []int   //array of int homework values
	testAvg        float64 //float test avg
	homeworkAvg    float64 //float homework avg
	gradeAvg       float64 //float grade average
}

func (s Student) NewStudent(firstName string, lastName string) Student { //create new student object and initializae first and last name
	s.firstName = firstName
	s.lastName = lastName
	return s //return the object
}

func (s Student) GetFirstName() string { //return firstname of student
	return s.firstName
}

func (s Student) GetLasttName() string { //return lastname of student
	return s.lastName
}

func (s Student) GetTestGrades() []int { //return list of test grades
	return s.testGrades
}

func (s Student) GetHomeworkGrades() []int { //return list of homework grades
	return s.homeworkGrades
}

func (s Student) setTestGrades(tests []int) Student { //set tests grades and return the object
	s.testGrades = tests
	return s
}

func (s Student) setHomeworkGrades(homeworks []int) Student { //set homework grades and return the object
	s.homeworkGrades = homeworks
	return s
}

func (s Student) GetTestAvg() float64 { //calcualate and return test average
	var sum float64
	for i := 0; i < len(s.GetTestGrades()); i++ { //step through objects list of scores
		sum += float64(s.testGrades[i]) //sum scroes
	}
	s.testAvg = sum / float64(len(s.GetTestGrades())) //divide by number of scores
	return s.testAvg
}

func (s Student) GetHomeworkAvg() float64 { //calculate and return homework average
	var sum float64
	for i := 0; i < len(s.GetHomeworkGrades()); i++ { //step through objects list of scores
		sum += float64(s.homeworkGrades[i]) //sum scores
	}
	s.homeworkAvg = sum / float64(len(s.GetHomeworkGrades())) //divide by number of scores
	return s.homeworkAvg
}

func (s Student) GetGradeAvg() float64 { //calculate and return overall student average
	var homeworkAvg = s.GetHomeworkAvg() //get homework and test avgs
	var testAvg = s.GetTestAvg()
	homeworkAvg *= homework_weight //apply homework and test weights
	testAvg *= test_weight
	s.gradeAvg = (homeworkAvg + testAvg) / 100.0 //add and divide by 100
	return s.gradeAvg
}

/*following used to implement sort interface*/
type ByLastName []Student //allows array of students to implement sort interface

func (ln ByLastName) Len() int { //implements length function of sort interface
	return len(ln)
}

func (ln ByLastName) Less(i, j int) bool { //implements less function of sort interface
	if ln[i].lastName == ln[j].lastName { //if last name same, sort by first name
		return ln[i].firstName < ln[j].firstName
	}
	return ln[i].lastName < ln[j].lastName //sort by last name
}

func (ln ByLastName) Swap(i, j int) { //implements swap function of sort interface
	ln[i], ln[j] = ln[j], ln[i] //swap i with j,  j with i
}

func findMaxAssignments(s []Student) { //find max number of homework assignments
	for i := 0; i < len(s); i++ { //step through array of students
		/*	if max_tests < len(s[i].GetTestGrades()) {
			max_tests = len(s[i].GetTestGrades())
		}*/
		if max_homework < len(s[i].GetHomeworkGrades()) { //if max is less than the current number of homework
			max_homework = len(s[i].GetHomeworkGrades()) //update max
		}
	}
}

func calculateOverallAverage(s []Student) { //calculate overall class averag
	for i := 0; i < len(s); i++ { //step through array of students
		overall_avg += s[i].GetGradeAvg() //sum each students averages together
	}
	overall_avg /= float64(len(s)) //divded by number of studets
}

func displayStatstics(s []Student) { //display class statistics
	fmt.Printf("GRADE REPORT --- %d STUDENTS FOUND IN FILE\n", len(s))
	fmt.Printf("TEST WEIGHT: %.1f%%\n", test_weight)
	fmt.Printf("HOMEWORK WEIGHT: %.1f%%\n", homework_weight)
	fmt.Printf("OVERALL AVERAGE is %.1f\n\n", overall_avg)
}

func displayStudents(s []Student) { //display statistics about each student
	fmt.Println("    STUDENT NAME\t:     TESTS      HOMEWORK     AVG")
	fmt.Println("------------------------------------------------------------")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%12v, ", s[i].GetLasttName())  //lastname
		fmt.Printf("%5v\t", s[i].GetFirstName())   //firstname
		fmt.Printf(":   %5.1f", s[i].GetTestAvg()) //test average
		fmt.Printf(" (")
		fmt.Printf("%v", len(s[i].GetTestGrades())) //number of tests
		fmt.Printf(")")
		fmt.Printf("\t %.1f", float64(s[i].GetHomeworkAvg())) //homework average
		fmt.Print(" (", len(s[i].GetHomeworkGrades()), ")")   //number of homeworks
		fmt.Printf("%9.1f", s[i].GetGradeAvg())               //grade avg
		fmt.Print("  ")
		if len(s[i].GetHomeworkGrades()) < max_homework { //check if students has less than max homework assignments
			fmt.Print("  ** fewer than max assignments **  ") //indicate if true
		}
		/*if len(s[i].GetTestGrades()) < max_tests {
			fmt.Print("  ** fewer than max tests **  ")
		}*/
		fmt.Println()
	}
}

//var max_tests int
//global variables used for class statistics
var max_homework int        // holds max num of homeworkd
var test_weight float64     //holds weight of tests
var homework_weight float64 //holds weight of homework
var overall_avg float64     //holds overall avg of class

//main function
func main() {
	var keyboard *bufio.Scanner // holds user input from stdin
	var newStu Student          //new student object
	var first string            //holds first name to initialize a student object
	var last string             //holds last name to initialize a student object
	var testSplice []int        //holds a splice of test scores from file
	var homeworkSplice []int    //holds a splice of homework scores from file
	var fileName string         //holds file name
	var inFile *os.File         //allows us to read through file
	var err error               //holds error return value
	var studentArray []Student  //splice to hold student objects
	var index int               //helper to index through studentArray

	//prompt user input
	fmt.Println("Welcome to the gradebook calculator test program.")
	fmt.Println("I am going to read students from an input data file.")
	fmt.Println("You will tell me the name of your input file.")
	keyboard = bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the name of your input file: ")
	keyboard.Scan()
	fileName = keyboard.Text()
	fmt.Print("Enter the % amount to weight test in overall avg: ")
	keyboard.Scan()
	test_weight, err = strconv.ParseFloat(keyboard.Text(), 64)
	homework_weight = 100.0 - test_weight
	fmt.Printf("Tests will be weighted %.1f%%, Homework weight %.1f%%\n", test_weight, homework_weight)

	//open file
	inFile, err = os.Open(fileName)
	if err != nil { //if error opening file
		fmt.Println("\nError: unable to open file: " + fileName)
		fmt.Println("Exiting program.")
		os.Exit(1) //exit on error
	}
	defer inFile.Close() //close file even if error

	//read file into splice
	var filelines = bufio.NewScanner(inFile) // lets us read through file line by line
	var line string                          //will hold a line of the file
	var fileSplice []string                  //will store all lines of the file
	for filelines.Scan() {                   //read until EOF
		line = filelines.Text()               //store line of file as string
		fileSplice = append(fileSplice, line) //add string to splice
	}

	/*at this point the whole file is read into a splice.
	  the splice can now be processed to appropriatly handle the data
	*/

	//extracting names from fileSplice
	//names will be used to created and initialize student objects
	//names always start at index 0 of file splice and a new name appears every 3 indices
	for i := 0; i < len(fileSplice); i += 3 { //step through filesplice and find only names
		nameCatcher := fileSplice[i]                 //catch and hold a name
		nameSplit := strings.Split(nameCatcher, " ") //split name at space into first and last names
		first = nameSplit[0]                         //hold first name
		last = nameSplit[1]                          //hold last name
		newStu = newStu.NewStudent(first, last)      //create and initialize a new student using first and last name
		studentArray = append(studentArray, newStu)  //add new student to studentArray splice
	}

	/*at this point all the names have
	been successfully extracted from the file splice
	and have been used to allocate a new student object for each student
	*/

	/*extracting test scores from fileSplice
	testscores will be saved to a splice
	and saved to each corresponding student
	testscores always start at index 1 of file splice
	and a new list of scores appears every 3 indices
	*/

	for i := 1; i < len(fileSplice); i += 3 { //step through filesplice and find only test scores
		testCatcher := fileSplice[i]                 //catch and hold a line of test scores
		testSplit := strings.Split(testCatcher, " ") //split the line into separate scores
		for j := 0; j < len(testSplit); j++ {        //convert scores from strings to ints
			scoreCatcher, err := strconv.Atoi(testSplit[j])
			if err != nil { //if error converting
				fmt.Printf("Error: Ascii value %v does not convert to integer\n", testSplit[j])
				fmt.Println("Exiting program.")
				os.Exit(1) //exit on error
			}
			testSplice = append(testSplice, scoreCatcher) //append integer score to testsplice
		}
		studentArray[index] = studentArray[index].setTestGrades(testSplice) //set student test scores to test splice
		index++                                                             //index through student array
		testSplice = nil                                                    //reset test splice so last student does not have everyones scores
	}

	/*at this point all the test scores have
	been successfully extracted from the file splice
	and each student object now holds there own test scores
	*/

	/*extracting home work scores from fileSplice
	home work scores will be saved to a splice
	and saved to each corresponding student
	home work cores always start at index 2 of file splice
	and a new list of scores appears every 3 indices
	*/

	index = 0                                 //reset index of studentArray to start at first student
	for i := 2; i < len(fileSplice); i += 3 { //step through filesplice and find only homework scores
		homeworkCatcher := fileSplice[i]                     //catch and hold a line of scores
		homeworkSplit := strings.Split(homeworkCatcher, " ") //split the line into separate scores
		for j := 0; j < len(homeworkSplit); j++ {            //convert scores from strings to ints
			scoreCatcher, err := strconv.Atoi(homeworkSplit[j])
			if err != nil { //if error converting
				fmt.Printf("Error: Ascii value %v does not convert to integer\n", homeworkSplit[j])
				fmt.Println("Exiting program.")
				os.Exit(1) //exit on error
			}
			homeworkSplice = append(homeworkSplice, scoreCatcher) //append integer score to homeworksplice
		}
		studentArray[index] = studentArray[index].setHomeworkGrades(homeworkSplice) //set student homework scores to homework splice
		index++                                                                     //index through student array
		homeworkSplice = nil                                                        //reset homework splice so last student does not have everyones scores
	}

	/*at this point all the home work scores have
	been successfully extracted from the file splice
	and each student object now holds there own homework scores
	now the student array can be sorted and have calculations performed on scores
	*/

	sort.Sort(ByLastName(studentArray))   //sort array
	findMaxAssignments(studentArray)      //find max number of homework assignments
	calculateOverallAverage(studentArray) //calculate overall class average
	displayStatstics(studentArray)        //display class statistics
	displayStudents(studentArray)         //display all students/statistics

}
