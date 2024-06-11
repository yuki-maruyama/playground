package snowflake

import (
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
				a: 7206211586589573138,
				b: 7206211586585567293,
			},
			want: 7206211586589573138,
		},
		{
			name:   "newer b (timestamp)",
			fields: fields{},
			args: args{
				a: 7206211586589573138,
				b: 7206218025523597399,
			},
			want: 7206218025523597399,
		},
		{
			name:   "newer a (seq no)",
			fields: fields{},
			args: args{
				a: 7206219147414360159,
				b: 7206219147415093342,
			},
			want: 7206219147414360159,
		},
		{
			name:   "newer b (seq no)",
			fields: fields{},
			args: args{
				a: 7206219147414360159,
				b: 7206219147416113251,
			},
			want: 7206219147416113251,
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
