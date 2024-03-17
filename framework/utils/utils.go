package utils

import (
	"strconv"
)

func ParseStringToInt(v string) (float64, error) {
	res, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return res, err
	}
	return res, nil
}
