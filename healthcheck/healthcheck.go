package healthcheck

import (
	"encoding/json"
	"strconv"
	"time"
)

func New(healthy bool, time time.Time, message string) HealthCheck {
	return HealthCheck{Healthy: healthy, Time: time, Message: message}
}

type HealthCheck struct {
	Healthy bool
	Time    time.Time
	Message string
}

func (hc HealthCheck) ToJson() ([]byte, error) {
	return json.Marshal(hc)
}

func (hc HealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Healthy string `json:"healthy"`
		Time    string `json:"time"`
		Message string `json:"message"`
	}{
		Healthy: strconv.FormatBool(hc.Healthy),
		Time:    hc.Time.Format(time.RFC3339),
		Message: hc.Message,
	})
}

// Move to own file
// Untested
type HealthCheckRepository struct {
	// make private and Constructor
	CurrentState HealthCheck
}

func (r *HealthCheckRepository) SetState(hc HealthCheck) {
	r.CurrentState = hc
}

func (r *HealthCheckRepository) GetState() HealthCheck {
	return r.CurrentState
}

func NewRepository(hc HealthCheck) HealthCheckRepository {
	return HealthCheckRepository{CurrentState: hc}
}
