package configuration

import (
    "errors"
    "net/url"
)

type Target struct {
	Kind   Kind
	Target string
}

type Kind string

const (
	URL         Kind = "URL"
	kubeService Kind = "kubeService"
)

func (t *Target) Validate() error {
    if err := t.validateKind(); err != nil {
        return err
    }

    if err := t.validateTarget(); err != nil {
        return err
    }

    return nil
}

func (t *Target) validateKind() error {
	validTargets := []string{"URL", "kubeService"}

	for _, target := range validTargets {

		if string(t.Target) == target {
			return nil
		}

	}

	return errors.New("invalid Method value")
}

func (t *Target) validateTarget() error {
    if t.Kind == "URL" {
        _, err := url.Parse(t.Target)

        if err != nil {
            return errors.New("invalid Target value: not a valid URL path")
        }
    }

    return nil
}
