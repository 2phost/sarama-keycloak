package saramakeycloak

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/Nerzal/gocloak/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	keycloackRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "keycloak_request_duration_seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "result"},
	)
)

// Keycloak wraps method to interact with keycloak server.
type Keycloak interface {
	// LoginClient sends a request to the token endpoint using client credentials.
	LoginClient(clientID, clientSecret, realm string) (*gocloak.JWT, error)

	// RefreshToken used to refresh the token.
	RefreshToken(refreshToken string, clientID, clientSecret, realm string) (*gocloak.JWT, error)

	RestyClient() *resty.Client
}

type keycloakWithMetrics struct {
	k Keycloak
}

func (k *keycloakWithMetrics) RestyClient() *resty.Client {
	return k.k.RestyClient()
}

func (k *keycloakWithMetrics) LoginClient(clientID, clientSecret, realm string) (*gocloak.JWT, error) {
	start := time.Now()

	jwt, err := k.k.LoginClient(clientID, clientSecret, realm)

	keycloackRequestDuration.WithLabelValues("LoginClient", resultLabel(err)).Observe(time.Since(start).Seconds())

	return jwt, err
}

func (k *keycloakWithMetrics) RefreshToken(refreshToken string, clientID, clientSecret, realm string) (*gocloak.JWT, error) {
	start := time.Now()

	jwt, err := k.k.RefreshToken(refreshToken, clientID, clientSecret, realm)

	keycloackRequestDuration.WithLabelValues("RefreshToken", resultLabel(err)).Observe(time.Since(start).Seconds())

	return jwt, err
}

func resultLabel(err error) string {
	if err != nil {
		return "error"
	}

	return "ok"
}
