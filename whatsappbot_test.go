package whatsappbot_test

import (
	"reflect"
	"testing"

	wb "github.com/JeremyJalpha/WhatsAppBot/whatsappbot"
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
	}

	for _, test := range tests {
		result, err := wb.ParseUpdateOrderCommand(test.commandText)
		if (err != nil) != test.expectError {
			t.Errorf("ParseUpdateOrderCommand(%q) error = %v, expectError %v", test.commandText, err, test.expectError)
			continue
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ParseUpdateOrderCommand(%q) = %v, want %v", test.commandText, result, test.expected)
		}
	}
}
