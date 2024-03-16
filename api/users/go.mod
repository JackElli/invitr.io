module invitr.io.com/users

go 1.21.0

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	go.uber.org/zap v1.26.0
	gotest.tools v2.2.0+incompatible
	invitr.io.com/responder v0.0.0
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)

replace invitr.io.com/responder v0.0.0 => ../responder
