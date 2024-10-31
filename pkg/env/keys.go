package env

import "os"

const (
	Secret Key = "SECRET"
)

type Key string

func (key Key) OrNull() *string {
	return GetOrNull(string(key))
}

func (key Key) Get() string {
	return os.Getenv(string(key))
}

func (key Key) Or(def string) string {
	return GetDefault(string(key), def)
}

func (key Key) MustGet() string {
	value := key.Get()
	if value == "" {
		panic(key.String() + " property has no value.")
	}
	return value
}

func (key Key) String() string {
	return string(key)
}
