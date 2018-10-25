// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/copilot/api"
	"code.cloudfoundry.org/copilot/models"
)

type BackendSetRepo struct {
	GetInternalBackendsStub        func(guid models.DiegoProcessGUID) *api.BackendSet
	getInternalBackendsMutex       sync.RWMutex
	getInternalBackendsArgsForCall []struct {
		guid models.DiegoProcessGUID
	}
	getInternalBackendsReturns struct {
		result1 *api.BackendSet
	}
	getInternalBackendsReturnsOnCall map[int]struct {
		result1 *api.BackendSet
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BackendSetRepo) GetInternalBackends(guid models.DiegoProcessGUID) *api.BackendSet {
	fake.getInternalBackendsMutex.Lock()
	ret, specificReturn := fake.getInternalBackendsReturnsOnCall[len(fake.getInternalBackendsArgsForCall)]
	fake.getInternalBackendsArgsForCall = append(fake.getInternalBackendsArgsForCall, struct {
		guid models.DiegoProcessGUID
	}{guid})
	fake.recordInvocation("GetInternalBackends", []interface{}{guid})
	fake.getInternalBackendsMutex.Unlock()
	if fake.GetInternalBackendsStub != nil {
		return fake.GetInternalBackendsStub(guid)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getInternalBackendsReturns.result1
}

func (fake *BackendSetRepo) GetInternalBackendsCallCount() int {
	fake.getInternalBackendsMutex.RLock()
	defer fake.getInternalBackendsMutex.RUnlock()
	return len(fake.getInternalBackendsArgsForCall)
}

func (fake *BackendSetRepo) GetInternalBackendsArgsForCall(i int) models.DiegoProcessGUID {
	fake.getInternalBackendsMutex.RLock()
	defer fake.getInternalBackendsMutex.RUnlock()
	return fake.getInternalBackendsArgsForCall[i].guid
}

func (fake *BackendSetRepo) GetInternalBackendsReturns(result1 *api.BackendSet) {
	fake.GetInternalBackendsStub = nil
	fake.getInternalBackendsReturns = struct {
		result1 *api.BackendSet
	}{result1}
}

func (fake *BackendSetRepo) GetInternalBackendsReturnsOnCall(i int, result1 *api.BackendSet) {
	fake.GetInternalBackendsStub = nil
	if fake.getInternalBackendsReturnsOnCall == nil {
		fake.getInternalBackendsReturnsOnCall = make(map[int]struct {
			result1 *api.BackendSet
		})
	}
	fake.getInternalBackendsReturnsOnCall[i] = struct {
		result1 *api.BackendSet
	}{result1}
}

func (fake *BackendSetRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getInternalBackendsMutex.RLock()
	defer fake.getInternalBackendsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BackendSetRepo) recordInvocation(key string, args []interface{}) {
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