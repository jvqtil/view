package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/fatih/color"
	"github.com/jvqtil/view"
)

func main() {
	// example help text, probably as you would do in your app
	var helpText = `This is an example help of *some* app
  [j] scroll down
  [k] scroll up
  [enter] confirm
  [h/?] open help
  [q] quit

And some lorem ipsum down here (formatted as one sentence per line)
  Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
  Vivamus interdum quis velit eu pretium. 
  Donec venenatis pellentesque magna, faucibus imperdiet massa porttitor eu. 
  Morbi sodales odio a sapien imperdiet, sed pulvinar augue cursus.
  Sed consectetur posuere mi, at tincidunt risus interdum et. 
  Maecenas sed purus in lectus volutpat commodo. 
  Aliquam eget lorem vehicula, porta lectus ut, aliquet quam. 
  Vivamus quis felis nunc. 
  Etiam suscipit dolor a massa malesuada sagittis. 
  Fusce facilisis dui mollis, porttitor lectus sed, consequat ante. 
  Aenean lorem arcu, efficitur a varius id, vestibulum non felis.
  Vestibulum porttitor urna vestibulum blandit condimentum. 
  Phasellus iaculis sed sem non venenatis. 
  Suspendisse ultrices pulvinar est eu ornare. 
  In eleifend justo id mi euismod semper. 
  Nunc vitae ornare elit, eu molestie justo. 
  Maecenas fringilla mi non ipsum finibus volutpat. 
  Vivamus placerat lectus eget interdum imperdiet. 
  Cras eu libero vel eros cursus ultricies. 
  Sed ac luctus augue, vel malesuada eros. 
  Pellentesque ut lacus tristique, venenatis enim ac, auctor ante. 
  Proin dapibus ipsum vel nibh ullamcorper convallis. 
  Etiam euismod odio elit, vel egestas nunc volutpat id. 
  Donec venenatis maximus ex, eget dictum erat mattis nec. 
  Vestibulum consequat odio congue sagittis egestas.`

	fmt.Println("raw")
	view.Show(helpText)

	fmt.Println("fatih/color")
	view.Show(helpText, color.New(color.FgYellow).SprintFunc())

	fmt.Println("lipgloss")
	styledHelp := lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render(helpText)
	view.Show(styledHelp)
}
