package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/config"
	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		Validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var Validate = validator.New()

func (v XValidator) ValidateStruct(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := Validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field()               // Export struct field name
			elem.Tag = MsgForTag(err.Tag(), err.Value()) // Export struct tag
			elem.Value = err.Value()                     // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func MsgForTag(tag string, data interface{}) string {
	if strings.Contains(tag, "invalid character") {
		tag = "invalid character"
	}

	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "ip":
		return "Invalid format ip address"
	case "min":
		return "This field below min character"
	case "max":
		return "This field reach max character"
	case "DateOnlyWithDash":
		return "This field format yyyy-mm-dd"
	case "DateLessThan":
		return "Date must less than " + config.NFS_DAYS + " days"
	case "gtfield":
		return "This field must be greater than " + fmt.Sprintf("%v", data)
	case "DateTimeWithDash":
		return "This field format yyyy-mm-dd hh:mm:ss"
	case "gt":
		return "this field must be greater than " + fmt.Sprintf("%v", data)
	case "lt":
		return "this field reach maximum 100.000"
	case "invalid character":
		return "this field contain invalid character"
	case "cannot unmarshal":
		return "this request is cannot be proceed."
	case "sql: no rows in result set":
		return "data not found " + fmt.Sprintf("%v", data)
	}

	return "This field required " + tag
}

func DateOnlyWithDash(fl validator.FieldLevel) bool {
	ISO8601DateRegexString := `^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
}

func DateTimeWithDash(fl validator.FieldLevel) bool {
	ISO8601DateTimeRegexString := `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$`
	ISO8601DateTimeRegex := regexp.MustCompile(ISO8601DateTimeRegexString)
	return ISO8601DateTimeRegex.MatchString(fl.Field().String())
}

func DateEndLessThanStart(fl validator.FieldLevel) bool {
	dates, err := time.Parse("2006-01-02 15:04:05", fl.Field().String())
	if err != nil {
		return false
	}

	nfs_days, err := strconv.ParseFloat(config.NFS_DAYS, 64)
	if err != nil {
		return false
	}

	diff := time.Now().Local().Sub(dates).Hours() / 24
	return diff < nfs_days
}

func DateLessThan(fl validator.FieldLevel) bool {
	dates, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	nfs_days, err := strconv.ParseFloat(config.NFS_DAYS, 64)
	if err != nil {
		return false
	}

	diff := time.Now().Local().Sub(dates).Hours() / 24
	return diff < nfs_days
}
