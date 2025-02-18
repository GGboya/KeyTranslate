package apiservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const maxBackoff = 32

// doRequest 发送通用 HTTP 请求
func doRequest(method, url string, body interface{}) ([]byte, error) {
	var requestBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		requestBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	err = godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables from .env file: %w", err)
	}
	key := os.Getenv("KEY")
	if key == "" {
		return nil, fmt.Errorf("API key not found in environment variables")
	}
	req.Header.Set("Authorization", "Bearer "+key)
	// 添加一个重试机制，指数回避
	var resp *http.Response
	backoff := 1

	for backoff <= maxBackoff {
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("HTTP request failed: %w", err)
		}
		if resp.StatusCode == 429 { // Too Many Requests
			waitDuration := time.Second * time.Duration(backoff) // Default wait time
			time.Sleep(waitDuration)                             // Wait before retrying
			backoff <<= 1
			continue
		}
		break
	}

	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP request failed with status %d: %s", resp.StatusCode, resp.Status)
	}

	return io.ReadAll(resp.Body)
}
