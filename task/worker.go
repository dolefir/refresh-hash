package task

import (
	"context"
	"time"

	"github.com/dolefir/refresh-hash/logger"
)

type Refresher interface {
	Refresh(ctx context.Context) error
}

type refreshTicker struct {
	timer        time.Duration
	queryTimeout time.Duration
	refresher    Refresher
	log          logger.Logger
}

// RefreshTicker is the service interface that
// describes business logic for working with ticker.
type RefreshTicker interface {
	Start(ctx context.Context) error
}

// NewRefreshTicker returns a new Ticker for refresh hash.
func NewRefreshTicker(
	timer time.Duration,
	queryTimeout time.Duration,
	refresher Refresher,
	log logger.Logger,
) RefreshTicker {
	return &refreshTicker{
		timer:        timer,
		queryTimeout: queryTimeout,
		refresher:    refresher,
		log:          log,
	}
}

func (r *refreshTicker) refresh(ctx context.Context) error {
	timeOut, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()
	err := r.refresher.Refresh(timeOut)
	if err != nil {
		r.log.Errorf("task.Refresh: %s", err)
		return err
	}

	return nil
}

// Start timer to rework the hash.
func (r *refreshTicker) Start(ctx context.Context) error {
	ticker := time.NewTicker(r.timer)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			r.log.Debugf("start ticker")
			if err := r.refresh(ctx); err != nil {
				r.log.Errorf("error refresh %v", err)
				return err
			}
			r.log.Debugf("done")

		case <-ctx.Done():
			return nil
		}
	}
}
