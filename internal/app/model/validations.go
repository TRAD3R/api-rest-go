package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// requiredIf - выполнение валидации только при cond true
// необходимо для редактирования пользователя, когда пароль не меняется
func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
