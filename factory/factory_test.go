package factory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {

	s.adapter = &Factory{}
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSamble() {
	f := New()
	total := 15
	aSpots := f.StartAssemblingProcess(total)
	for _, aSpot := range aSpots {
		if aSpot == nil {
			continue
		}
		v := aSpot.GetAssembledVehicle()
		if v == nil {
			continue
		}

		fmt.Println(fmt.Sprintf("id: %v, testingLog: %v, assemblyLog: %v",
			v.Id,
			v.TestingLog,
			aSpot.GetAssembledLogs()))
	}
	s.Assert().Equal(len(aSpots), total)
}
