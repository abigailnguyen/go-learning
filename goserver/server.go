package main

import (
	"context"
	"crypto/tls"
	"io"
	"mime"
	"net/http"
	"os"
	"time"

	"bitbucket.org/papercutsoftware/gopapercut/log"
	"github.com/gobuffalo/packr"
	"github.com/goji/httpauth"
	"github.com/gorilla/handlers"
)

// Resources to understand further the patterns in golang
// https://www.integralist.co.uk/posts/understanding-golangs-func-type/
// https://www.alexedwards.net/blog/making-and-using-middleware
// https://go.dev/blog/context

func main() {
	mux := http.NewServeMux()

	box := packr.NewBox("../frontend/connector-ui/build")

	log.Info(box.List())
	mux.Handle("/", http.FileServer(box))

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Error(err.Error())
	}

	loggingHandler := newLoggingHandler(logFile)

	authenticationHandler := httpauth.SimpleBasicAuth("alice", "pa$$word")
	finalHandler := http.HandlerFunc(final)

	mux.Handle("/", loggingHandler(authenticationHandler(enforceJSONHandler(finalHandler))))

	// authMw := &AuthorizationMiddleware{AuthToken: "someauthtoken", allowLocal: true}
	authMw := new(AuthorizationMiddleware)
	authMw.AuthToken = "someauthtoken"
	authMw.allowLocal = true

	serverReadTimeout, serverWriteTimeout := 30*time.Second, 30*time.Second

	log.Info("Listening on :3000...")
	s := &http.Server{
		Addr: ":3000",
		// http.Handler interface implements ServeHTTP, if you require that all routes must authorize
		Handler:      authMw.Handler(mux),
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
			MinVersion: tls.VersionTLS12,
		},
	}
	// err = http.ListenAndServe(":3000", mux)
	err = s.ListenAndServe()
	log.Error(err.Error())

}

func newLoggingHandler(dst io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(dst, h)
	}
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// one way to return the handler is to pass the previous handler
func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// Context example
func handleSearch(w http.ResponseWriter, req *http.Request) {
	// ctx is the Context for this handler. Calling cancel closes the
	// ctx.Done channel, which is the cancellation signal for requests
	// started by this handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires.
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Cancel ctx as soon as handleSearch returns.

	// Check the search query.
	query := req.FormValue("q")

	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	// Store the user IP in ctx for use by code in other packages.
	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)

	// Run the Google search and print the results.
	start := time.Now()
	results, err := google.Search(ctx, query)
	elapsed := time.Since(start)

	if err := resultsTemplate.Execute(w, struct {
		Results          google.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
}
