package interpreter

import (
	"math"
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MAX_DEPTH = math.MaxInt16

func TestPush(t *testing.T) {
	assert, stack := setup(t)
	letter := rune(rand.Int31())

	stack.Push(letter)

	retrieved, err := stack.Pop()
	assert.Nil(err)
	assert.Equal(retrieved, letter)
}

func TestPushN(t *testing.T) {
	assert, stack := setup(t)
	letters := randNLetters(rand.Intn(MAX_DEPTH))

	stack.PushN(letters...)
	assert.Equal(stack.Length(), len(letters))

	// Reverse the input array to match pop order and compare
	slices.Reverse(letters)
	for _, val := range letters {
		stackVal, _ := stack.Pop()
		assert.Equal(val, stackVal)
	}
}

func TestPop(t *testing.T) {
	assert, stack := setup(t)
	letter := randLetter()

	stack.Push(letter)
	popped, _ := stack.Pop()
	assert.Equal(letter, popped)
}

func TestPopN(t *testing.T) {
	assert, stack := setup(t)
	letters := randNLetters(rand.Intn(MAX_DEPTH))

	stack.PushN(letters...)
	assert.Equal(len(letters), stack.Length())

	// Reverse the input array to match pop order and compare
	slices.Reverse(letters)
	received, _ := stack.PopN(len(letters))

	for idx, val := range letters {
		assert.Equal(val, received[idx])
	}
}

func TestDuplicate(t *testing.T) {
	assert, stack := setup(t)
	letter := randLetter()

	stack.Push(letter)
	assert.Equal(1, stack.Length())
	stack.Duplicate()
	assert.Equal(2, stack.Length())

	var (
		expected    = []rune{letter, letter}
		received, _ = stack.PopN(2)
	)
	assertSlicesMatch(assert, expected, received)
}

func TestSwap(t *testing.T) {
	assert, stack := setup(t)
	expected := randNLetters(2)

	stack.PushN(expected...)
	stack.Swap()

	slices.Reverse(expected)
	received, _ := stack.PopN(2)
	assertSlicesMatch(assert, expected, received)
}

func TestRshift(t *testing.T) {
	assert, stack := setup(t)
	expected := randNLetters(3)

	stack.PushN(expected...)
	stack.Rshift() // Dropping a (hopefully) impossible error

	// manual rshift + reverse to account for LIFO
	expected = []rune{expected[1], expected[0], expected[2]}
	received, _ := stack.PopN(3)
	assertSlicesMatch(assert, expected, received)
}

/**** Utility ****/
func setup(t *testing.T) (*assert.Assertions, *stack) {
	return assert.New(t), new(stack)
}

func randLetter() rune {
	return rune(rand.Int31())
}

func randNLetters(n int) (out []rune) {
	out = make([]rune, n)
	for idx := range out {
		out[idx] = randLetter()
	}
	return
}

// testify/assert only has ElementsMatch which does not account for order
// As a stack is an ordered data structure, this is a problem
func assertSlicesMatch[T comparable](assert *assert.Assertions, expected []T, actual []T) {
	for idx, elem := range expected {
		assert.Equal(elem, actual[idx])
	}
}
