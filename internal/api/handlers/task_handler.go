package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gauravst/todo-backend-go/internal/models"
	"github.com/gauravst/todo-backend-go/internal/services"
	"github.com/gauravst/todo-backend-go/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func CreateTask(taskService services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Request validation
		err = validator.New().Struct(task)
		if err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		// call here services

		err = taskService.CreateTask(&task)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// return response
		response.WriteJson(w, http.StatusCreated, task)
	}
}

func GetTask(taskService services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		if id == "" {
			data, err := taskService.GetAllTask()
			if err != nil {
				response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
				return
			}
			response.WriteJson(w, http.StatusOK, data)
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		data, err := taskService.GetTaskByID(idInt)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, data)
	}
}

func UpdateTask(taskService services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("id parms not found")))
			return
		}

		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Request validation
		err = validator.New().Struct(task)
		if err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		err = taskService.UpdateTask(idInt, &task)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// return response
		response.WriteJson(w, http.StatusOK, task)
	}
}

func DeleteTask(taskService services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		if id == "" {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("id parms not found")))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		err = taskService.DeleteTask(idInt)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// return response
		response.WriteJson(w, http.StatusOK, "task deleted")

	}
}
