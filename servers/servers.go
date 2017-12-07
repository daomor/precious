package servers

type ServersData struct {
	Servers 	[]*ServerEntity		`yaml:"servers"`
}

func (s *ServersData) GetServerByName(name string) *ServerEntity {

	// Loop the servers.
	for _, v := range s.Servers {
		if(v.Name == name) {
			return v
		}
	}
	return nil
}