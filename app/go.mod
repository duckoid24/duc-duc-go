module check24.de/app

go 1.22.1

replace check24.de/hello => ../hello

replace check24.de/greeting => ../greeting

require check24.de/hello v0.0.0-00010101000000-000000000000

require check24.de/greeting v0.0.0-00010101000000-000000000000 // indirect
