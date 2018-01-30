package conf

// Conf set of ENV configs for app
type Conf struct {
	Addr string `env:"ADDR" envDefault:":22000"`
}
