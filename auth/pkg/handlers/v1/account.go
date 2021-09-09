package v1

import (
	"errors"
	"naerarauth/models"
	"naerarshared/constants"
	"naerarshared/helpers"
	appModel "naerarshared/models"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (handler *NaerarAuthRouteHandler) LoginAccount(w http.ResponseWriter, r *http.Request) {
	//Settle Cors challenge
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		helpers.RespondWithText(w, http.StatusOK, "")
		return
	}

	//Read Request Body & Validate
	var loginPayload models.LoginRequestPayload
	err := helpers.DecodeJSONBody(w, r, &loginPayload)
	if err != nil {
		response := helpers.CreateResponse(constants.ERROR, constants.INVALIDREQUEST, nil)
		helpers.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	_, err = valid.ValidateStruct(loginPayload)
	if err != nil {
		response := helpers.CreateResponse(constants.ERROR, err.Error(), nil)
		helpers.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	user, err := handler.Db.FindUserAccount(&appModel.UserAccount{Email: loginPayload.Email})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response := helpers.CreateResponse(constants.ERROR, constants.INVALIDUSER, nil)
		helpers.RespondWithJSON(w, http.StatusUnauthorized, response)
		return
	}
	if err != nil {
		response := helpers.CreateResponse(constants.ERROR, constants.SOMETHINGWENTWRONG, nil)
		helpers.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(loginPayload.Password))
	if err != nil {
		response := helpers.CreateResponse(constants.ERROR, constants.INVALIDUSER, nil)
		helpers.RespondWithJSON(w, http.StatusUnauthorized, response)
		return
	}

	ts, err := helpers.CreateToken(user.Id)
	if err != nil {
		response := helpers.CreateResponse(constants.ERROR, constants.SOMETHINGWENTWRONG, nil)
		helpers.RespondWithJSON(w, http.StatusUnauthorized, response)
		return
	}

	saveErr := helpers.CreateAuth(user.Id, ts, handler.MemStore)
	if saveErr != nil {
		response := helpers.CreateResponse(constants.ERROR, constants.SOMETHINGWENTWRONG, nil)
		helpers.RespondWithJSON(w, http.StatusUnauthorized, response)
		return
	}

	tokens := map[string]string{
		"data":          user.Id,
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	response := helpers.CreateResponse(constants.SUCCESS, constants.SUCCESS, tokens)
	helpers.RespondWithJSON(w, http.StatusOK, response)
}

func (handler *NaerarAuthRouteHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	//Settle Cors challenge
	helpers.SetupCors(&w, r)
	if r.Method == "OPTIONS" {
		helpers.RespondWithText(w, http.StatusOK, "")
		return
	}

	//Read Request Body & Validate
	var registerPayload models.RegisterUserRequestPayload
	err := helpers.DecodeJSONBody(w, r, &registerPayload)
	if err != nil {
		response := helpers.CreateResponse(constants.ERROR, constants.INVALIDREQUEST, nil)
		helpers.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	_, err = valid.ValidateStruct(registerPayload)
	if err != nil {
		response := helpers.CreateResponse(constants.ERROR, err.Error(), nil)
		helpers.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	user, err := handler.Db.FindUserAccount(&appModel.UserAccount{Email: registerPayload.Email})

	if errors.Is(err, gorm.ErrRecordNotFound) {

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(registerPayload.Password), bcrypt.DefaultCost)
		if err != nil {
			response := helpers.CreateResponse(constants.ERROR, constants.SOMETHINGWENTWRONG, nil)
			helpers.RespondWithJSON(w, http.StatusUnauthorized, response)
			return
		}
		var newAccount appModel.UserAccount
		newAccount.Password = hashedPass
		newAccount.Id = helpers.GuidId()
		err = copier.Copy(&newAccount, &registerPayload)

		if err != nil {
			response := helpers.CreateResponse(constants.ERROR, constants.INVALIDREQUEST, nil)
			helpers.RespondWithJSON(w, http.StatusInternalServerError, response)
			return
		}
		userId, err := handler.Db.CreateUserAccount(&newAccount)
		if err != nil {
			response := helpers.CreateResponse(constants.ERROR, err.Error(), nil)
			helpers.RespondWithJSON(w, http.StatusUnauthorized, response)
			return
		}

		response := helpers.CreateResponse(constants.SUCCESS, constants.SUCCESS, userId)
		helpers.RespondWithJSON(w, http.StatusOK, response)
		return
	}

	if user != nil {
		response := helpers.CreateResponse(constants.ERROR, constants.DUPLICATEACCOUNT, nil)
		helpers.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}
	response := helpers.CreateResponse(constants.ERROR, constants.SOMETHINGWENTWRONG, nil)
	helpers.RespondWithJSON(w, http.StatusInternalServerError, response)

}
