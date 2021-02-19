package main

import "fmt"

/*
{
	from:"https://leetcode-cn.com/problems/implement-trie-prefix-tree",
	reference:[],
	language:"Go",
	description:"实现前缀树",
	solved:true,
}
*/

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
type Trie struct {
	root Node
}

type Node struct {
	existsMatch bool
	child       map[rune]Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: Node{
			false,
			make(map[rune]Node),
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var now = this.root
	var l = len(word)
	for i, v := range word {
		if _, exi := now.child[v]; !exi {
			var exiMatch bool
			if i+1 < l {
				exiMatch = false
			} else {
				exiMatch = true
			}
			now.child[v] = Node{exiMatch, make(map[rune]Node)}
			now = now.child[v]
		} else {
			if i+1 == l {
				now.child[v] = Node{true, now.child[v].child}
			}
			now = now.child[v]
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var now = this.root
	var l = len(word)
	for i, v := range word {
		if _, exi := now.child[v]; exi {
			if i+1 == l && now.child[v].existsMatch {
				return true
			}
			now = now.child[v]
		} else {
			return false
		}
	}
	return false
}

///** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var now = this.root
	for _, v := range prefix {
		if _, exi := now.child[v]; exi {
			now = now.child[v]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func Print(values ...bool) {
	for _, v := range values {
		fmt.Println(v)
	}
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	Print(
		trie.Search("apple"),
		trie.Search("app"),
		trie.StartsWith("app"))
	trie.Insert("app")
	Print(trie.Search("app"))
}
