package src

import (
	// Import underlying generator use
	_ "github.com/gokums/core"
	_ "github.com/gokums/core/net/httpx"
	_ "google.golang.org/grpc"
)

// GitSHA is set build time.
var GitSHA string
