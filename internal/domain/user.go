package domain

import libuser "github.com/yeencloud/lib-user"

type User libuser.User

type AuthInformation struct {
	ID             string
	HashedPassword string
}
