package router

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func ForwardHandler(w http.ResponseWriter, r *http.Request) {
	// Get the URL and other information from the incoming request
	// urlString := r.URL.String()
	method := r.Method
	headers := r.Header
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, log.Default().Prefix()+" ERROR while reading request body", http.StatusBadRequest)
		return
	}

	// Parse the target URL
	targetUrl, err := url.Parse("https://google.com")
	if err != nil {
		http.Error(w, log.Default().Prefix()+" ERROR while parsing targetUrl", http.StatusFailedDependency)
		return
	}

	// Create a new request
	targetRequest := &http.Request{
		Method: method,
		URL:    targetUrl,
		Header: headers,
		Body:   ioutil.NopCloser(bytes.NewBuffer(body)),
	}

	// Create a new client
	httpClient := &http.Client{}

	// Process the response
	targetResponse, err := httpClient.Do(targetRequest)
	if err != nil {
		http.Error(w, log.Default().Prefix()+" ERROR while processing the response from target", http.StatusFailedDependency)
		return
	}
	defer targetResponse.Body.Close()

	// Set the status code and headers for the original response
	targetResponseBody, err := ioutil.ReadAll(targetResponse.Body)
	if err != nil {
		http.Error(w, log.Default().Prefix()+" ERROR while reading response body", http.StatusFailedDependency)
		return
	}

	// Write the response body back to original client
	w.Write(targetResponseBody)
}
