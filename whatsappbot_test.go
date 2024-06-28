package whatsappbot_test

import (
	"reflect"
	"testing"

	wb "github.com/JeremyJalpha/WhatsAppBot/whatsappbot"
	"github.com/stretchr/testify/assert"
)

func Test_ParseUpdateOrderCommand(t *testing.T) {
	tests := []struct {
		commandText string
		expected    []wb.MenuIndication
		expectError bool
	}{
		{
			commandText: "update order 6:0",
			expected: []wb.MenuIndication{
				{ItemMenuNum: 6, ItemAmount: "0"},
			},
			expectError: false,
		},
		{
			commandText: "update order: 6:0",
			expected: []wb.MenuIndication{
				{ItemMenuNum: 6, ItemAmount: "0"},
			},
			expectError: false,
		},
		{
			commandText: "update order 9:12, 10: 1x3, 3x2, 2x1, 6:5",
			expected: []wb.MenuIndication{
				{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
				{ItemMenuNum: 9, ItemAmount: "12"},
				{ItemMenuNum: 6, ItemAmount: "5"},
			},
			expectError: false,
		},
		// {
		// 	commandText: "update order 6-0",
		// 	expected: []wb.MenuIndication{
		// 		{ItemMenuNum: 6, ItemAmount: "0"},
		// 	},
		// 	expectError: true,
		// },
	}

	for _, test := range tests {
		result, err := wb.ParseUpdateOrderCommand(test.commandText)

		if err != nil {
			if test.expectError {
				t.Errorf("ParseUpdateOrderCommand(%q) error = %v, expectError %v", test.commandText, err, test.expectError)
				continue
			} else {
				t.Errorf("ParseUpdateOrderCommand(%q) received error = %v, but expected Error: %v", test.commandText, err, test.expectError)
				continue
			}
		} else {
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("ParseUpdateOrderCommand(%q) = %v, want %v", test.commandText, result, test.expected)
			}
		}
	}
}

// func (c *CustomerOrder) updateCustOrdItems(update OrderItems) error {
// baseline := "Update order 9:12, 10: 1x3, 3x2, 2x1, 6: 5"
func Test_UpdateCustOrdItems(t *testing.T) {
	var err error

	tests := []struct {
		given       wb.CustomerOrder
		update      wb.OrderItems
		expected    wb.OrderItems
		expectError bool
	}{
		{
			given: wb.CustomerOrder{
				OrderItems: wb.OrderItems{
					MenuIndications: []wb.MenuIndication{
						{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
						{ItemMenuNum: 9, ItemAmount: "12"},
						{ItemMenuNum: 6, ItemAmount: "5"},
					},
				},
			},
			update: wb.OrderItems{
				MenuIndications: []wb.MenuIndication{
					{ItemMenuNum: 6, ItemAmount: "0"},
				},
			},
			expected: wb.OrderItems{
				MenuIndications: []wb.MenuIndication{
					{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
					{ItemMenuNum: 9, ItemAmount: "12"},
				},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		err = test.given.UpdateCustOrdItems(test.update)
		assert.NoError(t, err)

		if (err != nil) != test.expectError {
			t.Errorf("UpdateCustOrdItems(%q) error = %v, expectError %v", test.update, err, test.expectError)
			continue
		}
		if !reflect.DeepEqual(test.given, test.expected) {
			t.Errorf("UpdateCustOrdItems(%q) = %v, want %v", test.update, test.given, test.expected)
		}
	}
}
