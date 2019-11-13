package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"github.com/mehdou92/vote-api/api/auth"
	"github.com/mehdou92/vote-api/api/models"
	"github.com/mehdou92/vote-api/api/responses"
	"github.com/mehdou92/vote-api/api/utils/formaterror"
)

func (server *Server) CreateVote(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vote := models.Vote{}
	err = json.Unmarshal(body, &vote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vote.Prepare()
	err = vote.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != vote.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	voteCreated, err := vote.SaveVote(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	type ResponseStruct struct{
		gorm.Model
		Id uint64 `json:id`
		Title string `json:title`
		Desc string `json:desc`
	}

	responseData := ResponseStruct{
		Id:voteCreated.ID, 
		Title:voteCreated.Title, 
		Desc:voteCreated.Desc} 
		

	resp, err := json.Marshal(responseData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (server *Server) GetVotes(w http.ResponseWriter, r *http.Request) {

	vote := models.Vote{}

	votes, err := vote.FindAllVotes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, votes)
}

func (server *Server) GetVote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	vote := models.Vote{}

	voteReceived, err := vote.FindVoteByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, voteReceived)
}

func (server *Server) UpdateVote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the vote id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the vote exist
	vote := models.Vote{}
	err = server.DB.Debug().Model(models.Vote{}).Where("id = ?", pid).Take(&vote).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Vote not found"))
		return
	}

	// If a user attempt to update a vote not belonging to him
	if uid != vote.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data voteed
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	voteUpdate := models.Vote{}
	err = json.Unmarshal(body, &voteUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != voteUpdate.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	voteUpdate.Prepare()
	err = voteUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	voteUpdate.ID = vote.ID //this is important to tell the model the vote id to update, the other update field are set above

	voteUpdated, err := voteUpdate.UpdateAVote(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, voteUpdated)
}

func (server *Server) DeleteVote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid vote id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the vote exist
	vote := models.Vote{}
	err = server.DB.Debug().Model(models.Vote{}).Where("id = ?", pid).Take(&vote).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this vote?
	if uid != vote.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = vote.DeleteAVote(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
