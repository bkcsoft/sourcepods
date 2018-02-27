package repository

import (
	"context"
	"time"

	"github.com/gitpods/gitpods/storage"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type loggingService struct {
	logger  log.Logger
	service Service
}

// NewLoggingService wraps the Service and provides logging for its methods.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger: logger, service: s}
}

func (s *loggingService) List(ctx context.Context, owner string) ([]*Repository, string, error) {
	start := time.Now()

	repositories, owner, err := s.service.List(ctx, owner)

	logger := log.With(s.logger,
		"method", "List",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log(
			"msg", "failed to list repositories by owner's username",
			"err", err,
		)
	} else {
		level.Debug(logger).Log()
	}

	return repositories, owner, err

}

func (s *loggingService) Find(ctx context.Context, owner string, name string) (*Repository, string, error) {
	start := time.Now()

	repository, owner, err := s.service.Find(ctx, owner, name)

	logger := log.With(s.logger,
		"method", "Find",
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log(
			"msg", "failed to find repository by owner & name",
			"err", err,
		)
	} else {
		level.Debug(logger).Log()
	}

	return repository, owner, err
}

func (s *loggingService) Create(ctx context.Context, owner string, repository *Repository) (*Repository, error) {
	start := time.Now()

	repository, err := s.service.Create(ctx, owner, repository)

	logger := log.With(s.logger,
		"method", "Create",
		"duration", time.Since(start),
	)

	if err != nil {
		if err != AlreadyExistsError {
			level.Warn(logger).Log(
				"msg", "failed to create repository",
				"err", err,
			)
		}
	} else {
		level.Debug(logger).Log(
			"owner", owner,
			"name", repository.Name,
		)
	}

	return repository, err
}
func (s *loggingService) Branches(ctx context.Context, owner string, name string) ([]*Branch, error) {
	start := time.Now()

	branches, err := s.service.Branches(ctx, owner, name)

	logger := log.With(s.logger,
		"method", "Branches",
		"duration", time.Since(start),
	)

	if err != nil {
		if err != AlreadyExistsError {
			level.Warn(logger).Log(
				"msg", "failed to list branches",
				"err", err,
			)
		}
	} else {
		level.Debug(logger).Log(
			"owner", owner,
			"name", name,
		)
	}

	return branches, err
}

func (s *loggingService) Commit(ctx context.Context, owner string, name string, rev string) (storage.Commit, error) {
	start := time.Now()

	commit, err := s.service.Commit(ctx, owner, name, rev)

	logger := log.With(s.logger,
		"method", "Commit",
		"owner", owner,
		"name", name,
		"rev", rev,
		"duration", time.Since(start),
	)

	if err != nil {
		level.Warn(logger).Log(
			"msg", "failed to get the commit for repository",
			"err", err,
		)
	} else {
		level.Debug(logger).Log()
	}

	return commit, err
}
