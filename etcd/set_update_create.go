package etcd

// Set sets the given key to the given value.
// It will create a new key value pair or replace the old one.
// It will not replace a existing directory.
func (c *Client) Set(key string, value string) (*Response, error) {
	raw, err := c.RawSet(key, value)

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}

// SetDir sets the given key to a directory.
// It will create a new directory or replace the old key value pair by a directory.
// It will not replace a existing directory.
func (c *Client) SetDir(key string) (*Response, error) {
	raw, err := c.RawSetDir(key)

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}

// CreateDir creates a directory. It succeeds only if
// the given key does not yet exist.
func (c *Client) CreateDir(key string) (*Response, error) {
	raw, err := c.RawCreateDir(key)

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}

// UpdateDir updates the given directory. It succeeds only if the
// given key already exists.
func (c *Client) UpdateDir(key string) (*Response, error) {
	raw, err := c.RawUpdateDir(key)

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}

// Create creates a file with the given value under the given key.  It succeeds
// only if the given key does not yet exist.
func (c *Client) Create(key string, value string) (*Response, error) {
	raw, err := c.RawCreate(key, value)

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}

// CreateInOrder creates a file with a key that's guaranteed to be higher than other
// keys in the given directory. It is useful for creating queues.
func (c *Client) CreateInOrder(dir string, value string) (*Response, error) {
	raw, err := c.RawCreateInOrder(dir, value)

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}

// Update updates the given key to the given value.  It succeeds only if the
// given key already exists.
func (c *Client) Update(key string, value string) (*Response, error) {
	raw, err := c.RawUpdate(key, value)

	if err != nil {
		return nil, err
	}

	return raw.Unmarshal()
}

func (c *Client) RawUpdateDir(key string) (*RawResponse, error) {
	ops := Options{
		"prevExist": true,
		"dir":       true,
	}

	return c.put(key, "", ops)
}

func (c *Client) RawCreateDir(key string) (*RawResponse, error) {
	ops := Options{
		"prevExist": false,
		"dir":       true,
	}

	return c.put(key, "", ops)
}

func (c *Client) RawSet(key string, value string) (*RawResponse, error) {
	return c.put(key, value, nil)
}

func (c *Client) RawSetDir(key string) (*RawResponse, error) {
	ops := Options{
		"dir": true,
	}

	return c.put(key, "", ops)
}

func (c *Client) RawUpdate(key string, value string) (*RawResponse, error) {
	ops := Options{
		"prevExist": true,
	}

	return c.put(key, value, ops)
}

func (c *Client) RawCreate(key string, value string) (*RawResponse, error) {
	ops := Options{
		"prevExist": false,
	}

	return c.put(key, value, ops)
}

func (c *Client) RawCreateInOrder(dir string, value string) (*RawResponse, error) {
	return c.post(dir, value)
}
