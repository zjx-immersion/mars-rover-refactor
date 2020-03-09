package marsrover

import "testing"

func Test_should_return_x_plus_1_given_given_command_is_M_and_facing_is_E(t *testing.T) {
	initLocation(0, 0, "E")
	x, _, _ := sendCommands([]byte{'M'})
	assertEqual(1, x, t)

}

func Test_should_return_x_minus_1_given_given_command_is_M_and_facing_is_W(t *testing.T) {
	initLocation(2, 0, "W")
	x, _, _ := sendCommands([]byte{'M'})
	assertEqual(1, x, t)
}

func Test_should_return_y_plus_1_given_given_command_is_M_and_facing_is_N(t *testing.T) {
	initLocation(0, 0, "N")
	_, y, _ := sendCommands([]byte{'M'})
	assertEqual(1, y, t)
}

func Test_should_return_y_minus_1_given_given_command_is_M_and_facing_is_S(t *testing.T) {
	initLocation(0, 2, "S")
	_, y, _ := sendCommands([]byte{'M'})
	assertEqual(1, y, t)
}

func Test_should_return_facing_N_given_given_command_is_L_and_facing_is_E(t *testing.T) {
	initLocation(0, 0, "E")
	_, _, direction := sendCommands([]byte{'L'})
	assertEqual("N", direction, t)
}

func Test_should_return_facing_W_given_given_command_is_L_and_facing_is_N(t *testing.T) {
	initLocation(0, 0, "N")
	_, _, direction := sendCommands([]byte{'L'})
	assertEqual("W", direction, t)
}

func Test_should_return_facing_S_given_given_command_is_L_and_facing_is_W(t *testing.T) {
	initLocation(0, 0, "W")
	_, _, direction := sendCommands([]byte{'L'})
	assertEqual("S", direction, t)
}

func Test_should_return_facing_E_given_given_command_is_L_and_facing_is_S(t *testing.T) {
	initLocation(0, 0, "S")
	_, _, direction := sendCommands([]byte{'L'})
	assertEqual("E", direction, t)
}

func Test_should_return_facing_S_given_given_command_is_R_and_facing_is_E(t *testing.T) {
	initLocation(0, 0, "E")
	_, _, direction := sendCommands([]byte{'R'})
	assertEqual("S", direction, t)
}

func Test_should_return_facing_W_given_given_command_is_R_and_facing_is_S(t *testing.T) {
	initLocation(0, 0, "S")
	_, _, direction := sendCommands([]byte{'R'})
	assertEqual("W", direction, t)
}

func Test_should_return_facing_N_given_given_command_is_R_and_facing_is_W(t *testing.T) {
	initLocation(0, 0, "W")
	_, _, direction := sendCommands([]byte{'R'})
	assertEqual("N", direction, t)
}

func Test_should_return_facing_E_given_given_command_is_R_and_facing_is_N(t *testing.T) {
	initLocation(0, 0, "N")
	_, _, direction := sendCommands([]byte{'R'})
	assertEqual("E", direction, t)
}

func Test_should_return_correct_location_and_direction_given_command_is_MLMMRMMML_when_facing_N(t *testing.T) {
	initLocation(0, 0, "N")
	x, y, direction := sendCommands([]byte("MLMMRMMML"))
	assertEqual(-2, x, t)
	assertEqual(4, y, t)
	assertEqual("W", direction, t)
}
