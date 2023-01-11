package handler

import (
	"ecomm-alpha/models"
	"net/mail"
	"reflect"
	"regexp"
)

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validateSellerSignUpInput(ssd *SellerSignUpDetails) (bool, []string) {
	errorFields := []string{}

	if ssd.ConfirmPassword != ssd.Password {
		errorFields = append(errorFields, "confirmPassword")
	}
	if isEmail(ssd.Email) == false {
		errorFields = append(errorFields, "email")
	}
	if ssd.FullName == "" {
		errorFields = append(errorFields, "fullName")
	}
	if ssd.Password == "" || len(ssd.Password) < 6 {
		errorFields = append(errorFields, "password")
	}

	if len(errorFields) > 0 {
		return false, errorFields
	}
	return true, errorFields
}

func validateStoreInput(s *models.Store) (bool, []string) {
	errorFields := []string{}

	if s.Name == "" {
		errorFields = append(errorFields, "name")
	}
	if s.Description == "" {
		errorFields = append(errorFields, "description")
	}

	if len(errorFields) > 0 {
		return false, errorFields
	}
	return true, errorFields
}

func validateSellerLoginInput(sc *SellerCredentials) (bool, []string) {
	errorFields := []string{}

	if isEmail(sc.Email) == false {
		errorFields = append(errorFields, "email")
	}
	if sc.Password == "" {
		errorFields = append(errorFields, "password")
	}

	if len(errorFields) > 0 {
		return false, errorFields
	}
	return true, errorFields
}

func validateAddressInput(address *models.Address) (bool, []string) {
	errorFields := []string{}
	validNum := regexp.MustCompile("(0|91)?[6-9][0-9]{9}")
	if !validNum.MatchString(address.Mobile_No) {
		errorFields = append(errorFields, "mobile num")
	}
	if address.City == "" {
		errorFields = append(errorFields, "city")
	}
	if address.State == "" {
		errorFields = append(errorFields, "state")
	}
	if address.Country_Code == "" {
		errorFields = append(errorFields, "country code")
	}
	if address.Country == "" {
		errorFields = append(errorFields, "country")
	}
	if address.Address == "" {
		errorFields = append(errorFields, "address")
	}

	if len(errorFields) > 0 {
		return false, errorFields
	}
	return true, errorFields
}
func GetStatusCodeFromError(err error) int {
	var statusCode int
	val := reflect.ValueOf(err).Elem()
	statusCode = -1
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		fieldValue := val.Field(i)
		if fieldName == "Code" {
			v := fieldValue.Int()
			statusCode = int(v)
			break
		}
	}
	if statusCode == -1 {
		statusCode = 400
	}
	return statusCode
}
