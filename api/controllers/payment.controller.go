package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"dendrix.io/nayalabs/reportserver/models"
	"dendrix.io/nayalabs/reportserver/services"
	"dendrix.io/nayalabs/reportserver/utils"
	"github.com/gorilla/mux"
)

type payment struct {
	services.PaymentService
}

var (
	encoder  utils.Encoder
	decoder  utils.Decoder
	timeUtil utils.TimeService
)

func init() {
	encoder = utils.NewJSONEncoder()
	decoder = utils.NewJSONDecoder()
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
	//r.HandleFunc(sb.String(), p.getAllHandler).Methods("GET")
	//r.HandleFunc(sb.String()+fmt.Sprintf("/{%s}", pathParam), p.getPathTextHandler).Methods("GET")
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
		err0 := encoder.Encode(w, payments)
		if err0 != nil {
			errMsg := fmt.Sprintf("{ error: %s }", err0.Error())
			w.Write([]byte(errMsg))
		}
	}
}

func (p *payment) getByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payment, err := p.PaymentService.GetByID(r.Context(), params[pathParam])
	if err != nil {
		errMsg := fmt.Sprintf("{ error: %s }", err.Error())
		w.Write([]byte(errMsg))
	} else {
		err0 := encoder.Encode(w, payment)
		if err0 != nil {
			errMsg := fmt.Sprintf("{ error: %s }", err0.Error())
			w.Write([]byte(errMsg))
		}
	}
}

func (p *payment) createHandler(w http.ResponseWriter, r *http.Request) {
	var in models.Payment
	if err := decoder.Decode(r.Body, &in); err != nil {
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
		err0 := encoder.Encode(w, payment)
		if err0 != nil {
			errMsg := fmt.Sprintf("{ error: %s }", err0.Error())
			w.Write([]byte(errMsg))
		}
	}
}
