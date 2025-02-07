package data

import "testing"

func Test_getContentDetailsTable(t *testing.T) {
	type args struct {
		contentID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "正常",
			args: args{
				contentID: "1",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getContentDetailsTable(tt.args.contentID); got != tt.want {
				t.Errorf("getContentDetailsTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
