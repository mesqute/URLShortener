package models

import (
	"sync"
	"testing"
)

func Test_inMemory_Insert(t *testing.T) {
	type fields struct {
		items map[string]value
		lock  sync.RWMutex
	}
	type args struct {
		URL string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantLink string
		wantErr  bool
	}{
		{
			name: "нормальная вставка",
			fields: fields{
				items: make(map[string]value),
				lock:  sync.RWMutex{},
			},
			args:    args{URL: "dsadas"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				items: tt.fields.items,
				lock:  tt.fields.lock,
			}
			_, err := i.Insert(tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_inMemory_Get(t *testing.T) {
	type fields struct {
		items map[string]value
		lock  sync.RWMutex
	}
	type args struct {
		link string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantURL string
		wantErr bool
	}{
		{
			name: "нормальное получение",
			fields: fields{
				items: map[string]value{
					"adu2781_1s": {url: "url"},
				},
				lock: sync.RWMutex{},
			},
			args: args{
				link: "adu2781_1s",
			},
			wantURL: "url",
			wantErr: false,
		},
		{
			name: "несуществующий link",
			fields: fields{
				items: map[string]value{
					"adu2781_1s": {url: "url"},
				},
				lock: sync.RWMutex{},
			},
			args: args{
				link: "s",
			},
			wantURL: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				items: tt.fields.items,
				lock:  tt.fields.lock,
			}
			gotURL, err := i.Get(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotURL != tt.wantURL {
				t.Errorf("Get() gotURL = %v, want %v", gotURL, tt.wantURL)
			}
		})
	}
}
