package business

import(
	"fmt"
	"time"
	"encoding/json"
	"go-api/models"
)

func init(){

}

func GetUser() ([]byte){

	profile := models.Profile{"Developer", []string{"snowboarding", "programming"}}

 	user := models.User{"Vinicius", time.Date(1993, time.December, 21, 0, 0, 0, 0, time.UTC), 22, profile}

 	js, err := json.Marshal(user)

	if err != nil {
		fmt.Print(err)
	}

 	return js

}