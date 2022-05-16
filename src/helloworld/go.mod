module helloworld

go 1.18

require github.com/perocha/calculator v0.0.0
replace github.com/perocha/calculator => ../calculator

require github.com/perocha/mypackage v0.0.0
replace github.com/perocha/mypackage => ../mypackage

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)
