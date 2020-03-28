package handler

import (
	"encoding/json"
	"net/http"
	_"github.com/jinzhu/gorm"
	_"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"io"
	"github.com/discover_services/app/model"
	"github.com/discover_services/logger"
	"time"
	_"fmt"
	
)

var logs *logger.Logger


func Init() {
	logs = logger.Logs.GetLogger("Handler")
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
	
}


func NotImplemented(w http.ResponseWriter, r *http.Request) {

	respondJSON(w, http.StatusOK, "API Not Implemented")
}


func getResponseData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	return contents, nil
}

func getPostData(data io.Reader)(model.PostData, error){

        postData := model.PostData{}
        decoder := json.NewDecoder(data)
        err := decoder.Decode(&postData)
        return postData, err

}

func request(url string) (*http.Response, error){
	client := http.Client{
                Timeout: 10 * time.Second,
        }
	resp, err := client.Get(url)
        return resp, err	
}

