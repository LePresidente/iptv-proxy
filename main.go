/*
 * Iptv-Proxy is a project to proxyfie an m3u file and to proxyfie an Xtream iptv service (client API).
 * Copyright (C) 2020  Pierre-Emmanuel Jacquier
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import "github.com/pierre-emmanuelJ/iptv-proxy/cmd"
import (
   "fmt"
   "time"
   "log"	
)

func main() {
	cmd.Execute()
   // manually set time zone
   if tz := os.Getenv("TZ"); tz != "" {
      var err error
      time.Local, err = time.LoadLocation(tz)
      if err != nil {
         log.Printf("error loading location '%s': %v\n", tz, err)
      }
   }

   // output current time zone
   fmt.Print("Local time zone ")
   fmt.Println(time.Now().Zone())
   fmt.Println(time.Now().Format("2006-01-02T15:04:05.000 MST"))
}
