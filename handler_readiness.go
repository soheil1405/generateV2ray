package main

import (
	"net/http"  
	// "fmt"
)

func handlerReadiness(w http.ResponseWriter , r *http.Request) {
	 
	// fmt.Println("asdasd")
	respondWithJson(w , 200 , struct{}{})

}