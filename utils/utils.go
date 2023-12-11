package utils

import (
	"github.com/google/uuid"
	"strings"
)

func NewID() string {
	id := uuid.New().String()
	id = strings.ReplaceAll(id, "-", "")
	return id
}
