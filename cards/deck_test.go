package main

import (
	"reflect"
	"testing"
)

func Test_newDeck(t *testing.T) {
	type args struct {
		filename []string
	}
	tests := []struct {
		name    string
		args    args
		wantD   deck
		wantErr bool
	}{
		{name: "File not found", args: args{filename: []string{""}}, wantD: deck{}, wantErr: true},
		{name: "Full new deck", args: args{filename: nil}, wantD: generateDeck(), wantErr: false},
		{name: "Read from file", args: args{filename: []string{"cards.txt"}}, wantD: generateDeck(), wantErr: false},
		{name: "Attempt multiple file read", args: args{filename: []string{"cards.txt", "cards1.txt"}}, wantD: deck{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := newDeck(tt.args.filename...)
			if (err != nil) != tt.wantErr {
				t.Errorf("newDeck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("newDeck() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func Test_generateDeck(t *testing.T) {
	tests := []struct {
		name  string
		wantD deck
	}{
		{name: "Deck generation", wantD: deck{"Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs", "Ace of Diamonds", "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotD := generateDeck(); !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("generateDeck() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func Test_deck_getDeckString(t *testing.T) {
	type args struct {
		sep string
	}
	tests := []struct {
		name  string
		d     deck
		args  args
		wantS string
	}{
		{name: "Comma separated", d: deck{"Test1", "Test2", "Test3", "Test4"}, args: args{","}, wantS: "Test1,Test2,Test3,Test4"},
		{name: "Space separated", d: deck{"Test1", "Test2", "Test3", "Test4"}, args: args{" "}, wantS: "Test1 Test2 Test3 Test4"},
		{name: "Empty string", d: deck{""}, args: args{","}, wantS: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := tt.d.getDeckString(tt.args.sep); gotS != tt.wantS {
				t.Errorf("deck.getDeckString() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_deal(t *testing.T) {
	type args struct {
		d        deck
		handSize int
	}
	tests := []struct {
		name  string
		args  args
		wantH deck
		wantR deck
	}{
		{name: "Deal", args: args{d: deck{"Test1", "Test2", "Test3", "Test4"}, handSize: 1}, wantH: deck{"Test1"}, wantR: deck{"Test2", "Test3", "Test4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotH, gotR := deal(tt.args.d, tt.args.handSize)
			if !reflect.DeepEqual(gotH, tt.wantH) {
				t.Errorf("deal() gotH = %v, want %v", gotH, tt.wantH)
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("deal() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func Test_deck_saveToFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		d       deck
		args    args
		wantErr bool
	}{
		{name: "Sucessful write", d: deck{"Test1", "Test2", "Test3", "Test4"}, args: args{"cards2.txt"}, wantErr: false},
		{name: "Unsucessful write", d: deck{"Test1", "Test2", "Test3", "Test4"}, args: args{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.saveToFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("deck.saveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
