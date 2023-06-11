package user

import (
	"igiannoulas/golang-microservices/src/pkg/rda"
	"igiannoulas/golang-microservices/src/pkg/rabbitmq"
	"fmt"
	"encoding/json"
	"time"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func StartUsersMachine() {
    userRMQ := rabbitmq.NewRMQ(
        "guest",
        "guest",
        "rabbitmq",
        "5672",
        "user_ex",
        "users",
        "user",
    )
    go userRMQ.Initialize()
	time.Sleep(8 * time.Second)
	//go userRMQ.Read()

	ch := make(chan []byte)
	endpoint := "https://random-data-api.com/api/v2/users"
	rapi := rda.NewRandomAPI(endpoint)
	go rapi.GetDataWithRateLoop(1, 1 , ch)
	for ;; {
		select {
		case data := <-ch:
			var result User
			if err := json.Unmarshal(data, &result); err != nil {
				fmt.Println("Can not unmarshal JSON")
			}
			//fmt.Println(PrettyPrint(result))
			userRMQ.Send(data)
		}
	}

}