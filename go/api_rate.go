/*
 * GSES BTC application
 * API version: 1.0.0
 */

package exchangerate

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Rate represents the USD to UAH conversion rate
type RateCurrency struct {
	Rate float64 `json:"rate"` // Specify the field name for JSON encoding
}

func Rate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rate, err := GetRate("USD")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	// Simulate fetching the rate from a third-party service (replace with actual implementation)
	//currentRate := 27.894 // Example rate (replace with actual logic)

	response := RateCurrency{Rate: rate} // Create a Rate object
	// Encode the Rate object to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "Error encoding response: %v", err)
		return
	}

	fmt.Printf("rate: %v", rate)
	//w.Write([]byte(fmt.Sprintf("%v", rate))) // nolint:errrate)

}
