package ahocorasick

import "github.com/tonyshaw/queue"

const minilmt = 3
const fulllmt = 128

type node struct {
	childch []rune
	minichild, fullchild []*node
	fail *node
	prefix *string
	index int
	fullmode bool
}

func newNode() *node {
	return &node{
		childch : make([]rune,minilmt,minilmt),
		minichild : make([]*node,minilmt,minilmt),
		fullchild : nil,
		prefix : nil,
		fail : nil,
		index : -1,
		fullmode : false,
	}
}

type Matcher struct {
	root,curNode *node
	count int
}

func NewMatcher() *Matcher {
	return &Matcher {
		root : newNode(),
		curNode : nil,
		count : 1,
	}
}

func (this *Matcher) Search(ch rune) int {
	for ;; {
		if this.curNode.fullmode {
			if this.curNode.fullchild[int(ch)] != nil {
				this.curNode = this.curNode.fullchild[int(ch)]
				return this.curNode.index
			}
		} else {
			for i, tmp := range this.curNode.childch {
				if tmp == ch {
					this.curNode = this.curNode.minichild[i]
					return this.curNode.index
				}
			}
		}
		if this.curNode == this.root {
			return this.curNode.index
		}
		this.curNode = this.curNode.fail
	}
}		

func (this *Matcher) Build(dictionary []string) {
	//set init state to root
	this.curNode = this.root
	//first build the trie
	for i, str := range dictionary {
		curNode := this.root
		for _, ch := range str {
			curNode = curNode.insert(ch, &this.count)
		}
		curNode.index = i
	}
	//then build the automathon
	this.build()
	//finally compress the branches and merge the nodes
	//this.compress()
}

func (this *Matcher) build() {
	q := queue.NewQueue(this.count)
	this.root.fail = this.root
	q.Push(this.root); 
	var curNode *node
	var x interface{}
	for ;q.Length()!=0; {
		x, _ = q.Pop()
		curNode = x.(*node)
		if !curNode.fullmode {
			for i,ch := range curNode.childch {
				if int(ch) != 0 {
					buildFailPointer(curNode.fail, curNode.minichild[i], int(ch))
					q.Push(curNode.minichild[i])
				}
			}
		} else {
			for i, child := range curNode.fullchild {
				if child != nil {
					buildFailPointer(curNode.fail, curNode.fullchild[i], i)
					q.Push(curNode.fullchild[i])
				}
			}
		}
	}			
}

func buildFailPointer(fail *node, curNode *node, ch int) {
	for ;;fail = fail.fail {
		if fail.fullmode {
			if fail.fullchild[ch] != nil && fail.fullchild[ch] != curNode{
				curNode.fail = fail.fullchild[ch]
				return
			}
		} else {
			for i,tmp := range fail.childch {
				if tmp == rune(ch) && fail.minichild[i] != curNode{
					curNode.fail = fail.minichild[i]
					return
				}
			}
		}
		if fail == fail.fail {//fail is root & we can't find a match in root. point this.fail to root
			curNode.fail = fail
			return
		}
	}
}

func (this *node) insert(ch rune, count *int) *node{
	if !this.fullmode {
		for i, tmp := range this.childch {
			if tmp == ch {
				return this.minichild[i]
			}
			if tmp == rune(0) {
				this.childch[i] = ch
				this.minichild[i] = newNode()
				*count++
				return this.minichild[i]
			}
		}
		//can't insert the char due to minilmt slots is full. copy all into next mode
		this.fullmode = true //change the flag
		this.fullchild = make([]*node,fulllmt,fulllmt)
		for i, tmp := range this.childch {
			this.fullchild[int(tmp)] = this.minichild[i]
		}
		this.minichild = nil
		this.childch = nil
	}
	if this.fullchild[int(ch)] != nil {
		return this.fullchild[int(ch)]
	}
	this.fullchild[int(ch)] = newNode()
	*count++
	return this.fullchild[int(ch)]
}
			
