package calculator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Банальный тестовый случай, для сложения чисел
func TestAdd(t *testing.T) {
	t.Run("test Add function", func(t *testing.T) {
		actualResult := Add(10, 20)
		expectedResult := 30
		assert.Equal(t, actualResult, expectedResult, "Result of adding 10 and 20 is 30")
	})
}

// Банальный тестовый случай, для вычитания чисел
func TestSub(t *testing.T) {
	t.Run("test Sub function", func(t *testing.T) {
		actualResult := Sub(20, 10)
		expectedResult := -10
		assert.Equal(t, actualResult, expectedResult, fmt.Sprintf("Result of subtracting %d from %d is %d", 20, 10, actualResult))
	})
}
