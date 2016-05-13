package main

import "backend/sensor"
import "fmt"
import "github.com/gorilla/mux"
import "net/http"
import "encoding/json"
import "strconv"

var port int = 6573
var serverport1 int = 7311
var serverport2 int = 6123
var count int = 0

func main() {

	//statements to be removed when deploying on AWS
	go sensor.StartServer(serverport1)
	go sensor.StartServer(serverport2)
	//

	fmt.Println("Starting Load Balancer : " + strconv.Itoa(port))

	startLBServer(port)
}

func startLBServer(port int) {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/sensor/data/current/{sensorid}", getSensorDataService).Methods("GET")
	// rtr.HandleFunc("/sensor/add", addSensorService).Methods("POST")
	// rtr.HandleFunc("/sensor/remove/{sensorid}", removeSensorService).Methods("DELETE")
	// rtr.HandleFunc("/sensor/status/{sensorid}/{status}", changeSensorStatusService).Methods("PUT")
	http.Handle("/", rtr)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

// webservice which adds a new sensor and return sensor data
// func addSensorService(rw http.ResponseWriter, req *http.Request) {

// 	//creating an empty input struct
// 	asreq := AddSensorReq{}

// 	// Populate the user data
// 	json.NewDecoder(req.Body).Decode(&asreq)

// 	res := addSensor(asreq)

// 	resJson, _ := json.Marshal(res)

// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.WriteHeader(201)
// 	fmt.Fprintf(rw, "%s", resJson)
// }

// webservice which gets sensor data for a particular sensor id
func getSensorDataService(rw http.ResponseWriter, req *http.Request) {
	p := mux.Vars(req)

	sensorid := p["sensorid"]

	client := &http.Client{}
	request, _ := http.NewRequest("GET", getServerUrlGET(sensorid), nil)
	response, _ := client.Do(request)

	//creating an empty input struct
	svres := sensor.SensorValueRes{}

	// Populate the user data

	json.NewDecoder(response.Body).Decode(&svres)

	resJson, _ := json.Marshal(svres)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	fmt.Fprintf(rw, "%s", resJson)
}

// webservice which removes sensor
// func removeSensorService(rw http.ResponseWriter, req *http.Request) {
// 	p := mux.Vars(req)

// 	sensorId := p["sensorid"]

// 	removeSensor(sensorId)

// 	rw.WriteHeader(204)
// }

// webservice which sets sensor status for a particular sensor id
// func changeSensorStatusService(rw http.ResponseWriter, req *http.Request) {

// 	p := mux.Vars(req)

// 	sensorId := p["sensorid"]
// 	sensorStatus, _ := strconv.ParseBool(p["status"])

// 	res := changeSensorStatus(sensorId, sensorStatus)

// 	resJson, _ := json.Marshal(res)

// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.WriteHeader(200)
// 	fmt.Fprintf(rw, "%s", resJson)
// }

func getServerUrlGET(sensorid string) string {

	var url string

	if count%2 == 0 && count%5 == 0 {
		url = "http://localhost:" + strconv.Itoa(serverport1) + "/sensor/data/current/" + sensorid
	} else if count%2 != 0 && count%5 != 0 {
		url = "http://localhost:" + strconv.Itoa(serverport2) + "/sensor/data/current/" + sensorid
	} else {
		url = "http://localhost:" + strconv.Itoa(serverport1) + "/sensor/data/current/" + sensorid
	}

	count++

	return url
}
