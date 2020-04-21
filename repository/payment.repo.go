package repository

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"dendrix.io/nayalabs/reportserver/models"
	"dendrix.io/nayalabs/reportserver/services"
)

type paymentRepository struct {
	tx *sql.Tx
}

//NewPaymentRepository return the payments repo
func NewPaymentRepository() services.IDataService {
	return new(paymentRepository)
}

func (repo *paymentRepository) Initialize() error {
	return nil
}

func (repo *paymentRepository) BeginTx() error {
	if err := db.Ping(); err != nil {
		log.Println("xxxxxxxxx  db.Ping() failed  xxxxxxxxxxx")
		log.Fatal(err)
	} else {
		log.Println("Pinged DB successfully....")
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	repo.tx = tx
	return nil
}

func (repo *paymentRepository) Commit() error {
	return repo.tx.Commit()
}

func (repo *paymentRepository) Create(r interface{}) (interface{}, error) {
	request := r.(models.Payment)
	query := "INSERT INTO payments (trxnref,senderid,receiverid,amount,currency,narration,createdon) VALUES( $1, $2, $3, $4, $5, $6, $7);"
	result, err := repo.tx.Exec(query, request.TrxnRef, request.SenderID, request.ReceiverID, request.Amount,
		request.Currency, request.Narration, request.CreatedOn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err0 := result.RowsAffected()
	if err0 != nil {
		return nil, err0
	}
	//idx, err1 := result.LastInsertId()
	//if err1 != nil {
	//	return nil, err1
	//}

	//request.ID = strconv.FormatInt(idx, 10)
	idx := 0
	log.Println(fmt.Sprintf("Inserted rows=%v and record with ID=%v", rows, idx))
	response := request
	return &response, nil
}

func (repo *paymentRepository) Update(id string, r interface{}) (interface{}, error) {
	request := r.(models.Payment)
	query := `UPDATE payments SET trxnref = ?, senderid = ?, receiverid = ?, amount = ?, currency = ?, narration = ? WHERE id = ?`
	result, err := repo.tx.Exec(query, request.TrxnRef, request.SenderID, request.ReceiverID, request.Amount,
		request.Currency, request.Narration, request.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err0 := result.RowsAffected()
	if err0 != nil {
		return nil, err0
	}
	idx, err1 := result.LastInsertId()
	if err1 != nil {
		return nil, err1
	}
	request.ID = strconv.FormatInt(idx, 10)
	response := request
	log.Println(fmt.Sprintf("Updated affected %v rows and record with ID=%v", rows, id))
	return &response, nil
}

func (repo *paymentRepository) GetById(id string) (interface{}, error) {
	response := new(models.Payment)
	query := "SELECT id, trxnref, senderid, receiverid, amount, currency, narration, createdon FROM payments WHERE id=$1"
	//idx, _ := strconv.Atoi(id)
	err := repo.tx.QueryRow(query, id).Scan(&response.ID, &response.TrxnRef, &response.SenderID, &response.ReceiverID,
		&response.Amount, &response.Currency, &response.Narration, &response.CreatedOn)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (repo *paymentRepository) GetAll(params string) (interface{}, error) {
	var payments []*models.Payment
	query := "SELECT id, trxnref, senderid, receiverid, amount, currency, narration, createdon FROM payments"
	rows, err := repo.tx.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := new(models.Payment)
		if err := rows.Scan(&p.ID, &p.TrxnRef, &p.SenderID, &p.ReceiverID, &p.Amount, &p.Currency, &p.Narration, &p.CreatedOn); err != nil {
			log.Println(err)
			return nil, err
		}
		payments = append(payments, p)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}
	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Println("Last known error in Rows scan...")
		log.Println(err)
	}
	return payments, nil
}

func (repo *paymentRepository) GetAllCSV() (interface{}, error) {
	var filePath string
	query := "SELECT id, trxnref, senderid, receiverid, amount, currency, narration, createdon FROM payments"
	rows, err := repo.tx.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//Get Write to CSV file
	staticDir := os.Getenv("STATIC_FILEPATH")
	filePath = path.Join(staticDir, "payments_report.csv")
	csvFile, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	csvWriter := csv.NewWriter(csvFile)
	if werr := csvWriter.Write([]string{"id", "trxnref", "senderid", "receiverid", "amount", "currency", "narration", "createdon"}); werr != nil {
		return nil, werr
	}
	for rows.Next() {
		var (
			id         string
			trxnRef    string
			senderID   string
			receiverID string
			amount     string
			currency   string
			narration  string
			createdOn  *time.Time
		)
		if err := rows.Scan(&id, &trxnRef, &senderID, &receiverID, &amount, &currency, &narration, &createdOn); err != nil {
			log.Println(err)
			return nil, err
		}
		if werr := csvWriter.Write([]string{id, trxnRef, senderID, receiverID, amount, currency, narration, createdOn.String()}); werr != nil {
			return nil, werr
		}
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}
	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Println("Last known error in Rows scan...")
		log.Println(err)
	}
	csvWriter.Flush()
	defer csvFile.Close()
	defer rows.Close()
	return filePath, nil
}
