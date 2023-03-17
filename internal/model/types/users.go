package types

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

type UserID uuid.UUID

func (u UserID) IsNil() bool {
	return uuid.UUID(u) == uuid.Nil
}

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

func (u *UserID) Scan(src interface{}) error {
	rawUuid := uuid.UUID{}
	err := rawUuid.Scan(src)
	if err != nil {
		return err
	}

	userUuid := MakeUserIDByUUID(rawUuid)
	*u = userUuid

	return nil
}

func (u UserID) Value() (driver.Value, error) {
	return u.String(), nil
}

func MakeUserIDByUUID(uuid uuid.UUID) UserID {
	return UserID(uuid)
}

var NilUserID = UserID(uuid.Nil)
