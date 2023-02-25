package commonData

import (
	"ecomm-alpha/handler"
)

var (
	Seller handler.SellerSignUpDetails
	Buyer  handler.BuyerSignUpDetails
)

func init() {
	//change credentials according to your needs
	Seller.Email = "seller1@gmail.com"
	Seller.FullName = "seller asd"
	Seller.Password = "123456"

	Buyer.Email = "buyer1@gmail.com"
	Buyer.FullName = "buyer asd"
	Buyer.Password = "!Aa123456"

}
