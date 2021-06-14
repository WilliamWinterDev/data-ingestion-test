package models

// Metric : The metric which is stored in the Mongo database
type Metric struct {
	Timestamp   int32   `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	CpuLoad     float64 `json:"cpu_load,omitempty" bson:"cpu_load,omitempty"`
	Concurrency float64 `json:"concurrency,omitempty" bson:"concurrency,omitempty"`
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}
