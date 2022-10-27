package backend

type Backend struct {
	Name string `json:"name"`
	Spec Spec   `json:"spec"`
}

type Spec struct {
	CertificateName string          `json:"certificateName"`
	Http2Enabled    bool            `json:"http2Enabled"`
	Timeout         int             `json:"timeout"`
	Credentials     Credentials     `json:"credentials"`
	RetryConfig     RetryConfig     `json:"retryConfig"`
	CircuitBreakers CircuitBreakers `json:"circuitBreakers"`
}

type Credentials struct {
	Type   string
	Secret string
}

type RetryConfig struct {
	Count      int
	StatusCode int
}

type CircuitBreakers struct {
	MaxConnections     int
	MaxRequests        int
	MaxPendingRequests int
	MaxRetries         int
	MaxConnectionPools int
}
