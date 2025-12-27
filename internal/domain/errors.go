package domain

import "github.com/yeencloud/lib-shared/apperr"

type DisabledRegistrationError struct {
}

func (e DisabledRegistrationError) Error() string {
	return "service registration is disabled"
}

func (e DisabledRegistrationError) Unwrap() error {
	return apperr.UnauthorizedError{}
}
