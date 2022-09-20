package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for corse - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

//model for author - file

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var courses []Course

//middleware, helper - file

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Building API in GO!")
	r := mux.NewRouter()
	//seeding of data

	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "ReactJS",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Unknown Author",
			Website:  "unknown.abc",
		},
	},
		Course{
			CourseId:    "2",
			CourseName:  "NodeJS",
			CoursePrice: 399,
			Author: &Author{
				Fullname: "Unknown Author",
				Website:  "unknown.abc",
			},
		},
		Course{
			CourseId:    "3",
			CourseName:  "MERN stack",
			CoursePrice: 499,
			Author: &Author{
				Fullname: "Unknown Author",
				Website:  "unknown.abc",
			},
		},
	)
	r.HandleFunc("/", serveHomeRoute).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	http.ListenAndServe("localhost:8080", r)

}

//controllers - file

func serveHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>/ route</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course by id")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses
	//find matching id
	//return response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No course found with given id : " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//what about {}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	for _, currentCourse := range courses {
		if currentCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course name already exists.")
			return
		}
	}

	//generate unique id which is in string
	//append new course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from request
	params := mux.Vars(r)

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	for _, currentCourse := range courses {
		if currentCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course name already exists.")
			return
		}
	}

	var index int = -1

	for i, c := range courses {
		if c.CourseId == params["id"] {
			index = i
		}
	}

	if index != -1 {
		course.CourseId = params["id"]
		courses[index] = course
		json.NewEncoder(w).Encode(course)
		return
	}

	//when course with given id is not present
	json.NewEncoder(w).Encode("Course doesn't exist with given id")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var index int = -1

	for i, course := range courses {
		if course.CourseId == params["id"] {
			index = i
			break
		}
	}

	if index != -1 {
		courses = append(courses[:index], courses[index+1:]...)
		json.NewEncoder(w).Encode("Course deleted with id : " + params["id"])
		return
	}

	//when course with given id is not present
	json.NewEncoder(w).Encode("Course doesn't exist with given id")
	return
}
