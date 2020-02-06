package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyTree(t *testing.T) {
	t.Run("Empty binary tree", emptyBinaryTree)
}

//empty tree test
func emptyBinaryTree(t *testing.T) {
	tree := BinaryTree{}
	result := PreOrderTraversal(tree.root)
	var expected []int64
	assert.Equal(t, expected, result)

}

func TestOneNodeTree(t *testing.T) {
	t.Run("One node tree", ondeNodeTree)
}

func ondeNodeTree(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(1)
	result := PreOrderTraversal(tree.root)
	var expected []int64
	expected = append(expected, 1)
	assert.Equal(t, expected, result)
}

func TestOnlyLeftNodeTree(t *testing.T) {
	t.Run("Only left node tree", onlyLeftNodeTree)
}

func onlyLeftNodeTree(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(8)
	tree.Insert(7)
	result := PreOrderTraversal(tree.root)
	var expected []int64
	expected = append(expected, 10, 9, 8, 7)
	assert.Equal(t, expected, result)
}

func TestOnlyRightNodeTree(t *testing.T) {
	t.Run("Only left node tree", onlyRightNodeTree)
}

func onlyRightNodeTree(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(13)
	result := PreOrderTraversal(tree.root)
	var expected []int64
	expected = append(expected, 10, 11, 12, 13)
	assert.Equal(t, expected, result)
}

func TestTree(t *testing.T) {
	t.Run("Only left node tree", tree)
}

func tree(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(11)
	tree.Insert(8)
	tree.Insert(12)
	tree.Insert(13)
	result := PreOrderTraversal(tree.root)
	var expected []int64
	expected = append(expected, 10, 9, 8, 11, 12, 13)
	assert.Equal(t, expected, result)
}

func TestDirectNodeInsert(t *testing.T) {
	t.Run("Try direct node insert", directNodeInsert)
}

func directNodeInsert(t *testing.T) {
	tree := BinaryTree{}
	tree.root.Insert(10)
	result := PreOrderTraversal(tree.root)
	var expected []int64
	assert.Equal(t, expected, result)
}
