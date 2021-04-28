package rest

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (handler *BillHandler) LiveCategories(w http.ResponseWriter, r *http.Request) {
	var opts []grpc.CallOption

	res, err := handler.GrpcPlug.GetBillCategories(r.Context(), &emptypb.Empty{}, opts...)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res.Categories)
}
