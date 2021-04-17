package rest

import (
	"authentication/models/v1"
	"authentication/pkg/helpers"


	//"encoding/hex"
	"fmt"
	"log"
	"net/http"
	//"net/url"
	//"encoding/json"
	//"os"
	//"strings"
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
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

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
		if err != nil {
			if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
				respondWithError(w, http.StatusNotFound, UserNotFound)
				return
			}
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
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	helpers.SetupCors(&w, r)
	key := helpers.ExtractToken(r)

	params := mux.Vars(r)
	email := params["email"]
	var opts []grpc.CallOption

	u, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Id: key, Email: email}, opts...)

	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if u.EmailVerifiedAt > 1000 {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Your email has been validated since %s", time.Unix(u.EmailVerifiedAt, 0)))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Verification mail has been sent to %v", u.Email)})

	token := helpers.RandUpperAlpha(7)
	handler.RedisService.Client.Set(email, token, time.Hour)
	log.Println(token)
	/*msg := events.ResendEmailEvent{
		ID:    hex.EncodeToString([]byte(u.Id)),
		Email: email,
		Token: token,
	}

	handler.EventEmitter.Emit(&msg, "NaeraExchange")*/

}

func (handler *AuthHandler) SendVerificationSMS(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	helpers.SetupCors(&w, r)
	key := helpers.ExtractToken(r)

	params := mux.Vars(r)
	phone := params["phone"]
	var opts []grpc.CallOption

	u, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{Id: key, PhoneNumber: phone}, opts...)

	if err != nil {
		if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
			respondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}
		respondWithError(w, http.StatusInternalServerError, InternalServerError)
		return
	}

	if u.PhoneVerifiedAt > 1000 {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Your phone number has been verified since %s", time.Unix(u.PhoneVerifiedAt, 0)))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Verification sms has been sent to %v", u.Email)})

	token := helpers.RandUpperAlpha(7)
	handler.RedisService.Client.Set(phone, token, time.Hour)
	log.Println(token)
/*
	NUMBER_FROM := os.Getenv("TwilloPhone")
	accountSid := os.Getenv("TwilloSID")
	authToken := os.Getenv("TwilloToken")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", fmt.Sprintf("+234%s", removeFirstChar(u.PhoneNumber)))
	msgData.Set("From", NUMBER_FROM)
	msgData.Set("Body", fmt.Sprintf("Your Naera pay verification code is: %v", token))
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}

	log.Println(token)
	msg := events.ResendEmailEvent{
		ID:    hex.EncodeToString([]byte(u.Id)),
		Email: email,
		Token: token,
	}

	handler.EventEmitter.Emit(&msg, "NaeraExchange")*/

}

func removeFirstChar(input string) string {
	if len(input) <= 1 {
		return ""
	}
	return input[1:]
}



func (handler *AuthHandler) VerifyPhone(w http.ResponseWriter, r *http.Request) {
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		respondWithJSON(w, http.StatusOK, nil)
		return
	}

	//Get route parameters
	params := mux.Vars(r)
	phone := params["phone"]
	reqToken := params["token"]

	storedToken, err := handler.RedisService.Client.Get(phone).Result()
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

		u, err := handler.GrpcPlug.FindAccount(r.Context(), &models.Account{PhoneNumber: phone}, opts...)
		if err != nil {
			if grpc.ErrorDesc(err) == gorm.ErrRecordNotFound.Error() {
				respondWithError(w, http.StatusNotFound, UserNotFound)
				return
			}
			respondWithError(w, http.StatusOK, InternalServerError)
			return
		}
		currentTime := time.Now()

		_, err = handler.GrpcPlug.UpdateAccount(r.Context(), &models.UpdateAccountRequest{Old: u, New: &models.Account{PhoneVerifiedAt: currentTime.Unix()}}, opts...)
		if err != nil {
			log.Fatal(err)
			return
		}
		handler.RedisService.Client.Del(phone)
		respondWithJSON(w, http.StatusOK, map[string]interface{}{"message": PhoneVerificationSuccessful})
		return
	} else {
		respondWithError(w, http.StatusNotFound, fmt.Errorf("Invalid or incorrect token").Error())
		return
	}
}