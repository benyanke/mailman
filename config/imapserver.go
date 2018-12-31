package config

type ImapServer struct {
	Port int
	Host string
	User string
	Pass string

	// Seconds
	Timeout int

}
