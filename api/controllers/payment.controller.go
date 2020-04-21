package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"dendrix.io/nayalabs/reportserver/models"
	"dendrix.io/nayalabs/reportserver/services"
	"dendrix.io/nayalabs/reportserver/utils"
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
		errMsg := fmt.Sprintf("{ error: %s }", err.Error())
		w.Write([]byte(errMsg))
	} else {
		err0 := jsonEncoder.Encode(w, payments)
		if err0 != nil {
			errMsg := fmt.Sprintf("{ error: %s }", err0.Error())
			w.Write([]byte(errMsg))
		}
	}
}

func (p *payment) getAllCSVHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	month, cerr := strconv.Atoi(params["month"])
	if cerr != nil {
		log.Fatal(cerr)
	}
	filename, err := p.PaymentService.GetAllCSV(r.Context(), month, timeUtil.Now().Year())
	if err != nil {
		errMsg := fmt.Sprintf("{ error: %s }", err.Error())
		log.Fatal(errMsg)
	}
	file, err := os.Open(filename)
	if err != nil {
		errMsg := fmt.Sprintf("{ error: %s }", err.Error())
		log.Fatal(errMsg)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		errMsg := fmt.Sprintf("{ error: %s }", err.Error())
		log.Fatal(errMsg)
	}

	FileSize := strconv.FormatInt(fileInfo.Size(), 10)
	//Send the headers before sending the file
	m := time.Month(month)
	downloadFilename := fmt.Sprintf("payments_%v_%v.csv", m.String(), timeUtil.NowTimestamp().Seconds)
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
		errMsg := fmt.Sprintf("{ error: %s }", err.Error())
		w.Write([]byte(errMsg))
	} else {
		err0 := jsonEncoder.Encode(w, payment)
		if err0 != nil {
			errMsg := fmt.Sprintf("{ error: %s }", err0.Error())
			w.Write([]byte(errMsg))
		}
	}
}

func (p *payment) createHandler(w http.ResponseWriter, r *http.Request) {
	var in models.Payment
	if err := jsonDecoder.Decode(r.Body, &in); err != nil {
		errMsg := fmt.Sprintf("{ error: %s }", err.Error())
		w.Write([]byte(errMsg))
	}
	t := timeUtil.Now()
	in.CreatedOn = &t
	payment, err := p.PaymentService.Create(r.Context(), in)
	if err != nil {
		errMsg := fmt.Sprintf("{ error: %s }", err.Error())
		w.Write([]byte(errMsg))
	} else {
		err0 := jsonEncoder.Encode(w, payment)
		if err0 != nil {
			errMsg := fmt.Sprintf("{ error: %s }", err0.Error())
			w.Write([]byte(errMsg))
		}
	}
}
