package main

import (
	"fmt"
	"math"
	"testing"

	
	"github.com/stretchr/testify/assert"
)

func main() {
    v := Abs(3)
    fmt.Println(v)
}

// Abs возвращает абсолютное значение.
// Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
// Покрыть тестами нужно эту функцию.
func Abs(value float64) float64 {
    return math.Abs(value)
} 

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		value float64
		want float64
	}{
		{
			name: "Положительное число",
			value: 5,
			want:  5,
		},{
			name:  "Отрицательное число",
			value: -17,
			want:  17,
		},{
			name:  "Отрицательное с дробной частью",
			value: -2.007656,
			want:  2.007656,
		},{
			name:  "Положитеьное с дробной частью",
			value: 5.007656,
			want:  5.007656,
		},
	}
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			v := Abs(tt.value)
			assert.Equal(t, tt.want, v)
		})
	}

}
