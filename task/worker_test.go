package task

import (
	"context"
	"testing"
	"time"

	"github.com/dolefir/refresh-hash/config"
	"github.com/dolefir/refresh-hash/logger"
)

type RefresherMock func(ctx context.Context) error

func (r RefresherMock) Refresh(ctx context.Context) error {
	return r(ctx)
}

func Test_refreshTicker_StartAndCancel(t *testing.T) {
	const (
		tickerDuration     = time.Millisecond * 5
		expectedCallsCount = 2
	)

	counter := 0
	refresherMock := RefresherMock(func(ctx context.Context) error {
		counter++
		return nil
	})

	cfg := config.NewConfig("")
	log := logger.NewLogger((*logger.CFGLogger)(&cfg.Logger), nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ticker := NewRefreshTicker(tickerDuration, time.Second, refresherMock, log)

	go func() {
		const timeBuffer = time.Millisecond
		time.Sleep(tickerDuration*2 + timeBuffer)
		cancel()
	}()

	if err := ticker.Start(ctx); err != nil {
		t.Fatal(err)
	}

	if counter != expectedCallsCount {
		t.Errorf("wrong number of calls (%d), expected - %d", counter, expectedCallsCount)
	}
}
