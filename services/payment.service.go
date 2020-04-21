package services

import (
	"context"
	"log"

	"dendrix.io/nayalabs/reportserver/models"
)

//PaymentService implements the payment service functions
type PaymentService interface {
	GetAll(ctx context.Context) ([]*models.Payment, error)
	GetByID(ctx context.Context, trxnRef string) (*models.Payment, error)
	Create(ctx context.Context, p models.Payment) (*models.Payment, error)
	GetAllCSV(ctx context.Context, month int, year int) (string, error)
}

type paymentServiceImpl struct {
	IDataService
}

func (srv *paymentServiceImpl) GetAll(ctx context.Context) ([]*models.Payment, error) {
	log.Println("Called PaymentService.GetAll() function")
	out, err := srv.IDataService.GetAll("")
	if err != nil {
		log.Println("GetAll error occurred:- ")
		log.Printf("%v", err)
		return nil, err
	}
	response := out.([]*models.Payment)
	return response, nil
}

func (srv *paymentServiceImpl) GetByID(ctx context.Context, id string) (*models.Payment, error) {
	log.Println("Called PaymentService.GetByID() function")
	out, err := srv.IDataService.GetById(id)
	if err != nil {
		log.Println("GetByID error occurred:- ")
		log.Printf("%v", err)
		return nil, err
	}
	response := out.(*models.Payment)
	return response, nil
}

func (srv *paymentServiceImpl) Create(ctx context.Context, p models.Payment) (*models.Payment, error) {
	log.Println("Called PaymentService.Create() function")
	srv.IDataService.BeginTx()
	log.Println("Called IDataService.BeginTx() function")
	out, err := srv.IDataService.Create(p)
	if err != nil {
		log.Println("Create error occurred:- ")
		log.Printf("%v", err)
		return nil, err
	}
	if err := srv.IDataService.Commit(); err != nil {
		log.Println("Commit error occurred:- ")
		log.Printf("%v", err)
		return nil, err
	}
	response := out.(*models.Payment)
	return response, nil
}

func (srv *paymentServiceImpl) GetAllCSV(ctx context.Context, month int, year int) (string, error) {
	log.Println("Called PaymentService.GetAllCSV() function")
	out, err := srv.IDataService.GetAllCSV(month, year)
	if err != nil {
		log.Println("GetAllCSV error occurred:- ")
		log.Printf("%v", err)
		return "", err
	}
	response := out.(string)
	return response, nil
}

//NewPaymentService returns an implementation of PaymentService
func NewPaymentService(data IDataService) PaymentService {
	ps := new(paymentServiceImpl)
	ps.IDataService = data
	return ps
}
