package client

type options struct {
	deviceFlow   bool
	clientSecret string
	username     string
	password     string
	token        string
}

func newOptions(opts ...Option) (*options, error) {
	o := &options{}
	for _, opt := range opts {
		if err := opt(o); err != nil {
			return nil, err
		}
	}
	return o, nil
}

type Option func(o *options) error

func WithPasswordGrant(
	username string,
	password string,
) Option {
	return func(o *options) error {
		o.deviceFlow = false
		o.username = username
		o.password = password
		return nil
	}
}

func WithDeviceFlow() Option {
	return func(o *options) error {
		o.deviceFlow = true
		return nil
	}
}

func WithToken(
	token string,
) Option {
	return func(o *options) error {
		o.deviceFlow = false
		o.token = token
		return nil
	}
}
