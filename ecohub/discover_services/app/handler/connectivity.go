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

func TestElasticsearch(w http.ResponseWriter, postData model.PostData) {
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
	client, err := elastic.NewClient(elastic.SetURL(esURL), elastic.SetHealthcheck(false), elastic.SetSniff(false), elastic.SetHealthcheckTimeout(5))
	if err != nil {
		err_msg := "Error Instantiating Elasticsearch Client. Error : " + string(err.Error())
		logs.Errorf(err_msg)
		respondError(w,  http.StatusBadRequest, err_msg)
		return 
	}
	// Ping the Elasticsearch server to get.
	_, _, err = client.Ping(esURL).Do(context.TODO())
	if err != nil {
		// Handle error
		logs.Errorf("Error pinging Elasticsearch . Error : ", err)
		err_msg := "Elasticsearch is not running, Ping to the ES host failed with the error : " + string(err.Error())
		status_code = http.StatusNotFound
                respondError(w,  status_code, err_msg)
		return
	}
}

func TestFlink(w http.ResponseWriter, postData model.PostData){
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
		return
	}
	if resp != nil{
		statusCode = resp.StatusCode
		//responseData, _ := ioutil.ReadAll(resp.Body)
	}
}

func TestKafka(w http.ResponseWriter, postData model.PostData) {
	var (
		kafkaHost string
		kafkaPort string
		statusCode int
		ProducerTopic string
		ConsumerTopic string
	)
	for _, val := range postData.Parameters{
                if val.Name == "kafka_host"{
                        kafkaHost = val.Value
                }
                if val.Name == "kafka_port"{
                        kafkaPort = val.Value
                }
		if val.Name == "producer_topic"{
			ProducerTopic = val.Value
		}
		if val.Name == "consumer_topic"{
			ConsumerTopic = val.Value
		}
        }
	statusCode, producerTopicFound, consumerTopicFound := kafka_connect(kafkaHost , kafkaPort,  ProducerTopic, ConsumerTopic)
        if statusCode == http.StatusNotFound{
                logs.Errorf("Not able to connect to kafka host")
                respondError(w, statusCode, "Not able to connect to kafka host")
                return
        }
	if ProducerTopic != ""{
		if producerTopicFound == false {
			logs.Errorf("Producer Topic not found") 
                	respondError(w, http.StatusBadRequest, "Producer/Consumer Topic not found")
                	return
		}
	}
	if ConsumerTopic != ""{
		if consumerTopicFound == false {
			logs.Errorf("Consumer Topic not found")
	                respondError(w, http.StatusBadRequest, "Consumer/Producer Topic not found")
        	        return
		}
	}
}

func TestConnectivity(w http.ResponseWriter, r *http.Request){
	postData, err := getPostData(r.Body)
        if err != nil {
                logs.Errorf("Request body decode Error")
                respondError(w,  http.StatusBadRequest, "Request body decode Error")
        }
	if postData.Type == "Flink"{
		TestFlink(w, postData)
		respondJSON(w, http.StatusOK, "Flink server is running successfully")
	}
	if postData.Type == "elasticsearch"{
		TestElasticsearch(w, postData)
                respondJSON(w, http.StatusOK, "Elasticsearch server is running successfully")
	}
	if postData.Type == "kafka"{
		TestKafka(w, postData)
                respondJSON(w, http.StatusOK, "Kafka server is running successfully")
	}

	if postData.Type == "aci"{
		if postData.Connection_type == "userpass"{
			TestAciUserpass(w, postData)
                       	respondJSON(w, http.StatusOK, "ACI test connectivity Done")
		}
		if postData.Connection_type == "certificate"{
			TestAciCertificate(w, postData)
			respondJSON(w, http.StatusOK, "ACI certificate not Implemented")
			}
		}
}

func TestAciUserpass(w http.ResponseWriter, postData model.PostData) {
	var (
		aciHost string
		aciUser string
		aciPass string
		tenantName string
		apName string
	)
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
		if val.Name == "tenant_name"{
			tenantName = val.Value
		}
		if val.Name == "ap_name"{
			apName = val.Value
		}
        }
	ApicInit(aciHost, aciUser, aciPass)
	login(w)
	if tenantName != ""{
		checkTenantExistence(w, tenantName)
	}
	if apName != ""{
		checkAppExistence(w, apName)
	}
	logout(w)
}

func TestAciCertificate(w http.ResponseWriter, postData model.PostData) {
	return
}
