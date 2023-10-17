package api

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// The Config struct holds general configuration options for the application
type Config struct {
	AuthenticationToken string
	Port                string
	PhraseBookLocation  string
}

// A global Config struct for use in bootstrapping and authenticating
var config *Config

// A context key for auth
type key string

const appContextAuthenticationFound key = "key"

// LoadConfig will load up a new configuration struct with sane defaults if none provided
func LoadConfig() *Config {
	if config != nil {
		return config
	}
	config = &Config{}
	config.AuthenticationToken = envHelper("GO_EO_AUTHTOKEN", "")
	config.Port = envHelper("GO_EO_API_PORT", "8081")
	config.PhraseBookLocation = envHelper("GO_EO_PHRASEBOOK_DIR", "./api/")

	if config.AuthenticationToken == "" {
		//randomize it with bcrypt on each server start up and prompt the user to specify one
		r1 := rand.Intn(100000000)
		r2 := rand.Intn(20000000)
		plain := fmt.Sprintf("%s-%d-%d", "go-esperanto", r1, r2)
		h := md5.New()
		h.Write([]byte(plain))
		code := string(fmt.Sprintf("%x", h.Sum(nil)))
		config.AuthenticationToken = code

		// write to std out what the token is so that is can be used
		fmt.Println("")
		fmt.Printf("===================================================================\n")
		fmt.Printf("\t\t\tGenerated Token: %s\n", config.AuthenticationToken)
		fmt.Printf("If you did not expect to see this message, check your configuration\n")
		fmt.Printf("===================================================================\n")
	}

	config.Port = fmt.Sprintf(":%s", config.Port)

	// ensure a trailing slash on the path
	if !strings.HasSuffix(config.PhraseBookLocation, "/") {
		config.PhraseBookLocation = config.PhraseBookLocation + "/"
	}

	// load the phrase book
	if len(phrases) == 0 {
		LoadPhrasebook()
	}

	return config
}

// SetupRouter sets up and returns a new chi router
func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// TODO: add rate limiter

	// setup some middlewares
	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(120 * time.Second))
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Types", "X-CSRF-TOKEN", "X-API-TOKEN", "X-API-SERVICE", "RANGE", "ACCEPT-RANGE"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
	r.Use(checkTokenMiddleware)

	// routes

	r.Get("/", GetAllPhrasesRoute)
	r.Get("/random", GetRandomPhraseRoute)

	return r
}

// checkTokenMiddleware is a helper to check for the API token in the HTTP headers
func checkTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticated := false
		key := r.Header.Get("X-API-TOKEN")
		// TODO: allow checking in basic auth or cookies

		if key == config.AuthenticationToken {
			authenticated = true
		}

		ctx := context.WithValue(r.Context(), appContextAuthenticationFound, authenticated)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// checkAuthenticatedHelper is a helper to check the authentication context
func checkAuthenticatedHelper(w *http.ResponseWriter, r *http.Request) (valid bool) {
	valid = r.Context().Value(appContextAuthenticationFound).(bool)
	return
}

// envHelper is a helper for getting environment variables
func envHelper(key, defaultMissing string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultMissing
	}
	return val
}
