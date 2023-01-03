package commondata

import (
	"ecomm-alpha/handler"
)

var (
	Seller handler.SellerSignUpDetails
)

func init() {
	Seller.Email = "harry@gmail.com"
	Seller.FullName = "harry potter"
	Seller.Password = "expecto patronum"
	Seller.ConfirmPassword = "expecto patronum"
}
