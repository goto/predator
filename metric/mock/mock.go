package mock

import (
	"github.com/goto/predator/protocol"
	"github.com/goto/predator/protocol/meta"
	"github.com/goto/predator/protocol/metric"
	"github.com/stretchr/testify/mock"
)

type mockMetricSpecGenerator struct {
	mock.Mock
}

func NewMetricSpecGenerator() *mockMetricSpecGenerator {
	return &mockMetricSpecGenerator{}
}

func (m *mockMetricSpecGenerator) GenerateMetricSpec(urn string) ([]*metric.Spec, error) {
	args := m.Called(urn)
	return args.Get(0).([]*metric.Spec), args.Error(1)
}

func (m *mockMetricSpecGenerator) Generate(tableSpec *meta.TableSpec, tolerances []*protocol.Tolerance) ([]*metric.Spec, error) {
	args := m.Called(tableSpec, tolerances)
	return args.Get(0).([]*metric.Spec), args.Error(1)
}
