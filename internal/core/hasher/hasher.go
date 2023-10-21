package hasher

type Hasher interface {
	Hash(password string) (string, error)
	Compare(candidate, hash string) bool
}
