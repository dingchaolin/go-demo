package intersection

import (
	"testing"
	"log"
)

func BenchmarkIntersection(b *testing.B) {
	for i := 0; i < b.N; i ++{
		strArr1 := []string{"aaa", "bbb"}
		strArr2 := []string{"bbb"}
		Intersection(strArr1,strArr2, "1", 3, []string{"bbb", "2"}, []string{"bbb", "2"})
	}
}

func BenchmarkTwoStrArr2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intersection([]string{"a", "b", "c", "d"}[:], []string{"c", "d", "e", "f", "g", "h"}[:])
	}
}

func TestIntersection(t *testing.T) {
	log.Println( Intersection([]string{"a", "b", "c", "d"}[:], []string{"c", "d", "e", "f", "g", "h"}[:]))
}