package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elenduuche/reportserver/models"
	"github.com/elenduuche/reportserver/services"
	"github.com/elenduuche/reportserver/utils"
	"github.com/gorilla/mux"
)

type payment struct {
	services.PaymentService
}

var (
	jsonEncoder    utils.Encoder
	jsonDecoder    utils.Decoder
	timeUtil       utils.TimeService
	gocsvEncoder   utils.Encoder
	csvutilEncoder utils.Encoder
)

func init() {
	jsonEncoder = utils.NewJSONEncoder()
	jsonDecoder = utils.NewJSONDecoder()
	gocsvEncoder = utils.NewGoCSVEncoder()
	csvutilEncoder = utils.NewCSVUtilEncoder()
	timeUtil = utils.NewTimeService()
}

const (
	path      = "/payments"
	pathParam = "id"
)

func (p *payment) registerRoutes(basePath string, r *mux.Router) {
	var sb strings.Builder
	if len(basePath) > 0 {
		sb.WriteString(basePath)
		sb.WriteString(path)
	}
	r.HandleFunc(sb.String(), p.getAllHandler).Methods("GET")
	r.HandleFunc(sb.String(), p.createHandler).Methods("POST")
	paymentRoute := r.PathPrefix(sb.String()).Subrouter()
	paymentRoute.HandleFunc(fmt.Sprintf("/{%s}", pathParam), p.getByIDHandler).Methods("GET")
	paymentRoute.HandleFunc("/month/{month}", p.getAllCSVHandler).Methods("GET")
}

func (p *payment) registerServices(data services.IDataService) {
	p.PaymentService = services.NewPaymentService(data)
}

func (p *payment) getAllHandler(w http.ResponseWriter, r *http.Request) {
	payments, err := p.PaymentService.GetAll(r.Context())
	if err != nil {
		handleError(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		err0 := jsonEncoder.Encode(w, payments)
		if err0 != nil {
			handleError(w, err0)
		}
	}
}

func (p *payment) getAllCSVHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	month, cerr := strconv.Atoi(params["month"])
	if cerr != nil {
		handleError(w, cerr)
	}
	filename, err := p.PaymentService.GetAllCSV(r.Context(), month, timeUtil.Now().Year())
	if err != nil {
		handleError(w, err)
	}
	file, err := os.Open(filename)
	if err != nil {
		handleError(w, err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		handleError(w, err)
	}

	FileSize := strconv.FormatInt(fileInfo.Size(), 10)
	//Send the headers before sending the file
	m := time.Month(month)
	downloadFilename := strings.ToLower(fmt.Sprintf("payments_%v_%v.csv", m.String(), timeUtil.NowTimestamp().Seconds))
	w.Header().Set("Content-Disposition", "attachment; filename="+downloadFilename)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	io.Copy(w, file)
}

func (p *payment) getByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payment, err := p.PaymentService.GetByID(r.Context(), params[pathParam])
	if err != nil {
		handleError(w, err)
	} else {
		err0 := jsonEncoder.Encode(w, payment)
		if err0 != nil {
			handleError(w, err)
		}
	}
}

func (p *payment) createHandler(w http.ResponseWriter, r *http.Request) {
	var in models.Payment
	if err := jsonDecoder.Decode(r.Body, &in); err != nil {
		handleError(w, err)
	}
	t := timeUtil.Now()
	in.CreatedOn = &t
	payment, err := p.PaymentService.Create(r.Context(), in)
	if err != nil {
		handleError(w, err)
	} else {
		err0 := jsonEncoder.Encode(w, payment)
		if err0 != nil {
			handleError(w, err)
		}
	}
}
