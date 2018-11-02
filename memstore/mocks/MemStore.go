//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import memstore "github.com/uber/aresdb/memstore"
import mock "github.com/stretchr/testify/mock"

// MemStore is an autogenerated mock type for the MemStore type
type MemStore struct {
	mock.Mock
}

// Archive provides a mock function with given fields: table, shardID, cutoff, reporter
func (_m *MemStore) Archive(table string, shardID int, cutoff uint32, reporter memstore.ArchiveJobDetailReporter) error {
	ret := _m.Called(table, shardID, cutoff, reporter)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, uint32, memstore.ArchiveJobDetailReporter) error); ok {
		r0 = rf(table, shardID, cutoff, reporter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backfill provides a mock function with given fields: table, shardID, reporter
func (_m *MemStore) Backfill(table string, shardID int, reporter memstore.BackfillJobDetailReporter) error {
	ret := _m.Called(table, shardID, reporter)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, memstore.BackfillJobDetailReporter) error); ok {
		r0 = rf(table, shardID, reporter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchSchema provides a mock function with given fields:
func (_m *MemStore) FetchSchema() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMemoryUsageDetails provides a mock function with given fields:
func (_m *MemStore) GetMemoryUsageDetails() (map[string]memstore.TableShardMemoryUsage, error) {
	ret := _m.Called()

	var r0 map[string]memstore.TableShardMemoryUsage
	if rf, ok := ret.Get(0).(func() map[string]memstore.TableShardMemoryUsage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]memstore.TableShardMemoryUsage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScheduler provides a mock function with given fields:
func (_m *MemStore) GetScheduler() memstore.Scheduler {
	ret := _m.Called()

	var r0 memstore.Scheduler
	if rf, ok := ret.Get(0).(func() memstore.Scheduler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(memstore.Scheduler)
		}
	}

	return r0
}

// GetSchema provides a mock function with given fields: table
func (_m *MemStore) GetSchema(table string) (*memstore.TableSchema, error) {
	ret := _m.Called(table)

	var r0 *memstore.TableSchema
	if rf, ok := ret.Get(0).(func(string) *memstore.TableSchema); ok {
		r0 = rf(table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*memstore.TableSchema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(table)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemas provides a mock function with given fields:
func (_m *MemStore) GetSchemas() map[string]*memstore.TableSchema {
	ret := _m.Called()

	var r0 map[string]*memstore.TableSchema
	if rf, ok := ret.Get(0).(func() map[string]*memstore.TableSchema); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*memstore.TableSchema)
		}
	}

	return r0
}

// GetTableShard provides a mock function with given fields: table, shardID
func (_m *MemStore) GetTableShard(table string, shardID int) (*memstore.TableShard, error) {
	ret := _m.Called(table, shardID)

	var r0 *memstore.TableShard
	if rf, ok := ret.Get(0).(func(string, int) *memstore.TableShard); ok {
		r0 = rf(table, shardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*memstore.TableShard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(table, shardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleIngestion provides a mock function with given fields: table, shardID, upsertBatch
func (_m *MemStore) HandleIngestion(table string, shardID int, upsertBatch *memstore.UpsertBatch) error {
	ret := _m.Called(table, shardID, upsertBatch)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, *memstore.UpsertBatch) error); ok {
		r0 = rf(table, shardID, upsertBatch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitShards provides a mock function with given fields: schedulerOff
func (_m *MemStore) InitShards(schedulerOff bool) {
	_m.Called(schedulerOff)
}

// Lock provides a mock function with given fields:
func (_m *MemStore) Lock() {
	_m.Called()
}

// Purge provides a mock function with given fields: table, shardID, batchIDStart, batchIDEnd, reporter
func (_m *MemStore) Purge(table string, shardID int, batchIDStart int, batchIDEnd int, reporter memstore.PurgeJobDetailReporter) error {
	ret := _m.Called(table, shardID, batchIDStart, batchIDEnd, reporter)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, int, int, memstore.PurgeJobDetailReporter) error); ok {
		r0 = rf(table, shardID, batchIDStart, batchIDEnd, reporter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RLock provides a mock function with given fields:
func (_m *MemStore) RLock() {
	_m.Called()
}

// RUnlock provides a mock function with given fields:
func (_m *MemStore) RUnlock() {
	_m.Called()
}

// Snapshot provides a mock function with given fields: table, shardID, reporter
func (_m *MemStore) Snapshot(table string, shardID int, reporter memstore.SnapshotJobDetailReporter) error {
	ret := _m.Called(table, shardID, reporter)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, memstore.SnapshotJobDetailReporter) error); ok {
		r0 = rf(table, shardID, reporter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unlock provides a mock function with given fields:
func (_m *MemStore) Unlock() {
	_m.Called()
}