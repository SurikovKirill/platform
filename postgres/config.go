package postgres

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
	SSL      string
}
