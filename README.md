## ascii-art

# description
    this is a simple go program that receives a string as an argument and outputs the string in a graphic represeantion using ASCII

# packages
    only standard go packages are allowed
    
## installation

    clone into this repo by entering this into the terminal

    git clone https://learn.zone01kisumu.ke/git/cliffootieno/ascii-art

## Running the progam

# usage

    student$ go run . "" | cat -e
    student$ go run . "\n" | cat -e
    $
    student$ go run . "Hello\n" | cat -e
     _    _          _   _          $
    | |  | |        | | | |         $
    | |__| |   ___  | | | |   ___   $
    |  __  |  / _ \ | | | |  / _ \  $
    | |  | | |  __/ | | | | | (_) | $
    |_|  |_|  \___| |_| |_|  \___/  $
                                    $
                                    $
    $  
    student$ go run . "hello" | cat -e
     _              _   _          $
    | |            | | | |         $
    | |__     ___  | | | |   ___   $
    |  _ \   / _ \ | | | |  / _ \  $
    | | | | |  __/ | | | | | (_) | $
    |_| |_|  \___| |_| |_|  \___/  $
                                   $
                                   $
## things to note
    using \\(double slashes) in a string represents escaping a character for special uses such as \\n for newline which is the only special case handled here.
    Therefore use "\\" lightly