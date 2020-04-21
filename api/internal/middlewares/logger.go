package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/elenduuche/reportserver/logging"
	"github.com/elenduuche/reportserver/utils"
)

var logger logging.Logger
var timeUtil utils.TimeService

func init() {
	logger = logging.NewLogger()
	timeUtil = utils.NewTimeService()
}

func LoggingMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		service := os.Getenv("APP_NAME")
		ctx := context.WithValue(r.Context(), "start-time", timeUtil.Now())
		req := r.WithContext(ctx)
		deadline, _ := ctx.Deadline()
		startTime := time.Now()
		dur := time.Since(startTime)
		go func() {
			log.Printf("[logger] Service:%s	Method:%s	RequestURI:%s	Request-deadline:%s	start:%s	duration:%v \n", service, r.Method, r.URL.RequestURI(), deadline.UTC().String(), startTime.UTC().String(), dur)
		}()
		next.ServeHTTP(w, req)
	})
}
