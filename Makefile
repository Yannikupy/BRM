generate_auth_swagger:
	cd back/auth && swag fmt && swag init -g cmd/server/main.go

generate_auth_proto_server:
	cd back/auth && go generate auth/internal/ports/grpcserver

generate_ads_proto_clients:
	cd back/brm-ads && go generate brm-ads/internal/adapters/grpccore
	cd back/brm-ads && go generate brm-ads/internal/adapters/grpcleads

generate_ads_proto_server:
	cd back/brm-ads && go generate brm-ads/internal/ports/grpcserver

generate_core_proto_server:
	cd back/brm-core && go generate brm-core/internal/ports/grpcserver

generate_core_proto_clients:
	cd back/brm-core && go generate brm-core/internal/adapters/grpcauth

generate_leads_proto_server:
	cd back/brm-leads && go generate brm-leads/internal/ports/grpcserver

generate_leads_proto_clients:
	cd back/brm-leads && go generate brm-leads/internal/adapters/grpcads
	cd back/brm-leads && go generate brm-leads/internal/adapters/grpccore

go_generate_registration_swagger:
	cd back/registration && swag fmt && swag init -g cmd/server/main.go

generate_registration_proto_clients:
	cd back/registration && go generate registration/internal/adapters/grpccore

generate_stats_proto_server:
	cd back/stats && go generate stats/internal/ports/grpcserver

generate_transport_swagger:
	cd back/transport-api && swag fmt && swag init -g cmd/server/main.go

generate_transport_proto_clients:
	cd back/transport-api && go generate transport-api/internal/adapters/grpcads
	cd back/transport-api && go generate transport-api/internal/adapters/grpccore
	cd back/transport-api && go generate transport-api/internal/adapters/grpcleads
	cd back/transport-api && go generate transport-api/internal/adapters/grpcstats

run_dev_front:
	docker-compose -p brm --profile dev_front up -d

run_dev_back:
	docker-compose -p brm --profile dev_back up -d
