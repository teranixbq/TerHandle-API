package gmaps

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type DistanceMatrixResponse struct {
	Rows []struct {
		Elements []struct {
			Distance struct {
				Value float64 `json:"value"`
			} `json:"distance"`
		} `json:"elements"`
	} `json:"rows"`
}

func CalculateDistanceFromGMaps(userLat, userLng, teknisiLat, teknisiLng float64) (float64, error) {
	godotenv.Load()
	url := buildDistanceURL(userLat, userLng, teknisiLat, teknisiLng)

	response, err := getDistanceResponse(url)
	if err != nil {
		return 0, err
	}

	distance, err := extractDistance(response)
	if err != nil {
		return 0, err
	}

	return distance, nil
}

func buildDistanceURL(userLat, userLng, teknisiLat, teknisiLng float64) string {
	apiKey := os.Getenv("API_KEY")
	return fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%f,%f&destinations=%f,%f&key=%s", userLat, userLng, teknisiLat, teknisiLng, apiKey)
}

func getDistanceResponse(url string) (DistanceMatrixResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DistanceMatrixResponse{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return DistanceMatrixResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return DistanceMatrixResponse{}, fmt.Errorf("gagal menghitung jarak. Status code: %d", resp.StatusCode)
	}

	var response DistanceMatrixResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return DistanceMatrixResponse{}, err
	}

	return response, nil
}

func extractDistance(response DistanceMatrixResponse) (float64, error) {
	if len(response.Rows) == 0 || len(response.Rows[0].Elements) == 0 {
		return 0, fmt.Errorf("gagal menghitung jarak. Invalid response from Google Maps API")
	}

	distanceValue := response.Rows[0].Elements[0].Distance.Value

	distanceInKilometers := distanceValue / 1000
	roundedDistance := math.Round(distanceInKilometers*10) / 10

	return roundedDistance, nil
}
