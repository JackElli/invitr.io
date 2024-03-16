module invitr.io.com/invites

go 1.21.0

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gorilla/mux v1.8.1
	go.uber.org/zap v1.27.0
	invitr.io.com/qr-codes v0.0.0
	invitr.io.com/users v0.0.0
	invitr.io.com/responder v0.0.0
)

require (
	github.com/google/uuid v1.6.0
	go.uber.org/multierr v1.11.0 // indirect
)

replace invitr.io.com/users v0.0.0 => ../users

replace invitr.io.com/qr-codes v0.0.0 => ../qr-codes

replace invitr.io.com/responder v0.0.0 => ../responder
