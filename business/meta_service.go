package business

import(
	"fmt"
	"encoding/json"
	"go-api/models"
)

func init(){

}

func GetInfo() ([]byte){

	meta := models.Meta{"Bem vindo ao Go!"}

	js, err := json.Marshal(meta)

	if err != nil {
		fmt.Print(err)
	}

 	return js

}