#Introduction
This ac-automathon is designed for GoSpell project & writen in go. **Notice** Since it's designed for GoSpell, the behaviour of this automathon is not same as general one :

1. This automathon will return the longest match.
2. This automathon will ignore the match if current match is partly overlap with previous ones.

*Example* : Pattern: "he", "her", "erk". String: "herk". The automathon will only find "her" in the string, because "he" is not the longest and "erk" is partly overlapped. **To change this behaviour** Please change the Search function in source code. 

The basic struct:

```go
type Matcher struct {
     root,curNode *node
     count int
}

type Result struct {
     StrIndex, DictIndex int
}

```
#NewMatcher

```go
func NewMatcher() *Matcher
```

This function will return a new matcher. The Matcher should be build before use.

##Build Matcher
```go
func (this *Matcher) Build(dictionary []string)
```

This function will init the matcher with your dictionary strings (patterns).

##Search
```go
func (this *Matcher) Search(ch string) []Result
```
This function will accept a string and return a slice of result. .DictIndex is the pattern index in dictionary. .StrIndex is the start position of that pattern in the string.

##Example
```go
matcher := NewMatcher()
dictionary := []string{"say","she","shr","he","her"}
matcher.Build(dictionary)
result := matcher.Search("yasherhs")
fmt.Println(result)
```

Output should be : [{2, 1}] 2 is StrIndex, means the start position in "yasherhs". And 1 is DictIndex, means the pattern is dictionary[1]. 