// Code generated by counterfeiter. DO NOT EDIT.
package dockerdriverfakes

import (
	"sync"

	"code.cloudfoundry.org/dockerdriver"
	"code.cloudfoundry.org/dockerdriver/invoker"
)

type FakeInvoker struct {
	InvokeStub        func(env dockerdriver.Env, executable string, args []string) ([]byte, error)
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		env        dockerdriver.Env
		executable string
		args       []string
	}
	invokeReturns struct {
		result1 []byte
		result2 error
	}
	invokeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInvoker) Invoke(env dockerdriver.Env, executable string, args []string) ([]byte, error) {
	var argsCopy []string
	if args != nil {
		argsCopy = make([]string, len(args))
		copy(argsCopy, args)
	}
	fake.invokeMutex.Lock()
	ret, specificReturn := fake.invokeReturnsOnCall[len(fake.invokeArgsForCall)]
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		env        dockerdriver.Env
		executable string
		args       []string
	}{env, executable, argsCopy})
	fake.recordInvocation("Invoke", []interface{}{env, executable, argsCopy})
	fake.invokeMutex.Unlock()
	if fake.InvokeStub != nil {
		return fake.InvokeStub(env, executable, args)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.invokeReturns.result1, fake.invokeReturns.result2
}

func (fake *FakeInvoker) InvokeCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *FakeInvoker) InvokeArgsForCall(i int) (dockerdriver.Env, string, []string) {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return fake.invokeArgsForCall[i].env, fake.invokeArgsForCall[i].executable, fake.invokeArgsForCall[i].args
}

func (fake *FakeInvoker) InvokeReturns(result1 []byte, result2 error) {
	fake.InvokeStub = nil
	fake.invokeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInvoker) InvokeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.InvokeStub = nil
	if fake.invokeReturnsOnCall == nil {
		fake.invokeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.invokeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInvoker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInvoker) recordInvocation(key string, args []interface{}) {
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

var _ invoker.Invoker = new(FakeInvoker)
