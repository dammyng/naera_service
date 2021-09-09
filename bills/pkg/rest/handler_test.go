package rest_test

import (
	"bills/billsgrpc"
	"bills/config"
	"bills/internals/db"
	"bills/models/v1"
	"bills/pkg"
	"bills/pkg/router"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"shared/amqp/sender"
	"testing"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"gopkg.in/stretchr/testify.v1/require"
)

var TestBills pkg.NaeraBill

func TestMain(m *testing.M) {
	os.Setenv("Environment", "test")
	os.Setenv("CMD_PATH", "/Users/kd/src/naera/naera_service/bills/cmd/")
	os.Setenv("test_token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjRhYzJjZmQzLWVhMjctNDBkMC05MDA2LWQ1MmEyMGEzYzkwZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYyMTEzMjQyNCwidXNlcl9pZCI6ImVkNDFmYmIzLWY5MzItNGJiOC1hNjIzLTAxOWNhNGM3NGNmNyJ9.dvaI0BCX6OHTSr63Ol0caXKIKnxhdMysuflH324o5Gg")
	GRPC_PORT := "0.0.0.0:9999"

	env := config.NewApConfig()

	dbLayer := db.NewSqlLayer(env.DSN)

	grpcServer := grpc.NewServer()
	naeraBill := billsgrpc.NewNaeraBillsRpcServer(dbLayer)
	models.RegisterNaeraBillingServiceServer(grpcServer, naeraBill)
	lis, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting GRPC Server for tests on port %v", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	grpcClient, err := billsgrpc.NewNaeraRPClient(GRPC_PORT)
	if err != nil {
		log.Panicln("Test setup failed")
	}

		//AMQP
		conn, err := amqp.Dial(env.AmqpBroker)
		if err != nil {
			log.Panicln("Test setup failed -- " + err.Error())

		}
		eventEmitter, err := sender.NewAmqpEventEmitter(conn, "NaeraExchange")
		if err != nil {
			log.Panicln("Test setup failed")

		}
	
	testRouter := router.InitServiceRouter(grpcClient, eventEmitter)

	TestBills.Router = testRouter

	code := m.Run()
	grpcClient.Conn.Close()
	grpcServer.Stop()
	lis.Close()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	TestBills.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponse(t *testing.T, expected int, actual *httptest.ResponseRecorder) {
	require.Equal(t, expected, actual.Code , actual.Body.String())
	//require.Equal(t, actual.Code, expected)
}
