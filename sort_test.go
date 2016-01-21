package natsort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoVersion(t *testing.T) {
	input := []string{"1.6beta1", "1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4", "1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.5.2", "1.5alpha1"}
	expect := []string{"1.0.2", "1.0.3", "1.1", "1.1.1", "1.1.2", "1.2rc1", "1.2rc2", "1.2rc3", "1.2rc4", "1.2rc5", "1.2", "1.2.1", "1.2.2", "1.3beta1", "1.3beta2", "1.3rc1", "1.3rc2", "1.3", "1.3.1", "1.3.2", "1.3.3", "1.4beta1", "1.4rc1", "1.4rc2", "1.4", "1.4.1", "1.4.2", "1.5alpha1", "1.5beta1", "1.5beta2", "1.5rc1", "1.5", "1.5.1", "1.5.2", "1.6beta1"}

	Convey("sort golang version should be success", t, func() {
		So(Sort(input), ShouldResemble, expect)
	})
}

func TestVersion(t *testing.T) {
	input := []string{"a2", "a5", "a9", "a1", "a4", "a10", "a6"}
	expect := []string{"a1", "a2", "a4", "a5", "a6", "a9", "a10"}
	Convey("sort golang version should be success", t, func() {
		So(Sort(input), ShouldResemble, expect)
	})
}

func TestOtherVersion(t *testing.T) {
	input := []string{"num-1", "num-0"}
	expect := []string{"num-0", "num-1"}
	Convey("sort golang version should be success", t, func() {
		So(Sort(input), ShouldResemble, expect)
	})

	input = []string{"1.1", "11.0.1", "10.1.1", "10.1beta10", "10.1beta2", "num-1"}
	expect = []string{"num-1", "1.1", "10.1beta2", "10.1beta10", "10.1.1", "11.0.1"}
	Convey("sort golang version should be success", t, func() {
		So(Sort(input), ShouldResemble, expect)
	})

	input = []string{"num-1", "num+1"}
	expect = []string{"num+1", "num-1"}
	Convey("sort golang version should be success", t, func() {
		So(Sort(input), ShouldResemble, expect)
	})
}

func TestEmptyVersion(t *testing.T) {
	input := []string{"", "a0"}
	expect := []string{"", "a0"}
	Convey("sort golang version should be success", t, func() {
		So(Sort(input), ShouldResemble, expect)
	})

	input = []string{"a0", ""}
	expect = []string{"", "a0"}
	Convey("sort golang version should be success", t, func() {
		So(Sort(input), ShouldResemble, expect)
	})
}

func TestRSort(t *testing.T) {
	input := []string{"a2", "a5", "a9", "a1", "a4", "a10", "a6"}
	expect := []string{"a10", "a9", "a6", "a5", "a4", "a2", "a1"}
	Convey("sort golang version should be success", t, func() {
		So(RSort(input), ShouldResemble, expect)
	})
}

func BenchmarkSort(b *testing.B) {
	input := []string{"1.6beta1", "1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4", "1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.5.2", "1.5alpha1"}
	for i := 0; i < b.N; i++ {
		Sort(input)
	}
}
