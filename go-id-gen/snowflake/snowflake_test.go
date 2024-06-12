package snowflake

import (
	"reflect"
	"testing"
	"time"
)

func BenchmarkSnowflake_Gen_FixedTime(b *testing.B) {
	s, _ := New(nil)
	t := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Gen(&t)
	}
}

func BenchmarkSnowflake_Gen(b *testing.B) {
	s, _ := New(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Gen(nil)
	}
}

func Test_snowflake_CompareNewer(t *testing.T) {
	type fields struct {
		now int64
		seq int16
	}
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name:   "newer a (timestamp)",
			fields: fields{},
			args: args{
				a: 7206460897189949499,
				b: 7206460897185755137,
			},
			want: 7206460897189949499,
		},
		{
			name:   "newer b (timestamp)",
			fields: fields{},
			args: args{
				a: 7206460897177366540,
				b: 7206460897189949442,
			},
			want: 7206460897189949442,
		},
		{
			name:   "newer a (seq no)",
			fields: fields{},
			args: args{
				a: 7206460897189949498,
				b: 7206460897189949495,
			},
			want: 7206460897189949498,
		},
		{
			name:   "newer b (seq no)",
			fields: fields{},
			args: args{
				a: 7206460897189949499,
				b: 7206460897189949508,
			},
			want: 7206460897189949508,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &snowflake{
				now: tt.fields.now,
				seq: tt.fields.seq,
			}
			if got := s.CompareNewer(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("snowflake.CompareNewer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		nodeId *int64
	}
	tests := []struct {
		name    string
		args    args
		want    *snowflake
		wantErr bool
	}{
		{
			name:    "nodeid lower bound exception",
			args:    args{nodeId: getInt64Ptr(int64(-1))},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "nodeid upper bound exception",
			args:    args{nodeId: getInt64Ptr(int64(1024))},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.nodeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_snowflake_Gen(t *testing.T) {
	type fields struct {
		now    int64
		seq    int16
		nodeId *int64
	}
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *int64
	}{
		{
			name: "gen specific time",
			fields: fields{
				now:    time.Unix(1560000000, 0).UnixMilli() << 22,
				seq:    1,
				nodeId: getInt64Ptr(int64(0)),
			},
			args: args{
				t: getTimePtr(time.Unix(1560000000, 0)),
			},
			want: getInt64Ptr(int64(4190209)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &snowflake{
				now:    tt.fields.now,
				seq:    tt.fields.seq,
				nodeId: tt.fields.nodeId,
			}
			if got := s.Gen(tt.args.t); got != tt.want {
				t.Errorf("snowflake.Gen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getInt64Ptr(i int64) *int64 {
	return &i
}

func getTimePtr(t time.Time) *time.Time {
	return &t
}
