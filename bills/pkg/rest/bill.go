package rest

import (
	"bills/models/v1"
	"bills/pkg/helpers"
	"bills/pkg/restclient"
	"bills/pkg/services"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

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

	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var u models.Bill
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var opts []grpc.CallOption

	bill := &models.Bill{
		Biller:          u.Biller,
		Cart:            u.Cart,
		Reoccurring:     u.Reoccurring,
		NextPaymentDate: u.NextPaymentDate,
		Active:          u.Active,
		PayingWith:      u.PayingWith,
		Title:           u.Title,
		Id:              uuid.NewV4().String(),
		By:              access.UserId,
		UpdatedAt:       time.Now().Unix(),
		CreatedAt:       time.Now().Unix(),
		LastPaymentDate: time.Now().Unix(),
	}
	tRes, err := handler.GrpcPlug.CreateBill(r.Context(), bill, opts...)
	if err != nil {
		err = errors.New("Error creating the bill record")
	}
	respondWithJSON(w, http.StatusCreated, tRes.Id)
}

func (handler *BillHandler) PayForBill(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var u models.Bill
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var opts []grpc.CallOption
	params := mux.Vars(r)
	billID := params["bill_id"]

	bill, err := handler.GrpcPlug.FindBill(r.Context(), &models.Bill{
		Id: billID,
	}, opts...)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}

	key := os.Getenv("FL_SECRETKEY_LIVE")
	// Verify transaction
	data, err := restclient.VerifyFwTransaction(key, u.TransactionId)
	if err != nil {
		err = errors.New(err.Error())
		respondWithError(w, http.StatusBadGateway, err.Error()+" -- "+key+" -- "+u.TransactionId)
		return
	}

	transaction := models.Transaction{
		Id:        uuid.NewV4().String(),
		Biller:    access.UserId,
		Title:     bill.Title,
		Amount:    float32(data.Data.Amount),
		TransRef:  u.TransactionId,
		Bill:      bill.Id,
		CreatedAt: time.Now().Unix(),
	}
	tRes, err := handler.GrpcPlug.CreateTransaction(r.Context(), &transaction, opts...)
	if err != nil {
		err = errors.New(err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error()+" -- "+key+" -- "+u.TransactionId)
	}

	var items []models.InCartItem
	if err := json.Unmarshal([]byte(bill.Cart), &items); err != nil {
		panic(err)
	}

	fatalErrors := make(chan error, len(items))
	orderErrors := make(chan error, len(items))
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
			_request, _ := json.Marshal(&request)
		

			_, err := restclient.ServiceTransaction(key, string(_request))
			order := &models.Order{
				TransactionId: transaction.Id,
				CreatedAt:     time.Now().Unix(),
				Amount:        float32(data.Data.Amount),
				Id:            uuid.NewV4().String(),
				Title:         request.Type +" "+ request.Customer,
				Charged:       true,
				Fulfilled:     true,
			}

			if err != nil {
				fatalErrors <- errors.New("We are unable to service your transaction for - " + err.Error() + request.Type + "_" + request.Customer)
				order.Fulfilled = false
			}
			_, err = handler.GrpcPlug.CreateOrder(r.Context(), order, opts...)
			if err != nil {
				orderErrors <- errors.New("Your bills were serviced but we could not crete your order record because - " + err.Error() + request.Type + "_ for _" + request.Customer)
				err = errors.New(err.Error())
				respondWithError(w, http.StatusInternalServerError, err.Error()+" -- "+key+" -- "+u.TransactionId)
			}
			wg.Done()
		}(v)

	}

	wg.Wait()
	close(fatalErrors)
	var total string

	resErrors := ""

	for msg := range fatalErrors {
		resErrors += fmt.Sprintf("%v, ", msg)
	}
	for msg := range orderErrors {
		resErrors += fmt.Sprintf("%v, ", msg)
	}

	if resErrors != "" {
		respondWithJSON(w, http.StatusBadRequest, total)
		return
	}
	respondWithJSON(w, http.StatusCreated, tRes.Id)
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

func (handler *BillHandler) VerifyCart(w http.ResponseWriter, r *http.Request) {

	var u models.Bill
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var items []models.InCartItem
	if err := json.Unmarshal([]byte(bill.Cart), &items); err != nil {
		panic(err)
	}

	res := services.ValidateOrderItem()
}



/**func createUnservedTransaction(bill *models.Bill, data *restclient.FlwVerifiedTransaction) models.Transaction {
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
}**/
