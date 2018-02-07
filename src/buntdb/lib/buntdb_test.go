package lib

import (
	"testing"
	"fmt"
)

func TestRead(t *testing.T) {
	name := Read( "name")
	fmt.Println( name )

}

func TestWrite(t *testing.T) {
	Write( "name", "dcl")
}