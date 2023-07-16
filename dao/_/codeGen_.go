package build

//go:generate cmd /c "dao-gen -d %cd%/../ -c %cd%/conf.toml"

//go:generate sh -c "dao-gen -d `pwd`/../ -c `pwd`/conf.toml"
