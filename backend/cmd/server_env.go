package cmd

func (server *ApiServer) loadEnv() []string {
	var errors []string

	// server.Dsn = os.Getenv("API_SERVER_DSN")
	// if server.Dsn == "" {
	// 	errors = append(errors, "dsn is required")
	// }

	// server.Username = os.Getenv("API_SERVER_USERNAME")
	// if server.Username == "" {
	// 	errors = append(errors, "username is required")
	// }

	// server.Password = os.Getenv("API_SERVER_PASSWORD")
	// if server.Password == "" {
	// 	errors = append(errors, "password is required")
	// }

	server.Dsn = "host=localhost port=5432 dbname=asau sslmode=disable"
	server.Username = "postgres"
	server.Password = "123456"
	server.SecretKey = "test"

	return errors
}
