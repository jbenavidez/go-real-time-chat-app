package main

import "net/http"

func (app *application) Tester(w http.ResponseWriter, r *http.Request) {
	resp := JSONResponse{
		Error:   false,
		Message: "Welcome",
	}
	_ = app.writeJSON(w, http.StatusOK, resp)
}
