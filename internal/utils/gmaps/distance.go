package gmaps

import (
    "encoding/json"
    "fmt"
    "io"
    "math"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

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

func getDistanceResponse(url string) (map[string]interface{}, error) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("gagal menghitung jarak. Status code: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var response map[string]interface{}
    err = json.Unmarshal(body, &response)
    if err != nil {
        return nil, err
    }

    return response, nil
}

func extractDistance(response map[string]interface{}) (float64, error) {
    rows, rowsExist := response["rows"].([]interface{})
    if !rowsExist || len(rows) == 0 {
        return 0, fmt.Errorf("gagal menghitung jarak. Invalid response from Google Maps API")
    }

    elements, elementsExist := rows[0].(map[string]interface{})["elements"].([]interface{})
    if !elementsExist || len(elements) == 0 {
        return 0, fmt.Errorf("gagal menghitung jarak. Invalid response from Google Maps API")
    }

    distance, distanceExist := elements[0].(map[string]interface{})["distance"].(map[string]interface{})
    if !distanceExist {
        return 0, fmt.Errorf("gagal menghitung jarak. Invalid response from Google Maps API")
    }

    distanceValue, valueExist := distance["value"].(float64)
    if !valueExist {
        return 0, fmt.Errorf("gagal menghitung jarak. Invalid response from Google Maps API")
    }

    distanceInKilometers := float64(distanceValue) / 1000
    roundedDistance := math.Round(distanceInKilometers*10) / 10

    return roundedDistance, nil
}
