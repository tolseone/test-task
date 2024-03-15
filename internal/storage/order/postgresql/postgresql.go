package db

import (
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/pkg/client/postgresql"
)

type Storage struct {
	client postgresql.Client
	log    *logger.Logger
}
