// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/baggageclaim/volume"
)

type FakeFilesystemLiveVolume struct {
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns     struct {
		result1 string
	}
	DataPathStub        func() string
	dataPathMutex       sync.RWMutex
	dataPathArgsForCall []struct{}
	dataPathReturns     struct {
		result1 string
	}
	LoadPropertiesStub        func() (volume.Properties, error)
	loadPropertiesMutex       sync.RWMutex
	loadPropertiesArgsForCall []struct{}
	loadPropertiesReturns     struct {
		result1 volume.Properties
		result2 error
	}
	StorePropertiesStub        func(volume.Properties) error
	storePropertiesMutex       sync.RWMutex
	storePropertiesArgsForCall []struct {
		arg1 volume.Properties
	}
	storePropertiesReturns struct {
		result1 error
	}
	LoadTTLStub        func() (volume.TTL, time.Time, error)
	loadTTLMutex       sync.RWMutex
	loadTTLArgsForCall []struct{}
	loadTTLReturns     struct {
		result1 volume.TTL
		result2 time.Time
		result3 error
	}
	StoreTTLStub        func(volume.TTL) (time.Time, error)
	storeTTLMutex       sync.RWMutex
	storeTTLArgsForCall []struct {
		arg1 volume.TTL
	}
	storeTTLReturns struct {
		result1 time.Time
		result2 error
	}
	ParentStub        func() (volume.FilesystemLiveVolume, bool, error)
	parentMutex       sync.RWMutex
	parentArgsForCall []struct{}
	parentReturns     struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}
	DestroyStub        func() error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct{}
	destroyReturns     struct {
		result1 error
	}
	SizeInBytesStub        func() (int64, error)
	sizeInBytesMutex       sync.RWMutex
	sizeInBytesArgsForCall []struct{}
	sizeInBytesReturns     struct {
		result1 int64
		result2 error
	}
	NewSubvolumeStub        func(handle string) (volume.FilesystemInitVolume, error)
	newSubvolumeMutex       sync.RWMutex
	newSubvolumeArgsForCall []struct {
		handle string
	}
	newSubvolumeReturns struct {
		result1 volume.FilesystemInitVolume
		result2 error
	}
}

func (fake *FakeFilesystemLiveVolume) Handle() string {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	} else {
		return fake.handleReturns.result1
	}
}

func (fake *FakeFilesystemLiveVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) DataPath() string {
	fake.dataPathMutex.Lock()
	fake.dataPathArgsForCall = append(fake.dataPathArgsForCall, struct{}{})
	fake.dataPathMutex.Unlock()
	if fake.DataPathStub != nil {
		return fake.DataPathStub()
	} else {
		return fake.dataPathReturns.result1
	}
}

