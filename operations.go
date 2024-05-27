package main

type Operation struct {
	Name    string
	Comment string
}

var Operations = []Operation{
	{
		Name:    "i32_LOAD",
		Comment: "Load 32-bit integer from memory",
	},
	{
		Name:    "i64_LOAD",
		Comment: "Load 64-bit integer from memory",
	},
	{
		Name:    "f32_LOAD",
		Comment: "Load 32-bit float from memory",
	},
	{
		Name:    "f64_LOAD",
		Comment: "Load 64-bit float from memory",
	},
	{
		Name:    "i32_LOAD8_s",
		Comment: "Load 8-bit signed integer from memory",
	},
	{
		Name:    "i32_LOAD8_u",
		Comment: "Load 8-bit unsigned integer from memory",
	},
	{
		Name:    "i32_LOAD16_s",
		Comment: "Load 16-bit signed integer from memory",
	},
	{
		Name:    "i32_LOAD16_u",
		Comment: "Load 16-bit unsigned integer from memory",
	},
	{
		Name:    "i64_LOAD8_s",
		Comment: "Load 8-bit signed integer from memory",
	},
	{
		Name:    "i64_LOAD8_u",
		Comment: "Load 8-bit unsigned integer from memory",
	},
	{
		Name:    "i64_LOAD16_s",
		Comment: "Load 16-bit signed integer from memory",
	},
	{
		Name:    "i64_LOAD16_u",
		Comment: "Load 16-bit unsigned integer from memory",
	},
	{
		Name:    "i64_LOAD32_s",
		Comment: "Load 32-bit signed integer from memory",
	},
	{
		Name:    "i64_LOAD32_u",
		Comment: "Load 32-bit unsigned integer from memory",
	},
	{
		Name:    "i32_STORE",
		Comment: "Store 32-bit integer to memory",
	},
	{
		Name:    "i64_STORE",
		Comment: "Store 64-bit integer to memory",
	},
	{
		Name:    "f32_STORE",
		Comment: "Store 32-bit float to memory",
	},
	{
		Name:    "f64_STORE",
		Comment: "Store 64-bit float to memory",
	},
	{
		Name:    "i32_STORE8",
		Comment: "Store 8-bit integer to memory",
	},
	{
		Name:    "i32_STORE16",
		Comment: "Store 16-bit integer to memory",
	},
	{
		Name:    "i64_STORE8",
		Comment: "Store 8-bit integer to memory",
	},
	{
		Name:    "i64_STORE16",
		Comment: "Store 16-bit integer to memory",
	},
	{
		Name:    "i64_STORE32",
		Comment: "Store 32-bit integer to memory",
	},
	{
		Name:    "SIZE",
		Comment: "Size of memory",
	},
	{
		Name:    "GROW",
		Comment: "Grow memory",
	},
	{
		Name:    "i32_CONST",
		Comment: "32-bit integer constant",
	},
	{
		Name:    "i64_CONST",
		Comment: "64-bit integer constant",
	},
	{
		Name:    "f32_CONST",
		Comment: "32-bit float constant",
	},
	{
		Name:    "f64_CONST",
		Comment: "64-bit float constant",
	},
	{
		Name:    "i32_EQZ",
		Comment: "Test if 32-bit integer is zero",
	},
	{
		Name:    "i32_EQ",
		Comment: "Test if two 32-bit integers are equal",
	},
	{
		Name:    "i32_NE",
		Comment: "Test if two 32-bit integers are not equal",
	},
	{
		Name:    "i32_LT_s",
		Comment: "Test if first 32-bit integer is less than second (signed)",
	},
	{
		Name:    "i32_LT_u",
		Comment: "Test if first 32-bit integer is less than second (unsigned)",
	},
	{
		Name:    "i32_GT_s",
		Comment: "Test if first 32-bit integer is greater than second (signed)",
	},
	{
		Name:    "i32_GT_u",
		Comment: "Test if first 32-bit integer is greater than second (unsigned)",
	},
	{
		Name:    "i32_LE_s",
		Comment: "Test if first 32-bit integer is less than or equal to second (signed)",
	},
	{
		Name:    "i32_LE_u",
		Comment: "Test if first 32-bit integer is less than or equal to second (unsigned)",
	},
	{
		Name:    "i32_GE_s",
		Comment: "Test if first 32-bit integer is greater than or equal to second (signed)",
	},
	{
		Name:    "i32_GE_u",
		Comment: "Test if first 32-bit integer is greater than or equal to second (unsigned)",
	},
	{
		Name:    "i64_EQZ",
		Comment: "Test if 64-bit integer is zero",
	},
	{
		Name:    "i64_EQ",
		Comment: "Test if two 64-bit integers are equal",
	},
	{
		Name:    "i64_NE",
		Comment: "Test if two 64-bit integers are not equal",
	},
	{
		Name:    "i64_LT_s",
		Comment: "Test if first 64-bit integer is less than second (signed)",
	},
	{
		Name:    "i64_LT_u",
		Comment: "Test if first 64-bit integer is less than second (unsigned)",
	},
	{
		Name:    "i64_GT_s",
		Comment: "Test if first 64-bit integer is greater than second (signed)",
	},
	{
		Name:    "i64_GT_u",
		Comment: "Test if first 64-bit integer is greater than second (unsigned)",
	},
	{
		Name:    "i64_LE_s",
		Comment: "Test if first 64-bit integer is less than or equal to second (signed)",
	},
	{
		Name:    "i64_LE_u",
		Comment: "Test if first 64-bit integer is less than or equal to second (unsigned)",
	},
	{
		Name:    "i64_GE_s",
		Comment: "Test if first 64-bit integer is greater than or equal to second (signed)",
	},
	{
		Name:    "i64_GE_u",
		Comment: "Test if first 64-bit integer is greater than or equal to second (unsigned)",
	},
	{
		Name:    "f32_EQ",
		Comment: "Test if two 32-bit floats are equal",
	},
	{
		Name:    "f32_NE",
		Comment: "Test if two 32-bit floats are not equal",
	},
	{
		Name:    "f32_LT",
		Comment: "Test if first 32-bit float is less than second",
	},
	{
		Name:    "f32_GT",
		Comment: "Test if first 32-bit float is greater than second",
	},
	{
		Name:    "f32_LE",
		Comment: "Test if first 32-bit float is less than or equal to second",
	},
	{
		Name:    "f32_GE",
		Comment: "Test if first 32-bit float is greater than or equal to second",
	},
	{
		Name:    "f64_EQ",
		Comment: "Test if two 64-bit floats are equal",
	},
	{
		Name:    "f64_NE",
		Comment: "Test if two 64-bit floats are not equal",
	},
	{
		Name:    "f64_LT",
		Comment: "Test if first 64-bit float is less than second",
	},
	{
		Name:    "f64_GT",
		Comment: "Test if first 64-bit float is greater than second",
	},
	{
		Name:    "f64_LE",
		Comment: "Test if first 64-bit float is less than or equal to second",
	},
	{
		Name:    "f64_GE",
		Comment: "Test if first 64-bit float is greater than or equal to second",
	},
	{
		Name:    "i32_CLZ",
		Comment: "Count leading zeros in 32-bit integer",
	},
	{
		Name:    "i32_CTZ",
		Comment: "Count trailing zeros in 32-bit integer",
	},
	{
		Name:    "i32_POPCNT",
		Comment: "Count number of one bits in 32-bit integer",
	},
	{
		Name:    "i32_ADD",
		Comment: "Add two 32-bit integers",
	},
	{
		Name:    "i32_SUB",
		Comment: "Subtract one 32-bit integer from another",
	},
	{
		Name:    "i32_MUL",
		Comment: "Multiply two 32-bit integers",
	},
	{
		Name:    "i32_DIV_s",
		Comment: "Divide one 32-bit integer by another (signed)",
	},
	{
		Name:    "i32_DIV_u",
		Comment: "Divide one 32-bit integer by another (unsigned)",
	},
	{
		Name:    "i32_REM_s",
		Comment: "Calculate remainder of dividing one 32-bit integer by another (signed)",
	},
	{
		Name:    "i32_REM_u",
		Comment: "Calculate remainder of dividing one 32-bit integer by another (unsigned)",
	},
	{
		Name:    "i32_AND",
		Comment: "Bitwise AND of two 32-bit integers",
	},
	{
		Name:    "i32_OR",
		Comment: "Bitwise OR of two 32-bit integers",
	},
	{
		Name:    "i32_XOR",
		Comment: "Bitwise XOR of two 32-bit integers",
	},
	{
		Name:    "i32_SHL",
		Comment: "Shift left logical of 32-bit integer",
	},
	{
		Name:    "i32_SHR_s",
		Comment: "Shift right arithmetic of 32-bit integer",
	},
	{
		Name:    "i32_SHR_u",
		Comment: "Shift right logical of 32-bit integer",
	},
	{
		Name:    "i32_ROTL",
		Comment: "Rotate left of 32-bit integer",
	},
	{
		Name:    "i32_ROTR",
		Comment: "Rotate right of 32-bit integer",
	},
	{
		Name:    "i64_CLZ",
		Comment: "Count leading zeros in 64-bit integer",
	},
	{
		Name:    "i64_CTZ",
		Comment: "Count trailing zeros in 64-bit integer",
	},
	{
		Name:    "i64_POPCNT",
		Comment: "Count number of one bits in 64-bit integer",
	},
	{
		Name:    "i64_ADD",
		Comment: "Add two 64-bit integers",
	},
	{
		Name:    "i64_SUB",
		Comment: "Subtract one 64-bit integer from another",
	},
	{
		Name:    "i64_MUL",
		Comment: "Multiply two 64-bit integers",
	},
	{
		Name:    "i64_DIV_s",
		Comment: "Divide one 64-bit integer by another (signed)",
	},
	{
		Name:    "i64_DIV_u",
		Comment: "Divide one 64-bit integer by another (unsigned)",
	},
	{
		Name:    "i64_REM_s",
		Comment: "Calculate remainder of dividing one 64-bit integer by another (signed)",
	},
	{
		Name:    "i64_REM_u",
		Comment: "Calculate remainder of dividing one 64-bit integer by another (unsigned)",
	},
	{
		Name:    "i64_AND",
		Comment: "Bitwise AND of two 64-bit integers",
	},
	{
		Name:    "i64_OR",
		Comment: "Bitwise OR of two 64-bit integers",
	},
	{
		Name:    "i64_XOR",
		Comment: "Bitwise XOR of two 64-bit integers",
	},
	{
		Name:    "i64_SHL",
		Comment: "Shift left logical of 64-bit integer",
	},
	{
		Name:    "i64_SHR_s",
		Comment: "Shift right arithmetic of 64-bit integer",
	},
	{
		Name:    "i64_SHR_u",
		Comment: "Shift right logical of 64-bit integer",
	},
	{
		Name:    "i64_ROTL",
		Comment: "Rotate left of 64-bit integer",
	},
	{
		Name:    "i64_ROTR",
		Comment: "Rotate right of 64-bit integer",
	},
	{
		Name:    "f32_ABS",
		Comment: "Absolute value of 32-bit float",
	},
	{
		Name:    "f32_NEG",
		Comment: "Negation of 32-bit float",
	},
	{
		Name:    "f32_CEIL",
		Comment: "Round up to nearest integer for 32-bit float",
	},
	{
		Name:    "f32_FLOOR",
		Comment: "Round down to nearest integer for 32-bit float",
	},
	{
		Name:    "f32_TRUNC",
		Comment: "Round towards zero to nearest integer for 32-bit float",
	},
	{
		Name:    "f32_NEAREST",
		Comment: "Round to nearest integer, ties to even, for 32-bit float",
	},
	{
		Name:    "f32_SQRT",
		Comment: "Square root of 32-bit float",
	},
	{
		Name:    "f32_ADD",
		Comment: "Add two 32-bit floats",
	},
	{
		Name:    "f32_SUB",
		Comment: "Subtract one 32-bit float from another",
	},
	{
		Name:    "f32_MUL",
		Comment: "Multiply two 32-bit floats",
	},
	{
		Name:    "f32_DIV",
		Comment: "Divide one 32-bit float by another",
	},
	{
		Name:    "f32_MIN",
		Comment: "Minimum of two 32-bit floats",
	},
	{
		Name:    "f32_MAX",
		Comment: "Maximum of two 32-bit floats",
	},
	{
		Name:    "f32_COPYSIGN",
		Comment: "Copy sign from one 32-bit float to another",
	},
	{
		Name:    "f64_ABS",
		Comment: "Absolute value of 64-bit float",
	},
	{
		Name:    "f64_NEG",
		Comment: "Negation of 64-bit float",
	},
	{
		Name:    "f64_CEIL",
		Comment: "Round up to nearest integer for 64-bit float",
	},
	{
		Name:    "f64_FLOOR",
		Comment: "Round down to nearest integer for 64-bit float",
	},
	{
		Name:    "f64_TRUNC",
		Comment: "Round towards zero to nearest integer for 64-bit float",
	},
	{
		Name:    "f64_NEAREST",
		Comment: "Round to nearest integer, ties to even, for 64-bit float",
	},
	{
		Name:    "f64_SQRT",
		Comment: "Square root of 64-bit float",
	},
	{
		Name:    "f64_ADD",
		Comment: "Add two 64-bit floats",
	},
	{
		Name:    "f64_SUB",
		Comment: "Subtract one 64-bit float from another",
	},
	{
		Name:    "f64_MUL",
		Comment: "Multiply two 64-bit floats",
	},
	{
		Name:    "f64_DIV",
		Comment: "Divide one 64-bit float by another",
	},
	{
		Name:    "f64_MIN",
		Comment: "Minimum of two 64-bit floats",
	},
	{
		Name:    "f64_MAX",
		Comment: "Maximum of two 64-bit floats",
	},
	{
		Name:    "f64_COPYSIGN",
		Comment: "Copy sign from one 64-bit float to another",
	},
	{
		Name:    "i32_WRAP_i64",
		Comment: "Wrap 64-bit integer to 32-bit integer",
	},
	{
		Name:    "i32_TRUNC_f32_s",
		Comment: "Truncate signed 32-bit float to 32-bit integer",
	},
	{
		Name:    "i32_TRUNC_f32_u",
		Comment: "Truncate unsigned 32-bit float to 32-bit integer",
	},
	{
		Name:    "i32_TRUNC_f64_s",
		Comment: "Truncate signed 64-bit float to 32-bit integer",
	},
	{
		Name:    "i32_TRUNC_f64_u",
		Comment: "Truncate unsigned 64-bit float to 32-bit integer",
	},
	{
		Name:    "i64_EXTEND_i32_s",
		Comment: "Extend signed 32-bit integer to 64-bit integer",
	},
	{
		Name:    "i64_EXTEND_i32_u",
		Comment: "Extend unsigned 32-bit integer to 64-bit integer",
	},
	{
		Name:    "i64_TRUNC_f32_s",
		Comment: "Truncate signed 32-bit float to 64-bit integer",
	},
	{
		Name:    "i64_TRUNC_f32_u",
		Comment: "Truncate unsigned 32-bit float to 64-bit integer",
	},
	{
		Name:    "i64_TRUNC_f64_s",
		Comment: "Truncate signed 64-bit float to 64-bit integer",
	},
	{
		Name:    "i64_TRUNC_f64_u",
		Comment: "Truncate unsigned 64-bit float to 64-bit integer",
	},
	{
		Name:    "f32_CONVERT_i32_s",
		Comment: "Convert signed 32-bit integer to 32-bit float",
	},
	{
		Name:    "f32_CONVERT_i32_u",
		Comment: "Convert unsigned 32-bit integer to 32-bit float",
	},
	{
		Name:    "f32_CONVERT_i64_s",
		Comment: "Convert signed 64-bit integer to 32-bit float",
	},
	{
		Name:    "f32_CONVERT_i64_u",
		Comment: "Convert unsigned 64-bit integer to 32-bit float",
	},
	{
		Name:    "f32_DEMOTE_f64",
		Comment: "Demote 64-bit float to 32-bit float",
	},
	{
		Name:    "f64_CONVERT_i32_s",
		Comment: "Convert signed 32-bit integer to 64-bit float",
	},
	{
		Name:    "f64_CONVERT_i32_u",
		Comment: "Convert unsigned 32-bit integer to 64-bit float",
	},
	{
		Name:    "f64_CONVERT_i64_s",
		Comment: "Convert signed 64-bit integer to 64-bit float",
	},
	{
		Name:    "f64_CONVERT_i64_u",
		Comment: "Convert unsigned 64-bit integer to 64-bit float",
	},
	{
		Name:    "f64_PROMOTE_f32",
		Comment: "Promote 32-bit float to 64-bit float",
	},
	{
		Name:    "i32_REINTERPRET_f32",
		Comment: "Reinterpret bits of 32-bit float as 32-bit integer",
	},
	{
		Name:    "i64_REINTERPRET_f64",
		Comment: "Reinterpret bits of 64-bit float as 64-bit integer",
	},
	{
		Name:    "f32_REINTERPRET_i32",
		Comment: "Reinterpret bits of 32-bit integer as 32-bit float",
	},
	{
		Name:    "f64_REINTERPRET_i64",
		Comment: "Reinterpret bits of 64-bit integer as 64-bit float",
	},
	{
		Name:    "LOCAL_GET",
		Comment: "Get the value of a local variable",
	},
	{
		Name:    "LOCAL_SET",
		Comment: "Set the value of a local variable",
	},
	{
		Name:    "GLOBAL_GET",
		Comment: "Get the value of a global variable",
	},
	{
		Name:    "GLOBAL_SET",
		Comment: "Set the value of a global variable",
	},
	{
		Name:    "TEE",
		Comment: "Set the value of a local variable and keep the value on the stack",
	},
	{
		Name:    "UNREACHABLE",
		Comment: "Indicate unreachable code",
	},
	{
		Name:    "NOP",
		Comment: "No operation",
	},
	{
		Name:    "BLOCK",
		Comment: "Begin a sequence of instructions with an implied end label",
	},
	{
		Name:    "LOOP",
		Comment: "Begin a block that can be jumped to with an implicit label at the end",
	},
	{
		Name:    "IF",
		Comment: "Conditional branching",
	},
	{
		Name:    "ELSE",
		Comment: "Alternate block for conditional branching",
	},
	{
		Name:    "END",
		Comment: "End of a block, loop, or if instruction",
	},
	{
		Name:    "BR",
		Comment: "Branch to a given label",
	},
	{
		Name:    "BR_IF",
		Comment: "Conditional branch to a given label",
	},
	{
		Name:    "BR_TABLE",
		Comment: "Branch to a label in a table of labels",
	},
	{
		Name:    "RETURN",
		Comment: "Return from the current function",
	},
	{
		Name:    "CALL",
		Comment: "Call a function by index",
	},
	{
		Name:    "CALL_INDIRECT",
		Comment: "Call a function through an indirect function table",
	},
	{
		Name:    "DROP",
		Comment: "Remove the top value from the stack",
	},
	{
		Name:    "SELECT",
		Comment: "Select between two values based on a condition",
	},
}
