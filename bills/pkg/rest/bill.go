package rest

import (
	m "bills/models"
	"bills/models/v1"
	"bills/pkg/helpers"
	"bills/pkg/restclient"
	"bills/pkg/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shared/amqp/events"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func (handler *BillHandler) MyBills(w http.ResponseWriter, r *http.Request) {
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
	var opts []grpc.CallOption

	tRes, err := handler.GrpcPlug.GetBillerBills(r.Context(), &models.GetBillerBillsRequest{BillerID: access.UserId}, opts...)
	if err != nil {
		err = errors.New("Error creating the bill record")
	}
	respondWithJSON(w, http.StatusCreated, tRes)
}

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
		Biller:          access.UserId,
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
		respondWithError(w, http.StatusBadRequest, err.Error())

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

	if u.Active != bill.Active && u.Active {
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

func (handler *BillHandler) VerifyNewCart(w http.ResponseWriter, r *http.Request) {
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
	var opts []grpc.CallOption

	biller, err := handler.GrpcPlug.FindBiller(r.Context(), &models.Biller{Id: access.UserId}, opts...)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	var items []models.InCartItem
	var verifiedItems []m.DisplayVerifyBill
	if err := json.Unmarshal([]byte(biller.Cart), &items); err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(items))
	ItemCode := ""
	BillerCode := ""
	for _, v := range items {
		go func(v models.InCartItem) {
			if v.Category == "airtime" {
				ItemCode = "AT099"
				BillerCode = "BIL099"
			}
			order := m.OrderRequest{
				ItemCode:   ItemCode,
				Customer:   v.Beneficiary,
				BillerCode: BillerCode,
			}
			res, err := services.ValidateOrderItem(order)

			if err != nil {
				_item := m.DisplayVerifyBill{
					Id:          v.ID,
					Name:        v.Provider,
					Amount:      0,
					Title:       fmt.Sprintf("%s for %s", v.Provider, v.Beneficiary),
					Beneficiary: v.Beneficiary,
					Status:      err.Error(),
				}
				verifiedItems = append(verifiedItems, _item)
			} else {
				_item := m.DisplayVerifyBill{
					Id:          v.ID,
					Name:        "Airtime",
					Amount:      v.Amount,
					Title:       fmt.Sprintf("%s for %s", "Airtime", res.Data.Customer),
					Beneficiary: res.Data.Customer,
					Status:      res.Data.ResponseMessage,
				}
				verifiedItems = append(verifiedItems, _item)
			}

			wg.Done()
		}(v)
	}
	wg.Wait()
	respondWithJSON(w, http.StatusOK, verifiedItems)
}

func (handler *BillHandler) ChargeCard(w http.ResponseWriter, r *http.Request) {
	// Avoid corde errror
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	// Get Authorization token
	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Get Request body and parse to bill struct
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

	// Find the card referenced in the bill
	card, err := handler.GrpcPlug.FindCard(r.Context(), &models.Card{Id: u.CardId}, opts...)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not find complete transaction - Invalid card")
		return
	}

	// Creade the bill
	bill := &models.Bill{
		Biller:          access.UserId,
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
		CardId:          card.Id,
	}
	_, err = handler.GrpcPlug.CreateBill(r.Context(), bill, opts...)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not find complete transaction - Unable to create bill record")
		return
	}

	// Get list of items to be paid for
	var items []models.InCartItem
	if err := json.Unmarshal([]byte(u.Cart), &items); err != nil {
		panic(err)
	}
	// Get total sum to charge card
	totalAmount := 0.0
	for _, v := range items {
		totalAmount += v.Amount
	}

	payload := fmt.Sprintf(`{
		"token": "%s",
		"currency": "NGN",
		"country": "NG",
		"amount": %v,
		"email": "%s",
		"first_name": "%s",
		"last_name": "Customer",
		"narration": "%s",
		"tx_ref": "%s"
	}`, card.Token, totalAmount, card.Email, bill.Title, access.UserId, fmt.Sprintf("%s_%v", helpers.RandAlpha(5), time.Now().Unix()))
	vTran, err := restclient.ChargeCard(payload)
	if err != nil {
		err = errors.New(err.Error())
		respondWithError(w, http.StatusBadGateway, "Unable to charge card - "+err.Error())
		return
	}

	transaction := models.Transaction{
		Id:        uuid.NewV4().String(),
		Biller:    access.UserId,
		Title:     bill.Title,
		Amount:    float32(totalAmount),
		Bill:      bill.Id,
		CreatedAt: time.Now().Unix(),
		FlRef:     vTran.Data.TxRef,
		TransRef:  vTran.Data.TxRef,
	}
	tRes, err := handler.GrpcPlug.CreateTransaction(r.Context(), &transaction, opts...)
	if err != nil {
		err = errors.New(err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error()+" _Error creating transactiom")
	}

	cartEvent := []events.Event{}
	for _, v := range items {
		v.Transaction = tRes.Id
		cartEvent = append(cartEvent, v.CreateMsg())
	}

	for _, v := range cartEvent {
		handler.EventEmitter.Emit(v, "NaeraExchange")
	}

	// pass idem array to background

	respondWithJSON(w, http.StatusCreated, tRes.Id)

}

