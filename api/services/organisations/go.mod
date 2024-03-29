module invitr.io.com/services/organisations

go 1.21.0

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gorilla/mux v1.8.1
	go.uber.org/zap v1.27.0
	invitr.io.com/cors v0.0.0
	invitr.io.com/responder v0.0.0

)

require filippo.io/edwards25519 v1.1.0 // indirect

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/google/uuid v1.6.0
	github.com/gorilla/handlers v1.5.2 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)

replace invitr.io.com/responder v0.0.0 => ../../pkg/responder

replace invitr.io.com/cors v0.0.0 => ../../pkg/cors
