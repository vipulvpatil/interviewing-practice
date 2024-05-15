package nary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreSameTrue(t *testing.T) {
	tree1 := Node{
		Value: "A",
		Children: []*Node{
			{
				Value: "B",
			},
			{
				Value: "C",
			},
			{
				Value: "D",
			},
			{
				Value: "E",
			},
		},
	}

	tree2 := Node{
		Value: "A",
		Children: []*Node{
			{
				Value: "B",
				Children: []*Node{
					{
						Value: "C",
						Children: []*Node{
							{
								Value: "D",
							},
						},
					},
				},
			},
			{
				Value: "E",
			},
		},
	}

	tree3 := Node{
		Value: "A",
		Children: []*Node{
			{
				Value: "B",
				Children: []*Node{
					{
						Value: "C",
						Children: []*Node{
							{
								Value: "D",
							},
							{
								Value: "",
							},
							{
								Value: "",
							},
						},
					},
				},
			},
			{
				Value: "",
			},
			{
				Value: "E",
			},
			{
				Value: "",
			},
		},
	}

	t.Run("match successes", func(t *testing.T) {
		result := AreSame(&tree1, &tree2)
		assert.True(t, result, "represent same string in tree1 and tree2")
		result = AreSame(&tree1, &tree3)
		assert.True(t, result, "represent same string in tree1 and tree3")
	})
}

func TestAreSameFalse(t *testing.T) {
	tree1 := Node{
		Value: "A",
		Children: []*Node{
			{
				Value: "B",
			},
			{
				Value: "C",
			},
			{
				Value: "D",
			},
			{
				Value: "E",
			},
		},
	}

	tree2 := Node{
		Value: "A",
		Children: []*Node{
			{
				Value: "B",
				Children: []*Node{
					{
						Value: "D",
						Children: []*Node{
							{
								Value: "C",
							},
						},
					},
				},
			},
			{
				Value: "E",
			},
		},
	}

	tree3 := Node{
		Value: "A",
		Children: []*Node{
			{
				Value: "B",
				Children: []*Node{
					{
						Value: "C",
					},
				},
			},
			{
				Value: "",
			},
			{
				Value: "E",
			},
			{
				Value: "",
			},
		},
	}

	t.Run("match successes", func(t *testing.T) {
		result := AreSame(&tree1, &tree2)
		assert.False(t, result, "represent diff string in tree1 and tree2")
		result = AreSame(&tree1, &tree3)
		assert.False(t, result, "represent diff string in tree1 and tree3")
	})
}
