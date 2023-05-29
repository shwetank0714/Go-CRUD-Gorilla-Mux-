package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"authorname"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake Database

var courses []Course

// middleware/helper functions - file
// method for model
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	//
	fmt.Println("Welcome to CRUD API using Gorilla Mux")
	seedFakeData()

	// routing
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course/create", createCourse).Methods("POST")
	router.HandleFunc("/course/update/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/delete/{id}", deleteOneCourse).Methods("DELETE")

	
	

	log.Fatal(http.ListenAndServe(":4000", router))
}

func seedFakeData()  {
	// seeding the Fake Data
	courses = append(courses, Course{
		CourseId: "23",
		CourseName: "C++",
		CoursePrice: 100,
		Author: &Author{
			FullName: "Shwetank",
			Website: "abc.go.in",
		},
	})

	courses = append(courses, Course{
		CourseId: "33",
		CourseName: "python",
		CoursePrice: 150,
		Author: &Author{
			FullName: "John",
			Website: "python.go.in",
		},
	})
}

// controllers - file
// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Weclome to Crud API using Go Muxrilla</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// getting the course through some ID passed in the request----
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting One Course")
	w.Header().Set("Content-Type", "application/json")
	// grab the ID from request
	params := mux.Vars(r)

	// loop through the courses
	// find the matching id , and Return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course available with given id")
}

// adding the Data to Fake DB

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty ---
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Data Not allowed as body")
	}

	// what if body is {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Empty Data {} Not allowed as body")
		return
	}

	// generate unique id, string
	// append the course into courses slice

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}


func updateCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("update One Course")
	w.Header().Set("Content-Type", "application/json")
	
	// grab id from request

	params := mux.Vars(r)

	// loop to get id
	// remove the id from slices
	// add the updated data with the same id that is coming in params

	for index, course := range(courses){
		if course.CourseId == params["id"] {
			// removing the course from the given index
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode("Course updated")
			return
		}
	}
	// Send a response when id is not found
	json.NewEncoder(w).Encode("Course Not found with the given id")
}


func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")
	
	// grab id from request

	params := mux.Vars(r)

	// loop to get id
	// remove the id from slices
	// add the updated data with the same id that is coming in params

	for index, course := range(courses){
		if course.CourseId == params["id"] {
			// removing the course from the given index
			courses = append(courses[:index], courses[index+1:]...)

			json.NewEncoder(w).Encode("Course Delete succesfully")
			return
		}
	}
	// Send a response when id is not found
	json.NewEncoder(w).Encode("Course Not found with the given id")
}