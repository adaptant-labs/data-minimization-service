package main

import (
	"encoding/json"
	minimizers "github.com/adaptant-labs/go-minimizer"
	"github.com/etherlabsio/healthcheck"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"html/template"
	"net/http"
)

var (
	tmpl            = template.Must(template.ParseFiles("index.html"))
	minimizerTypes  []string
	minimizerLevels []string
)

type DataMinimizationRequest struct {
	Input interface{}                  `json:"input,omitempty"`
	Type  string                       `json:"type"`
	Level minimizers.MinimizationLevel `json:"level"`
}

type DataMinimizationResponse struct {
	Result interface{} `json:"result"`
}

func dataInputHandler(w http.ResponseWriter, r *http.Request) {
	var min DataMinimizationRequest

	d := json.NewDecoder(r.Body)
	err := d.Decode(&min)
	if err != nil {
		panic(err)
	}

	handler := minimizers.TagMap[min.Type]
	if handler != nil {
		var response DataMinimizationResponse

		processRequestMetrics(min.Level)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response.Result = handler(min.Level, min.Input)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "no matching handler found", http.StatusNotFound)
	}
}

type FormResponse struct {
	SupportedTypes  []string
	SupportedLevels []string
	Input           string
	Level           minimizers.MinimizationLevel
	Type            string
	Result          interface{}
}

func (f FormResponse) validate() bool {
	if f.Level == minimizers.MinimizationAnonymize {
		return true
	}

	if f.Input == "" && f.Type == "" {
		return false
	}

	return true
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var data FormResponse

	data.SupportedTypes = minimizerTypes
	data.SupportedLevels = minimizerLevels

	data.Input = r.FormValue("input")
	data.Level = minimizers.LevelFromString(r.FormValue("level"))
	data.Type = r.FormValue("type")

	if data.validate() == true {
		handler := minimizers.TagMap[data.Type]
		if handler != nil {
			data.Result = handler(data.Level, data.Input)
		}
	}

	tmpl.Execute(w, &data)
}

func NewServiceRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/data", dataInputHandler).Methods("POST")
	r.HandleFunc("/", indexHandler)

	r.Handle("/healthcheck", healthcheck.Handler())
	r.Handle("/metrics", promhttp.Handler())

	return r
}

func init() {
	for k, _ := range minimizers.StringLevelMap {
		minimizerLevels = append(minimizerLevels, k)
	}

	for k, _ := range minimizers.TagMap {
		minimizerTypes = append(minimizerTypes, k)
	}
}
