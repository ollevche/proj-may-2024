package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	s := NewStorage()

	trips := TripResource{
		s: s,
	}

	users := UserResource{
		s: s,
	}

	auth := Auth{
		s: s,
	}

	mux.HandleFunc("POST /users", users.CreateOne)
	mux.HandleFunc("GET /trips", auth.checkAuth(trips.GetAll))
	mux.HandleFunc("POST /trips", auth.checkAuth(trips.CreateOne))
	mux.HandleFunc("DELETE /trips/{id}", auth.checkAuth(trips.DeleteOne))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Failed to listen and serve: %v\n", err)
	}
}

// TODO: filter trips by auth user from auth middleware
type TripResource struct {
	s *Storage
}

func (t *TripResource) GetAll(w http.ResponseWriter, r *http.Request) {
	trips := t.s.GetAllTrips()

	err := json.NewEncoder(w).Encode(trips)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *TripResource) CreateOne(w http.ResponseWriter, r *http.Request) {
	var trip Trip

	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	trip.DestinationWeather, err = getWeatherByCity(trip.Destination)
	if err != nil {
		fmt.Printf("Failed to get weather: %v\n", err)
	}

	trip.ID = t.s.CreateOneTrip(trip)

	err = json.NewEncoder(w).Encode(trip)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (t *TripResource) DeleteOne(w http.ResponseWriter, r *http.Request) {
	idVal := r.PathValue("id")

	// Parsing: "123" -> 123
	tripID, err := strconv.Atoi(idVal)
	if err != nil {
		fmt.Printf("Invalid id param: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := t.s.DeleteTripByID(tripID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

type UserResource struct {
	s *Storage
}

func (ur *UserResource) CreateOne(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := ur.s.CreateOneUser(user)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func getWeatherByCity(city string) (string, error) {
	const baseURL = "https://api.weatherbit.io/v2.0/current"

	r, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}

	token := os.Getenv("WEATHER_API_TOKEN")
	if token == "" {
		return "", errors.New("no WEATHER_API_TOKEN set")
	}

	q := r.URL.Query()
	q.Add("key", token)
	q.Add("city", city)
	r.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", fmt.Errorf("doing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("got non-ok response: %v", resp.Status)
	}

	type weather struct {
		Description string `json:"description"`
	}

	type dataArrEntry struct {
		Weather weather `json:"weather"`
	}

	type weatherResp struct {
		Data []dataArrEntry `json:"data"`
	}

	var respBody weatherResp

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return "", fmt.Errorf("decoding resp: %w", err)
	}

	if len(respBody.Data) == 0 {
		return "", fmt.Errorf("no data in weather response")
	}

	return respBody.Data[0].Weather.Description, nil
}
