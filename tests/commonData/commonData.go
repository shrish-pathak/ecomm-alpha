package commonData

import (
	"ecomm-alpha/handler"
)

var (
	Seller handler.SellerSignUpDetails
)

func init() {
	//change credentials according to your needs
	Seller.Email = "harry2@gmail.com"
	Seller.FullName = "harry potter2"
	Seller.Password = "123456"
	
}
