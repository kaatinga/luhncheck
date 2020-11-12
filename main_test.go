package luhncheck

import (
	"reflect"
	"testing"
)

const (
	notSoGoodCard = "4561261212345464"
	badCard1      = "1"
)

func TestNewCard(t *testing.T) {

	tests := []struct {
		name     string
		number   string
		bankCard *BankCard
		wantErr  bool
	}{
		{notSoGoodCard,
			notSoGoodCard,
			&BankCard{
				Number:  notSoGoodCard,
				Valid:   false,
				Checked: false,
			},
			false,
		},
		{badCard1, badCard1, nil, true},
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

func TestBankCard_Check(t *testing.T) {

	tests := []struct {
		name    string
		card  BankCard
		want    bool
		wantErr bool
	}{
		{"ok", BankCard{Number: notSoGoodCard}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.card.Check()
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
		})
	}
}