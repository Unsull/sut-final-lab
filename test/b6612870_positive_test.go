package test

import (
	"LabTest/model"

	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Validate(t *testing.T) {
	g := NewWithT(t)

	// all suc
	t.Run("Corent to all", func(t *testing.T) {
		book := model.Books{
			Title: "text",
			Price: 500,
			Code:  "BK1232456",
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

}
