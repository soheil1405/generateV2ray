package main

import (
	"net/http"  
	// "fmt"
)

func handlerErr(w http.ResponseWriter , r *http.Request) {
	 
	// fmt.Println("asdasd")
	respondWithError(w , 400 , "smthing went wrong")

}