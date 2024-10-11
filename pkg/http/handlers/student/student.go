package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

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

		//json data to serialize means struct ke ander daal paaye
		response.WriteJson(w, http.StatusCreated, map[string]string{"sucess": "student created"})
	}
}
