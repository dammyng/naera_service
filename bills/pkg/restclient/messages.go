package restclient

import (
	"errors"
	"fmt"
)

func flutterError(data interface{}) error  {
	return errors.New(fmt.Sprintf("%v", data ))  
}