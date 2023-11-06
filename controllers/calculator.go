package controllers

import (
	"401k_calculator/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Calculate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inputData models.Calculator401kSimple
		err := json.NewDecoder(r.Body).Decode(&inputData)

		if err != nil {
			// Log the error and respond with a detailed error message
			errorMessage := fmt.Sprintf("Failed to parse request body: %v", err)
			fmt.Println(errorMessage)

			// You can also log the content of inputData for debugging purposes
			inputDataJSON, _ := json.Marshal(inputData)
			fmt.Printf("Request Body: %s\n", inputDataJSON)

			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}
		inputData.CalculatePercentToHitMaxContribution()
		responseData := inputData.BuildPercentResponse()

		json.NewEncoder(w).Encode(responseData)
	}
}
