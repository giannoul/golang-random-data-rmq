package processor

import (
	"igiannoulas/golang-microservices/src/pkg/rabbitmq"
	"fmt"
	"encoding/json"
	"time"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func StartProcessorMachine() {
    userRMQ := rabbitmq.NewRMQ(
        "guest",
        "guest",
        "rabbitmq",
        "5672",
        "user_ex",
        "users",
        "user",
    )
    ccardRMQ := rabbitmq.NewRMQ(
        "guest",
        "guest",
        "rabbitmq",
        "5672",
        "ccard_ex",
        "ccards",
        "ccard",
    )

	userChannel := make(chan []byte)
	creditCardChannel := make(chan []byte)
	time.Sleep(10 * time.Second)
	go userRMQ.GetMessage(&userChannel)
	go ccardRMQ.GetMessage(&creditCardChannel)
	
	combo := Combination{}
	for ;; {
		
		select {
		case userData := <-userChannel:
			if err := json.Unmarshal(userData, &combo.User); err != nil {
				fmt.Println("Can not unmarshal user JSON")
			}
		case creditCardData := <-creditCardChannel:
			if len(combo.User.Uid) > 1 {
				if err := json.Unmarshal(creditCardData, &combo.CreditCard); err != nil {
					fmt.Println("Can not unmarshal user JSON")
				}
			}else{
				combo = Combination{}
			}

		}
		if len(combo.User.Uid) > 1 {
			fmt.Println("**********")
			fmt.Println(PrettyPrint(combo))
		}
	}
}