package main
/* Go Fuck Yourself, 
*      a BrainFuck interpreter in Go! ( http://golang.org )
*      by Glenn Franxman
*      written from the spec at http://en.wikipedia.org/wiki/Brainfuck
*
*/
import (
    "os" ; 
    "io" ;
)


const (
    MAX_PROG_LEN = 30000 ;
)


func main() {

    // for each prog
    for args := 1; args < len( os.Args ); args++ { 
        // Read in the source
        program, _ := io.ReadFile( os.Args[args] ) ;

        // prepare our vm
        data := make( []uint8, MAX_PROG_LEN ) ;
        data_ptr := 0;
        loop_depth := 0;
        instruction_ptr := 0;

        // execution loop
        for instruction_ptr = 0; instruction_ptr < len(program); instruction_ptr++ {

            // '>' increment the data pointer (to point to the next cell to the right).
            if program[instruction_ptr] == '>' {
                data_ptr++;
            }

            // '<' decrement the data pointer (to point to the next cell to the left).
            else if program[instruction_ptr] == '<' {
                data_ptr--;
            }


            // '+' increment (increase by one) the byte at the data pointer.
            else if program[instruction_ptr] == '+' {
                data[data_ptr]++;
            }

            // '-' decrement (decrease by one) the byte at the data pointer.
            else if program[instruction_ptr] == '-' {
                data[data_ptr]--;
            }

            // '.' output the value of the byte at the data pointer.
            else if program[instruction_ptr] == '.' {
                print( string( uint8(data[data_ptr]) ) ) ;
            }

            // ',' accept one byte of input, storing its value in the byte at the data pointer.
            else if program[instruction_ptr] == ',' {
                var b =  make( []uint8, 1 ) ;
                _,_ = os.Stdin.Read(b);
                data[data_ptr] = b[0];
            }

            // '[' if the byte at the data pointer is zero, then 
            //    instead of moving the instruction pointer forward to the next command, 
            //    jump it forward to the command after the matching ] command.
            // * interesting note -- the spec does not mention that the jump instructions should be nestable,
            //    but without this feature my test suite fails.
            else if program[instruction_ptr] == '[' {
                if data[data_ptr] == 0 {
                    instruction_ptr++;
                    // allow nested [ ] pairs by looping until we hit the end-jump for our loop depth
                    for loop_depth > 0 || program[instruction_ptr] != 93 {
                        if program[instruction_ptr] == '[' { 
                            loop_depth++; 
                        }
                        else if program[instruction_ptr] == ']' { 
                            loop_depth--; 
                        }
                        instruction_ptr++;
                    }
                }
            }

            // ']' if the byte at the data pointer is nonzero, 
            //    then instead of moving the instruction pointer forward to the next command, 
            //    jump it back to the command after the matching [ command.
            // * interesting note -- the spec calls for a check that the datapointer is pointing to a non-zero value,
            //    but doing so causes my test suite of bf programs to fail.
            else if program[instruction_ptr] == ']' {
                instruction_ptr--;
                // allow nested [ ] pairs by looping until we hit the end-jump for our loop depth
                for loop_depth > 0 || program[instruction_ptr] != '[' {
                        if program[instruction_ptr] == ']' { 
                            loop_depth++; 
                        }
                        else if program[instruction_ptr] == '[' { 
                            loop_depth--; 
                        }
                        instruction_ptr--;
                }
                instruction_ptr--;
            }

        }  
    print("\n");  
    }
}
