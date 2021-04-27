package main

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func Test_entities_Mid(t *testing.T) {
	type args struct {
		e entity
	}
	tests := []struct {
		name string
		c    entities
		args args
		want int
	}{
		{
			name: "zero arr",
			c:    nil,
			args: args{e: entity{someData: "a"}},
			want: 0,
		},
		{
			name: "6 els",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "g"},
			},
			args: args{e: entity{someData: "f"}},
			want: 5,
		},
		{
			name: "first",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
				{someData: "g"},
			},
			args: args{e: entity{someData: "a"}},
			want: 0,
		},
		{
			name: "first 2",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
				{someData: "g"},
			},
			args: args{e: entity{someData: "0"}},
			want: 0,
		},
		{
			name: "last",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
				{someData: "g"},
			},
			args: args{e: entity{someData: "g"}},
			want: 6,
		},
		{
			name: "last 2",
			c: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
			},
			args: args{e: entity{someData: "g"}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Mid(tt.args.e); got != tt.want {
				t.Errorf("Mid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedEntities_Add(t *testing.T) {
	type fields struct {
		entities entities
	}
	type args struct {
		e entity
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sortedEntities
	}{
		{
			name: "add 1",
			fields: fields{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
			args: args{e: entity{someData: "d"}},
			want: sortedEntities{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "d"},
				{someData: "e"},
			}},
		},
		{
			name: "add last",
			fields: fields{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
			args: args{e: entity{someData: "f"}},
			want: sortedEntities{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
				{someData: "f"},
			}},
		},
		{
			name: "add first 1",
			fields: fields{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
			args: args{e: entity{someData: "a"}},
			want: sortedEntities{entities: entities{
				{someData: "a"},
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
		},
		{
			name: "add first 2",
			fields: fields{entities: entities{
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
			args: args{e: entity{someData: "a"}},
			want: sortedEntities{entities: entities{
				{someData: "a"},
				{someData: "b"},
				{someData: "c"},
				{someData: "d"},
				{someData: "e"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := sortedEntities{
				entities: tt.fields.entities,
			}
			c.Add(tt.args.e)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("matching error\nwant: %+v\ngot: %+v", tt.want, c)
			}
		})
	}
}

func Test_sortedEntities_Find(t *testing.T) {
	type testR struct {
		someData string
		want     int
	}
	var (
		collection = sortedEntities{}
		tests      = make([]testR, 0, 10)
	)
	for nn := 0; nn < 1000; nn++ {
		collection.Add(entity{someData: strconv.FormatInt(int64(nn), 16) + "_" + strconv.FormatInt(rand.Int63(), 16)})
	}
	for nn := 0; nn < cap(tests); nn++ {
		r := rand.Intn(len(collection.entities))
		tests = append(tests, testR{
			someData: collection.entities[r].someData,
			want:     r,
		})
	}
	for _, tt := range tests {
		t.Run(tt.someData, func(t *testing.T) {
			if got := collection.Find(tt.someData); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
	// fmt.Printf("%+v", collection.entities[180:190])
}
