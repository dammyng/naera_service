package restclient

import (
	"fmt"
)

func flutterError(data interface{}) error  {
	return fmt.Errorf("%v", data )  
}