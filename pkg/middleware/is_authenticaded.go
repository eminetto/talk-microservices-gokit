package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var ErrUnauthorized = errors.New("unauthorized")

func IsAuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		errorMessage := "Authentication error"
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			respondWithError(rw, http.StatusUnauthorized, ErrUnauthorized.Error(), errorMessage)
			return
		}
		payload := `{
			"token": "` + tokenString + `"
		}`
		//@TODO use env var
		req, err := http.Post("http://localhost:8081/v1/validate-token", "text/plain", strings.NewReader(payload))
		if err != nil {
			respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
			return
		}
		if req.StatusCode == http.StatusUnauthorized {
			respondWithError(rw, http.StatusUnauthorized, ErrUnauthorized.Error(), errorMessage)
			return
		}
		defer req.Body.Close()
		type result struct {
			Email string `json:"email"`
		}
		var res result
		err = json.NewDecoder(req.Body).Decode(&res)
		if err != nil {
			respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
			return
		}
		r.Header.Add("email", res.Email)

		next.ServeHTTP(rw, r)
	})
}

//RespondWithError return a http error
func respondWithError(w http.ResponseWriter, code int, e string, message string) {
	respondWithJSON(w, code, map[string]string{"code": strconv.Itoa(code), "error": e, "message": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
