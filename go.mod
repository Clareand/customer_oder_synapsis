module github.com/Clareand/customer_oder_synapsis

go 1.13

require (
	github.com/globocom/echo-prometheus v0.1.2
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo/v4 v4.6.1
	github.com/prometheus/client_golang v1.10.0
	github.com/rs/zerolog v1.27.0
	github.com/sirupsen/logrus v1.6.0
	go.elastic.co/apm/module/apmechov4/v2 v2.1.0
	go.elastic.co/apm/module/apmgormv2/v2 v2.2.0
	go.elastic.co/apm/module/apmlogrus/v2 v2.2.0
	go.elastic.co/apm/v2 v2.2.0
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	gorm.io/driver/postgres v1.4.5 // indirect
	gorm.io/gorm v1.24.2
)
