package db

type Connection struct {
	Host     string
	Port     int
	UserName string
	Password string
	Database string
}

type ConnectionOption func(*Connection)

func NewConnection(options ...ConnectionOption) *Connection {
	conn := &Connection{
		Host:     "127.0.0.1",
		Port:     3306,
		UserName: "root",
	}

	for _, option := range options {
		option(conn)
	}
	return conn
}

func WithHost(h string) ConnectionOption {
	return func(c *Connection) {
		c.Host = h
	}
}

func WithPort(p int) ConnectionOption {
	return func(c *Connection) {
		c.Port = p
	}
}

func WithUsername(n string) ConnectionOption {
	return func(c *Connection) {
		c.UserName = n
	}
}

func WithPassword(p string) ConnectionOption {
	return func(c *Connection) {
		c.Password = p
	}
}

func WithDatabase(d string) ConnectionOption {
	return func(c *Connection) {
		c.Database = d
	}
}
