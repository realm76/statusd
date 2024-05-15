package entities

type NotificationType int

const (
	Webhook NotificationType = iota + 1
)

func (nt NotificationType) String() string {
	return [...]string{"Webhook"}[nt-1]
}
