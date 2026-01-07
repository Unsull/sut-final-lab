package test

import (
	"LabTest/model"

	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func ValidateNegativeCode(t *testing.T) {
	g := NewWithT(t)

	// price
	t.Run("Code must start with BK followed by 6 digits (0-9)", func(t *testing.T) {
		book := model.Books{
			Title: "text",
			Price: 5,
			Code:  "BK1232456",
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})

}
