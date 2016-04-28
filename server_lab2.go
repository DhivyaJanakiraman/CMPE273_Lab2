package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	//"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//Map to store key value pairs
var my_map = make(map[string]string)
var my_map1 = make(map[string]string)
var my_map2 = make(map[string]string)
var my_map3 = make(map[string]string)
var my_map4 = make(map[string]string)

type KeyValPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func PutKey(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	//storing id and value from the request
	key_id := param.ByName("id")
	value := param.ByName("value")
	//storing the values in map - my_map
	my_map[key_id] = value
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)

	fmt.Println("KEY saved successfully !!!")
	fmt.Printf("Key %s and value : %s \n", key_id, value)

}

func GetKey(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeyValPair KeyValPair
	key_id := param.ByName("id")
	if val, ok := my_map[key_id]; ok {

		myKeyValPair.Key = key_id
		myKeyValPair.Value = val

		// Marshal provided interface into JSON structure
		myJsonObj, _ := json.Marshal(myKeyValPair)
		// Write content-type, statuscode, payload
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "%s", myJsonObj)
		fmt.Println("Key retrieved successfully!")

	}

}

func GetAllKeys(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeys []KeyValPair

	for key, value := range my_map {
		myKeyValPair := KeyValPair{
			Key:   key,
			Value: value,
		}
		myKeys = append(myKeys, myKeyValPair)
	}

	resp, _ := json.MarshalIndent(myKeys, "", "    ")

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resp)
	fmt.Printf("%s", resp)
	fmt.Println("All Keys retrieved successfully!")

}

func PutKey1(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	//storing id and value from the request
	key_id := param.ByName("id")
	value := param.ByName("value")
	//storing the values in map - my_map
	my_map1[key_id] = value
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)

	fmt.Println("KEY saved successfully !!!")
	fmt.Printf("Key %s and value : %s \n", key_id, value)

}

func GetKey1(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeyValPair KeyValPair
	key_id := param.ByName("id")
	if val, ok := my_map1[key_id]; ok {

		myKeyValPair.Key = key_id
		myKeyValPair.Value = val

		// Marshal provided interface into JSON structure
		myJsonObj, _ := json.Marshal(myKeyValPair)
		// Write content-type, statuscode, payload
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "%s", myJsonObj)
		fmt.Println("Key retrieved successfully!")

	}

}

func GetAllKeys1(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeys []KeyValPair

	for key, value := range my_map1 {
		myKeyValPair := KeyValPair{
			Key:   key,
			Value: value,
		}
		myKeys = append(myKeys, myKeyValPair)
	}

	resp, _ := json.MarshalIndent(myKeys, "", "    ")

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resp)
	fmt.Printf("%s", resp)
	fmt.Println("All Keys retrieved successfully!")

}

func PutKey2(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	//storing id and value from the request
	key_id := param.ByName("id")
	value := param.ByName("value")
	//storing the values in map - my_map
	my_map2[key_id] = value
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)

	fmt.Println("KEY saved successfully !!!")
	fmt.Printf("Key %s and value : %s \n", key_id, value)

}

func GetKey2(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeyValPair KeyValPair
	key_id := param.ByName("id")
	if val, ok := my_map2[key_id]; ok {

		myKeyValPair.Key = key_id
		myKeyValPair.Value = val

		// Marshal provided interface into JSON structure
		myJsonObj, _ := json.Marshal(myKeyValPair)
		// Write content-type, statuscode, payload
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "%s", myJsonObj)
		fmt.Println("Key retrieved successfully!")

	}

}

func GetAllKeys2(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeys []KeyValPair

	for key, value := range my_map2 {
		myKeyValPair := KeyValPair{
			Key:   key,
			Value: value,
		}
		myKeys = append(myKeys, myKeyValPair)
	}

	resp, _ := json.MarshalIndent(myKeys, "", "    ")

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resp)
	fmt.Printf("%s", resp)
	fmt.Println("All Keys retrieved successfully!")

}

func PutKey3(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	//storing id and value from the request
	key_id := param.ByName("id")
	value := param.ByName("value")
	//storing the values in map - my_map
	my_map3[key_id] = value
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)

	fmt.Println("KEY saved successfully !!!")
	fmt.Printf("Key %s and value : %s \n", key_id, value)

}

func GetKey3(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeyValPair KeyValPair
	key_id := param.ByName("id")
	if val, ok := my_map3[key_id]; ok {

		myKeyValPair.Key = key_id
		myKeyValPair.Value = val

		// Marshal provided interface into JSON structure
		myJsonObj, _ := json.Marshal(myKeyValPair)
		// Write content-type, statuscode, payload
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "%s", myJsonObj)
		fmt.Println("Key retrieved successfully!")

	}

}

