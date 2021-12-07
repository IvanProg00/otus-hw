package validators

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ValidateIfString(v string, params map[string]string) error {
	val, ok := params[ValidateParamLen]
	if ok {
		length, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		if err := ValidateStringLength(v, length); err != nil {
			return err
		}
	}

	val, ok = params[ValidateParamRegexp]
	if ok {
		if err := ValidateStringRegexp(v, val); err != nil {
			return err
		}
	}

	val, ok = params[ValidateParamIn]
	if ok {
		if err := validateStringIn(v, val); err != nil {
			return err
		}
	}

	return nil
}

func ValidateIfSliceString(vals []string, params map[string]string) error {
	val, ok := params[ValidateParamLen]
	if ok {
		length, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		for _, v := range vals {
			if err := ValidateStringLength(v, length); err != nil {
				return err
			}
		}
	}

	val, ok = params[ValidateParamRegexp]
	if ok {
		for _, v := range vals {
			if err := ValidateStringRegexp(v, val); err != nil {
				return err
			}
		}
	}

	val, ok = params[ValidateParamIn]
	if ok {
		for _, v := range vals {
			if err := validateStringIn(v, val); err != nil {
				return err
			}
		}
	}

	return nil
}

func ValidateStringLength(s string, length int) error {
	if len(s) != length {
		return fmt.Errorf("length not equals %d", length)
	}
	return nil
}

func ValidateStringRegexp(s, pattern string) error {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if match := reg.MatchString(s); !match {
		return fmt.Errorf("value not match the pattern \"%s\"", pattern)
	}
	return nil
}

func validateStringIn(s, paramIn string) error {
	inValues := strings.Split(paramIn, ",")

	for _, in := range inValues {
		if s == in {
			return nil
		}
	}

	return fmt.Errorf("value not in %s", strings.Join(inValues, ", "))
}
