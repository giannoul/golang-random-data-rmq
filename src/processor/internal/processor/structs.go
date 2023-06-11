package processor

import (
	"igiannoulas/golang-microservices/src/internal/user"
	"igiannoulas/golang-microservices/src/internal/ccard"
)

type Combination struct {
	User user.User
	CreditCard ccard.CreditCard
}