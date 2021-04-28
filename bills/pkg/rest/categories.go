package rest

import (
	"bills/pkg/helpers"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (handler *BillHandler) LiveCategories(w http.ResponseWriter, r *http.Request) {
	
	
	helpers.SetupCors(&w, r)
	var opts []grpc.CallOption

	res, err := handler.GrpcPlug.GetBillCategories(r.Context(), &emptypb.Empty{}, opts...)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	if len(res.Categories) ==0 {
		respondWithJSON(w, http.StatusOK, make([]string, 0))
		return
	}
	respondWithJSON(w, http.StatusOK, res.Categories)
}
