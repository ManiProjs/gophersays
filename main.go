package gophersays

import (
	"fmt"
)

func say(msg string) {
	fmt.Printf(`
	<%s>
	 \
	  \
	   \
						,_---~~~~~----._         
		_,,_,*^____      _____ *g*\"*, 
		/ __/ /'     ^.  /      \ ^@q   f 
		[  @f | @))    |  | @))   l  0 _/  
		\ /   \~____ / __ \_____/    \   
		|           _l__l_           I   
		}          [______]           I  
		]            | | |            |  
	    ]             ~ ~             |  
	    |                             |   
        |                             |
	`, msg)
	fmt.Println("")
}
