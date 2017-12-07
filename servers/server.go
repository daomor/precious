package servers

import "database/sql"

// Server details - Contains list of databases.
type ServerEntity struct {
	Name 		string				`yaml:"name"`
	Host 		string				`yaml:"host"`
	User 		string				`yaml:"user"`
	Password 	string				`yaml:"pass"`
	Port 		string				`yaml:"port"`
	Databases 	[]string			`yaml:"databases"`
}

//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
//username:password@tcp(127.0.0.1:3307)/dbname?param=value
func (s *ServerEntity) GetConnectionString() string {

	// Start the connection string.
	connection := s.User

	// Check if we need a password.
	if(s.Password != "") {
		connection = connection + ":" + s.Password
	}

	// Figure out the port - set default to 3306 in case one isn't set.
	port := "3306"
	if(s.Port != "") {
		port = s.Port
	}

	// Add the rest!
	connection = connection + "@tcp(" + s.Host + ":"+ port + ")/"

	// Return the connection string.
	return connection
}

func (s *ServerEntity) CheckConnection() (bool, error) {

	// Get the servers connection string.
	dbString := s.GetConnectionString()

	// Create mysql connection (not open).
	db, err := sql.Open("mysql", dbString)

	// Log error.
	if err != nil {
		return false, err
	}

	// Open mysql connection and test the server details.
	err = db.Ping()

	// Log error.
	if err != nil {
		return false, err
	}

	// If we get here, we connected successfully.
	return true, nil
}