package config

// Configurations exported
type Configurations struct {
	Server ServerConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
	GrpcPort int
}
