package pkg

import (
	"github.com/google/uuid"
)

func GenerateUId() string {
	newUUID:= uuid.New()
	return string(newUUID.String())
}