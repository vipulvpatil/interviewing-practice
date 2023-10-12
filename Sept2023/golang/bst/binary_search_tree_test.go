package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	t.Run("insert", func(t *testing.T) {
		var b = constructRandomBST()
		expected := constructExpectedBST()
		actual := b.InOrder()
		assert.Equal(t, len(expected), len(actual))
		for i := range expected {
			assert.Equal(t, expected[i].key, actual[i].key)
			assert.Equal(t, expected[i].value, actual[i].value)
		}
	})
	t.Run("search", func(t *testing.T) {})
	t.Run("min", func(t *testing.T) {})
	t.Run("max", func(t *testing.T) {})
	t.Run("floor", func(t *testing.T) {})
	t.Run("ceil", func(t *testing.T) {})
	t.Run("inorder", func(t *testing.T) {})
}

func constructExpectedBST() []*TreeNode[string, string] {
	return []*TreeNode[string, string]{
		{key: "A", value: "A"},
		{key: "AA", value: "AA"},
		{key: "AAA", value: "AAA"},
		{key: "AAAA", value: "AAAA"},
		{key: "AAAAAA", value: "AAAAAA"},
		{key: "AAAAAAAA", value: "AAAAAAAA"},
		{key: "AAAAAAAAAAA", value: "AAAAAAAAAAA"},
		{key: "B", value: "B"},
		{key: "BB", value: "BB"},
		{key: "BBB", value: "BBB"},
		{key: "BBBB", value: "BBBB"},
		{key: "BBBBBB", value: "BBBBBB"},
		{key: "BBBBBBBB", value: "BBBBBBBB"},
		{key: "BBBBBBBBBBB", value: "BBBBBBBBBBB"},
		{key: "C", value: "C"},
		{key: "CC", value: "CC"},
		{key: "CCC", value: "CCC"},
		{key: "CCCC", value: "CCCC"},
		{key: "CCCCCC", value: "CCCCCC"},
		{key: "CCCCCCCC", value: "CCCCCCCC"},
		{key: "CCCCCCCCCCC", value: "CCCCCCCCCCC"},
		{key: "D", value: "D"},
		{key: "DD", value: "DD"},
		{key: "DDD", value: "DDD"},
		{key: "DDDD", value: "DDDD"},
		{key: "DDDDDD", value: "DDDDDD"},
		{key: "DDDDDDDD", value: "DDDDDDDD"},
		{key: "DDDDDDDDDDD", value: "DDDDDDDDDDD"},
		{key: "E", value: "E"},
		{key: "EE", value: "EE"},
		{key: "EEE", value: "EEE"},
		{key: "EEEE", value: "EEEE"},
		{key: "EEEEEE", value: "EEEEEE"},
		{key: "EEEEEEEE", value: "EEEEEEEE"},
		{key: "EEEEEEEEEEE", value: "EEEEEEEEEEE"},
		{key: "F", value: "F"},
		{key: "FF", value: "FF"},
		{key: "FFF", value: "FFF"},
		{key: "FFFF", value: "FFFF"},
		{key: "FFFFFF", value: "FFFFFF"},
		{key: "FFFFFFFF", value: "FFFFFFFF"},
		{key: "FFFFFFFFFFF", value: "FFFFFFFFFFF"},
		{key: "G", value: "G"},
		{key: "GG", value: "GG"},
		{key: "GGG", value: "GGG"},
		{key: "GGGG", value: "GGGG"},
		{key: "GGGGGG", value: "GGGGGG"},
		{key: "GGGGGGGG", value: "GGGGGGGG"},
		{key: "GGGGGGGGGGG", value: "GGGGGGGGGGG"},
		{key: "H", value: "H"},
		{key: "HH", value: "HH"},
		{key: "HHH", value: "HHH"},
		{key: "HHHH", value: "HHHH"},
		{key: "HHHHHH", value: "HHHHHH"},
		{key: "HHHHHHHH", value: "HHHHHHHH"},
		{key: "HHHHHHHHHHH", value: "HHHHHHHHHHH"},
		{key: "I", value: "I"},
		{key: "II", value: "II"},
		{key: "III", value: "III"},
		{key: "IIII", value: "IIII"},
		{key: "IIIIII", value: "IIIIII"},
		{key: "IIIIIIII", value: "IIIIIIII"},
		{key: "IIIIIIIIIII", value: "IIIIIIIIIII"},
		{key: "J", value: "J"},
		{key: "JJ", value: "JJ"},
		{key: "JJJ", value: "JJJ"},
		{key: "JJJJ", value: "JJJJ"},
		{key: "JJJJJJ", value: "JJJJJJ"},
		{key: "JJJJJJJJ", value: "JJJJJJJJ"},
		{key: "JJJJJJJJJJJ", value: "JJJJJJJJJJJ"},
		{key: "K", value: "K"},
		{key: "KK", value: "KK"},
		{key: "KKK", value: "KKK"},
		{key: "KKKK", value: "KKKK"},
		{key: "KKKKKK", value: "KKKKKK"},
		{key: "KKKKKKKK", value: "KKKKKKKK"},
		{key: "KKKKKKKKKKK", value: "KKKKKKKKKKK"},
		{key: "L", value: "L"},
		{key: "LL", value: "LL"},
		{key: "LLL", value: "LLL"},
		{key: "LLLL", value: "LLLL"},
		{key: "LLLLLL", value: "LLLLLL"},
		{key: "LLLLLLLL", value: "LLLLLLLL"},
		{key: "LLLLLLLLLLL", value: "LLLLLLLLLLL"},
		{key: "M", value: "M"},
		{key: "MM", value: "MM"},
		{key: "MMM", value: "MMM"},
		{key: "MMMM", value: "MMMM"},
		{key: "MMMMMM", value: "MMMMMM"},
		{key: "MMMMMMMM", value: "MMMMMMMM"},
		{key: "MMMMMMMMMMM", value: "MMMMMMMMMMM"},
		{key: "N", value: "N"},
		{key: "NN", value: "NN"},
		{key: "NNN", value: "NNN"},
		{key: "NNNN", value: "NNNN"},
		{key: "NNNNNN", value: "NNNNNN"},
		{key: "NNNNNNNN", value: "NNNNNNNN"},
		{key: "NNNNNNNNNNN", value: "NNNNNNNNNNN"},
		{key: "O", value: "O"},
		{key: "OO", value: "OO"},
		{key: "OOO", value: "OOO"},
		{key: "OOOO", value: "OOOO"},
		{key: "OOOOOO", value: "OOOOOO"},
		{key: "OOOOOOOO", value: "OOOOOOOO"},
		{key: "OOOOOOOOOOO", value: "OOOOOOOOOOO"},
		{key: "P", value: "P"},
		{key: "PP", value: "PP"},
		{key: "PPP", value: "PPP"},
		{key: "PPPP", value: "PPPP"},
		{key: "PPPPPP", value: "PPPPPP"},
		{key: "PPPPPPPP", value: "PPPPPPPP"},
		{key: "PPPPPPPPPPP", value: "PPPPPPPPPPP"},
		{key: "Q", value: "Q"},
		{key: "QQ", value: "QQ"},
		{key: "QQQ", value: "QQQ"},
		{key: "QQQQ", value: "QQQQ"},
		{key: "QQQQQQ", value: "QQQQQQ"},
		{key: "QQQQQQQQ", value: "QQQQQQQQ"},
		{key: "QQQQQQQQQQQ", value: "QQQQQQQQQQQ"},
		{key: "R", value: "R"},
		{key: "RR", value: "RR"},
		{key: "RRR", value: "RRR"},
		{key: "RRRR", value: "RRRR"},
		{key: "RRRRRR", value: "RRRRRR"},
		{key: "RRRRRRRR", value: "RRRRRRRR"},
		{key: "RRRRRRRRRRR", value: "RRRRRRRRRRR"},
		{key: "S", value: "S"},
		{key: "SS", value: "SS"},
		{key: "SSS", value: "SSS"},
		{key: "SSSS", value: "SSSS"},
		{key: "SSSSSS", value: "SSSSSS"},
		{key: "SSSSSSSS", value: "SSSSSSSS"},
		{key: "SSSSSSSSSSS", value: "SSSSSSSSSSS"},
		{key: "T", value: "T"},
		{key: "TT", value: "TT"},
		{key: "TTT", value: "TTT"},
		{key: "TTTT", value: "TTTT"},
		{key: "TTTTTT", value: "TTTTTT"},
		{key: "TTTTTTTT", value: "TTTTTTTT"},
		{key: "TTTTTTTTTTT", value: "TTTTTTTTTTT"},
		{key: "U", value: "U"},
		{key: "UU", value: "UU"},
		{key: "UUU", value: "UUU"},
		{key: "UUUU", value: "UUUU"},
		{key: "UUUUUU", value: "UUUUUU"},
		{key: "UUUUUUUU", value: "UUUUUUUU"},
		{key: "UUUUUUUUUUU", value: "UUUUUUUUUUU"},
		{key: "V", value: "V"},
		{key: "VV", value: "VV"},
		{key: "VVV", value: "VVV"},
		{key: "VVVV", value: "VVVV"},
		{key: "VVVVVV", value: "VVVVVV"},
		{key: "VVVVVVVV", value: "VVVVVVVV"},
		{key: "VVVVVVVVVVV", value: "VVVVVVVVVVV"},
		{key: "W", value: "W"},
		{key: "WW", value: "WW"},
		{key: "WWW", value: "WWW"},
		{key: "WWWW", value: "WWWW"},
		{key: "WWWWWW", value: "WWWWWW"},
		{key: "WWWWWWWW", value: "WWWWWWWW"},
		{key: "WWWWWWWWWWW", value: "WWWWWWWWWWW"},
		{key: "X", value: "X"},
		{key: "XX", value: "XX"},
		{key: "XXX", value: "XXX"},
		{key: "XXXX", value: "XXXX"},
		{key: "XXXXXX", value: "XXXXXX"},
		{key: "XXXXXXXX", value: "XXXXXXXX"},
		{key: "XXXXXXXXXXX", value: "XXXXXXXXXXX"},
		{key: "Y", value: "Y"},
		{key: "YY", value: "YY"},
		{key: "YYY", value: "YYY"},
		{key: "YYYY", value: "YYYY"},
		{key: "YYYYYY", value: "YYYYYY"},
		{key: "YYYYYYYY", value: "YYYYYYYY"},
		{key: "YYYYYYYYYYY", value: "YYYYYYYYYYY"},
		{key: "Z", value: "Z"},
		{key: "ZZ", value: "ZZ"},
		{key: "ZZZ", value: "ZZZ"},
		{key: "ZZZZ", value: "ZZZZ"},
		{key: "ZZZZZZ", value: "ZZZZZZ"},
		{key: "ZZZZZZZZ", value: "ZZZZZZZZ"},
		{key: "ZZZZZZZZZZZ", value: "ZZZZZZZZZZZ"},
	}
}

