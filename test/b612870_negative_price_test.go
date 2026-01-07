package test

import (
	"LabTest/model"

	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func ValidateNegativePrice(t *testing.T) {
	g := NewWithT(t)

	// price
	t.Run("Price must be between 50 and 5000", func(t *testing.T) {
		book := model.Books{
			Title: "text",
			Price: 5,
			Code:  "B6612870",
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))
	})

}
