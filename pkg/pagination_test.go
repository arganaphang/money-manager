package pkg

import (
	"reflect"
	"testing"
)

func TestPagination_TransformToLimitOffset(t *testing.T) {
	type fields struct {
		Page    uint
		PerPage uint
		Total   uint
	}
	tests := []struct {
		name   string
		fields fields
		want   LimitOffset
	}{
		{
			name: "test 1",
			fields: fields{
				Page:    1,
				PerPage: 10,
			},
			want: LimitOffset{
				Limit:  10,
				Offset: 0,
			},
		},
		{
			name: "test 2",
			fields: fields{
				Page:    5,
				PerPage: 7,
			},
			want: LimitOffset{
				Limit:  7,
				Offset: 28,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pagination{
				Page:    tt.fields.Page,
				PerPage: tt.fields.PerPage,
				Total:   tt.fields.Total,
			}
			if got := p.TransformToLimitOffset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pagination.TransformToLimitOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
