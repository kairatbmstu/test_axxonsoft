package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"example.com/test_axxonsoft/v2/dto"
)

func PostTask(taskDto dto.TaskDTO) {

}

func GetTask(id string) (dto.TaskDTO, error) {
	fmt.Println("id : " + id)
	var taskDto = dto.TaskDTO{
		Id:             id,
		Method:         "GET",
		HttpStatusCode: "OK",
		Url:            "http://google.com",
		ResponseBody:   "{message:ok}",
	}
	return taskDto, nil
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		if r.URL.Path == "/task" {
			http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
			return
		}
		id := strings.TrimPrefix(r.URL.Path, "/task/")
		if id == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}
		taskDto, err := GetTask(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody, err := json.Marshal(taskDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	case "POST":
		decoder := json.NewDecoder(r.Body)
		taskDto := dto.TaskDTO{}
		err := decoder.Decode(&taskDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		PostTask(taskDto)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}
