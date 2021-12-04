package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// type MovieInterface interface {
// 	getTitle() string
// 	getRating() string
// 	getBoxOffice() float32
// 	setTitle(title string)
// 	setRating(rating string)
// 	setBoxOffice(box float32)
// 	NewMovie(title, rating string, boxOffice float32)
// }

// type Movie struct {
// 	title string
// 	rating string
// 	boxOffice float32
// }

// func (m *Movie) NewMovie(title, rating string, boxoffice float32) {
// 	m.title = title
// 	m.boxOffice  = boxoffice
// 	m.rating = rating
// }

// func (m *Movie) getTitle() string {
// 	return m.title
// }
// func (m *Movie) getRating() string {
// 	return m.rating
// }

// func (m *Movie) getBoxOffice() float32 {
// 	return m.boxOffice
// }

// func (m *Movie) setTitle(title string)  {
// 	 m.title = title
// }
// func (m *Movie) setRating(rating string)  {
// 	 m.rating = rating
// }

// func (m *Movie) setBoxOffice(box float32)  {
// 	 m.boxOffice = box
// }

// type empl struct {
// 	id       int
// 	fname    string
// 	lastname string
// }

// type cust struct {
// 	id      int
// 	fname   string
// 	lname   string
// 	company string
// }

func Reflections() {
	// ncust1 := cust{
	// 	1,
	// 	"ASDasd",
	// 	"3123123",
	// 	"dasdasdas",
	// }

	// ne1 := empl{
	// 	123,
	// 	"ASdasd",
	// 	"czxcz",
	// }

	timed := MakeTimedfunc(propertitle).(func(string) string)
	
	newtitle := timed("testing title")
	fmt.Println(newtitle)

	// addPerson(ncust1)
	// addPerson(ne1)

}

func propertitle(s string) string {
	return strings.ToUpper(s)
}

func MakeTimedfunc(f interface{}) interface{}{
	returnfunc := reflect.TypeOf(f)

	if returnfunc.Kind() != reflect.Func {
		panic("expects a function")
	}

	valuef := reflect.ValueOf(f)
	wrapperf := reflect.MakeFunc(returnfunc, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()
		outp := valuef.Call(args)
		end := time.Now()

		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(valuef.Pointer()).Name(), end.Sub(start))
		return outp
	})
	return wrapperf.Interface()

}
// func addPerson(p interface{}) bool {
// 	if reflect.ValueOf(p).Kind() == reflect.Struct {
// 		v := reflect.ValueOf(p)

// 		switch reflect.TypeOf(p).Name() {
// 		case "empl":
// 			fmt.Printf("empl struct: %v\n", v)
// 		case "cust":
// 			fmt.Printf("cust struct: %v\n", v)
// 		}
// 	}
// 	return false
// }


