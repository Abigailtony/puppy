package puppy

import "github.com/Abigailtony/dog"

func Bark() string {
	return "WOOF!"
}

func Barks() string {
	return "WOOF! WOOF! WOOF!"
}

func Bigbark() string {
	return dog.WhenAult(Bark())
}

func BigbarkS() string {
	return dog.WhenAult(Bark())
}
