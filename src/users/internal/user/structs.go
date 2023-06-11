package user

type Employment struct {
	Title    string `json:"title"`
	KeySkill string `json:"key_skill"`
}

type Address struct {
	City          string      `json:"city"`
	StreeName     string      `json:"street_name"`
	StreetAddress string      `json:"street_address"`
	ZipCode       string      `json:"zip_code"`
	State         string      `json:"state"`
	Country       string      `json:"country"`
	Coordinates   Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type CreditCard struct {
	CCNumber string `json:"cc_number"`
}

type Subscription struct {
	Plan          string `json:"plan"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
	Term          string `json:"term"`
}

type User struct {
	Id                    int          `json:"id"`
	Uid                   string       `json:"uid"`
	Password              string       `json:"password"`
	FirstName             string       `json:"first_name"`
	LastName              string       `json:"last_name"`
	UserName              string       `json:"username"`
	Email                 string       `json:"email"`
	Avatar                string       `json:"avatar"`
	Gender                string       `json:"gender"`
	PhoneNumber           string       `json:"phone_number"`
	SocialInsuranceNumber string       `json:"social_insurance_number"`
	DateOfBirth           string       `json:"date_of_birth"`
	Employment            Employment   `json:"employment"`
	Address               Address      `json:"address"`
	CreditCard            CreditCard   `json:"credit_card"`
	Subscription          Subscription `json:"subscription"`
}
