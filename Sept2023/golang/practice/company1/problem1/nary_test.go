package nary

import "testing"

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
			Value: "C",
		},
		{
			Value: "D",
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
				},
				Children: []*Node{
					{
						Value: "D",
					},
				},
			},
		},
	},
}

func testMatch(t *testing.T) {
	t.Run("match successes", func(t *testing.T){
		ch1 := make(chan string)
		ch2 := make(chan string)

		PreOrderNode(tree1, ch1)
		PreOrderNode(tree2, ch2)

		val1, ok1 := <-ch1
		val2, ok2 := <-ch2
		for ok1 && ok2 {
			if val1 != val2 {
				fmt.Println("match not found")
			}
			val1, ok1 = <-ch1
			val2, ok2 = <-ch2
		}
		if ok1 {
			fmt.Println("match not found")
		}
		if ok2 {
			fmt.Println("match not found")
		}
		fmt.Println("match found")
	})
}
