package flag_test

import (
	. "code.cloudfoundry.org/cli/v8/command/flag"
	flags "github.com/jessevdk/go-flags"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BindingName", func() {
	var bindingName BindingName

	BeforeEach(func() {
		bindingName = BindingName{}
	})

	When("the value provided to the --binding-name flag is the empty string", func() {
		It("returns a ErrMarshal error that the binding name must be greater than 1 character long", func() {
			Expect(bindingName.UnmarshalFlag("")).To(MatchError(&flags.Error{
				Type:    flags.ErrMarshal,
				Message: "--binding-name must be at least 1 character in length",
			}))
		})
	})

	When("the value provided to the --binding-name flag is greater than 0 characters long", func() {
		It("stores the binding name and does not return an error", func() {
			err := bindingName.UnmarshalFlag("some-name")
			Expect(err).NotTo(HaveOccurred())
			Expect(bindingName.Value).To(Equal("some-name"))
		})
	})

	Describe("Complete", func() {
		When("BindingNameCompleteFunc is not set", func() {
			It("returns nil", func() {
				result := bindingName.Complete("test-")
				Expect(result).To(BeNil())
			})
		})

		When("BindingNameCompleteFunc is set", func() {
			It("calls BindingNameCompleteFunc with the provided prefix and returns the result", func() {
				expectedCompletions := []flags.Completion{
					{Item: "binding-1"},
					{Item: "binding-2"},
				}
				BindingNameCompleteFunc = func(prefix string) []flags.Completion {
					return expectedCompletions
				}
				defer func() {
					BindingNameCompleteFunc = func(prefix string) []flags.Completion { return nil }
				}()

				result := bindingName.Complete("binding-")
				Expect(result).To(Equal(expectedCompletions))
			})
		})

		When("Complete is called with different prefixes", func() {
			It("passes the prefix correctly to BindingNameCompleteFunc", func() {
				var capturedPrefix string
				BindingNameCompleteFunc = func(prefix string) []flags.Completion {
					capturedPrefix = prefix
					return nil
				}
				defer func() {
					BindingNameCompleteFunc = func(prefix string) []flags.Completion { return nil }
				}()

				bindingName.Complete("my-prefix")
				Expect(capturedPrefix).To(Equal("my-prefix"))
			})
		})
	})
})
