package models

//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/gocommerce/constants"
//	"github.com/vavilen84/gocommerce/validator"
//	"testing"
//)
//
//func TestOrderProduct_ValidateOnCreate(t *testing.T) {
//	m := OrderProduct{}
//	err := validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//	assert.NotEmpty(t, err[constants.OrderOrderIdField])
//	assert.NotEmpty(t, err[constants.OrderProductIdField])
//	assert.NotEmpty(t, err[constants.OrderQuantityField])
//
//	m = OrderProduct{
//		OrderId:   1,
//		ProductId: 1,
//		Quantity:  1,
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//}
