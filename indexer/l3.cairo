// Address: 0x00806778f9b06746fffd6ca567e0cfea9b3515432d9ba39928201d18c8dc9fdf

%lang starknet

from starkware.cairo.common.cairo_builtins import HashBuiltin

//
// Event
//
@event
func value_change(value: felt) {
}

//
// Storage
//
@storage_var
func value() -> (res: felt) {
}

//
// Getter
//
@view
func get_value{syscall_ptr: felt*, pedersen_ptr: HashBuiltin*, range_check_ptr}() -> (res: felt) {
    return value.read();
}

//
// Setter
//
@external
func set_value{syscall_ptr: felt*, pedersen_ptr: HashBuiltin*, range_check_ptr}(new_value: felt) {
    value.write(new_value);

    value_change.emit(new_value);

    return ();
}
