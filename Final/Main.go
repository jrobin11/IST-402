package main

import (
	"fmt"
	"strings"
)

// Rotor is a struct representing a single rotor in the Enigma machine.
type Rotor struct {
	shift int // the current shift value of the rotor
}

// NewRotor creates a new Rotor with the given initial shift.
func NewRotor(shift int) *Rotor {
	return &Rotor{shift: shift} // return a new Rotor with the given initial shift
}

// Rotate advances the rotor by one position, wrapping around after 26 positions.
func (r *Rotor) Rotate() {
	r.shift = (r.shift + 1) % 26 // advance the rotor by one position, wrapping around after 26 positions
}

// ShiftChar applies the rotor's shift to the given character. If forward is true, it
// shifts the character forward; if false, it shifts the character backward.
func (r *Rotor) ShiftChar(c byte, forward bool) byte {
	if forward {
		// shift the character forward by the rotor's shift value
		return byte((int(c-'A')+r.shift)%26 + 'A')
	}
	// shift the character backward by the rotor's shift value
	return byte((int(c-'A')-r.shift+26)%26 + 'A')
}

// Enigma is a struct representing the Enigma machine.
type Enigma struct {
	rotors []*Rotor // an array of rotors in the Enigma machine
}

// NewEnigma creates a new Enigma machine with the given rotor shifts.
func NewEnigma(rotorShifts []int) *Enigma {
	// create a new Enigma machine with the given rotor shifts
	rotors := make([]*Rotor, len(rotorShifts))
	for i, shift := range rotorShifts {
		rotors[i] = NewRotor(shift)
	}
	return &Enigma{rotors: rotors}
}

// Encrypt encrypts the given text using the Enigma machine.
func (e *Enigma) Encrypt(text string) string {
	upperText := strings.ToUpper(text) // convert the input text to uppercase
	encrypted := make([]byte, len(upperText))

	// Process each character in the input text
	for i, c := range upperText {
		// Only process alphabetic characters
		if c >= 'A' && c <= 'Z' {
			// Apply each rotor's shift to the character, moving forward
			for _, rotor := range e.rotors {
				c = rune(rotor.ShiftChar(byte(c), true))
			}
			// Advance the rotors after processing the character
			e.advanceRotors()
		}
		encrypted[i] = byte(c) // Store the processed character
	}

	return string(encrypted) // return the encrypted message as a string
}

// Decrypt decrypts the given text using the Enigma machine.
func (e *Enigma) Decrypt(text string) string {
	upperText := strings.ToUpper(text) // convert the input text to uppercase
	decrypted := make([]byte, len(upperText))

	// Process each character in the input text
	for i, c := range upperText {
		// Only process alphabetic characters
		if c >= 'A' && c <= 'Z' {
			// Apply each rotor's shift to the character, moving backward
			for j := len(e.rotors) - 1; j >= 0; j-- {
				c = rune(e.rotors[j].ShiftChar(byte(c), false))
			}
			// Advance the rotors after processing the character
			e.advanceRotors()
		}
		decrypted[i] = byte(c) // Store the processed character
	}

	return string(decrypted)
}

// advanceRotors advances the rotors in the Enigma machine.
// advanceRotors advances the positions of the rotors in the Enigma machine.
func (e *Enigma) advanceRotors() {
	// Loop through each rotor
	for i := 0; i < len(e.rotors); i++ {
		// Rotate the current rotor by one position
		e.rotors[i].Rotate()

		// If the current rotor's shift is not at its initial position (0),
		// then stop rotating the remaining rotors and exit the loop.
		if e.rotors[i].shift != 0 {
			break
		}
	}
}

// ResetRotors resets the positions of the rotors to the given shifts.
func (e *Enigma) ResetRotors(shifts []int) {
	for i, shift := range shifts {
		e.rotors[i].shift = shift
	}
}

func main() {
	// Create a new Enigma machine with rotor shifts of 1, 3, and 5
	rotorShifts := []int{1, 3, 5}
	enigma := NewEnigma(rotorShifts)

	// Encrypt and decrypt a message
	message := "IST 402 personalized Enigma Machine"
	encrypted := enigma.Encrypt(message)
	fmt.Printf("Encrypted message: %s\n", encrypted)

	enigma.ResetRotors(rotorShifts)
	decrypted := enigma.Decrypt(encrypted)
	fmt.Printf("Decrypted message: %s\n", decrypted)
}
