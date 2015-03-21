package ahocorasick

import "testing"
import "fmt"

func printNode(n *node) {
	fmt.Println("Is FullMode? ", n.fullmode)
	fmt.Print("Childs: ")
	for _,ch := range n.childch {
		fmt.Printf("%c ",ch)
	}
	fmt.Println()
}

func PrintNode(n *node) {
	fmt.Println("-------Node:  ")
	printNode(n)
	fmt.Println("-Fail :")
	printNode(n.fail)
}	


func TestBuild(t *testing.T) {
	matcher := NewMatcher()
	dictionary := []string{"say","she","shr","he","her"}
	matcher.Build(dictionary)
	//check all nodes status
	if matcher.count != 10 {
		t.Errorf("Tree node number isn't correct")
	}
	//Below part is sesigned to check tree shape manullaycheck the shape of tree
	/*fmt.Println(matcher.count)
	PrintNode(matcher.root)
	PrintNode(matcher.root.minichild[0])
	PrintNode(matcher.root.minichild[1])
	PrintNode(matcher.root.minichild[0].minichild[0])
	PrintNode(matcher.root.minichild[0].minichild[1])
	PrintNode(matcher.root.minichild[1].minichild[0])
	PrintNode(matcher.root.minichild[0].minichild[1].minichild[0])
	PrintNode(matcher.root.minichild[0].minichild[1].minichild[1])*/
	teststr := "yasherhs"
	for i,ch := range teststr {
		index := matcher.Search(ch)
		if i == 4  {
			if index != 1 {
				t.Errorf("Match 'he' position is not right")
			}
			continue
		}
		if i == 5 {
			if index != 4 {
				t.Errorf("Match 'her' position is not right")
			}
			continue
		}
		if index != -1 {
			t.Errorf("Match is not right in : ", i)
		}
	}
}

