package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/repository"
)

func TestInventoryService_Create(t *testing.T) {
	type fields struct {
		repository repository.IInventoryRepository
	}
	type args struct {
		inventory domain.InventoryItem
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create inventory happy scenario",
			fields: fields{
				repository: repository.NewInventoryInMemoryRepository(),
			},
			args: args{
				inventory: domain.InventoryItem{
					Slug:        "test",
					Name:        "test",
					Description: "test",
					Deleted:     false,
				},
			},
			wantErr: false,
		},
		{
			name: "Create inventory with non-alphanumeric slug",
			fields: fields{
				repository: repository.NewInventoryInMemoryRepository(),
			},
			args: args{
				inventory: domain.InventoryItem{
					Slug:        "test 123",
					Name:        "test",
					Description: "test",
					Deleted:     false,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := InventoryService{
				repository: tt.fields.repository,
			}
			if err := service.Create(tt.args.inventory); (err != nil) != tt.wantErr {
				fmt.Println("Heeeere", err)
				t.Errorf("InventoryService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInventoryService_List(t *testing.T) {
	type fields struct {
		repository repository.IInventoryRepository
	}
	type args struct {
		deleted bool
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		fixtures []domain.InventoryItem
		want     []domain.InventoryItem
		wantErr  bool
	}{
		{
			name: "List undeleted inventory",
			fields: fields{
				repository: repository.NewInventoryInMemoryRepository(),
			},
			args: args{
				deleted: false,
			},
			wantErr: false,
			fixtures: []domain.InventoryItem{
				{
					Slug:        "test",
					Name:        "test",
					Description: "test",
					Deleted:     false,
				},
				{
					Slug:        "test2",
					Name:        "test2",
					Description: "test2",
					Deleted:     true,
				},
			},
			want: []domain.InventoryItem{
				{
					Slug:        "test",
					Name:        "test",
					Description: "test",
					Deleted:     false,
				},
			},
		},
		{
			name: "List deleted inventory",
			fields: fields{
				repository: repository.NewInventoryInMemoryRepository(),
			},
			args: args{
				deleted: true,
			},
			wantErr: false,
			fixtures: []domain.InventoryItem{
				{
					Slug:        "test",
					Name:        "test",
					Description: "test",
					Deleted:     false,
				},
				{
					Slug:        "test2",
					Name:        "test2",
					Description: "test2",
					Deleted:     true,
				},
			},
			want: []domain.InventoryItem{
				{
					Slug:        "test2",
					Name:        "test2",
					Description: "test2",
					Deleted:     true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := InventoryService{
				repository: tt.fields.repository,
			}

			service.repository.Create(tt.fixtures[0])
			service.repository.Create(tt.fixtures[1])

			got, err := service.List(tt.args.deleted)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryService.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventoryService_GET(t *testing.T) {
	type fields struct {
		repository repository.IInventoryRepository
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		fixture domain.InventoryItem
		want    domain.InventoryItem
		wantErr bool
	}{
		{
			name: "Get item",
			fields: fields{
				repository: repository.NewInventoryInMemoryRepository(),
			},
			args: args{
				slug: "test",
			},
			fixture: domain.InventoryItem{
				Slug:        "test",
				Name:        "test",
				Description: "test",
				Deleted:     false,
			},
			want: domain.InventoryItem{
				Slug:        "test",
				Name:        "test",
				Description: "test",
				Deleted:     false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := InventoryService{
				repository: tt.fields.repository,
			}

			service.repository.Create(tt.fixture)

			got, err := service.GET(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.GET() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InventoryService.GET() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventoryService_Update(t *testing.T) {
	type fields struct {
		repository repository.IInventoryRepository
	}
	type args struct {
		item domain.InventoryItem
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		fixtures   []domain.InventoryItem
		want       domain.InventoryItem
		wantErr    bool
		wantGetErr bool
	}{
		{
			name: "Update item happy scenario",
			fields: fields{
				repository: repository.NewInventoryInMemoryRepository(),
			},
			args: args{
				item: domain.InventoryItem{
					Slug:        "test",
					Name:        "test updated",
					Description: "test updated",
					Deleted:     false,
				},
			},
			fixtures: []domain.InventoryItem{
				{
					Slug:        "test",
					Name:        "test",
					Description: "test",
					Deleted:     false,
				},
			},
			want: domain.InventoryItem{
				Slug:        "test",
				Name:        "test updated",
				Description: "test updated",
				Deleted:     false,
			},
			wantErr: false,
		},
		{
			name: "Update item with not found slug",
			fields: fields{
				repository: repository.NewInventoryInMemoryRepository(),
			},
			args: args{
				item: domain.InventoryItem{
					Slug:        "I_am_not_in_the_database",
					Name:        "test updated",
					Description: "test updated",
					Deleted:     false,
				},
			},
			fixtures: []domain.InventoryItem{
				{
					Slug:        "test",
					Name:        "test",
					Description: "test",
					Deleted:     false,
				},
			},
			want: domain.InventoryItem{
				Slug:        "test",
				Name:        "test updated",
				Description: "test updated",
				Deleted:     false,
			},
			wantErr:    true,
			wantGetErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := InventoryService{
				repository: tt.fields.repository,
			}

			service.repository.Create(tt.fixtures[0])

			if err := service.Update(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := service.GET(tt.args.item.Slug)
			if (err != nil) != tt.wantGetErr {
				t.Errorf("InventoryService.GET() error = %v, wantErr %v", err, tt.wantGetErr)
				return
			}

			if !tt.wantGetErr {
				// make the autogenrated dates the same
				tt.want.CreatedAt = got.CreatedAt
				tt.want.UpdatedAt = got.UpdatedAt

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("InventoryService.GET() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}

func TestInventoryService_Delete(t *testing.T) {
	type fields struct {
		repository repository.IInventoryRepository
	}
	type args struct {
		slug    string
		comment string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		fixtures   []domain.InventoryItem
		want       domain.InventoryItem
		wantGetErr bool
		wantErr    bool
	}{
		{
			name: "Delete item happy scenario",
			fields: fields{
				repository: repository.NewInventoryInMemoryRepository(),
			},
			args: args{
				slug:    "test",
				comment: "test comment",
			},
			fixtures: []domain.InventoryItem{
				{
					Slug:        "test",
					Name:        "test",
					Description: "test",
					Deleted:     false,
				},
			},
			want: domain.InventoryItem{
				Slug:          "test",
				Name:          "test",
				Description:   "test",
				Deleted:       true,
				DeleteComment: "test comment",
			},
			wantErr:    false,
			wantGetErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := InventoryService{
				repository: tt.fields.repository,
			}

			service.repository.Create(tt.fixtures[0])

			if err := service.Delete(tt.args.slug, tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := service.GET(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("InventoryService.GET() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantGetErr {
				// make the autogenrated dates the same
				tt.want.CreatedAt = got.CreatedAt
				tt.want.UpdatedAt = got.UpdatedAt

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("InventoryService.GET() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
