a small utility to find and replace strings in binary files

###Usage of stringreplace.exe:
  -find string
        string to find (default "q264")
  -prefix string
        prefix to new file (default "M")
  -replace string
        string to replace found string, must be same length > 0 (default "H264")


###Example:

stringreplace -find=ABCD -replace=XYZX -prefix=X *.bin

- this will open all bin files in current directory find string "ABCD" and replace it with "XYZX" add prefix "X" to file and save it as new X*.bin file.
