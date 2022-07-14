package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"contact.com/src/models"
	"log"
	"context"

)

//postea un programa dado su nombre y exportación draw en el endpoint http://localhost:3000/postProgram
func PostProgram (w http.ResponseWriter, r *http.Request){

	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)
	program := models.Program{
		Name : request["Name"],
		Data : request["Data"],
	}

	// inicialización la base de datos (por optimizar)
	var ctx = context.TODO()
	var collection *mongo.Collection

	clientOptions := options.Client().ApplyURI("mongodb+srv://luisDB:luisDB@cluster0.l1jjl.mongodb.net/test?authSource=admin&replicaSet=atlas-qqjtvi-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
        log.Fatal(err)
    }
  
    err = client.Ping(ctx, nil)
    if err != nil {
      log.Fatal(err)
    }else{
        fmt.Println("conección exitosa")
    }

	//

	//llamado a la colección
	collection = client.Database("proyectNodesBack").Collection("programs")
	//post de programa
    id, err := collection.InsertOne(ctx, program)

    if(err == nil){
        fmt.Println(id)
    }
	
	json.NewEncoder(w).Encode(program)
}

func GetPrograms (w http.ResponseWriter, r *http.Request){
	prog := models.Program{
		Name : "programaUno",
		Data : "dataProgramaUno",
	}

	progjosn, _ := json.Marshal(prog)
	progstring := string(progjosn)
	fmt.Println(progstring)
	json.NewEncoder(w).Encode(prog)
}

func DeleteProgram (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("se eliminó "))
}

