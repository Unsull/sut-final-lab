package test

import (
	"LabTest/model"

	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Validate(t *testing.T) {
	g := NewWiht(t)

	// all suc
	g.Run("Corent to all", func(t *testing.T) {
		g.book := model.Books{
			Title: "text",
			Price: 500,
			Code: "B6612870",
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeTrue)
		g.Expect(err).To(BeNil)
	}) 
}