func GetAllKeys3(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeys []KeyValPair

	for key, value := range my_map3 {
		myKeyValPair := KeyValPair{
			Key:   key,
			Value: value,
		}
		myKeys = append(myKeys, myKeyValPair)
	}

	resp, _ := json.MarshalIndent(myKeys, "", "    ")

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resp)
	fmt.Printf("%s", resp)
	fmt.Println("All Keys retrieved successfully!")

}

func PutKey4(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	//storing id and value from the request
	key_id := param.ByName("id")
	value := param.ByName("value")
	//storing the values in map - my_map
	my_map4[key_id] = value
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)

	fmt.Println("KEY saved successfully !!!")
	fmt.Printf("Key %s and value : %s \n", key_id, value)

}

func GetKey4(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeyValPair KeyValPair
	key_id := param.ByName("id")
	if val, ok := my_map4[key_id]; ok {

		myKeyValPair.Key = key_id
		myKeyValPair.Value = val

		// Marshal provided interface into JSON structure
		myJsonObj, _ := json.Marshal(myKeyValPair)
		// Write content-type, statuscode, payload
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "%s", myJsonObj)
		fmt.Println("Key retrieved successfully!")

	}

}

func GetAllKeys4(rw http.ResponseWriter, req *http.Request, param httprouter.Params) {

	var myKeys []KeyValPair

	for key, value := range my_map4 {
		myKeyValPair := KeyValPair{
			Key:   key,
			Value: value,
		}
		myKeys = append(myKeys, myKeyValPair)
	}

	resp, _ := json.MarshalIndent(myKeys, "", "    ")

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resp)
	fmt.Printf("%s", resp)
	fmt.Println("All Keys retrieved successfully!")

}

//Setting up server

func main() {

	mux := httprouter.New()
	mux1 := httprouter.New()
	mux2 := httprouter.New()
	mux3 := httprouter.New()
	mux4 := httprouter.New()

	mux.PUT("/keys/:id/:value", PutKey)
	mux.GET("/keys/:id", GetKey)
	mux.GET("/keys", GetAllKeys)

	mux1.PUT("/keys/:id/:value", PutKey1)
	mux1.GET("/keys/:id", GetKey1)
	mux1.GET("/keys", GetAllKeys1)

	mux2.PUT("/keys/:id/:value", PutKey2)
	mux2.GET("/keys/:id", GetKey2)
	mux2.GET("/keys", GetAllKeys2)

	mux3.PUT("/keys/:id/:value", PutKey3)
	mux3.GET("/keys/:id", GetKey3)
	mux3.GET("/keys", GetAllKeys3)

	mux4.PUT("/keys/:id/:value", PutKey4)
	mux4.GET("/keys/:id", GetKey4)
	mux4.GET("/keys", GetAllKeys4)

	fmt.Println("All servers launched")

	myAddr := os.Args[1]
	fmt.Printf("Addr: %s \n", myAddr)
	ports := strings.Split(myAddr, "-")
	start, _ := strconv.Atoi(ports[0])
	end, _ := strconv.Atoi(ports[1])
	//port1, _ := strconv.ParseInt(start, 10, 0)
	//port2, _ := strconv.ParseInt(end, 10, 0)
	fmt.Printf("start : %d end : %d \n", start, end)
	count := (end - start) + 1
	fmt.Printf("count: %d", count)

	for i := 0; i < count; i++ {
		fmt.Println(i)
		/*go func(i int) {
			fmt.Println("Listening...")
			var port = fmt.Sprintf(":%d", i)
			log.Println(i)
			http.ListenAndServe(port, nil)
		}(i)*/

	}

	go func() {
		server1 := http.Server{
			Addr:    "0.0.0.0:3002",
			Handler: mux1,
		}
		server1.ListenAndServe()
	}()

	go func() {
		server2 := http.Server{
			Addr:    "0.0.0.0:3003",
			Handler: mux2,
		}
		server2.ListenAndServe()
	}()

	go func() {
		server3 := http.Server{
			Addr:    "0.0.0.0:3004",
			Handler: mux3,
		}
		server3.ListenAndServe()
	}()

	go func() {
		server4 := http.Server{
			Addr:    "0.0.0.0:3005",
			Handler: mux4,
		}
		server4.ListenAndServe()
	}()

	server := http.Server{
		Addr:    "0.0.0.0:3001",
		Handler: mux,
	}
	server.ListenAndServe()

}
