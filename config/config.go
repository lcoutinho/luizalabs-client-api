package config

import "time"

const (
	DB_NAME                 = "customer_db"
	DB_COLLECTION_PRODUCTS  = "products"
	DB_COLLECTION_CUSTOMERS = "customers"
	DB_COLLECTION_USERS     = "users"

	//JWT
	JWT_KEY         = "HIzQMrY0UVmPAiYRd0ZHaWNhOt7m8qe7h4xuz6bRVwEjnbSiRxVE5Fbeed4KDopyx6OVMKMa8rNyxlM5"
	JWT_TIME_EXPIRE = time.Minute * 5

	// TLS_CERTIFICATE_JWT or SIMPLE_JWT
	AUTH_STRATEGY = "TLS_CERTIFICATE_JWT"
)
