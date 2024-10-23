package userhandler

import (
	"Backend-Go/dbhandler"
	"Backend-Go/model"

	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10" //add this
	//add this

	//add this
	//add this
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = dbhandler.GetCollection(dbhandler.DB, "users")
var validate = validator.New()

// signUpHandler handles the user signup
func SignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		// Decode the request body into the user struct
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newUser := model.User{
			Id:       primitive.NewObjectID(),
			Email:    user.Email,
			Password: user.Password,
		}

		// Insert user into the database using DB_handler
		collection := userCollection // Adjust collection name as needed
		_, err := collection.InsertOne(context.TODO(), newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond to the client
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User registered successfully"))
	}
}
