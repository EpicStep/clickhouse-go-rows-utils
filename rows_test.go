package clickhouse_go_rows_utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/EpicStep/clickhouse-go-rows-utils/mock"
)

func TestCollectRows(t *testing.T) {
	t.Parallel()

	type args struct {
		prepareRowsMock func(rows *mock.MockRows)
		fn              RowToFunc[string]
	}
	type testCase struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}
	tests := []testCase{
		{
			name: "Ok",
			args: args{
				prepareRowsMock: func(rows *mock.MockRows) {
					rows.EXPECT().Next().Return(true)
					rows.EXPECT().Next().Return(false)

					rows.EXPECT().Scan(gomock.Any()).Return(nil)
					rows.EXPECT().Err().Return(nil)
					rows.EXPECT().Close().Return(nil)
				},
				fn: func(row CollectableRow) (string, error) {
					var res string
					err := row.Scan(&res)
					return res, err
				},
			},
			wantLen: 1,
		},
		{
			name: "Err at func",
			args: args{
				prepareRowsMock: func(rows *mock.MockRows) {
					rows.EXPECT().Next().Return(true)
					rows.EXPECT().Close().Return(nil)
				},
				fn: func(_ CollectableRow) (string, error) {
					return "", errors.New("test")
				},
			},
			wantErr: true,
		},
		{
			name: "Err at rows",
			args: args{
				prepareRowsMock: func(rows *mock.MockRows) {
					rows.EXPECT().Next().Return(true)
					rows.EXPECT().Next().Return(false)
					rows.EXPECT().Err().Return(errors.New("test"))
					rows.EXPECT().Close().Return(nil)
				},
				fn: func(_ CollectableRow) (string, error) {
					return "", nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			rowsMock := mock.NewMockRows(ctrl)

			if tt.args.prepareRowsMock != nil {
				tt.args.prepareRowsMock(rowsMock)
			}

			got, err := CollectRows(rowsMock, tt.args.fn)
			assert.Len(t, got, tt.wantLen)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAppendRows(t *testing.T) {
	t.Parallel()

	type args struct {
		prepareRowsMock func(rows *mock.MockRows)
		fn              RowToFunc[string]
		slice           []string
	}
	type testCase struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}
	tests := []testCase{
		{
			name: "Ok",
			args: args{
				slice: []string{"test-1"},
				prepareRowsMock: func(rows *mock.MockRows) {
					rows.EXPECT().Next().Return(true)
					rows.EXPECT().Next().Return(false)

					rows.EXPECT().Scan(gomock.Any()).Return(nil)
					rows.EXPECT().Err().Return(nil)
					rows.EXPECT().Close().Return(nil)
				},
				fn: func(row CollectableRow) (string, error) {
					var res string
					err := row.Scan(&res)
					return res, err
				},
			},
			wantLen: 2,
		},
		{
			name: "Err at func",
			args: args{
				slice: []string{"test-1"},
				prepareRowsMock: func(rows *mock.MockRows) {
					rows.EXPECT().Next().Return(true)
					rows.EXPECT().Close().Return(nil)
				},
				fn: func(_ CollectableRow) (string, error) {
					return "", errors.New("test")
				},
			},
			wantErr: true,
		},
		{
			name: "Err at rows",
			args: args{
				slice: []string{"test-1"},
				prepareRowsMock: func(rows *mock.MockRows) {
					rows.EXPECT().Next().Return(true)
					rows.EXPECT().Next().Return(false)
					rows.EXPECT().Err().Return(errors.New("test"))
					rows.EXPECT().Close().Return(nil)
				},
				fn: func(_ CollectableRow) (string, error) {
					return "", nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			rowsMock := mock.NewMockRows(ctrl)

			if tt.args.prepareRowsMock != nil {
				tt.args.prepareRowsMock(rowsMock)
			}

			got, err := AppendRows(tt.args.slice, rowsMock, tt.args.fn)
			assert.Len(t, got, tt.wantLen)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestForEachRow(t *testing.T) {
	t.Parallel()

	type args struct {
		prepareRowsMock func(rows *mock.MockRows)
		fn              ForEachRowFunc
	}
	type testCase struct {
		args    args
		name    string
		wantLen int
		wantErr bool
	}
	tests := []testCase{
		{
			name: "Ok",
			args: args{
				prepareRowsMock: func(rows *mock.MockRows) {
					rows.EXPECT().Next().Return(true)
					rows.EXPECT().Next().Return(false)

					rows.EXPECT().Scan(gomock.Any()).Return(nil)
					rows.EXPECT().Err().Return(nil)
					rows.EXPECT().Close().Return(nil)
				},
				fn: func(row CollectableRow) error {
					var res string
					return row.Scan(&res)
				},
			},
			wantLen: 2,
		},
		{
			name: "Error",
			args: args{
				prepareRowsMock: func(rows *mock.MockRows) {
					rows.EXPECT().Next().Return(true)
					rows.EXPECT().Close().Return(nil)
				},
				fn: func(_ CollectableRow) error {
					return errors.New("test")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			rowsMock := mock.NewMockRows(ctrl)

			if tt.args.prepareRowsMock != nil {
				tt.args.prepareRowsMock(rowsMock)
			}

			assert.Equal(
				t,
				tt.wantErr,
				ForEachRow(rowsMock, tt.args.fn) != nil,
			)
		})
	}
}