func constructRandomBST() *BinarySearchTree[string, string] {
	root := BinarySearchTree[string, string]{}
	root.Insert("VV", "VV")
	root.Insert("MM", "MM")
	root.Insert("UUUU", "UUUU")
	root.Insert("A", "A")
	root.Insert("PPPPPPPPPPP", "PPPPPPPPPPP")
	root.Insert("KK", "KK")
	root.Insert("PPP", "PPP")
	root.Insert("M", "M")
	root.Insert("ZZZZ", "ZZZZ")
	root.Insert("OOOOOOOO", "OOOOOOOO")
	root.Insert("WWW", "WWW")
	root.Insert("C", "C")
	root.Insert("KKKK", "KKKK")
	root.Insert("HHH", "HHH")
	root.Insert("P", "P")
	root.Insert("LLLLLLLLLLL", "LLLLLLLLLLL")
	root.Insert("LL", "LL")
	root.Insert("HHHH", "HHHH")
	root.Insert("WW", "WW")
	root.Insert("BBBB", "BBBB")
	root.Insert("OO", "OO")
	root.Insert("III", "III")
	root.Insert("IIIIIIIIIII", "IIIIIIIIIII")
	root.Insert("TTTT", "TTTT")
	root.Insert("KKK", "KKK")
	root.Insert("I", "I")
	root.Insert("FFFFFF", "FFFFFF")
	root.Insert("B", "B")
	root.Insert("DDD", "DDD")
	root.Insert("VVVV", "VVVV")
	root.Insert("GGGGGG", "GGGGGG")
	root.Insert("JJJJJJJJJJJ", "JJJJJJJJJJJ")
	root.Insert("E", "E")
	root.Insert("EEEEEEEEEEE", "EEEEEEEEEEE")
	root.Insert("YY", "YY")
	root.Insert("KKKKKK", "KKKKKK")
	root.Insert("O", "O")
	root.Insert("ZZ", "ZZ")
	root.Insert("JJJJJJ", "JJJJJJ")
	root.Insert("V", "V")
	root.Insert("WWWWWWWWWWW", "WWWWWWWWWWW")
	root.Insert("BBB", "BBB")
	root.Insert("T", "T")
	root.Insert("LLLLLL", "LLLLLL")
	root.Insert("AAA", "AAA")
	root.Insert("N", "N")
	root.Insert("OOOOOOOOOOO", "OOOOOOOOOOO")
	root.Insert("U", "U")
	root.Insert("KKKKKKKKKKK", "KKKKKKKKKKK")
	root.Insert("W", "W")
	root.Insert("SSSSSSSS", "SSSSSSSS")
	root.Insert("Y", "Y")
	root.Insert("QQ", "QQ")
	root.Insert("YYYYYYYYYYY", "YYYYYYYYYYY")
	root.Insert("OOOOOO", "OOOOOO")
	root.Insert("Z", "Z")
	root.Insert("HHHHHHHH", "HHHHHHHH")
	root.Insert("RRRRRRRRRRR", "RRRRRRRRRRR")
	root.Insert("UU", "UU")
	root.Insert("EEEE", "EEEE")
	root.Insert("XXXXXXXXXXX", "XXXXXXXXXXX")
	root.Insert("QQQ", "QQQ")
	root.Insert("TTTTTTTTTTT", "TTTTTTTTTTT")
	root.Insert("EE", "EE")
	root.Insert("KKKKKKKK", "KKKKKKKK")
	root.Insert("IIIIII", "IIIIII")
	root.Insert("GG", "GG")
	root.Insert("IIIIIIII", "IIIIIIII")
	root.Insert("JJJ", "JJJ")
	root.Insert("BB", "BB")
	root.Insert("RRRRRRRR", "RRRRRRRR")
	root.Insert("II", "II")
	root.Insert("X", "X")
	root.Insert("WWWW", "WWWW")
	root.Insert("BBBBBBBB", "BBBBBBBB")
	root.Insert("SSS", "SSS")
	root.Insert("AA", "AA")
	root.Insert("TTT", "TTT")
	root.Insert("XXXX", "XXXX")
	root.Insert("JJ", "JJ")
	root.Insert("WWWWWWWW", "WWWWWWWW")
	root.Insert("JJJJ", "JJJJ")
	root.Insert("K", "K")
	root.Insert("TT", "TT")
	root.Insert("OOO", "OOO")
	root.Insert("IIII", "IIII")
	root.Insert("CCCCCCCC", "CCCCCCCC")
	root.Insert("PP", "PP")
	root.Insert("UUUUUUUUUUU", "UUUUUUUUUUU")
	root.Insert("DD", "DD")
	root.Insert("MMMMMMMM", "MMMMMMMM")
	root.Insert("HHHHHH", "HHHHHH")
	root.Insert("QQQQ", "QQQQ")
	root.Insert("VVV", "VVV")
	root.Insert("SSSSSSSSSSS", "SSSSSSSSSSS")
	root.Insert("LLLLLLLL", "LLLLLLLL")
	root.Insert("UUU", "UUU")
	root.Insert("PPPPPPPP", "PPPPPPPP")
	root.Insert("CC", "CC")
	root.Insert("ZZZZZZZZZZZ", "ZZZZZZZZZZZ")
	root.Insert("NNNNNNNN", "NNNNNNNN")
	root.Insert("FFFF", "FFFF")
	root.Insert("NN", "NN")
	root.Insert("CCCC", "CCCC")
	root.Insert("EEEEEEEE", "EEEEEEEE")
	root.Insert("DDDD", "DDDD")
	root.Insert("ZZZ", "ZZZ")
	root.Insert("PPPP", "PPPP")
	root.Insert("CCCCCC", "CCCCCC")
	root.Insert("XXX", "XXX")
	root.Insert("J", "J")
	root.Insert("FFFFFFFF", "FFFFFFFF")
	root.Insert("AAAAAA", "AAAAAA")
	root.Insert("SSSS", "SSSS")
	root.Insert("EEEEEE", "EEEEEE")
	root.Insert("RRR", "RRR")
	root.Insert("LLLL", "LLLL")
	root.Insert("BBBBBB", "BBBBBB")
	root.Insert("RRRR", "RRRR")
	root.Insert("L", "L")
	root.Insert("MMMM", "MMMM")
	root.Insert("CCC", "CCC")
	root.Insert("AAAA", "AAAA")
	root.Insert("F", "F")
	root.Insert("MMMMMM", "MMMMMM")
	root.Insert("NNN", "NNN")
	root.Insert("YYYYYYYY", "YYYYYYYY")
	root.Insert("LLL", "LLL")
	root.Insert("QQQQQQQQ", "QQQQQQQQ")
	root.Insert("YYYY", "YYYY")
	root.Insert("GGGGGGGG", "GGGGGGGG")
	root.Insert("MMM", "MMM")
	root.Insert("D", "D")
	root.Insert("GGGG", "GGGG")
	root.Insert("UUUUUU", "UUUUUU")
	root.Insert("QQQQQQQQQQQ", "QQQQQQQQQQQ")
	root.Insert("FF", "FF")
	root.Insert("OOOO", "OOOO")
	root.Insert("DDDDDD", "DDDDDD")
	root.Insert("ZZZZZZZZ", "ZZZZZZZZ")
	root.Insert("HH", "HH")
	root.Insert("Q", "Q")
	root.Insert("XXXXXXXX", "XXXXXXXX")
	root.Insert("AAAAAAAAAAA", "AAAAAAAAAAA")
	root.Insert("ZZZZZZ", "ZZZZZZ")
	root.Insert("BBBBBBBBBBB", "BBBBBBBBBBB")
	root.Insert("SSSSSS", "SSSSSS")
	root.Insert("CCCCCCCCCCC", "CCCCCCCCCCC")
	root.Insert("WWWWWW", "WWWWWW")
	root.Insert("FFF", "FFF")
	root.Insert("VVVVVV", "VVVVVV")
	root.Insert("H", "H")
	root.Insert("NNNN", "NNNN")
	root.Insert("TTTTTT", "TTTTTT")
	root.Insert("YYYYYY", "YYYYYY")
	root.Insert("VVVVVVVVVVV", "VVVVVVVVVVV")
	root.Insert("XXXXXX", "XXXXXX")
	root.Insert("DDDDDDDD", "DDDDDDDD")
	root.Insert("NNNNNN", "NNNNNN")
	root.Insert("JJJJJJJJ", "JJJJJJJJ")
	root.Insert("EEE", "EEE")
	root.Insert("FFFFFFFFFFF", "FFFFFFFFFFF")
	root.Insert("GGGGGGGGGGG", "GGGGGGGGGGG")
	root.Insert("R", "R")
	root.Insert("UUUUUUUU", "UUUUUUUU")
	root.Insert("HHHHHHHHHHH", "HHHHHHHHHHH")
	root.Insert("RR", "RR")
	root.Insert("GGG", "GGG")
	root.Insert("TTTTTTTT", "TTTTTTTT")
	root.Insert("YYY", "YYY")
	root.Insert("G", "G")
	root.Insert("MMMMMMMMMMM", "MMMMMMMMMMM")
	root.Insert("QQQQQQ", "QQQQQQ")
	root.Insert("NNNNNNNNNNN", "NNNNNNNNNNN")
	root.Insert("XX", "XX")
	root.Insert("AAAAAAAA", "AAAAAAAA")
	root.Insert("SS", "SS")
	root.Insert("PPPPPP", "PPPPPP")
	root.Insert("RRRRRR", "RRRRRR")
	root.Insert("VVVVVVVV", "VVVVVVVV")
	root.Insert("S", "S")
	root.Insert("DDDDDDDDDDD", "DDDDDDDDDDD")

	return &root
}
