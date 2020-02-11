package handler

import (
	//"github.com/elastic/go-elasticsearch/v5"
	"github.com/olivere/elastic"
	//"github.com/sparrc/go-ping"
	"github.com/discover_services/app/model"	
	"net/http"
	_"io/ioutil"
	_"reflect"
	_"time"
	"context"
	
)

func TestElasticsearch(w http.ResponseWriter, postData model.PostData) int{
	var esHost string
	var esPort string
	var status_code int
	for _, val := range postData.Parameters{
		if val.Name == "es_host"{
			esHost = val.Value
		}
		if val.Name == "es_port"{
			esPort = val.Value
		}
	}
	esURL := "http://" + esHost + ":" + esPort
	//ctx := context.Background()
	//client, err := elastic.NewClient(elastic.SetURL(esURL))
	client, err := elastic.NewClient(elastic.SetURL(esURL), elastic.SetHealthcheck(false), elastic.SetSniff(false))
	if err != nil {
		err_msg := "Error Instantiating Elasticsearch Client. Error : " + string(err.Error())
		logs.Errorf(err_msg)
		respondError(w,  http.StatusBadRequest, err_msg)
	}
	// Ping the Elasticsearch server to get.
	_, _, err = client.Ping(esURL).Do(context.TODO())
	if err != nil {
		// Handle error
		logs.Errorf("Error pinging Elasticsearch . Error : ", err)
		err_msg := "Elasticsearch is not running, Ping to the ES host failed with the error : " + string(err.Error())
		status_code = http.StatusNotFound
                respondError(w,  status_code, err_msg)
	} else {
		status_code = http.StatusOK
	}
	return status_code
}

func TestFlink(w http.ResponseWriter, postData model.PostData) int{
	var flinkHost string
	var flinkPort string
	var statusCode int
        for _, val := range postData.Parameters{
                if val.Name == "flink_host"{
                        flinkHost = val.Value
                }
                if val.Name == "flink_port"{
                        flinkPort = val.Value
                }
       	}
	flinkUrl := "http://" + flinkHost + ":" + flinkPort
	resp, err := request(flinkUrl)
	if err != nil{
		logs.Errorf("Error connecting flink server. Error: ", err.Error())
		err_msg := "Flink is not running. Error : " + string(err.Error())
		statusCode = http.StatusNotFound
		respondError(w,  statusCode, err_msg)
	}
	if resp != nil{
		statusCode = resp.StatusCode
		//responseData, _ := ioutil.ReadAll(resp.Body)
	}
	return statusCode
}

func TestKafka(w http.ResponseWriter, postData model.PostData) int{
	var kafkaHost string
	var kafkaPort string
	var statusCode int
	for _, val := range postData.Parameters{
                if val.Name == "kafka_host"{
                        kafkaHost = val.Value
                }
                if val.Name == "kafka_port"{
                        kafkaPort = val.Value
                }
        }
	kafkaUrl := "http://" + kafkaHost + ":" + kafkaPort
	resp, err := request(kafkaUrl)
	if err != nil{
		logs.Errorf("Error connecting kafka server. Error: ", err.Error())
                err_msg := "Kafka is not running. Error : " + string(err.Error())
		statusCode = http.StatusNotFound
                respondError(w, statusCode, err_msg)
        }
	if resp != nil{
		statusCode = resp.StatusCode
	}
	return statusCode
}

func TestConnectivity(w http.ResponseWriter, r *http.Request){
	postData, err := getPostData(r.Body)
        if err != nil {
                logs.Errorf("Request body decode Error")
                respondError(w,  http.StatusBadRequest, "Request body decode Error")
        }
	if postData.Type == "Flink"{
		status_code := TestFlink(w, postData)
		if status_code == 200{
			respondJSON(w, http.StatusOK, "Flink server is running successfully")
		}
	}
	if postData.Type == "elasticsearch"{
		status_code := TestElasticsearch(w, postData)
		if status_code == 200{
                        respondJSON(w, http.StatusOK, "Elasticsearch server is running successfully")
                }

	}
	if postData.Type == "kafka"{
		status_code := TestKafka(w, postData)
                if status_code == 200{
                        respondJSON(w, http.StatusOK, "Kafka server is running successfully")
                }
	}

	if postData.Type == "aci"{
		if postData.Connection_type == "userpass"{
			status_code := TestAciUserpass(w, postData)
	                if status_code == 200{ 
                        	respondJSON(w, http.StatusOK, "ACI test connectivity Done")
                	}
		}
		if postData.Connection_type == "certificate"{
			status_code := TestAciCertificate(w, postData)
			if status_code == 200{
				respondJSON(w, http.StatusOK, "ACI test connectivity Done")
			}
		}
	}
}

func TestAciUserpass(w http.ResponseWriter, postData model.PostData) int{
	var aciHost string
        var aciUser string
        var aciPass string
        for _, val := range postData.Parameters{
                if val.Name == "apic_hostname"{
                        aciHost = val.Value
                }
                if val.Name == "aci_username"{
                        aciUser = val.Value
                }
		if val.Name == "aci_password"{
                        aciPass = val.Value
                }
        }
	ApicInit(aciHost, aciUser, aciPass)
	login(w)
	logout(w)
	return 200
}

func TestAciCertificate(w http.ResponseWriter, postData model.PostData) int{
	return 200
}
