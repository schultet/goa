package opt

import "fmt"

func someOptions() *OptionSet {
	opts := NewOptionSet()
	opts.Add(NewOption(Bool, "foo", 'f', false, "a boolean option"))
	opts.Add(NewOption(String, "bar", 'b', "42", "a string option"))
	return opts
}

func Example() {
	opts := someOptions()
	args := []string{"_", "-f", "-x", "--bar", "thirtyfive"}

	opts.Parse(args)
	var foo bool = opts.GetBool("foo")
	var bar string = opts.GetString("bar")
	fmt.Println(foo, bar)
	fmt.Println(args)

	// Output:
	// true thirtyfive
	// [_ -f -x --bar thirtyfive]
}
