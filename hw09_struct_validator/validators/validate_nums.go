package validators

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateIfInt(num int, params map[string]string) error {
	val, ok := params[ValidateParamMin]
	if ok {
		min, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		if err = validateNumMin(num, min); err != nil {
			return err
		}
	}

	val, ok = params[ValidateParamMax]
	if ok {
		max, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		if err = validateNumMax(num, max); err != nil {
			return err
		}
	}

	val, ok = params[ValidateParamIn]
	if ok {
		if err := validateNumIn(num, val); err != nil {
			return err
		}
	}

	return nil
}

func ValidateIfSliceInt(nums []int, params map[string]string) error {
	val, ok := params[ValidateParamMin]
	if ok {
		min, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		for _, num := range nums {
			if err = validateNumMin(num, min); err != nil {
				return err
			}
		}
	}

	val, ok = params[ValidateParamMax]
	if ok {
		max, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		for _, num := range nums {
			if err = validateNumMax(num, max); err != nil {
				return err
			}
		}
	}

	val, ok = params[ValidateParamIn]
	if ok {
		for _, num := range nums {
			if err := validateNumIn(num, val); err != nil {
				return err
			}
		}
	}

	return nil
}

func validateNumMin(num int, min int) error {
	if num < min {
		return fmt.Errorf("minimum value is %d", min)
	}
	return nil
}

func validateNumMax(num int, max int) error {
	if num > max {
		return fmt.Errorf("maximum value is %d", max)
	}
	return nil
}

func validateNumIn(num int, paramIn string) error {
	inValues := strings.Split(paramIn, ",")

	for _, in := range inValues {
		inNum, err := strconv.Atoi(in)
		if err != nil {
			return err
		}
		if num == inNum {
			return nil
		}
	}

	return fmt.Errorf("number not in %s", strings.Join(inValues, ", "))
}
