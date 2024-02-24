module invitio.com/invites

go 1.21.0

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gorilla/mux v1.8.1
	go.uber.org/zap v1.26.0
	invitio.com/qr-codes v0.0.0
	invitio.com/users v0.0.0
)

require (
	github.com/google/uuid v1.6.0
	go.uber.org/multierr v1.11.0 // indirect
)

replace invitio.com/users v0.0.0 => ../users

replace invitio.com/qr-codes v0.0.0 => ../qr-codes
