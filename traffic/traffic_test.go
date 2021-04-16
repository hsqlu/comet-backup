package traffic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraffic(t *testing.T) {
	input := [][]int32{
		{1, 3, 2, 2},
		{2, 4, 3, 7},
		{3, 2, 1, 1},
		{2, 9, 7, 5},
		{2, 5, 4, 4},
	}
	expected := [][]int32{
		{8, 16, 7, 23, 15},
		{80, 75, 100, 90, 75},
	}

	expectedController := []Controller{
		new(Roundabout),
		new(TrafficLight),
		new(Roundabout),
		new(TrafficLight),
		new(TrafficLight),
	}
	for index, sub := range input {
		i := New(sub[0], sub[1], sub[2], sub[3])
		assert.Equal(t, expected[0][index], i.TotalCPM())
		j, c := i.EfficientController()
		fmt.Println(j, c.Name())
		assert.Equal(t, expected[1][index], j)
		assert.Equal(t, expectedController[index], c)
	}
}
