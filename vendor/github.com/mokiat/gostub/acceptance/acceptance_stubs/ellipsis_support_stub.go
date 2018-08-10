// Generated by 'github.com/mokiat/gostub'

package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/mokiat/gostub/acceptance"
	alias2 "github.com/mokiat/gostub/acceptance/external/external_dup"
)

type EllipsisSupportStub struct {
	StubGUID          int
	MethodStub        func(arg1 string, arg2 int, arg3 ...alias2.Address)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 []alias2.Address
	}
}

var _ alias1.EllipsisSupport = new(EllipsisSupportStub)

func (stub *EllipsisSupportStub) Method(arg1 string, arg2 int, arg3 ...alias2.Address) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 []alias2.Address
	}{arg1, arg2, arg3})
	if stub.MethodStub != nil {
		stub.MethodStub(arg1, arg2, arg3...)
	}
}
func (stub *EllipsisSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *EllipsisSupportStub) MethodArgsForCall(index int) (string, int, []alias2.Address) {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1, stub.methodArgsForCall[index].arg2, stub.methodArgsForCall[index].arg3
}