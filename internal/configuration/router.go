package configuration

import (
	"errors"
	"net/url"
)

type Router struct {
	Matcher    string
	Target     Target
	Scheme     Scheme
	Method     Method
	Middleware string
}

type Scheme string
type Method string

const (
	GET   Method = "GET"
	POST  Method = "POST"
	HTTP  Scheme = "HTTP"
	HTTPS Scheme = "HTTPS"
)

func (r *Router) Validate() error {
	if err := r.validateMatcher(); err != nil {
		return err
	}

	if err := r.validateScheme(); err != nil {
		return err
	}

	if err := r.Target.validateTarget(); err != nil {
		return err
	}

	if err := r.validateMethod(); err != nil {
		return err
	}

	if err := r.validateMiddleware(); err != nil {
		return err
	}

	return nil
}

func (r *Router) validateMatcher() error {
	if r.Matcher == "" {
		return errors.New("Matcher is missing")
	}

	_, err := url.ParseRequestURI(r.Matcher)

	if err != nil {
		return errors.New("invalid Matcher value: not a valid URL path")
	}

	return nil
}

func (r *Router) validateScheme() error {
	validSchemes := []string{"HTTP", "HTTPS"}

	for _, scheme := range validSchemes {

		if string(r.Scheme) == scheme {
			return nil
		}

	}

	return errors.New("invalid Scheme value")
}

func (r *Router) validateMethod() error {
	validMethods := []string{"GET", "POST"}

	for _, method := range validMethods {

		if string(r.Method) == method {
			return nil
		}

	}

	return errors.New("invalid Method value")
}

func (r *Router) validateMiddleware() error {
	if r.Middleware == "" {
		return nil
	}

	return nil
}
