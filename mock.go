package wireguardgrpc

//go:generate mockgen -source=pkg/exec/executable.go -destination=test/mock/executor.go -package mock
//go:generate mockgen -source=api/wireguard/v1/wireguard_service_grpc.pb.go -destination=test/mock/wireguardv1.go -package mock
