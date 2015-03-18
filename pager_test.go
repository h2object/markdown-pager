package pager

import (
	"log"
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

// func TestAuthorities(t *testing.T) {

func TestPager(t *testing.T) {
	toml, err1 := os.Open("testdata/toml_meta.md")
	assert.Nil(t, err1)

	page, err2 := ReadFrom(toml)
	assert.Nil(t, err2)

	m1, err3 := page.Metadata()
	assert.Nil(t, err3)
	log.Println("toml meta:", m1)
	log.Println("toml content:", string(page.Content()))
}
