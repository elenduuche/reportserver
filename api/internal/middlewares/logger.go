package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"dendrix.io/nayalabs/reportserver/logging"
)

var logger logging.Logger

func init() {
	logger = logging.NewLogger()
}

func LoggingMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//logger.Info("## Called interceptor (Logger) 1")
		service := os.Getenv("APP_NAME")
		ctx := context.WithValue(r.Context(), "start-time", time.Now())
		//msg := fmt.Sprintf("##STARTED## %s request from %s to %s", r.Method, r.RemoteAddr, r.RequestURI)
		//log.Println(msg)
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
