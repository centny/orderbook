package orderbook

import (
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

var testTimestamp = 123452342343
var testQuanity, _ = decimal.NewFromString("0.1")
var testPrice, _ = decimal.NewFromString("0.1")
var testOrderId = 1
var testTradeId = 1

var testTimestamp1 = 123452342345
var testQuanity1, _ = decimal.NewFromString("0.2")
var testPrice1, _ = decimal.NewFromString("0.1")
var testOrderId1 = 2
var testTradeId1 = 2

func TestNewOrder(t *testing.T) {
	var order_list OrderList
	dummyOrder := make(map[string]string)
	dummyOrder["timestamp"] = strconv.Itoa(testTimestamp)
	dummyOrder["quantity"] = testQuanity.String()
	dummyOrder["price"] = testPrice.String()
	dummyOrder["order_id"] = strconv.Itoa(testOrderId)
	dummyOrder["trade_id"] = strconv.Itoa(testTradeId)

	order := NewOrder(dummyOrder, &order_list)

	if !(order.timestamp == testTimestamp) {
		t.Errorf("Timesmape incorrect, got: %d, want: %d.", order.timestamp, testTimestamp)
	}

	if !(order.quantity.Equal(testQuanity)) {
		t.Errorf("quantity incorrect, got: %d, want: %d.", order.quantity, testQuanity)
	}

	if !(order.price.Equal(testPrice)) {
		t.Errorf("price incorrect, got: %d, want: %d.", order.price, testPrice)
	}

	if !(order.order_id == strconv.Itoa(testOrderId)) {
		t.Errorf("order id incorrect, got: %s, want: %d.", order.order_id, testOrderId)
	}

	if !(order.trade_id == strconv.Itoa(testTradeId)) {
		t.Errorf("trade id incorrect, got: %s, want: %d.", order.trade_id, testTradeId)
	}
}

func TestOrder(t *testing.T) {
	orderList := NewOrderList(testPrice)

	dummyOrder := make(map[string]string)
	dummyOrder["timestamp"] = strconv.Itoa(testTimestamp)
	dummyOrder["quantity"] = testQuanity.String()
	dummyOrder["price"] = testPrice.String()
	dummyOrder["order_id"] = strconv.Itoa(testOrderId)
	dummyOrder["trade_id"] = strconv.Itoa(testTradeId)

	order := NewOrder(dummyOrder, orderList)

	orderList.AppendOrder(order)

	order.UpdateQuantity(testQuanity1, testTimestamp1)

	if !(order.quantity.Equal(testQuanity1)) {
		t.Errorf("order id incorrect, got: %s, want: %d.", order.order_id, testOrderId)
	}

	if !(order.timestamp == testTimestamp1) {
		t.Errorf("trade id incorrect, got: %s, want: %d.", order.trade_id, testTradeId)
	}
}
