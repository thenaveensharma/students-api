package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/thenaveensharma/students-api/internal/types"
	"github.com/thenaveensharma/students-api/internal/utils/response"
)

func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		fmt.Println(student)
		if errors.Is(err, io.EOF) {
			w.WriteHeader(http.StatusBadRequest)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		slog.Info("creating new student")

		if err != nil {

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return

		}
		//validation
		if err := validator.New(validator.WithRequiredStructEnabled()).Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return

		}
		fmt.Println(err)
		response.WriteJson(w, http.StatusCreated, map[string]string{
			"success": "OK",
		})
	}
}
