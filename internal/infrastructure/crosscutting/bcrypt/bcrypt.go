package bcrypt

import (
	"authsvc/internal/core/hasher"

	"github.com/pkg/errors"
	cryptoBcrypt "golang.org/x/crypto/bcrypt"
)

type implementation struct{}

func New() hasher.Hasher {
	return &implementation{}
}

const cost = 14

func (i *implementation) Hash(password string) (string, error) {
	data, err := cryptoBcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", errors.Wrap(err, "failed to encrypt password")
	}

	return string(data), nil
}

func (i *implementation) Compare(candidate, hash string) bool {
	err := cryptoBcrypt.CompareHashAndPassword([]byte(hash), []byte(candidate))

	return nil == err
}
