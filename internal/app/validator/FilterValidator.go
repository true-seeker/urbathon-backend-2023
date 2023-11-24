package validator

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"urbathon-backend-2023/pkg/config"
	"urbathon-backend-2023/pkg/errorHandler"
)

func ValidateAndReturnIntField(field, fieldName string) (int, *errorHandler.HttpErr) {
	intField, err := strconv.Atoi(field)
	if err != nil {
		return 0, errorHandler.New(fmt.Sprintf("%s must be integer", fieldName), http.StatusBadRequest)
	}
	return intField, nil
}

func ValidateAndReturnFloatField(field, fieldName string, precision int) (float64, *errorHandler.HttpErr) {
	floatField, err := strconv.ParseFloat(field, precision)
	if err != nil {
		return 0, errorHandler.New(fmt.Sprintf("%s must be float", fieldName), http.StatusBadRequest)
	}
	return floatField, nil
}

func ValidateAndReturnId(idStr, fieldName string) (int32, *errorHandler.HttpErr) {
	id, httpErr := ValidateAndReturnIntField(idStr, fieldName)
	if httpErr != nil {
		return 0, httpErr
	}

	if id <= 0 {
		return 0, errorHandler.New(fmt.Sprintf("%s must be greater than 0", fieldName), http.StatusBadRequest)
	}
	return int32(id), nil
}

func ValidateAndReturnDateTime(field, fieldName string) (*time.Time, *errorHandler.HttpErr) {
	date, err := time.Parse(config.DateTimeLayout, field)
	if err != nil {
		return nil, errorHandler.New(fmt.Sprintf("%s must be in ISO-8601 format", fieldName), http.StatusBadRequest)
	}
	return &date, nil
}

func ValidateAndReturnDate(field, fieldName string) (*time.Time, *errorHandler.HttpErr) {
	date, err := time.Parse(config.DateLayout, field)
	if err != nil {
		return nil, errorHandler.New(fmt.Sprintf("%s must be in ISO-8601 format", fieldName), http.StatusBadRequest)
	}
	return &date, nil
}
