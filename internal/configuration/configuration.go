package configuration

type Config struct {
	Routers []Router
}

func (c *Config) Validate() error {
	for _, router := range c.Routers {
		if err := router.Validate(); err != nil {
			return err
		}
	}
	return nil
}
