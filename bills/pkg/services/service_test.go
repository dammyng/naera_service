package services_test

import (
	"bills/config"
	"os"
	"testing"
)


func TestMain(m *testing.M) {
	os.Setenv("Environment", "test")
	os.Setenv("CMD_PATH", "C:\\Users\\KD\\Desktop\\naera_service\\bills\\cmd\\")
	
	config.NewApConfig()

	code := m.Run()
	os.Exit(code)
}
