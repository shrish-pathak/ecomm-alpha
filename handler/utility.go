package handler

import (
	"ecomm-alpha/models"
	"net/mail"
	"reflect"
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
