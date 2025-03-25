package utills

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenerateUniqueTag(address, zone, workplaceType string, floor, number int64) string {
	info := fmt.Sprintf("%s|%s|%s|%d|%d", address, zone, workplaceType, floor, number)
	hash := sha256.Sum256([]byte(info))
	id := base64.RawURLEncoding.EncodeToString(hash[:])

	return id
}
