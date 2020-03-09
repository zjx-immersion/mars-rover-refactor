package mars_rover

import (
	"mars-rover-refactor/pkg/marsrover/assert"
	"testing"
)

func Test_should_return_x_plus_1_given_given_command_is_M_and_facing_is_E(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "E")
	x, _, _ := marsRover.sendCommand([]byte{'M'})
	assert.AssertEqual(1, x, t)

}

func Test_should_return_x_minus_1_given_given_command_is_M_and_facing_is_W(t *testing.T) {
	marsRover := NewMarsRover(2, 0, "W")
	x, _, _ := marsRover.sendCommand([]byte{'M'})
	assert.AssertEqual(1, x, t)
}

func Test_should_return_y_plus_1_given_given_command_is_M_and_facing_is_N(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "N")
	_, y, _ := marsRover.sendCommand([]byte{'M'})
	assert.AssertEqual(1, y, t)
}

func Test_should_return_y_minus_1_given_given_command_is_M_and_facing_is_S(t *testing.T) {
	marsRover := NewMarsRover(0, 2, "S")
	_, y, _ := marsRover.sendCommand([]byte{'M'})
	assert.AssertEqual(1, y, t)
}

func Test_should_return_facing_N_given_given_command_is_L_and_facing_is_E(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "E")
	_, _, direction := marsRover.sendCommand([]byte{'L'})
	assert.AssertEqual("N", direction, t)
}

func Test_should_return_facing_W_given_given_command_is_L_and_facing_is_N(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "N")
	_, _, direction := marsRover.sendCommand([]byte{'L'})
	assert.AssertEqual("W", direction, t)
}

func Test_should_return_facing_S_given_given_command_is_L_and_facing_is_W(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "W")
	_, _, direction := marsRover.sendCommand([]byte{'L'})
	assert.AssertEqual("S", direction, t)
}

func Test_should_return_facing_E_given_given_command_is_L_and_facing_is_S(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "S")
	_, _, direction := marsRover.sendCommand([]byte{'L'})
	assert.AssertEqual("E", direction, t)
}

func Test_should_return_facing_S_given_given_command_is_R_and_facing_is_E(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "E")
	_, _, direction := marsRover.sendCommand([]byte{'R'})
	assert.AssertEqual("S", direction, t)
}

func Test_should_return_facing_W_given_given_command_is_R_and_facing_is_S(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "S")
	_, _, direction := marsRover.sendCommand([]byte{'R'})
	assert.AssertEqual("W", direction, t)
}

func Test_should_return_facing_N_given_given_command_is_R_and_facing_is_W(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "W")
	_, _, direction := marsRover.sendCommand([]byte{'R'})
	assert.AssertEqual("N", direction, t)
}

func Test_should_return_facing_E_given_given_command_is_R_and_facing_is_N(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "N")
	_, _, direction := marsRover.sendCommand([]byte{'R'})
	assert.AssertEqual("E", direction, t)
}

func Test_should_return_correct_location_and_direction_given_command_is_MLMMRMMML_when_facing_N(t *testing.T) {
	marsRover := NewMarsRover(0, 0, "N")
	x, y, direction := marsRover.sendCommand([]byte("MLMMRMMML"))
	assert.AssertEqual(-2, x, t)
	assert.AssertEqual(4, y, t)
	assert.AssertEqual("W", direction, t)
}
