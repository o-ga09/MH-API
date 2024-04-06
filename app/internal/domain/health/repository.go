package health

type HealthRepository interface {
	GetStatus() error
}
