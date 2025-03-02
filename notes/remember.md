lazy nushell start
go run (ls **/*.go | get name | str join " ")