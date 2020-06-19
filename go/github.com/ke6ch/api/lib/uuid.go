package lib

import (
	"fmt"

	"github.com/google/uuid"
)

// CreateUUID create user token
func CreateUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return u.String()
}
