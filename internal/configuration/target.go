package configuration

import (
	"errors"
	"net/url"
)

type Destination struct {
	Kind   Kind
	Target string
}

type Kind string

const (
	URL         Kind = "URL"
	kubeService Kind = "kubeService"
)

func (d *Destination) Validate() error {
	if err := d.validateKind(); err != nil {
		return err
	}

	if err := d.validateTarget(); err != nil {
		return err
	}

	return nil
}

func (d *Destination) validateKind() error {
	validKinds := []string{"URL", "kubeService"}

	for _, kind := range validKinds {

		if string(d.Kind) == kind {
			return nil
		}

	}

	return errors.New("invalid Kind value")
}

func (d *Destination) validateTarget() error {
	if d.Target == "URL" {
		_, err := url.Parse(d.Target)

		if err != nil {
			return errors.New("invalid Target value: not a valid URL path")
		}
	}

	return nil
}
