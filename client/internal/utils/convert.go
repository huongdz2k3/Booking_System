package utils

import (
	"customer/internal/logger"
	"strconv"
)

func ConvertStringToInt(input string) (*int, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		logger.NewLogger().Error("Input is not number")
		return nil, err
	}
	return &i, nil
}
