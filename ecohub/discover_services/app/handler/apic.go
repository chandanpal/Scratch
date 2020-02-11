package handler

import (
	//"code.google.com/p/go.net/publicsuffix"
	"net/http/cookiejar"
	"net/url"
	"net/http"
	"io/ioutil"
	
)

type Apic struct{
	Host      string
	Username  string
	Password  string
	Base_url  string
	Client    http.Client
}
var APIC Apic

func ApicInit(hostname string, username string, password string){
	APIC.Host = hostname
	APIC.Username = username
	APIC.Password = password
	APIC.Base_url = "http://" + hostname
	
}
func login(w http.ResponseWriter) {
	apic_url := APIC.Base_url + "/api/aaaLogin.json"
	/*
	options := cookiejar.Options{
        	PublicSuffixList: publicsuffix.List,
        }
	*/
   	jar, err := cookiejar.New(nil)
    	if err != nil {
        	respondError(w, http.StatusBadRequest, err.Error())
    	}
    	client := http.Client{Jar: jar}
    	resp, err := client.PostForm(apic_url, url.Values{
        	"password": {APIC.Password},
        	"username" : {APIC.Username},        
        })
        if err != nil {
        	respondError(w, http.StatusUnauthorized, err.Error())
       	}
	defer resp.Body.Close()
	APIC.Client = client
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil{
		respondError(w, http.StatusBadRequest, err.Error())
	}
	
}

func logout(w http.ResponseWriter){
	url := APIC.Base_url + "/api/aaaLogout.json"
	client := APIC.Client
	resp, err := client.Get(url)
	if err != nil {
                respondError(w, http.StatusBadRequest, err.Error())
        }
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil{
		respondError(w, http.StatusBadRequest, err.Error())
	}
}

func request_apic(method string, url string){
	
}
