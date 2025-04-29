package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/thenaveensharma/students-api/internal/storage"
	"github.com/thenaveensharma/students-api/internal/types"
	"github.com/thenaveensharma/students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {

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
		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{
			"id": lastId,
		})
	}
}
func GetById(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		studentId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("error converting student id to int64", slog.String("id", id))

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id")))
			return
		}
		slog.Info("getting a student", slog.String("id", fmt.Sprint(studentId)))

		student, err := storage.GetStudentById(studentId)
		if err != nil {
			slog.Error("error getting student", slog.String("id", id))

			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}
