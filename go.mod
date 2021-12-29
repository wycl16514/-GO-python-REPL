module go_python_compiler

replace repl => ./repl

replace token => ./token

replace lexer => ./lexer

go 1.17

require repl v0.0.0-00010101000000-000000000000

require (
	lexer v0.0.0-00010101000000-000000000000 // indirect
	token v0.0.0-00010101000000-000000000000 // indirect
)
