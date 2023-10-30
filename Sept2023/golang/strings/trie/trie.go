package trie

type Node struct {
	children map[rune]*Node
	value    string
	hasValue bool
}

func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		value:    "",
		hasValue: false,
	}
}

type Trie struct {
	root *Node
}

func (t *Trie) Put(key, value string) {
	keyArr := []rune(key)
	t.root = t.root.put(keyArr, value)
}

func (t *Trie) Get(key string) string {
	keyArr := []rune(key)
	node := t.root.get(keyArr)
	if node != nil {
		return node.value
	}
	return ""
}

func (t *Trie) Delete(key string) {
	keyArr := []rune(key)
	t.root = t.root.delete(keyArr)
}

func (t *Trie) Keys() []string {
	return t.root.keys([]rune{})
}

func (t *Trie) KeysWithPrefix(key string) []string {
	keyArr := []rune(key)
	prefixNode := t.root.get(keyArr)
	if prefixNode != nil {
		return prefixNode.keys(keyArr)
	}
	return nil
}

// key may contain '.' that matches everything.
func (t *Trie) KeysThatMatch(key string) []string {
	keyArr := []rune(key)
	return t.root.keysThatMatch([]rune{}, keyArr)
}

func (t *Trie) LongestPrefixOf(key string) string {
	keyArr := []rune(key)
	node := t.root.longestPrefixOf(keyArr)
	if node != nil {
		return node.value
	}
	return ""
}

func (n *Node) put(key []rune, value string) *Node {
	if n == nil {
		n = NewNode()
	}
	if len(key) == 0 {
		n.value = value
		n.hasValue = true
	} else {
		n.children[key[0]] = n.children[key[0]].put(key[1:], value)
	}
	return n
}

func (n *Node) get(key []rune) *Node {
	if len(key) == 0 {
		if n.hasValue {
			return n
		}
		return nil
	}
	nextNode, ok := n.children[key[0]]
	if !ok {
		return nil
	}
	return nextNode.get(key[1:])
}

func (n *Node) delete(key []rune) *Node {
	if n == nil {
		return nil
	}
	if len(key) == 0 {
		n.value = ""
		n.hasValue = false
	} else {
		n.children[key[0]] = n.children[key[0]].delete(key[1:])
	}
	return n
}

func (n *Node) keys(currentPrefix []rune) []string {
	if n == nil {
		return nil
	}
	keysSoFar := []string{}
	if n.hasValue {
		keysSoFar = append(keysSoFar, string(currentPrefix))
	}
	for k, child := range n.children {
		newPrefix := append(currentPrefix, k)
		adding := child.keys(newPrefix)
		keysSoFar = append(keysSoFar, adding...)
	}

	return keysSoFar
}

func (n *Node) keysThatMatch(currentPrefix []rune, key []rune) []string {
	if len(key) == 0 {
		if n.hasValue {
			return []string{string(currentPrefix)}
		}
		return nil
	}
	keysSoFar := []string{}
	if key[0] == '.' {
		for k, child := range n.children {
			newPrefix := append(currentPrefix, k)
			adding := child.keysThatMatch(newPrefix, key[1:])
			keysSoFar = append(keysSoFar, adding...)
		}
	} else {
		nextNode, ok := n.children[key[0]]
		if !ok {
			return nil
		}
		newPrefix := append(currentPrefix, key[0])
		adding := nextNode.keysThatMatch(newPrefix, key[1:])
		keysSoFar = append(keysSoFar, adding...)
	}
	return keysSoFar
}

func (n *Node) longestPrefixOf(key []rune) *Node {
	var currentLongestPrefix *Node
	if n.hasValue {
		currentLongestPrefix = n
	}

	if len(key) != 0 {
		nextNode, ok := n.children[key[0]]
		if !ok {
			return currentLongestPrefix
		}
		nextLongestPrefix := nextNode.longestPrefixOf(key[1:])
		if nextLongestPrefix != nil {
			return nextLongestPrefix
		}
	}

	return currentLongestPrefix
}
