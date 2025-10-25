package merkel

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

type block struct {
	ID string
}

func (b *block) String() string {
	return b.ID
}

func (b *block) Hash() []byte {
	h := sha256.New()
	h.Write([]byte(b.ID))
	return h.Sum(nil)
}

func TestMerkel(t *testing.T) {

	items := []Raw{
		&block{
			ID: "test",
		},
		&block{
			ID: "test2",
		},
		&block{
			ID: "test3",
		},
	}
	tree := New(items)
	fmt.Println(tree.Root.HashString())
	fmt.Println(tree.String())

	items2 := []Raw{
		&block{
			ID: "rebuild1",
		},
		&block{
			ID: "rebuild2",
		},
		&block{
			ID: "rebuild3",
		},
	}

	tree.Rebuild(items2)
	fmt.Println(tree.Root.HashString())
	fmt.Println(tree.String())
}

func TestDiff(t *testing.T) {
	items1 := []Raw{
		&block{"Hello"},
		&block{"world"},
		&block{"uiui"},
	}
	items2 := []Raw{
		&block{"Hello"},
		&block{"world"},
		&block{"uiui"},
	}
	left, right := New(items1), New(items2)
	fmt.Println("left hash ", left.Root.HashString())
	fmt.Println("right hash ", right.Root.HashString())
	fmt.Println(FindDifference(left, right))
}
