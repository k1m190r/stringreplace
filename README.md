a small utility to find and replace strings in binary files

###Usage of stringreplace.exe:
* -find string: a string to find (default "q264")
* -replace string: a string to replace found string, must be same length > 0 (default "H264")
* -prefix string: a prefix to new file (default "M")


###Example:

stringreplace -find=ABCD -replace=XYZX -prefix=X *.bin

* this will open all bin files in current directory find string "ABCD" and replace it with "XYZX" add prefix "X" to file and save it as new X*.bin file.


###Installation:
Check under [releases](https://github.com/biosckon/stringreplace/releases)
download and move to some directory on your path "e.g. c:\Windows\" or "/usr/local/bin" (linux, MacOSX)

