package config

import "fmt"

type Value[T any] struct {
	value     T
	validator func(T) error
}

func NewValue[T any](value T, validator func(T) error) *Value[T] {
	return &Value[T]{
		value:     value,
		validator: validator,
	}
}

func (v *Value[T]) Get() T {
	return v.value
}

func (v *Value[T]) Set(newValue T) error {
	if v.validator != nil {
		if err := v.validator(newValue); err != nil {
			return err
		}
	}
	v.value = newValue
	return nil
}

func (v *Value[T]) String() string {
	return fmt.Sprintf("%v", v.value)
}
