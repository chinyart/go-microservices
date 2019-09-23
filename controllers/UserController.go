package controllers

import (
	"net/http"
)

func CreatePerson(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")
	// var person models.Pe
	// _ = json.NewDecoder(request.Body).Decode(&person)
	// collection := client.Database("thepolyglotdeveloper").Collection("people")
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// result, _ := collection.InsertOne(ctx, person)
	// json.NewEncoder(response).Encode(result)
}
