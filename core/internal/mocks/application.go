// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bridges "github.com/smartcontractkit/chainlink/core/bridges"
	bulletprooftxmanager "github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"

	config "github.com/smartcontractkit/chainlink/core/config"

	context "context"

	evm "github.com/smartcontractkit/chainlink/core/chains/evm"

	feeds "github.com/smartcontractkit/chainlink/core/services/feeds"

	gorm "gorm.io/gorm"

	health "github.com/smartcontractkit/chainlink/core/services/health"

	job "github.com/smartcontractkit/chainlink/core/services/job"

	keystore "github.com/smartcontractkit/chainlink/core/services/keystore"

	logger "github.com/smartcontractkit/chainlink/core/logger"

	mock "github.com/stretchr/testify/mock"

	null "gopkg.in/guregu/null.v4"

	packr "github.com/gobuffalo/packr"

	pipeline "github.com/smartcontractkit/chainlink/core/services/pipeline"

	postgres "github.com/smartcontractkit/chainlink/core/services/postgres"

	sessions "github.com/smartcontractkit/chainlink/core/sessions"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/types"

	uuid "github.com/satori/go.uuid"

	webhook "github.com/smartcontractkit/chainlink/core/services/webhook"

	zapcore "go.uber.org/zap/zapcore"
)

// Application is an autogenerated mock type for the Application type
type Application struct {
	mock.Mock
}

// AddJobV2 provides a mock function with given fields: ctx, _a1, name
func (_m *Application) AddJobV2(ctx context.Context, _a1 job.Job, name null.String) (job.Job, error) {
	ret := _m.Called(ctx, _a1, name)

	var r0 job.Job
	if rf, ok := ret.Get(0).(func(context.Context, job.Job, null.String) job.Job); ok {
		r0 = rf(ctx, _a1, name)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, job.Job, null.String) error); ok {
		r1 = rf(ctx, _a1, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BPTXMORM provides a mock function with given fields:
func (_m *Application) BPTXMORM() bulletprooftxmanager.ORM {
	ret := _m.Called()

	var r0 bulletprooftxmanager.ORM
	if rf, ok := ret.Get(0).(func() bulletprooftxmanager.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bulletprooftxmanager.ORM)
		}
	}

	return r0
}

// BridgeORM provides a mock function with given fields:
func (_m *Application) BridgeORM() bridges.ORM {
	ret := _m.Called()

	var r0 bridges.ORM
	if rf, ok := ret.Get(0).(func() bridges.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bridges.ORM)
		}
	}

	return r0
}

// DeleteJob provides a mock function with given fields: ctx, jobID
func (_m *Application) DeleteJob(ctx context.Context, jobID int32) error {
	ret := _m.Called(ctx, jobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EVMORM provides a mock function with given fields:
func (_m *Application) EVMORM() types.ORM {
	ret := _m.Called()

	var r0 types.ORM
	if rf, ok := ret.Get(0).(func() types.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ORM)
		}
	}

	return r0
}

// GetChainSet provides a mock function with given fields:
func (_m *Application) GetChainSet() evm.ChainSet {
	ret := _m.Called()

	var r0 evm.ChainSet
	if rf, ok := ret.Get(0).(func() evm.ChainSet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(evm.ChainSet)
		}
	}

	return r0
}

// GetConfig provides a mock function with given fields:
func (_m *Application) GetConfig() config.GeneralConfig {
	ret := _m.Called()

	var r0 config.GeneralConfig
	if rf, ok := ret.Get(0).(func() config.GeneralConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.GeneralConfig)
		}
	}

	return r0
}

// GetDB provides a mock function with given fields:
func (_m *Application) GetDB() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// GetEventBroadcaster provides a mock function with given fields:
func (_m *Application) GetEventBroadcaster() postgres.EventBroadcaster {
	ret := _m.Called()

	var r0 postgres.EventBroadcaster
	if rf, ok := ret.Get(0).(func() postgres.EventBroadcaster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(postgres.EventBroadcaster)
		}
	}

	return r0
}

// GetExternalInitiatorManager provides a mock function with given fields:
func (_m *Application) GetExternalInitiatorManager() webhook.ExternalInitiatorManager {
	ret := _m.Called()

	var r0 webhook.ExternalInitiatorManager
	if rf, ok := ret.Get(0).(func() webhook.ExternalInitiatorManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(webhook.ExternalInitiatorManager)
		}
	}

	return r0
}

// GetFeedsService provides a mock function with given fields:
func (_m *Application) GetFeedsService() feeds.Service {
	ret := _m.Called()

	var r0 feeds.Service
	if rf, ok := ret.Get(0).(func() feeds.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(feeds.Service)
		}
	}

	return r0
}

