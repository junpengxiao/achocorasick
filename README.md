#Introduction
This ac-automathon is designed for GoSpell project & writen in go. The basic struct:

```go
type Matcher struct {
        root,curNode *node
	count int
}
```
**Notice** the `Search` function below will only return the longest matches and won't return overlap matchers. For example if "she,he" exists, and the string bshe will only return she's index. If "she,heh" exists, and the string is shehee, then she will be return, curtent node pointer will point to root immediately. Thus the left string is "hee" which doesn't exist heh. 

##New Matcer

```go
func NewMatcher() *Matcher
```

This function will return a new matcher. The Matcher should be build before use.

##Build Matcher
```go
func (this *Matcher) Build(dictionary []string)
```

This function will init the matcher with your dictionary string, i.e., pattern string.

##Search
```go
func (this *Matcher) Search(ch char) int
```

This function will accept the next chat of current string and return the index of matches. If there is a match, then the index is same as that pattern string index in the dictionary. Otherwise the return value is -1.


