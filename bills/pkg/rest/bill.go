package rest

import (
	"bills/models/v1"
	"bills/pkg/helpers"
	"bills/pkg/restclient"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"net/http"
	"os"
	"time"
"fmt"
	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func (handler *BillHandler) CreateBill(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	var u models.Bill
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	var opts []grpc.CallOption

	bill := &models.Bill{
		Biller:          access.UserId,
		Cart:            u.Cart,
		Reoccurring:     u.Reoccurring,
		NextPaymentDate: u.NextPaymentDate,
		Active:          u.Active,
		PayingWith:      u.PayingWith,
		TransactionId:   u.TransactionId,
		Title:           u.Title,
		Id:              uuid.NewV4().String(),
		By:              access.UserId,
		UpdatedAt:       time.Now().Unix(),
		CreatedAt:       time.Now().Unix(),
		LastPaymentDate: time.Now().Unix(),
	}
	res, err := handler.GrpcPlug.CreateBill(r.Context(), bill, opts...)
	if err != nil {
		err = errors.New("Error creating the bill record")
	}

	key := os.Getenv("FL_SECRETKEY_LIVE")
	// Verify transaction
	data, err := restclient.VerifyFwTransaction(key, u.TransactionId)
	if err != nil {
		err = errors.New(err.Error())
		respondWithError(w, http.StatusBadGateway, err.Error()+" -- "+key+" -- "+u.TransactionId)
		return
	}
	var items []models.InCartItem
	if err := json.Unmarshal([]byte(bill.Cart), &items); err != nil {
		panic(err)
	}
	
	fatalErrors := make(chan error,len(items))
	var wg sync.WaitGroup
	wg.Add(len(items))
	for _, v := range items {

		go func(v models.InCartItem) {

			log.Printf("Request from cart Item %v", v)

			request := models.ServiceRequestPayload{
				Country:    "NG",
				Customer:   v.Beneficiary,
				Amount:     v.Amount,
				Recurrence: "ONCE",
				Type:       "AIRTIME",
				Reference:  v.ID,
			}
			_request, err := json.Marshal(&request)
			if err != nil {
				t := createUnservedTransaction(bill, data)
				_, err = handler.GrpcPlug.CreateTransaction(r.Context(), &t, opts...)
			}

			eer, err := restclient.ServiceTransaction(key, string(_request))
			transaction := &models.Transaction{
				Bill:      bill.Id,
				CreatedAt: time.Now().Unix(),
				TransRef:  data.Data.TxRef,
				Amount:    float32(data.Data.Amount),
				Id:        uuid.NewV4().String(),
				Biller:    bill.Biller,
				Title:     bill.Title,
				Charged:   true,
				Served:    true,
			}
			log.Printf("Service payment response for %v is %v, %v:  ", v.ID, eer, err)
			if err != nil {
				fatalErrors <- errors.New("We are unable to service your transaction for - " + err.Error() + request.Type + "_" + request.Customer) 
				transaction.Served = false
			}
			_, err = handler.GrpcPlug.CreateTransaction(r.Context(), transaction, opts...)
			if err != nil {
			}
			wg.Done()
		}(v)

	}

	wg.Wait()
	close(fatalErrors)
	var total string

	resErrors := ""

	for msg := range fatalErrors {
		resErrors += fmt.Sprintf("%v, ",msg)
    }

	if resErrors != "" {
		respondWithJSON(w, http.StatusBadRequest, total)
		return
	}
	respondWithJSON(w, http.StatusCreated, res.Id)
}

func (handler *BillHandler) UpdateBill(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}
	params := mux.Vars(r)
	billID := params["bill_id"]
	var u models.Bill
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var opts []grpc.CallOption
	_, err = helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	bill, err := handler.GrpcPlug.FindBill(r.Context(), &models.Bill{Id: billID}, opts...)

	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, BillNotFound)
			return
		}
		respondWithError(w, http.StatusBadRequest, err.Error()+"---"+InternalServerError)
		return
	}

	if u.Active != bill.Active && u.Active == true {
		respondWithError(w, http.StatusBadRequest, "You cannot activate a bill zou are not paying for")
		return
	}

	_, err = handler.GrpcPlug.UpdateBill(r.Context(), &models.UpdateBillRequest{Old: bill, New: &u}, opts...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, InternalServerError)
		return
	}
	respondWithJSON(w, http.StatusCreated, "")

}

func (handler *BillHandler) DeleteBill(w http.ResponseWriter, r *http.Request) {

}

func createUnservedTransaction(bill *models.Bill, data *restclient.FlwVerifiedTransaction) models.Transaction {
	return models.Transaction{
		Bill:      bill.Id,
		CreatedAt: time.Now().Unix(),
		TransRef:  data.Data.TxRef,
		Amount:    float32(data.Data.Amount),
		Id:        uuid.NewV4().String(),
		Biller:    bill.Biller,
		Title:     bill.Title,
		Charged:   true,
		Served:    false,
	}
}
