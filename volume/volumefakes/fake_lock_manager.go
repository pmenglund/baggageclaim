// This file was generated by counterfeiter
package volumefakes

import (
	"sync"

	"github.com/concourse/baggageclaim/volume"
)

type FakeLockManager struct {
	LockStub        func(key string)
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		key string
	}
	UnlockStub        func(key string)
	unlockMutex       sync.RWMutex
	unlockArgsForCall []struct {
		key string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockManager) Lock(key string) {
	fake.lockMutex.Lock()
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		key string
	}{key})
	fake.recordInvocation("Lock", []interface{}{key})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		fake.LockStub(key)
	}
}

func (fake *FakeLockManager) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeLockManager) LockArgsForCall(i int) string {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return fake.lockArgsForCall[i].key
}

func (fake *FakeLockManager) Unlock(key string) {
	fake.unlockMutex.Lock()
	fake.unlockArgsForCall = append(fake.unlockArgsForCall, struct {
		key string
	}{key})
	fake.recordInvocation("Unlock", []interface{}{key})
	fake.unlockMutex.Unlock()
	if fake.UnlockStub != nil {
		fake.UnlockStub(key)
	}
}

func (fake *FakeLockManager) UnlockCallCount() int {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return len(fake.unlockArgsForCall)
}

func (fake *FakeLockManager) UnlockArgsForCall(i int) string {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return fake.unlockArgsForCall[i].key
}

func (fake *FakeLockManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLockManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ volume.LockManager = new(FakeLockManager)