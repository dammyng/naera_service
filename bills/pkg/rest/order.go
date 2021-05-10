package rest

func (handler *BillHandler) CreateBill(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	var u models.Order
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

	order := &models.Order{
		
	}

	tRes, err := handler.GrpcPlug.CreateOrder(r.Context(), order, opts...)
	if err != nil {
		err = errors.New("Error creating the bill record")
		respondWithError(w, http.StatusBadRequest, err.Error())

	}
	respondWithJSON(w, http.StatusCreated, tRes.Id)
}
