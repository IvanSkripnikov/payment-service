package helpers

import (
	"database/sql"
	"errors"
	"net/http"

	logger "github.com/IvanSkripnikov/go-logger"
)

func checkError(w http.ResponseWriter, err error, category string) bool {
	httpStatusCode := http.StatusOK
	if err != nil {
		logger.Errorf("Runtime error %s", err.Error())

		var data ResponseData
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			httpStatusCode = http.StatusNotFound
			data = ResponseData{
				"error": "Data not found",
			}
		} else {
			httpStatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			data = ResponseData{
				"error": "Internal error",
			}
		}

		SendResponse(w, data, category, httpStatusCode)

		return true
	}

	return false
}
