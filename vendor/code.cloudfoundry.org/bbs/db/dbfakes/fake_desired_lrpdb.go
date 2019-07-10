// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeDesiredLRPDB struct {
	DesireLRPStub        func(context.Context, lager.Logger, *models.DesiredLRP) error
	desireLRPMutex       sync.RWMutex
	desireLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.DesiredLRP
	}
	desireLRPReturns struct {
		result1 error
	}
	desireLRPReturnsOnCall map[int]struct {
		result1 error
	}
	DesiredLRPByProcessGuidStub        func(context.Context, lager.Logger, string) (*models.DesiredLRP, error)
	desiredLRPByProcessGuidMutex       sync.RWMutex
	desiredLRPByProcessGuidArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	desiredLRPByProcessGuidReturns struct {
		result1 *models.DesiredLRP
		result2 error
	}
	desiredLRPByProcessGuidReturnsOnCall map[int]struct {
		result1 *models.DesiredLRP
		result2 error
	}
	DesiredLRPSchedulingInfosStub        func(context.Context, lager.Logger, models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error)
	desiredLRPSchedulingInfosMutex       sync.RWMutex
	desiredLRPSchedulingInfosArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 models.DesiredLRPFilter
	}
	desiredLRPSchedulingInfosReturns struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	desiredLRPSchedulingInfosReturnsOnCall map[int]struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}
	DesiredLRPsStub        func(context.Context, lager.Logger, models.DesiredLRPFilter) ([]*models.DesiredLRP, error)
	desiredLRPsMutex       sync.RWMutex
	desiredLRPsArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 models.DesiredLRPFilter
	}
	desiredLRPsReturns struct {
		result1 []*models.DesiredLRP
		result2 error
	}
	desiredLRPsReturnsOnCall map[int]struct {
		result1 []*models.DesiredLRP
		result2 error
	}
	RemoveDesiredLRPStub        func(context.Context, lager.Logger, string) error
	removeDesiredLRPMutex       sync.RWMutex
	removeDesiredLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	removeDesiredLRPReturns struct {
		result1 error
	}
	removeDesiredLRPReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateDesiredLRPStub        func(context.Context, lager.Logger, string, *models.DesiredLRPUpdate) (*models.DesiredLRP, error)
	updateDesiredLRPMutex       sync.RWMutex
	updateDesiredLRPArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 *models.DesiredLRPUpdate
	}
	updateDesiredLRPReturns struct {
		result1 *models.DesiredLRP
		result2 error
	}
	updateDesiredLRPReturnsOnCall map[int]struct {
		result1 *models.DesiredLRP
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDesiredLRPDB) DesireLRP(arg1 context.Context, arg2 lager.Logger, arg3 *models.DesiredLRP) error {
	fake.desireLRPMutex.Lock()
	ret, specificReturn := fake.desireLRPReturnsOnCall[len(fake.desireLRPArgsForCall)]
	fake.desireLRPArgsForCall = append(fake.desireLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 *models.DesiredLRP
	}{arg1, arg2, arg3})
	fake.recordInvocation("DesireLRP", []interface{}{arg1, arg2, arg3})
	fake.desireLRPMutex.Unlock()
	if fake.DesireLRPStub != nil {
		return fake.DesireLRPStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.desireLRPReturns
	return fakeReturns.result1
}

