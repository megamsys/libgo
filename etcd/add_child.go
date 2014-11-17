package etcd

// Add a new directory with a random etcd-generated key under the given path.
func (c *Client) AddChildDir(key string) (*Response, error) {
	raw, err := c.post(key, "")

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}

// Add a new file with a random etcd-generated key under the given path.
func (c *Client) AddChild(key string, value string) (*Response, error) {
	raw, err := c.post(key, value)

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}
