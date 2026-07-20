package asynq

import (
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type PeriodicTaskManager struct {
	s            *Scheduler
	p            PeriodicTaskConfigProvider
	syncInterval time.Duration
	done         chan (struct{})
	wg           sync.WaitGroup
	m            map[string]string
}

type PeriodicTaskManagerOpts struct {
	PeriodicTaskConfigProvider PeriodicTaskConfigProvider

	RedisConnOpt RedisConnOpt

	RedisUniversalClient redis.UniversalClient

	*SchedulerOpts

	SyncInterval time.Duration
}

const defaultSyncInterval = 3 * time.Minute

func NewPeriodicTaskManager(opts PeriodicTaskManagerOpts) (*PeriodicTaskManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PeriodicTaskConfigProvider interface {
	GetConfigs() ([]*PeriodicTaskConfig, error)
}

type PeriodicTaskConfig struct {
	Cronspec string
	Task     *Task
	Opts     []Option
}

func (c *PeriodicTaskConfig) hash() string { _ = "STUB: not implemented"; return "" }

func validatePeriodicTaskConfig(c *PeriodicTaskConfig) error { _ = "STUB: not implemented"; return nil }

func (mgr *PeriodicTaskManager) Start() error { _ = "STUB: not implemented"; return nil }

func (mgr *PeriodicTaskManager) Shutdown() { _ = "STUB: not implemented"; return }

func (mgr *PeriodicTaskManager) Run() error { _ = "STUB: not implemented"; return nil }

func (mgr *PeriodicTaskManager) initialSync() error { _ = "STUB: not implemented"; return nil }

func (mgr *PeriodicTaskManager) add(configs []*PeriodicTaskConfig) {
	_ = "STUB: not implemented"
	return
}

func (mgr *PeriodicTaskManager) remove(removed map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (mgr *PeriodicTaskManager) sync() { _ = "STUB: not implemented"; return }

func (mgr *PeriodicTaskManager) diffRemoved(configs []*PeriodicTaskConfig) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *PeriodicTaskManager) diffAdded(configs []*PeriodicTaskConfig) []*PeriodicTaskConfig {
	_ = "STUB: not implemented"
	return nil
}
