package tufer

import (
  "fmt"
  "strings"
  )

  func PackArray( s []string )( p string ){
    for i := 0; i < len(s); i++ {
      if len(s[i]) > 0 {
        p += fmt.Sprintf( "%s ", s[i] )
      }
    }
    return
  }

  func Pack( s ...string )( p string) {

    return PackArray( s )

  }

  func Unpack( p string )( messages []string ) {
    trimmed := strings.TrimSpace(p)
    messages = strings.Split( trimmed, " ")
    return
  }
