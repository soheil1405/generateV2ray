package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func respondWithError(w http.ResponseWriter , code int  ,msg string){
	if code > 499{
		log.Println("Response with 5XX error : " , msg)
		type errResponse struct {
			Error string `json : "error"`
		}

		respondWithJson( w , code , errResponse{
			Error : msg ,
		})
	}
}

func respondWithJson(w http.ResponseWriter , code int  , payload interface {}){
	data , err :=json.Marshal(payload)

	if err != nil{
		log.Print("failed  to marshal Json Response   ", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Contetn-type" , "application/json")

	w.WriteHeader(200)
	w.Write(data)


}	