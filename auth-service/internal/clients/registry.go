package clients

import "auth-service/config"

type Registry struct {
	User *UserClient
}

func NewRegistry(cfg *config.Config) (*Registry, error) {
	userClient, err := NewUserClient(cfg.UserServiceAddr)
	if err != nil {
		return nil, err
	}
	return &Registry{
		User: userClient,
	}, nil
}

func (r *Registry) Close() error {
	if r.User.conn != nil {
		if err := r.User.conn.Close(); err != nil {
			return err
		}
	}
	return nil
}
