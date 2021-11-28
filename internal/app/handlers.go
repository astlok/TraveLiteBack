package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"travalite/configs"
	"travalite/internal/app/profile"
	"travalite/internal/app/region"
	"travalite/internal/app/session"
	"travalite/internal/app/trek"
	"travalite/pkg/middleware"
)

func ConfigureRoute(config configs.Config, postgres *sqlx.DB) *mux.Router {
	sessionRepo := session.NewRepo(postgres)
	sessionUseCase := session.NewUseCase(*sessionRepo)
	sessionMiddleware := session.NewMiddleware(*sessionUseCase)
	sessionHandler := session.NewHandlers(*sessionUseCase)

	profileRepo := profile.NewRepo(postgres)
	profileUseCase := profile.NewUseCase(*profileRepo, *sessionRepo)
	profileHandler := profile.NewHandler(*profileUseCase, *sessionUseCase)

	regionRepo := region.NewRepo(postgres)
	regionUseCase := region.NewUseCase(*regionRepo)
	regionHandler := region.NewHandler(*regionUseCase)

	trekRepo := trek.NewRepo(postgres)
	trekUseCase := trek.NewUseCase(*trekRepo)
	trekHandler := trek.NewHandler(*trekUseCase)

	router := mux.NewRouter()

	apiV1Prefix := router.PathPrefix("/api/v1").Subrouter()
	apiV1Prefix.Use(middleware.LoggingRequest)

	apiV1Prefix.HandleFunc("/login", profileHandler.AuthProfile).Methods(http.MethodPost)
	apiV1Prefix.HandleFunc("/logout", sessionHandler.LogOut).Methods(http.MethodDelete)
	apiV1Prefix.HandleFunc("/profile", profileHandler.RegistrationProfile).Methods(http.MethodPost)

	profilePrefix := apiV1Prefix.PathPrefix("/profile").Subrouter()

	profilePrefix.Use(sessionMiddleware.CheckSession)
	profilePrefix.HandleFunc("", profileHandler.ChangeProfile).Methods(http.MethodPatch)
	profilePrefix.HandleFunc("/{id:[0-9]+}", profileHandler.GetProfile).Methods(http.MethodGet)
	profilePrefix.HandleFunc("/{id:[0-9]+/avatar", profileHandler.PutAvatar).Methods(http.MethodGet)

	trekPrefix := apiV1Prefix.PathPrefix("/trek").Subrouter()

	trekPrefix.Use(sessionMiddleware.CheckSession)
	trekPrefix.HandleFunc("", trekHandler.CreateTrek).Methods(http.MethodPost)
	trekPrefix.HandleFunc("", trekHandler.GetTreks).Methods(http.MethodGet)
	trekPrefix.HandleFunc("/{id:[0-9]+}", trekHandler.GetTrekInfo).Methods(http.MethodGet)
	trekPrefix.HandleFunc("/{id:[0-9]+}", trekHandler.ChangeTrek).Methods(http.MethodPatch)
	trekPrefix.HandleFunc("/{id:[0-9]+}", trekHandler.DelTrek).Methods(http.MethodDelete)
	trekPrefix.HandleFunc("/profile/{id:[0-9]+}", trekHandler.GetUsersTreks).Methods(http.MethodGet)
	trekPrefix.HandleFunc("/search", trekHandler.SearchTrek).Methods(http.MethodGet)
	trekPrefix.HandleFunc("/{id:[0-9]+}/comments", trekHandler.CreatComment).Methods(http.MethodPost)
	trekPrefix.HandleFunc("/{id:[0-9]+}/comments", trekHandler.GetTrekComments).Methods(http.MethodGet)
	trekPrefix.HandleFunc("/{id:[0-9]+}/rate", trekHandler.RateTrek).Methods(http.MethodPost)

	regionPrefix := apiV1Prefix.PathPrefix("/region").Subrouter()

	regionPrefix.HandleFunc("/{id:[0-9]+}", regionHandler.GetRegionInfo).Methods(http.MethodGet)
	regionPrefix.HandleFunc("", regionHandler.GetRegions).Methods(http.MethodGet)

	return router
}
