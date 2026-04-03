package main

const LETTERS_COUNT = 'z' - 'a' + 1

type MapSumNode struct {
	value      int
	letterPtrs [LETTERS_COUNT]*MapSumNode
}

type MapSum struct {
	mapSumNode *MapSumNode
}

func Constructor() MapSum {
	return MapSum{new(MapSumNode)}
}

func (this *MapSum) Insert(key string, val int) {
	node := this.mapSumNode
	for _, c := range key {
		if node.letterPtrs[c-'a'] == nil {
			node.letterPtrs[c-'a'] = new(MapSumNode)
		}
		node = node.letterPtrs[c-'a']
	}
	node.value = val
}

func sumAll(node *MapSumNode) int {
	if node == nil {
		return 0
	}

	sum := node.value
	for _, letterNode := range node.letterPtrs {
		sum += sumAll(letterNode)
	}

	return sum
}

func (this *MapSum) Sum(prefix string) int {
	node := this.mapSumNode
	for _, c := range prefix {
		if node.letterPtrs[c-'a'] == nil {
			return 0
		}
		node = node.letterPtrs[c-'a']
	}

	return sumAll(node)
}
