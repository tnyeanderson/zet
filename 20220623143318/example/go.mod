module mymodule

go 1.18

require commonmodule v0.0.0

require childmodule v0.0.0

replace commonmodule v0.0.0 => ./commonmodule

replace childmodule v0.0.0 => ./childmodule
