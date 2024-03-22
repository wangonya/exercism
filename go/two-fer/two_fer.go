package twofer

// If your friend likes cookies, and is named Do-yun, then you will say:
// One for Do-yun, one for me.
// else, say _you_
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
