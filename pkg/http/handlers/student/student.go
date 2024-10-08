package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/SAHIL-Sharma21/students-management/pkg/types"
	"github.com/SAHIL-Sharma21/students-management/pkg/utils/response"
)

// all crud operations
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}

		slog.Info("creating a student")

		//json data to serialize means struct ke ander daal paaye

		// w.Write([]byte("welcome to students management"))
		response.WriteJson(w, http.StatusCreated, student)
	}
}
