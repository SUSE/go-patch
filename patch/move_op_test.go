package patch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/SUSE/go-patch/patch"
)

var _ = Describe("MoveOp.Apply", func() {
	Describe("array item", func() {
		It("moves array item", func() {
			res, err := MoveOp{
				Path: MustNewPointerFromString("/-"),
				From: MustNewPointerFromString("/0"),
			}.Apply([]interface{}{1, 2, 3})
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal([]interface{}{2, 3, 1}))
		})
	})

	Describe("map key", func() {
		It("renames map key", func() {
			doc := map[interface{}]interface{}{
				"abc": "abc",
				"xyz": "xyz",
			}

			res, err := MoveOp{
				From: MustNewPointerFromString("/abc"),
				Path: MustNewPointerFromString("/def?"),
			}.Apply(doc)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(map[interface{}]interface{}{
				"def": "abc",
				"xyz": "xyz",
			}))
		})
	})
})