func (handler *BillHandler) ChargeLoan(w http.ResponseWriter, r *http.Request) {
	// Avoid corde errror
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	// Get Authorization token
	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Get Request body and parse to bill struct
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

	// Creade the bill
	bill := &models.Bill{
		Biller:          access.UserId,
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
		CardId:          "loan service",
	}
	_, err = handler.GrpcPlug.CreateBill(r.Context(), bill, opts...)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not find complete transaction - Unable to create bill record")
		return
	}

	// Get list of items to be paid for
	var items []models.InCartItem
	if err := json.Unmarshal([]byte(u.Cart), &items); err != nil {
		panic(err)
	}
	// Get total sum to charge card
	totalAmount := 0.0
	for _, v := range items {
		totalAmount += v.Amount
	}


	transaction := models.Transaction{
		Id:        uuid.NewV4().String(),
		Biller:    access.UserId,
		Title:     bill.Title,
		Amount:    float32(totalAmount),
		Bill:      bill.Id,
		CreatedAt: time.Now().Unix(),
		FlRef:     "loanID",
		TransRef:  "loanRef",
	}
	tRes, err := handler.GrpcPlug.CreateTransaction(r.Context(), &transaction, opts...)
	if err != nil {
		err = errors.New(err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error()+" _Error creating transactiom")
	}

	cartEvent := []events.Event{}
	for _, v := range items {
		v.Transaction = tRes.Id
		cartEvent = append(cartEvent, v.CreateMsg())
	}

	for _, v := range cartEvent {
		handler.EventEmitter.Emit(v, "NaeraExchange")
	}

	respondWithJSON(w, http.StatusCreated, tRes.Id)

}

func (handler *BillHandler) PayWithFL(w http.ResponseWriter, r *http.Request) {
	// Avoid cors error
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	// Get Authorization token
	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Get Request body and parse to bill struct
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

	// Find the card referenced in the bill
	verified, err := restclient.VerifyFwTransaction(u.CardId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not find complete transaction - Invalid transaction")
		return
	}
	// Creade the bill
	bill := &models.Bill{
		Biller:          access.UserId,
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
		CardId:          u.CardId,
	}
	_, err = handler.GrpcPlug.CreateBill(r.Context(), bill, opts...)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not find complete transaction - Unable to create bill record")
		return
	}

	// Get list of items to be paid for
	var items []models.InCartItem
	if err := json.Unmarshal([]byte(u.Cart), &items); err != nil {
		panic(err)
	}
	// Get total sum to charge card
	totalAmount := 0.0
	for _, v := range items {
		totalAmount += v.Amount
	}

	transaction := models.Transaction{
		Id:        uuid.NewV4().String(),
		Biller:    access.UserId,
		Title:     bill.Title,
		Amount:    float32(verified.Data.Amount),
		Bill:      bill.Id,
		CreatedAt: time.Now().Unix(),
		FlRef:     verified.Data.TxRef,
		TransRef:  verified.Data.TxRef,
	}
	tRes, err := handler.GrpcPlug.CreateTransaction(r.Context(), &transaction, opts...)
	if err != nil {
		err = errors.New(err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error()+" _Error creating transactiom")
	}

	cartEvent := []events.Event{}
	for _, v := range items {
		v.Transaction = tRes.Id
		cartEvent = append(cartEvent, v.CreateMsg())
	}

	for _, v := range cartEvent {
		handler.EventEmitter.Emit(v, "NaeraExchange")
	}
	respondWithJSON(w, http.StatusCreated, tRes.Id)
}

func (handler *BillHandler) BillTransactions(w http.ResponseWriter, r *http.Request) {
	// Avoid corde errror
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	// Get Authorization token
	_, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	params := mux.Vars(r)
	billID := params["bill_id"]

	var opts []grpc.CallOption

	transactions, err := handler.GrpcPlug.BillTransactions(r.Context(), &models.GetBillTransactionsRequest{BillID: billID}, opts...)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(transactions.Transactions) == 0 {
		respondWithJSON(w, http.StatusOK, make([]string, 0))
		return
	}

	respondWithJSON(w, http.StatusOK, transactions.Transactions)
}

func (handler *BillHandler) BillTransactionOrders(w http.ResponseWriter, r *http.Request) {
	// Avoid corde errror
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	// Get Authorization token
	_, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	params := mux.Vars(r)
	_ = params["bill_id"]
	transID := params["trans_id"]
	var opts []grpc.CallOption

	orders, err := handler.GrpcPlug.TransactionOrders(r.Context(), &models.GetTransactionOrdersRequest{TransactionID: transID}, opts...)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if len(orders.Orders) == 0 {
		respondWithJSON(w, http.StatusOK, make([]string, 0))
		return
	}

	respondWithJSON(w, http.StatusOK, orders.Orders)
}

func (handler *BillHandler) BillerTransactions(w http.ResponseWriter, r *http.Request) {
	// Avoid corde errror
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	// Get Authorization token
	access, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	var opts []grpc.CallOption

	transactioms, err := handler.GrpcPlug.BillerTransactions(r.Context(), &models.GetBillerTransactionsRequest{BillerID: access.UserId}, opts...)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if len(transactioms.Transactions) == 0 {
		respondWithJSON(w, http.StatusOK, make([]string, 0))
		return
	}

	respondWithJSON(w, http.StatusOK, transactioms.Transactions)
}
