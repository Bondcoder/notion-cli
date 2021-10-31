module notion-cli/notion

go 1.17

replace notion-cli/utils => ./utils

replace notion-cli/blocks => ./blocks

replace notion-cli/users => ./users

require (
	github.com/gookit/color v1.5.0 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	golang.org/x/sys v0.0.0-20210330210617-4fbd30eecc44 // indirect
	notion-cli/blocks v0.0.0-00010101000000-000000000000 // indirect
	notion-cli/users v0.0.0-00010101000000-000000000000 // indirect
	notion-cli/utils v0.0.0-00010101000000-000000000000 // indirect
)