// GetHealthChecker provides a mock function with given fields:
func (_m *Application) GetHealthChecker() health.Checker {
	ret := _m.Called()

	var r0 health.Checker
	if rf, ok := ret.Get(0).(func() health.Checker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(health.Checker)
		}
	}

	return r0
}

// GetKeyStore provides a mock function with given fields:
func (_m *Application) GetKeyStore() keystore.Master {
	ret := _m.Called()

	var r0 keystore.Master
	if rf, ok := ret.Get(0).(func() keystore.Master); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Master)
		}
	}

	return r0
}

// GetLogger provides a mock function with given fields:
func (_m *Application) GetLogger() logger.Logger {
	ret := _m.Called()

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func() logger.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// GetWebAuthnConfiguration provides a mock function with given fields:
func (_m *Application) GetWebAuthnConfiguration() sessions.WebAuthnConfiguration {
	ret := _m.Called()

	var r0 sessions.WebAuthnConfiguration
	if rf, ok := ret.Get(0).(func() sessions.WebAuthnConfiguration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sessions.WebAuthnConfiguration)
	}

	return r0
}

// JobORM provides a mock function with given fields:
func (_m *Application) JobORM() job.ORM {
	ret := _m.Called()

	var r0 job.ORM
	if rf, ok := ret.Get(0).(func() job.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(job.ORM)
		}
	}

	return r0
}

// JobSpawner provides a mock function with given fields:
func (_m *Application) JobSpawner() job.Spawner {
	ret := _m.Called()

	var r0 job.Spawner
	if rf, ok := ret.Get(0).(func() job.Spawner); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(job.Spawner)
		}
	}

	return r0
}

// NewBox provides a mock function with given fields:
func (_m *Application) NewBox() packr.Box {
	ret := _m.Called()

	var r0 packr.Box
	if rf, ok := ret.Get(0).(func() packr.Box); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(packr.Box)
	}

	return r0
}

// PipelineORM provides a mock function with given fields:
func (_m *Application) PipelineORM() pipeline.ORM {
	ret := _m.Called()

	var r0 pipeline.ORM
	if rf, ok := ret.Get(0).(func() pipeline.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pipeline.ORM)
		}
	}

	return r0
}

// ReplayFromBlock provides a mock function with given fields: chainID, number
func (_m *Application) ReplayFromBlock(chainID *big.Int, number uint64) error {
	ret := _m.Called(chainID, number)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, uint64) error); ok {
		r0 = rf(chainID, number)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResumeJobV2 provides a mock function with given fields: ctx, taskID, result
func (_m *Application) ResumeJobV2(ctx context.Context, taskID uuid.UUID, result pipeline.Result) error {
	ret := _m.Called(ctx, taskID, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, pipeline.Result) error); ok {
		r0 = rf(ctx, taskID, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunJobV2 provides a mock function with given fields: ctx, jobID, meta
func (_m *Application) RunJobV2(ctx context.Context, jobID int32, meta map[string]interface{}) (int64, error) {
	ret := _m.Called(ctx, jobID, meta)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int32, map[string]interface{}) int64); ok {
		r0 = rf(ctx, jobID, meta)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32, map[string]interface{}) error); ok {
		r1 = rf(ctx, jobID, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunWebhookJobV2 provides a mock function with given fields: ctx, jobUUID, requestBody, meta
func (_m *Application) RunWebhookJobV2(ctx context.Context, jobUUID uuid.UUID, requestBody string, meta pipeline.JSONSerializable) (int64, error) {
	ret := _m.Called(ctx, jobUUID, requestBody, meta)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string, pipeline.JSONSerializable) int64); ok {
		r0 = rf(ctx, jobUUID, requestBody, meta)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string, pipeline.JSONSerializable) error); ok {
		r1 = rf(ctx, jobUUID, requestBody, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionORM provides a mock function with given fields:
func (_m *Application) SessionORM() sessions.ORM {
	ret := _m.Called()

	var r0 sessions.ORM
	if rf, ok := ret.Get(0).(func() sessions.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sessions.ORM)
		}
	}

	return r0
}

// SetLogLevel provides a mock function with given fields: ctx, lvl
func (_m *Application) SetLogLevel(ctx context.Context, lvl zapcore.Level) error {
	ret := _m.Called(ctx, lvl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, zapcore.Level) error); ok {
		r0 = rf(ctx, lvl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetServiceLogLevel provides a mock function with given fields: ctx, service, level
func (_m *Application) SetServiceLogLevel(ctx context.Context, service string, level zapcore.Level) error {
	ret := _m.Called(ctx, service, level)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, zapcore.Level) error); ok {
		r0 = rf(ctx, service, level)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Application) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Application) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WakeSessionReaper provides a mock function with given fields:
func (_m *Application) WakeSessionReaper() {
	_m.Called()
}