func (fake *FakeFilesystemLiveVolume) DataPathCallCount() int {
	fake.dataPathMutex.RLock()
	defer fake.dataPathMutex.RUnlock()
	return len(fake.dataPathArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) DataPathReturns(result1 string) {
	fake.DataPathStub = nil
	fake.dataPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) LoadProperties() (volume.Properties, error) {
	fake.loadPropertiesMutex.Lock()
	fake.loadPropertiesArgsForCall = append(fake.loadPropertiesArgsForCall, struct{}{})
	fake.loadPropertiesMutex.Unlock()
	if fake.LoadPropertiesStub != nil {
		return fake.LoadPropertiesStub()
	} else {
		return fake.loadPropertiesReturns.result1, fake.loadPropertiesReturns.result2
	}
}

func (fake *FakeFilesystemLiveVolume) LoadPropertiesCallCount() int {
	fake.loadPropertiesMutex.RLock()
	defer fake.loadPropertiesMutex.RUnlock()
	return len(fake.loadPropertiesArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) LoadPropertiesReturns(result1 volume.Properties, result2 error) {
	fake.LoadPropertiesStub = nil
	fake.loadPropertiesReturns = struct {
		result1 volume.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) StoreProperties(arg1 volume.Properties) error {
	fake.storePropertiesMutex.Lock()
	fake.storePropertiesArgsForCall = append(fake.storePropertiesArgsForCall, struct {
		arg1 volume.Properties
	}{arg1})
	fake.storePropertiesMutex.Unlock()
	if fake.StorePropertiesStub != nil {
		return fake.StorePropertiesStub(arg1)
	} else {
		return fake.storePropertiesReturns.result1
	}
}

func (fake *FakeFilesystemLiveVolume) StorePropertiesCallCount() int {
	fake.storePropertiesMutex.RLock()
	defer fake.storePropertiesMutex.RUnlock()
	return len(fake.storePropertiesArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) StorePropertiesArgsForCall(i int) volume.Properties {
	fake.storePropertiesMutex.RLock()
	defer fake.storePropertiesMutex.RUnlock()
	return fake.storePropertiesArgsForCall[i].arg1
}

func (fake *FakeFilesystemLiveVolume) StorePropertiesReturns(result1 error) {
	fake.StorePropertiesStub = nil
	fake.storePropertiesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) LoadTTL() (volume.TTL, time.Time, error) {
	fake.loadTTLMutex.Lock()
	fake.loadTTLArgsForCall = append(fake.loadTTLArgsForCall, struct{}{})
	fake.loadTTLMutex.Unlock()
	if fake.LoadTTLStub != nil {
		return fake.LoadTTLStub()
	} else {
		return fake.loadTTLReturns.result1, fake.loadTTLReturns.result2, fake.loadTTLReturns.result3
	}
}

func (fake *FakeFilesystemLiveVolume) LoadTTLCallCount() int {
	fake.loadTTLMutex.RLock()
	defer fake.loadTTLMutex.RUnlock()
	return len(fake.loadTTLArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) LoadTTLReturns(result1 volume.TTL, result2 time.Time, result3 error) {
	fake.LoadTTLStub = nil
	fake.loadTTLReturns = struct {
		result1 volume.TTL
		result2 time.Time
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFilesystemLiveVolume) StoreTTL(arg1 volume.TTL) (time.Time, error) {
	fake.storeTTLMutex.Lock()
	fake.storeTTLArgsForCall = append(fake.storeTTLArgsForCall, struct {
		arg1 volume.TTL
	}{arg1})
	fake.storeTTLMutex.Unlock()
	if fake.StoreTTLStub != nil {
		return fake.StoreTTLStub(arg1)
	} else {
		return fake.storeTTLReturns.result1, fake.storeTTLReturns.result2
	}
}

func (fake *FakeFilesystemLiveVolume) StoreTTLCallCount() int {
	fake.storeTTLMutex.RLock()
	defer fake.storeTTLMutex.RUnlock()
	return len(fake.storeTTLArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) StoreTTLArgsForCall(i int) volume.TTL {
	fake.storeTTLMutex.RLock()
	defer fake.storeTTLMutex.RUnlock()
	return fake.storeTTLArgsForCall[i].arg1
}

func (fake *FakeFilesystemLiveVolume) StoreTTLReturns(result1 time.Time, result2 error) {
	fake.StoreTTLStub = nil
	fake.storeTTLReturns = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) Parent() (volume.FilesystemLiveVolume, bool, error) {
	fake.parentMutex.Lock()
	fake.parentArgsForCall = append(fake.parentArgsForCall, struct{}{})
	fake.parentMutex.Unlock()
	if fake.ParentStub != nil {
		return fake.ParentStub()
	} else {
		return fake.parentReturns.result1, fake.parentReturns.result2, fake.parentReturns.result3
	}
}

func (fake *FakeFilesystemLiveVolume) ParentCallCount() int {
	fake.parentMutex.RLock()
	defer fake.parentMutex.RUnlock()
	return len(fake.parentArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) ParentReturns(result1 volume.FilesystemLiveVolume, result2 bool, result3 error) {
	fake.ParentStub = nil
	fake.parentReturns = struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFilesystemLiveVolume) Destroy() error {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub()
	} else {
		return fake.destroyReturns.result1
	}
}

func (fake *FakeFilesystemLiveVolume) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) SizeInBytes() (int64, error) {
	fake.sizeInBytesMutex.Lock()
	fake.sizeInBytesArgsForCall = append(fake.sizeInBytesArgsForCall, struct{}{})
	fake.sizeInBytesMutex.Unlock()
	if fake.SizeInBytesStub != nil {
		return fake.SizeInBytesStub()
	} else {
		return fake.sizeInBytesReturns.result1, fake.sizeInBytesReturns.result2
	}
}

func (fake *FakeFilesystemLiveVolume) SizeInBytesCallCount() int {
	fake.sizeInBytesMutex.RLock()
	defer fake.sizeInBytesMutex.RUnlock()
	return len(fake.sizeInBytesArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) SizeInBytesReturns(result1 int64, result2 error) {
	fake.SizeInBytesStub = nil
	fake.sizeInBytesReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) NewSubvolume(handle string) (volume.FilesystemInitVolume, error) {
	fake.newSubvolumeMutex.Lock()
	fake.newSubvolumeArgsForCall = append(fake.newSubvolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.newSubvolumeMutex.Unlock()
	if fake.NewSubvolumeStub != nil {
		return fake.NewSubvolumeStub(handle)
	} else {
		return fake.newSubvolumeReturns.result1, fake.newSubvolumeReturns.result2
	}
}

func (fake *FakeFilesystemLiveVolume) NewSubvolumeCallCount() int {
	fake.newSubvolumeMutex.RLock()
	defer fake.newSubvolumeMutex.RUnlock()
	return len(fake.newSubvolumeArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) NewSubvolumeArgsForCall(i int) string {
	fake.newSubvolumeMutex.RLock()
	defer fake.newSubvolumeMutex.RUnlock()
	return fake.newSubvolumeArgsForCall[i].handle
}

func (fake *FakeFilesystemLiveVolume) NewSubvolumeReturns(result1 volume.FilesystemInitVolume, result2 error) {
	fake.NewSubvolumeStub = nil
	fake.newSubvolumeReturns = struct {
		result1 volume.FilesystemInitVolume
		result2 error
	}{result1, result2}
}

var _ volume.FilesystemLiveVolume = new(FakeFilesystemLiveVolume)
