package interfaces

import "time"

type MemStorage interface{
	Get(interface{}) (string, error)
	Set(key, value string, expires time.Duration) error
}