module invitr.io.com/services/invites

go 1.21.0

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gorilla/mux v1.8.1
	go.uber.org/zap v1.27.0
	gotest.tools v2.2.0+incompatible
	invitr.io.com/cors v0.0.0
	invitr.io.com/responder v0.0.0
	invitr.io.com/services/qr-codes v0.0.0
	invitr.io.com/services/users v0.0.0
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)

require (
	github.com/google/uuid v1.6.0
	go.uber.org/multierr v1.11.0 // indirect
)

replace invitr.io.com/services/users v0.0.0 => ../users

replace invitr.io.com/services/qr-codes v0.0.0 => ../qr-codes

replace invitr.io.com/responder v0.0.0 => ../../pkg/responder

replace invitr.io.com/cors v0.0.0 => ../../pkg/cors
