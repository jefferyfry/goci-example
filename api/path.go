package api

import (
"github.com/gorilla/mux"
"net/http"
)

func GetApiService() (http.Handler) {
	handler := GetMathHandler()

	apiService := mux.NewRouter();
	apiService.HandleFunc("/",handler.welcome);
	apiService.Methods(http.MethodGet).Path("/add").HandlerFunc(handler.addition)
	apiService.Methods(http.MethodGet).Path("/subtract").HandlerFunc(handler.subtraction)
	apiService.Methods(http.MethodGet).Path("/multiply").HandlerFunc(handler.multiplication)
	apiService.Methods(http.MethodGet).Path("/divide").HandlerFunc(handler.division)

	return apiService;
}

func StartApiService() error {
	return http.ListenAndServe(":8091", GetApiService())
}
