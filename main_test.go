package luhncheck

import (
	"reflect"
	"testing"
)

const (
	goodCard = "4561261212345464"
)

func TestNewCard(t *testing.T) {

	tests := []struct {
		name     string
		number   string
		bankCard *BankCard
		wantErr  bool
	}{
		{"ok",
			goodCard,
			&BankCard{
				Number:  goodCard,
				Valid:   false,
				Checked: false,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCard(tt.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.bankCard) {
				t.Errorf("NewCard() got = %v, want %v", got, tt.bankCard)
			}
		})
	}
}
