# chop

Extracting the useful things I need from log files

Given an example file, `receipts.txt` such as this (but much longer)

	Receipt 2018-05-01
	Coffee 2.50
	Bagel 3.00
	Receipt 2018-05-14
	Tea 1.75
	Chocolate bar 0.60
	Receipt 2018-05-15
	Coffee 2.50
	Bun 0.75
	...

How would you extract the receipts that include Coffee? Assuming you are not
a `sed` or `awk` master

	  chomp --header Receipt --wanted Coffee receipts.txt

Which will return

	Receipt 2018-05-01
	Coffee 2.50
	Bagel 3.00
	
	Receipt 2018-05-15
	Coffee 2.50
	Bun 0.75

A bit dull really but damn useful when wading through multi gigabyte log files :)
