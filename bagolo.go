// Naranza Bagolo, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package bagolo

import (
  "encoding/base64"
  "net/http"
  "strings"
  "errors"
)

const Version = "1.2025.2"

func Auth(r *http.Request) (string, string, error) {
  auth := r.Header.Get("Authorization")
  if auth == "" || !strings.HasPrefix(auth, "Basic ") {
    return "", "", errors.New("missing or invalid auth header")
  }

  payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
  if err != nil {
    return "", "", errors.New("malformed base64 in auth header")
  }

  parts := strings.SplitN(string(payload), ":", 2)
  if len(parts) != 2 {
    return "", "", errors.New("invalid auth format")
  }

  return parts[0], parts[1], nil
}
