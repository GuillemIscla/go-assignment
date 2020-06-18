package http

import (
	"gitlab.com/zenport.io/go-assignment/engine"
	"gitlab.com/zenport.io/go-assignment/domain"
	"net/http"
	"time"
	"context"
    "encoding/json"
    "io/ioutil"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"
)

type HTTPAdapter struct{
    Server http.Server
    Engine engine.Engine
}

type KnightReq struct {
    Name     *string `json:"name,omitempty"`
    Strength *int `json:"strength,omitempty"`
    WeaponPower *float64 `json:"weapon_power,omitempty"`
}


func (knightReq KnightReq) ToKnight() domain.Knight {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	return domain.Knight{
	    Id: strconv.Itoa(r1.Intn(100000)),
	    Name: *knightReq.Name,
	    Strength: *knightReq.Strength,
	    WeaponPower: *knightReq.WeaponPower}
}


type ErrorResponse struct {
    Code int `json:"code"`
    Message string `json:"message"`
}

func (adapter *HTTPAdapter) Start() {
	adapter.Server.ListenAndServe()
}

func (adapter *HTTPAdapter) Stop() {
	ctx, _ := context.WithTimeout(context.TODO(), 10 * time.Second)
	adapter.Server.Shutdown(ctx)
}

func NewHTTPAdapter(e engine.Engine) *HTTPAdapter {
    var r = mux.NewRouter()

    var s = http.Server {
        Addr:           ":8080",
        Handler:        r,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }


    adapter := &HTTPAdapter{Server: s, Engine: e}
    r.HandleFunc("/knight", adapter.HandlerGetKnights).Methods("GET")
    r.HandleFunc("/knight/{id}", adapter.HandlerGetSingleKnight).Methods("GET")
    r.HandleFunc("/knight", adapter.HandlerCreateKnight).Methods("POST")

	return adapter
}

func (adapter *HTTPAdapter) HandlerGetKnights(w http.ResponseWriter, r *http.Request) {
    knights := adapter.Engine.ListKnights()

    writeJsonResponse(w, knights, http.StatusOK)
}


func (adapter *HTTPAdapter) HandlerGetSingleKnight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

    knight, err := adapter.Engine.GetKnight(id)

    if err != nil {
        writeErrorResponse(w, 1, "Knight #" + id + " not found.", http.StatusNotFound)
    }
    if err == nil {
        writeJsonResponse(w, knight, http.StatusOK)
    }
}

func (adapter *HTTPAdapter) HandlerCreateKnight(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeJsonResponse(w, err, http.StatusBadRequest)
		return
	}

    var knightReq KnightReq
	err = json.Unmarshal(body, &knightReq)

	if err != nil {
		writeJsonResponse(w, err, http.StatusBadRequest)
		return
	}

	if knightReq.Name == nil {
        writeErrorResponse(w, 2, "Missing name", http.StatusBadRequest)
        return
    }

    if knightReq.Strength == nil {
        writeErrorResponse(w, 3, "Missing strength", http.StatusBadRequest)
        return
    }

    if knightReq.WeaponPower == nil {
        writeErrorResponse(w, 4, "Missing weapon_power", http.StatusBadRequest)
        return
    }

	var knight = knightReq.ToKnight()

	adapter.Engine.Save(knight)

	writeJsonResponse(w, knight, http.StatusCreated)
}


func writeJsonResponse(w http.ResponseWriter, data interface{}, status int){
	jsonData, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
	w.Write(jsonData)
}


func writeErrorResponse(w http.ResponseWriter, code int, message string, status int) {
    var errorResponse = ErrorResponse{Code: code, Message: message}
    jsonData, _ := json.Marshal(errorResponse)
    w.WriteHeader(status)
    w.Write(jsonData)
}

