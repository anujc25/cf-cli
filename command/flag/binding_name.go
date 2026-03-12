package flag

import flags "github.com/jessevdk/go-flags"

type BindingName struct {
	Value string
}

var BindingNameCompleteFunc = func(prefix string) []flags.Completion { return nil }

func (BindingName) Complete(prefix string) []flags.Completion {
	return BindingNameCompleteFunc(prefix)
}

func (b *BindingName) UnmarshalFlag(val string) error {
	if val == "" {
		return &flags.Error{
			Type:    flags.ErrMarshal,
			Message: "--binding-name must be at least 1 character in length",
		}
	}

	b.Value = val
	return nil
}
