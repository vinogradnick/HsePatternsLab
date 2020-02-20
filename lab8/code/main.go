package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"hse.pattern/lab8/code/smart_home/sensors"
	"log"
	"net/http"
)

type Metric struct {
	Metric string `json:"metric"`
}

func createStruct(data string) *Metric {
	return &Metric{Metric: data}
}

/*
motion sensor
pressure sensor
light sensor
electrical sensor
temperature sensor
*/
func main() {
	e := sensors.ElectricSensor{}
	th := sensors.ThermalMetric{}
	ps := sensors.LightMetric{}
	r := mux.NewRouter()

	r.HandleFunc("/electro", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(createStruct(e.CreateMetric()))
	})

	r.HandleFunc("/temp", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(createStruct(th.CreateMetric()))
	})
	r.HandleFunc("/light", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(createStruct(ps.CreateMetric()))
	})

	log.Fatal(http.ListenAndServe(":4000", r))
}
