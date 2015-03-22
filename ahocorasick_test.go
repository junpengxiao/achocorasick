package ahocorasick

import "testing"
import "fmt"

func printNode(n *node) {
	fmt.Println("Is FullMode? ", n.fullmode)
	fmt.Print("Childs: ")
	for _,ch := range n.childch {
		fmt.Printf("%c ",ch)
	}
	fmt.Println("Index ", n.index)
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
	

	result := matcher.Search("yasherhs")

	if len(result) != 1 {
		t.Errorf("1st : The length of result is not correct", len(result))
	}
	if result[0].StrIndex != 2 || result[0].DictIndex != 1 {
		t.Errorf("1st : The index of result is not correct", result[0])
	}

	result = matcher.Search("aahersfdf")
	if len(result) != 1 {
		t.Errorf("2nd : The length of result is not correct", len(result))
	}
	if result[0].StrIndex != 2 || result[0].DictIndex != 4 {
		t.Errorf("2nd : The index of result is not correct", result[0])
	}
}

