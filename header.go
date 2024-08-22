package chess

func (c *Chess) Header(headers map[string]string) map[string]string {
	for key, value := range headers {
		c.header[key] = value
	}

	return c.header
}
