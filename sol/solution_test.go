package sol

import (
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	arr := *input
	result := make([]*TreeNode, len(arr))
	for idx, val := range arr {
		if val != "null" {
			num, _ := strconv.Atoi(val)
			result[idx] = &TreeNode{Val: num}
		}
		if idx == 0 {
			tree = result[idx]
		}
	}
	for idx := range result {
		if result[idx] != nil {
			if 2*idx+1 < len(result) {
				result[idx].Left = result[2*idx+1]
			}
			if 2*idx+2 < len(result) {
				result[idx].Right = result[2*idx+2]
			}
		}
	}
	return tree
}

func BenchmarkMaxPathSum(b *testing.B) {
	root := CreateBinaryTree(&[]string{"-10", "9", "20", "null", "null", "15", "7"})
	for idx := 0; idx < b.N; idx++ {
		maxPathSum(root)
	}
}
func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [1,2,3]",
			args: args{CreateBinaryTree(&[]string{"1", "2", "3"})},
			want: 6,
		},
		{
			name: "root = [-10,9,20,null,null,15,7]",
			args: args{CreateBinaryTree(&[]string{"-10", "9", "20", "null", "null", "15", "7"})},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
