package main

import (
	"fmt" 
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"net/http"
	"github.com/go-chi/cors"

)

func  main ()  {
	

	godotenv.Load(".env")


	portString := os.Getenv("PORT")

	if(portString == ""){
		log.Fatal("PORT is not found in the enviroment")
	}

	router :=chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins :	 []string{"https://*" , "http://*"} ,
		AllowedMethods :	 []string{"GET" , "POST" , "PUT" , "DELETE" , "OPTIONS"} ,
		AllowedHeaders :	 []string{"*"} ,
		ExposedHeaders :	 []string{"Link"} ,
		AllowCredentials:  false ,
		MaxAge              : 300  ,

	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/" , handlerReadiness)
	v1Router.Get("/err" , handlerErr)

	router.Mount("/v1" ,v1Router)
	
	srv := &http.Server{
		Handler : router ,
		Addr : ":" + portString ,  
	}
	
	
	err :=  srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("PORT : " , portString)



}