package search

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
	runeArrays := t.root.keys([]rune{})
	str := []string{}
	for _, runeArr := range runeArrays {
		str = append(str, string(runeArr))
	}
	return str
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
