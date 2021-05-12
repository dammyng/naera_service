package services_test

import (
	"bills/config"
	"os"
	"testing"
)


func TestMain(m *testing.M) {
	os.Setenv("Environment", "test")
	os.Setenv("CMD_PATH", "/Users/kd/src/naera/naera_service/bills/cmd/")
	
	config.NewApConfig()

	code := m.Run()
	os.Exit(code)
}