func (fake *FakeDesiredLRPDB) DesireLRPCallCount() int {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	return len(fake.desireLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesireLRPCalls(stub func(context.Context, lager.Logger, *models.DesiredLRP) error) {
	fake.desireLRPMutex.Lock()
	defer fake.desireLRPMutex.Unlock()
	fake.DesireLRPStub = stub
}

func (fake *FakeDesiredLRPDB) DesireLRPArgsForCall(i int) (context.Context, lager.Logger, *models.DesiredLRP) {
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	argsForCall := fake.desireLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDesiredLRPDB) DesireLRPReturns(result1 error) {
	fake.desireLRPMutex.Lock()
	defer fake.desireLRPMutex.Unlock()
	fake.DesireLRPStub = nil
	fake.desireLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) DesireLRPReturnsOnCall(i int, result1 error) {
	fake.desireLRPMutex.Lock()
	defer fake.desireLRPMutex.Unlock()
	fake.DesireLRPStub = nil
	if fake.desireLRPReturnsOnCall == nil {
		fake.desireLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.desireLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuid(arg1 context.Context, arg2 lager.Logger, arg3 string) (*models.DesiredLRP, error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	ret, specificReturn := fake.desiredLRPByProcessGuidReturnsOnCall[len(fake.desiredLRPByProcessGuidArgsForCall)]
	fake.desiredLRPByProcessGuidArgsForCall = append(fake.desiredLRPByProcessGuidArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("DesiredLRPByProcessGuid", []interface{}{arg1, arg2, arg3})
	fake.desiredLRPByProcessGuidMutex.Unlock()
	if fake.DesiredLRPByProcessGuidStub != nil {
		return fake.DesiredLRPByProcessGuidStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.desiredLRPByProcessGuidReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidCallCount() int {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	return len(fake.desiredLRPByProcessGuidArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidCalls(stub func(context.Context, lager.Logger, string) (*models.DesiredLRP, error)) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	defer fake.desiredLRPByProcessGuidMutex.Unlock()
	fake.DesiredLRPByProcessGuidStub = stub
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	argsForCall := fake.desiredLRPByProcessGuidArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidReturns(result1 *models.DesiredLRP, result2 error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	defer fake.desiredLRPByProcessGuidMutex.Unlock()
	fake.DesiredLRPByProcessGuidStub = nil
	fake.desiredLRPByProcessGuidReturns = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPByProcessGuidReturnsOnCall(i int, result1 *models.DesiredLRP, result2 error) {
	fake.desiredLRPByProcessGuidMutex.Lock()
	defer fake.desiredLRPByProcessGuidMutex.Unlock()
	fake.DesiredLRPByProcessGuidStub = nil
	if fake.desiredLRPByProcessGuidReturnsOnCall == nil {
		fake.desiredLRPByProcessGuidReturnsOnCall = make(map[int]struct {
			result1 *models.DesiredLRP
			result2 error
		})
	}
	fake.desiredLRPByProcessGuidReturnsOnCall[i] = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfos(arg1 context.Context, arg2 lager.Logger, arg3 models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error) {
	fake.desiredLRPSchedulingInfosMutex.Lock()
	ret, specificReturn := fake.desiredLRPSchedulingInfosReturnsOnCall[len(fake.desiredLRPSchedulingInfosArgsForCall)]
	fake.desiredLRPSchedulingInfosArgsForCall = append(fake.desiredLRPSchedulingInfosArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 models.DesiredLRPFilter
	}{arg1, arg2, arg3})
	fake.recordInvocation("DesiredLRPSchedulingInfos", []interface{}{arg1, arg2, arg3})
	fake.desiredLRPSchedulingInfosMutex.Unlock()
	if fake.DesiredLRPSchedulingInfosStub != nil {
		return fake.DesiredLRPSchedulingInfosStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.desiredLRPSchedulingInfosReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosCallCount() int {
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	return len(fake.desiredLRPSchedulingInfosArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosCalls(stub func(context.Context, lager.Logger, models.DesiredLRPFilter) ([]*models.DesiredLRPSchedulingInfo, error)) {
	fake.desiredLRPSchedulingInfosMutex.Lock()
	defer fake.desiredLRPSchedulingInfosMutex.Unlock()
	fake.DesiredLRPSchedulingInfosStub = stub
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosArgsForCall(i int) (context.Context, lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	argsForCall := fake.desiredLRPSchedulingInfosArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosReturns(result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.desiredLRPSchedulingInfosMutex.Lock()
	defer fake.desiredLRPSchedulingInfosMutex.Unlock()
	fake.DesiredLRPSchedulingInfosStub = nil
	fake.desiredLRPSchedulingInfosReturns = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPSchedulingInfosReturnsOnCall(i int, result1 []*models.DesiredLRPSchedulingInfo, result2 error) {
	fake.desiredLRPSchedulingInfosMutex.Lock()
	defer fake.desiredLRPSchedulingInfosMutex.Unlock()
	fake.DesiredLRPSchedulingInfosStub = nil
	if fake.desiredLRPSchedulingInfosReturnsOnCall == nil {
		fake.desiredLRPSchedulingInfosReturnsOnCall = make(map[int]struct {
			result1 []*models.DesiredLRPSchedulingInfo
			result2 error
		})
	}
	fake.desiredLRPSchedulingInfosReturnsOnCall[i] = struct {
		result1 []*models.DesiredLRPSchedulingInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPs(arg1 context.Context, arg2 lager.Logger, arg3 models.DesiredLRPFilter) ([]*models.DesiredLRP, error) {
	fake.desiredLRPsMutex.Lock()
	ret, specificReturn := fake.desiredLRPsReturnsOnCall[len(fake.desiredLRPsArgsForCall)]
	fake.desiredLRPsArgsForCall = append(fake.desiredLRPsArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 models.DesiredLRPFilter
	}{arg1, arg2, arg3})
	fake.recordInvocation("DesiredLRPs", []interface{}{arg1, arg2, arg3})
	fake.desiredLRPsMutex.Unlock()
	if fake.DesiredLRPsStub != nil {
		return fake.DesiredLRPsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.desiredLRPsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDesiredLRPDB) DesiredLRPsCallCount() int {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	return len(fake.desiredLRPsArgsForCall)
}

func (fake *FakeDesiredLRPDB) DesiredLRPsCalls(stub func(context.Context, lager.Logger, models.DesiredLRPFilter) ([]*models.DesiredLRP, error)) {
	fake.desiredLRPsMutex.Lock()
	defer fake.desiredLRPsMutex.Unlock()
	fake.DesiredLRPsStub = stub
}

func (fake *FakeDesiredLRPDB) DesiredLRPsArgsForCall(i int) (context.Context, lager.Logger, models.DesiredLRPFilter) {
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	argsForCall := fake.desiredLRPsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDesiredLRPDB) DesiredLRPsReturns(result1 []*models.DesiredLRP, result2 error) {
	fake.desiredLRPsMutex.Lock()
	defer fake.desiredLRPsMutex.Unlock()
	fake.DesiredLRPsStub = nil
	fake.desiredLRPsReturns = struct {
		result1 []*models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) DesiredLRPsReturnsOnCall(i int, result1 []*models.DesiredLRP, result2 error) {
	fake.desiredLRPsMutex.Lock()
	defer fake.desiredLRPsMutex.Unlock()
	fake.DesiredLRPsStub = nil
	if fake.desiredLRPsReturnsOnCall == nil {
		fake.desiredLRPsReturnsOnCall = make(map[int]struct {
			result1 []*models.DesiredLRP
			result2 error
		})
	}
	fake.desiredLRPsReturnsOnCall[i] = struct {
		result1 []*models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRP(arg1 context.Context, arg2 lager.Logger, arg3 string) error {
	fake.removeDesiredLRPMutex.Lock()
	ret, specificReturn := fake.removeDesiredLRPReturnsOnCall[len(fake.removeDesiredLRPArgsForCall)]
	fake.removeDesiredLRPArgsForCall = append(fake.removeDesiredLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("RemoveDesiredLRP", []interface{}{arg1, arg2, arg3})
	fake.removeDesiredLRPMutex.Unlock()
	if fake.RemoveDesiredLRPStub != nil {
		return fake.RemoveDesiredLRPStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeDesiredLRPReturns
	return fakeReturns.result1
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPCallCount() int {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	return len(fake.removeDesiredLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPCalls(stub func(context.Context, lager.Logger, string) error) {
	fake.removeDesiredLRPMutex.Lock()
	defer fake.removeDesiredLRPMutex.Unlock()
	fake.RemoveDesiredLRPStub = stub
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	argsForCall := fake.removeDesiredLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPReturns(result1 error) {
	fake.removeDesiredLRPMutex.Lock()
	defer fake.removeDesiredLRPMutex.Unlock()
	fake.RemoveDesiredLRPStub = nil
	fake.removeDesiredLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) RemoveDesiredLRPReturnsOnCall(i int, result1 error) {
	fake.removeDesiredLRPMutex.Lock()
	defer fake.removeDesiredLRPMutex.Unlock()
	fake.RemoveDesiredLRPStub = nil
	if fake.removeDesiredLRPReturnsOnCall == nil {
		fake.removeDesiredLRPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeDesiredLRPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRP(arg1 context.Context, arg2 lager.Logger, arg3 string, arg4 *models.DesiredLRPUpdate) (*models.DesiredLRP, error) {
	fake.updateDesiredLRPMutex.Lock()
	ret, specificReturn := fake.updateDesiredLRPReturnsOnCall[len(fake.updateDesiredLRPArgsForCall)]
	fake.updateDesiredLRPArgsForCall = append(fake.updateDesiredLRPArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
		arg4 *models.DesiredLRPUpdate
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("UpdateDesiredLRP", []interface{}{arg1, arg2, arg3, arg4})
	fake.updateDesiredLRPMutex.Unlock()
	if fake.UpdateDesiredLRPStub != nil {
		return fake.UpdateDesiredLRPStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateDesiredLRPReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPCallCount() int {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	return len(fake.updateDesiredLRPArgsForCall)
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPCalls(stub func(context.Context, lager.Logger, string, *models.DesiredLRPUpdate) (*models.DesiredLRP, error)) {
	fake.updateDesiredLRPMutex.Lock()
	defer fake.updateDesiredLRPMutex.Unlock()
	fake.UpdateDesiredLRPStub = stub
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPArgsForCall(i int) (context.Context, lager.Logger, string, *models.DesiredLRPUpdate) {
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	argsForCall := fake.updateDesiredLRPArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPReturns(result1 *models.DesiredLRP, result2 error) {
	fake.updateDesiredLRPMutex.Lock()
	defer fake.updateDesiredLRPMutex.Unlock()
	fake.UpdateDesiredLRPStub = nil
	fake.updateDesiredLRPReturns = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) UpdateDesiredLRPReturnsOnCall(i int, result1 *models.DesiredLRP, result2 error) {
	fake.updateDesiredLRPMutex.Lock()
	defer fake.updateDesiredLRPMutex.Unlock()
	fake.UpdateDesiredLRPStub = nil
	if fake.updateDesiredLRPReturnsOnCall == nil {
		fake.updateDesiredLRPReturnsOnCall = make(map[int]struct {
			result1 *models.DesiredLRP
			result2 error
		})
	}
	fake.updateDesiredLRPReturnsOnCall[i] = struct {
		result1 *models.DesiredLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeDesiredLRPDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.desireLRPMutex.RLock()
	defer fake.desireLRPMutex.RUnlock()
	fake.desiredLRPByProcessGuidMutex.RLock()
	defer fake.desiredLRPByProcessGuidMutex.RUnlock()
	fake.desiredLRPSchedulingInfosMutex.RLock()
	defer fake.desiredLRPSchedulingInfosMutex.RUnlock()
	fake.desiredLRPsMutex.RLock()
	defer fake.desiredLRPsMutex.RUnlock()
	fake.removeDesiredLRPMutex.RLock()
	defer fake.removeDesiredLRPMutex.RUnlock()
	fake.updateDesiredLRPMutex.RLock()
	defer fake.updateDesiredLRPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDesiredLRPDB) recordInvocation(key string, args []interface{}) {
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

var _ db.DesiredLRPDB = new(FakeDesiredLRPDB)
