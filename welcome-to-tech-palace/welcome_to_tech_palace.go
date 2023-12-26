package techpalace
import "strings"
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {

	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border:=strings.Repeat("*", numStarsPerLine)
	border += "\n" + welcomeMsg + "\n"
	border+=strings.Repeat("*", numStarsPerLine)
	return border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	str:=strings.Trim(oldMsg, "* ");
	str=strings.Trim(str, "\n*");
	str=strings.Trim(str, " ");
	return str
}
