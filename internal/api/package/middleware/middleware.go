package middleware

import "gymshark-test/internal/api/package/service"

// Middleware describes a service middleware
type Middleware func(service service.PackageService) service.PackageService
