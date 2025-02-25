package main

import (
	"belajar/logic1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1(t *testing.T) {
	expected := []int{1, 3, 5, 7, 9}
	assert.Equal(t, expected, logic1.Soal1(5))
}

func TestSoal2(t *testing.T) {
	expected := []int{2, 4, 6, 8, 10}
	assert.Equal(t, expected, logic1.Soal2(5))
}

func TestSoal3(t *testing.T) {

}
