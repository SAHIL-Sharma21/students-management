package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/SAHIL-Sharma21/students-management/pkg/storage"
	"github.com/SAHIL-Sharma21/students-management/pkg/types"
	"github.com/SAHIL-Sharma21/students-management/pkg/utils/response"
	"github.com/go-playground/validator/v10"
)

// all crud operations
func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty request body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//validate the request -> in prodution we need to do this important
		if err := validator.New().Struct(student); err != nil {
			validateErr := err.(validator.ValidationErrors) //type caste
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErr))
			return
		}

		//create student
		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		slog.Info("Student created successfully", slog.String("userId", fmt.Sprint(lastId)))

		//json data to serialize means struct ke ander daal paaye
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		slog.Info("Getting user by Id", slog.String("userId", id))

		//converting str to int64
		userId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			slog.Error("Error converting string to int64", slog.String("id", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(userId)

		if err != nil {
			slog.Error("Error getting user by id", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		slog.Info("User found with the given id", slog.String("id", id))
		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetListOfStudents(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting list of students")

		students, err := storage.GetListOfStudents()

		if err != nil {
			slog.Error("Error getting list of students", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, err)
		}

		slog.Info("List of students fetched successfully", slog.String("count", fmt.Sprintf("%d", len(students))))
		response.WriteJson(w, http.StatusOK, students)
	}
}

func DeleteById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
