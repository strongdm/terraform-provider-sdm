package releases

import (
	"io"
	"log"
	"time"
)

var (
	defaultInstallTimeout = 30 * time.Second
	defaultListTimeout    = 10 * time.Second
	discardLogger         = log.New(io.Discard, "", 0)
)
