package configuration

import (
	"net/http"
	"time"

	"github.com/distribution/distribution/v3/configuration"
)

var defaultWantConfig = configuration.Configuration{
	Version: "0.1",
	Log: struct {
		// AccessLog configures access logging.
		AccessLog struct {
			// Disabled disables access logging.
			Disabled bool `yaml:"disabled,omitempty"`
		} `yaml:"accesslog,omitempty"`

		// Level is the granularity at which registry operations are logged.
		Level configuration.Loglevel `yaml:"level,omitempty"`

		// Formatter overrides the default formatter with another. Options
		// include "text", "json" and "logstash".
		Formatter string `yaml:"formatter,omitempty"`

		// Fields allows users to specify static string fields to include in
		// the logger context.
		Fields map[string]interface{} `yaml:"fields,omitempty"`

		// Hooks allows users to configure the log hooks, to enabling the
		// sequent handling behavior, when defined levels of log message emit.
		Hooks []configuration.LogHook `yaml:"hooks,omitempty"`
	}{
		Level: configuration.Loglevel("debug"),
		Fields: map[string]interface{}{
			"service":     "registry",
			"environment": "development",
		},
		Hooks: []configuration.LogHook{
			{
				Type:     "mail",
				Disabled: true,
				Levels: []string{
					"panic",
				},
				MailOptions: configuration.MailOptions{
					SMTP: struct {
						// Addr defines smtp host address
						Addr string `yaml:"addr,omitempty"`

						// Username defines user name to smtp host
						Username string `yaml:"username,omitempty"`

						// Password defines password of login user
						Password string `yaml:"password,omitempty"`

						// Insecure defines if smtp login skips the secure certification.
						Insecure bool `yaml:"insecure,omitempty"`
					}{
						Addr:     "mail.example.com:25",
						Username: "mailuser",
						Password: "password",
						Insecure: true,
					},
					From: "sender@example.com",
					To: []string{
						"errors@example.com",
					},
				},
			},
		},
	},
	Storage: configuration.Storage{
		"delete": configuration.Parameters{
			"enabled": true,
		},
		"cache": configuration.Parameters{
			"blobdescriptor": "redis",
		},
		"filesystem": configuration.Parameters{
			"rootdirectory": "/var/lib/registry",
		},
		"maintenance": configuration.Parameters{
			"uploadpurging": map[interface{}]interface{}{
				"enabled": false,
			},
		},
	},
	HTTP: struct {
		// Addr specifies the bind address for the registry instance.
		Addr string `yaml:"addr,omitempty"`

		// Net specifies the net portion of the bind address. A default empty value means tcp.
		Net string `yaml:"net,omitempty"`

		// Host specifies an externally-reachable address for the registry, as a fully
		// qualified URL.
		Host string `yaml:"host,omitempty"`

		Prefix string `yaml:"prefix,omitempty"`

		// Secret specifies the secret key which HMAC tokens are created with.
		Secret string `yaml:"secret,omitempty"`

		// RelativeURLs specifies that relative URLs should be returned in
		// Location headers
		RelativeURLs bool `yaml:"relativeurls,omitempty"`

		// Amount of time to wait for connection to drain before shutting down when registry
		// receives a stop signal
		DrainTimeout time.Duration `yaml:"draintimeout,omitempty"`

		// TLS instructs the http server to listen with a TLS configuration.
		// This only support simple tls configuration with a cert and key.
		// Mostly, this is useful for testing situations or simple deployments
		// that require tls. If more complex configurations are required, use
		// a proxy or make a proposal to add support here.
		TLS struct {
			// Certificate specifies the path to an x509 certificate file to
			// be used for TLS.
			Certificate string `yaml:"certificate,omitempty"`

			// Key specifies the path to the x509 key file, which should
			// contain the private portion for the file specified in
			// Certificate.
			Key string `yaml:"key,omitempty"`

			// Specifies the CA certs for client authentication
			// A file may contain multiple CA certificates encoded as PEM
			ClientCAs []string `yaml:"clientcas,omitempty"`

			// Specifies the lowest TLS version allowed
			MinimumTLS string `yaml:"minimumtls,omitempty"`

			// Specifies a list of cipher suites allowed
			CipherSuites []string `yaml:"ciphersuites,omitempty"`

			// LetsEncrypt is used to configuration setting up TLS through
			// Let's Encrypt instead of manually specifying certificate and
			// key. If a TLS certificate is specified, the Let's Encrypt
			// section will not be used.
			LetsEncrypt struct {
				// CacheFile specifies cache file to use for lets encrypt
				// certificates and keys.
				CacheFile string `yaml:"cachefile,omitempty"`

				// Email is the email to use during Let's Encrypt registration
				Email string `yaml:"email,omitempty"`

				// Hosts specifies the hosts which are allowed to obtain Let's
				// Encrypt certificates.
				Hosts []string `yaml:"hosts,omitempty"`
			} `yaml:"letsencrypt,omitempty"`
		} `yaml:"tls,omitempty"`

		// Headers is a set of headers to include in HTTP responses. A common
		// use case for this would be security headers such as
		// Strict-Transport-Security. The map keys are the header names, and
		// the values are the associated header payloads.
		Headers http.Header `yaml:"headers,omitempty"`

		// Debug configures the http debug interface, if specified. This can
		// include services such as pprof, expvar and other data that should
		// not be exposed externally. Left disabled by default.
		Debug struct {
			// Addr specifies the bind address for the debug server.
			Addr string `yaml:"addr,omitempty"`
			// Prometheus configures the Prometheus telemetry endpoint.
			Prometheus struct {
				Enabled bool   `yaml:"enabled,omitempty"`
				Path    string `yaml:"path,omitempty"`
			} `yaml:"prometheus,omitempty"`
		} `yaml:"debug,omitempty"`

		// HTTP2 configuration options
		HTTP2 struct {
			// Specifies whether the registry should disallow clients attempting
			// to connect via http2. If set to true, only http/1.1 is supported.
			Disabled bool `yaml:"disabled,omitempty"`
		} `yaml:"http2,omitempty"`
	}{
		Addr: ":5000",
		Debug: struct {
			// Addr specifies the bind address for the debug server.
			Addr string `yaml:"addr,omitempty"`
			// Prometheus configures the Prometheus telemetry endpoint.
			Prometheus struct {
				Enabled bool   `yaml:"enabled,omitempty"`
				Path    string `yaml:"path,omitempty"`
			} `yaml:"prometheus,omitempty"`
		}{
			Addr: ":5001",
			Prometheus: struct {
				Enabled bool   `yaml:"enabled,omitempty"`
				Path    string `yaml:"path,omitempty"`
			}{
				Enabled: false,
				Path:    "/metrics",
			},
		},
		Headers: http.Header{
			"X-Content-Type-Options": []string{"nosniff"},
		},
	},
	Redis: struct {
		// Addr specifies the the redis instance available to the application.
		Addr string `yaml:"addr,omitempty"`

		// Password string to use when making a connection.
		Password string `yaml:"password,omitempty"`

		// DB specifies the database to connect to on the redis instance.
		DB int `yaml:"db,omitempty"`

		// TLS configures settings for redis in-transit encryption
		TLS struct {
			Enabled bool `yaml:"enabled,omitempty"`
		} `yaml:"tls,omitempty"`

		DialTimeout  time.Duration `yaml:"dialtimeout,omitempty"`  // timeout for connect
		ReadTimeout  time.Duration `yaml:"readtimeout,omitempty"`  // timeout for reads of data
		WriteTimeout time.Duration `yaml:"writetimeout,omitempty"` // timeout for writes of data

		// Pool configures the behavior of the redis connection pool.
		Pool struct {
			// MaxIdle sets the maximum number of idle connections.
			MaxIdle int `yaml:"maxidle,omitempty"`

			// MaxActive sets the maximum number of connections that should be
			// opened before blocking a connection request.
			MaxActive int `yaml:"maxactive,omitempty"`

			// IdleTimeout sets the amount time to wait before closing
			// inactive connections.
			IdleTimeout time.Duration `yaml:"idletimeout,omitempty"`
		} `yaml:"pool,omitempty"`
	}{
		Addr: "localhost:6379",
		Pool: struct {
			// MaxIdle sets the maximum number of idle connections.
			MaxIdle int `yaml:"maxidle,omitempty"`

			// MaxActive sets the maximum number of connections that should be
			// opened before blocking a connection request.
			MaxActive int `yaml:"maxactive,omitempty"`

			// IdleTimeout sets the amount time to wait before closing
			// inactive connections.
			IdleTimeout time.Duration `yaml:"idletimeout,omitempty"`
		}{
			MaxIdle:     16,
			MaxActive:   64,
			IdleTimeout: 300 * time.Second,
		},
		DialTimeout:  10 * time.Millisecond,
		ReadTimeout:  10 * time.Millisecond,
		WriteTimeout: 10 * time.Millisecond,
	},
	Health: configuration.Health{
		StorageDriver: struct {
			// Enabled turns on the health check for the storage driver
			Enabled bool `yaml:"enabled,omitempty"`
			// Interval is the duration in between checks
			Interval time.Duration `yaml:"interval,omitempty"`
			// Threshold is the number of times a check must fail to trigger an
			// unhealthy state
			Threshold int `yaml:"threshold,omitempty"`
		}{
			Enabled:   true,
			Interval:  10 * time.Second,
			Threshold: 3,
		},
	},
}

