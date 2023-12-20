package config

type Server struct {
	Name string
	Addr string
}

var Servers = map[string]Server{
	"iot_server": {Name: "iot_server", Addr: ":8901"},
}

func GetServerInfo(serverName string) *Server {
	server, ok := Servers[serverName]
	if !ok {
		return nil
	}
	return &server
}
