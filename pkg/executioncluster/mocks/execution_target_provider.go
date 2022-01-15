// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	executioncluster "github.com/flyteorg/flyteadmin/pkg/executioncluster"

	interfaces "github.com/flyteorg/flyteadmin/pkg/runtime/interfaces"

	mock "github.com/stretchr/testify/mock"

	prometheus "github.com/prometheus/client_golang/prometheus"
)

// ExecutionTargetProvider is an autogenerated mock type for the ExecutionTargetProvider type
type ExecutionTargetProvider struct {
	mock.Mock
}

type ExecutionTargetProvider_GetExecutionTarget struct {
	*mock.Call
}

func (_m ExecutionTargetProvider_GetExecutionTarget) Return(_a0 *executioncluster.ExecutionTarget, _a1 error) *ExecutionTargetProvider_GetExecutionTarget {
	return &ExecutionTargetProvider_GetExecutionTarget{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ExecutionTargetProvider) OnGetExecutionTarget(initializationErrorCounter prometheus.Counter, k8sCluster interfaces.ClusterConfig) *ExecutionTargetProvider_GetExecutionTarget {
	c := _m.On("GetExecutionTarget", initializationErrorCounter, k8sCluster)
	return &ExecutionTargetProvider_GetExecutionTarget{Call: c}
}

func (_m *ExecutionTargetProvider) OnGetExecutionTargetMatch(matchers ...interface{}) *ExecutionTargetProvider_GetExecutionTarget {
	c := _m.On("GetExecutionTarget", matchers...)
	return &ExecutionTargetProvider_GetExecutionTarget{Call: c}
}

// GetExecutionTarget provides a mock function with given fields: initializationErrorCounter, k8sCluster
func (_m *ExecutionTargetProvider) GetExecutionTarget(initializationErrorCounter prometheus.Counter, k8sCluster interfaces.ClusterConfig) (*executioncluster.ExecutionTarget, error) {
	ret := _m.Called(initializationErrorCounter, k8sCluster)

	var r0 *executioncluster.ExecutionTarget
	if rf, ok := ret.Get(0).(func(prometheus.Counter, interfaces.ClusterConfig) *executioncluster.ExecutionTarget); ok {
		r0 = rf(initializationErrorCounter, k8sCluster)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executioncluster.ExecutionTarget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(prometheus.Counter, interfaces.ClusterConfig) error); ok {
		r1 = rf(initializationErrorCounter, k8sCluster)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
