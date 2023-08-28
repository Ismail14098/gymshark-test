package healthcheck

import (
	"fmt"
	"net/http"
)

// Health check func
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "ok")
}
