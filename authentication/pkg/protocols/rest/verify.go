package rest

import (
	"authentication/models/v1"
	"authentication/pkg/helpers"
	//"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"

	//"shared/amqp/events"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

const Redirection = `
<h5>
Your Email has been successfully verified. You will be redirected back in 3 seconds
</h5>
<h6>Click <a href="https://consumer.naerademo.com/signout"> here </a>to manually redirect</h6>
<script type="text/javascript"> setTimeout(function(){ window.location.replace("https://consumer.naerademo.com/signout") }, 2000)</script>
`

func (handler *AuthHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)

	//Get route parameters
	params := mux.Vars(r)
	email := params["email"]
	reqToken := params["token"]

	storedToken, err := handler.RedisService.Client.Get(email).Result()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("Invalid or incorrect token").Error())
		return
	}
	if storedToken == "" {
		respondWithError(w, http.StatusBadRequest, fmt.Errorf("Invalid or incorrect token").Error())
		return
	}
	if match := reqToken == storedToken; match {
		var opts []grpc.CallOption

		u, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Email: email}, opts...)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusOK, UserNotFound)
			return
		}
		if err != nil {
			respondWithError(w, http.StatusOK, InternalServerError)
			return
		}
		currentTime := time.Now()

		_, err = handler.GrpcPlug.UpdateAccount(r.Context(), &models.UpdateAccountRequest{Old: u, New: &models.Account{EmailVerifiedAt: currentTime.Unix()}}, opts...)
		if err != nil {
			log.Fatal(err)
			return
		}
		handler.RedisService.Client.Del(email)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(Redirection))
		return
	} else {
		respondWithError(w, http.StatusNotFound, fmt.Errorf("Invalid or incorrect token").Error())
		return
	}
}

func (handler *AuthHandler) SendVerification(w http.ResponseWriter, r *http.Request) {
	setupCors(&w, r)

	params := mux.Vars(r)
	email := params["email"]
	var opts []grpc.CallOption

	u, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Email: email}, opts...)
	
	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusOK, InternalServerError)
		return
	}

	if u.EmailVerifiedAt > 1000{
		respondWithError(w, http.StatusOK,  fmt.Sprintf("Your email has been validated since %s" ,time.Unix(u.EmailVerifiedAt, 0)))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Verification mail has been sent to %v", u.Email)})

	token := helpers.RandUpperAlpha(7)
	handler.RedisService.Client.Set(email, token, time.Hour)
	/*msg := events.ResendEmailEvent{
		ID:    hex.EncodeToString([]byte(u.Id)),
		Email: email,
		Token: token,
	}

	handler.EventEmitter.Emit(&msg, "NaeraExchange")*/

}
