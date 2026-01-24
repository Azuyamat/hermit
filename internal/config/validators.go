package config

import "fmt"

func InRange(min, max int) func(int) error {
	return func(i int) error {
		if i < min || i > max {
			return fmt.Errorf("value must be between %d and %d", min, max)
		}
		return nil
	}
}

func NonEmpty(s string) error {
	if s == "" {
		return fmt.Errorf("string value must be non-empty")
	}
	return nil
}

func ValidPath(s string) error {
	if s == "" {
		return fmt.Errorf("path must be non-empty")
	}
	return nil
}

func And[T any](validators ...func(T) error) func(T) error {
	return func(v T) error {
		for _, validator := range validators {
			if err := validator(v); err != nil {
				return err
			}
		}
		return nil
	}
}

func Or[T any](validators ...func(T) error) func(T) error {
	return func(v T) error {
		var errMessages []string
		for _, validator := range validators {
			if err := validator(v); err != nil {
				errMessages = append(errMessages, err.Error())
			}
		}
		if len(errMessages) == len(validators) {
			return fmt.Errorf("all validators failed: %v", errMessages)
		}
		return nil
	}
}