type wantConfigOption func(*configuration.Configuration)

func GetWantConfig(opts ...wantConfigOption) *configuration.Configuration {
	config := defaultWantConfig
	for _, o := range opts {
		o(&config)
	}
	return &config
}

func WithStorage(storage configuration.Storage) wantConfigOption {
	return func(config *configuration.Configuration) {
		config.Storage = storage
	}
}

func WithHeaders(headers http.Header) wantConfigOption {
	return func(config *configuration.Configuration) {
		config.HTTP.Headers = headers
	}
}

func WithHTTP(http HTTP) wantConfigOption {
	return func(config *configuration.Configuration) {
		config.HTTP.Addr = http.Addr
		config.HTTP.Net = http.Net
		config.HTTP.Host = http.Host
		config.HTTP.Headers = http.Headers
		config.HTTP.Secret = http.Secret
		// TODO: #1 Add more fields from struct
	}
}

type HTTP struct {
	// Addr specifies the bind address for the registry instance.
	Addr string `yaml:"addr,omitempty"`

	// Net specifies the net portion of the bind address. A default empty value means tcp.
	Net string `yaml:"net,omitempty"`

	// Host specifies an externally-reachable address for the registry, as a fully
	// qualified URL.
	Host string `yaml:"host,omitempty"`

	Prefix string `yaml:"prefix,omitempty"`

	// Secret specifies the secret key which HMAC tokens are created with.
	Secret string `yaml:"secret,omitempty"`

	// RelativeURLs specifies that relative URLs should be returned in
	// Location headers
	RelativeURLs bool `yaml:"relativeurls,omitempty"`

	// Amount of time to wait for connection to drain before shutting down when registry
	// receives a stop signal
	DrainTimeout time.Duration `yaml:"draintimeout,omitempty"`

	// TLS instructs the http server to listen with a TLS configuration.
	// This only support simple tls configuration with a cert and key.
	// Mostly, this is useful for testing situations or simple deployments
	// that require tls. If more complex configurations are required, use
	// a proxy or make a proposal to add support here.
	TLS struct {
		// Certificate specifies the path to an x509 certificate file to
		// be used for TLS.
		Certificate string `yaml:"certificate,omitempty"`

		// Key specifies the path to the x509 key file, which should
		// contain the private portion for the file specified in
		// Certificate.
		Key string `yaml:"key,omitempty"`

		// Specifies the CA certs for client authentication
		// A file may contain multiple CA certificates encoded as PEM
		ClientCAs []string `yaml:"clientcas,omitempty"`

		// Specifies the lowest TLS version allowed
		MinimumTLS string `yaml:"minimumtls,omitempty"`

		// Specifies a list of cipher suites allowed
		CipherSuites []string `yaml:"ciphersuites,omitempty"`

		// LetsEncrypt is used to configuration setting up TLS through
		// Let's Encrypt instead of manually specifying certificate and
		// key. If a TLS certificate is specified, the Let's Encrypt
		// section will not be used.
		LetsEncrypt struct {
			// CacheFile specifies cache file to use for lets encrypt
			// certificates and keys.
			CacheFile string `yaml:"cachefile,omitempty"`

			// Email is the email to use during Let's Encrypt registration
			Email string `yaml:"email,omitempty"`

			// Hosts specifies the hosts which are allowed to obtain Let's
			// Encrypt certificates.
			Hosts []string `yaml:"hosts,omitempty"`
		} `yaml:"letsencrypt,omitempty"`
	} `yaml:"tls,omitempty"`

	// Headers is a set of headers to include in HTTP responses. A common
	// use case for this would be security headers such as
	// Strict-Transport-Security. The map keys are the header names, and
	// the values are the associated header payloads.
	Headers http.Header `yaml:"headers,omitempty"`

	// Debug configures the http debug interface, if specified. This can
	// include services such as pprof, expvar and other data that should
	// not be exposed externally. Left disabled by default.
	Debug struct {
		// Addr specifies the bind address for the debug server.
		Addr string `yaml:"addr,omitempty"`
		// Prometheus configures the Prometheus telemetry endpoint.
		Prometheus struct {
			Enabled bool   `yaml:"enabled,omitempty"`
			Path    string `yaml:"path,omitempty"`
		} `yaml:"prometheus,omitempty"`
	} `yaml:"debug,omitempty"`

	// HTTP2 configuration options
	HTTP2 struct {
		// Specifies whether the registry should disallow clients attempting
		// to connect via http2. If set to true, only http/1.1 is supported.
		Disabled bool `yaml:"disabled,omitempty"`
	} `yaml:"http2,omitempty"`
}
