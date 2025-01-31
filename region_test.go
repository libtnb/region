package region

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RegionTestSuite struct {
	suite.Suite
	instance *Region
}

func TestRegionTestSuite(t *testing.T) {
	suite.Run(t, &RegionTestSuite{
		instance: NewRegion(),
	})
}

func (suite *RegionTestSuite) TestParseByCode_ValidCode() {
	province, city, area, street, err := suite.instance.ParseByCode("450305004")
	suite.NoError(err)
	suite.Equal("广西壮族自治区", province)
	suite.Equal("桂林市", city)
	suite.Equal("七星区", area)
	suite.Equal("漓东街道", street)
}

func (suite *RegionTestSuite) TestParseByCode_InvalidCode() {
	_, _, _, _, err := suite.instance.ParseByCode("invalidCode")
	suite.Error(err)
}

func (suite *RegionTestSuite) TestParseByName_ValidName() {
	code, err := suite.instance.ParseByName("广西壮族自治区", "桂林市", "七星区", "漓东街道")
	suite.NoError(err)
	suite.Equal("450305004", code)
}

func (suite *RegionTestSuite) TestParseByName_InvalidName() {
	_, err := suite.instance.ParseByName("invalidProvince", "invalidCity", "invalidArea", "invalidStreet")
	suite.Error(err)
}

func (suite *RegionTestSuite) TestSearch_ValidKeyword() {
	result := suite.instance.Search("广西壮族自治区")
	suite.NotEmpty(result)
	suite.Equal(14, len(result[0].Children))
	suite.Equal("广西壮族自治区", result[0].Name)
	suite.Equal("45", result[0].Code)
	result = suite.instance.Search("广西壮族自治区桂林市七星区")
	suite.NotEmpty(result)
	suite.Equal(6, len(result[0].Children))
	suite.Equal("七星区", result[0].Name)
	suite.Equal("450305", result[0].Code)
	result = suite.instance.Search("天津市市辖区和平区")
	suite.NotEmpty(result)
	suite.Equal(6, len(result[0].Children))
	suite.Equal("和平区", result[0].Name)
	suite.Equal("120101", result[0].Code)
	result = suite.instance.Search("天津市市辖区和平区南市街道")
	suite.NotEmpty(result)
	suite.Equal(0, len(result[0].Children))
	suite.Equal("南市街道", result[0].Name)
	suite.Equal("120101006", result[0].Code)
	result = suite.instance.Search("重庆市县城口县葛城街道")
	suite.NotEmpty(result)
	suite.Equal(0, len(result[0].Children))
	suite.Equal("葛城街道", result[0].Name)
	suite.Equal("500229001", result[0].Code)
}

func (suite *RegionTestSuite) TestSearch_InvalidKeyword() {
	result := suite.instance.Search("invalidKeyword")
	suite.Empty(result)
}
