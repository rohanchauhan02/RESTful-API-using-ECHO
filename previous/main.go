package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }
type login int
type welcome int

// func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	// fmt.Fprintf(w,"Request : %+v\n",r)
// 	// if r.Method=="GET"{
// 	// 	fmt.Fprintln(w,"Thank You!")
// 	// }
// 	switch r.Method{
// 	case "GET":{
// 		fmt.Fprintln(w,"Using Get")
// 	}
// 	case "POST":{
// 		fmt.Fprintln(w,"Using Post")
// 	}
// 	}
// 	fmt.Fprintln(w,"on login page")
// }
// func (wl welcome) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln(w,"on welcome page")
// }
func myLogin(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		fmt.Fprintln(w,"Using get for login endpoint")
	}
	fmt.Fprintln(w,"on login page")
}
func myWelcome(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"on welcome page")
}

func main() {
	http.HandleFunc("/welcome",myWelcome)
	http.HandleFunc("/login",myLogin)

	//Another way of doing this

	// http.Handle("/welcome",http.HandlerFunc(myWelcome))
	// http.Handle("/login",http.HandlerFunc(myLogin))

	//Another way of doing this

	// var i login
	// var j welcome
	// http.Handle("/login",i)
	// http.Handle("/welcome",j)

	fmt.Println("Listning on port 8000")
	http.ListenAndServe("localhost:8000",nil)
}