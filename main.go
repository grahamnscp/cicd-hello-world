/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START all]
package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "time"
)

var Version string
var Buildtime string
var Hostname, oserr = os.Hostname()


func main() {

  go func() {
    for range time.NewTicker(time.Second).C {
      log.Printf("This is a heartbeat..." + time.Now().String())
    }
  }()

  // use PORT environment variable, or default to 8080
  port := "8080"
  if fromEnv := os.Getenv("PORT"); fromEnv != "" {
    port = fromEnv
  }


  // register functions to handle all requests
  server := http.NewServeMux()
  server.HandleFunc("/", index)
  server.HandleFunc("/d2iq", d2iq)
  server.HandleFunc("/d2iq/", d2iq)
  server.HandleFunc("/docker", docker)
  server.HandleFunc("/docker/", docker)

  // start the web server on port and accept requests
  log.Printf("Server listening on port %s", port)
  err := http.ListenAndServe(":"+port, server)
  log.Fatal(err)
}

// hello responds to the request with a plain-text "Hello, world" message.
func index(w http.ResponseWriter, r *http.Request) {
  log.Printf("Serving request: %s", r.URL.Path)
  fmt.Fprint(w,"\n")
  fmt.Fprint(w," ____ ____  _  ___\n")
  fmt.Fprint(w,"|  _ \\___ \\(_)/ _ \\\n")
  fmt.Fprint(w,"| | | |__) | | | | |\n")
  fmt.Fprint(w,"| |_| / __/| | |_| |\n")
  fmt.Fprint(w,"|____/_____|_|\\__\\_\\\n")
  fmt.Fprint(w,"\n")
  fmt.Fprintf(w, "Hello!, from D2iQ Dispatch CI/CD GitOps!\n")
  fmt.Fprintf(w, "Version: %s\n", Version)
  fmt.Fprintf(w, "Build time: %s\n", Buildtime)
  fmt.Fprintf(w, "App Version: 20200317-1\n")
  fmt.Fprint(w,"\n")
  fmt.Fprint(w,  "Container hostname: ", Hostname, "\n")
  dt := time.Now()
  fmt.Fprint(w,  "Time: ", dt.Format("2006-01-02 15:04:05.00"), "\n")
}
func d2iq(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"\n")
  fmt.Fprint(w," _           _    _            ____ ____  _  ___\n")
  fmt.Fprint(w,"| |     ___ | |  | |    ___   |  _ \\___ \\(_)/ _ \\\n")
  fmt.Fprint(w,"| |___ / _ \\| |  | |   / _ \\  | | | |__) | | | | |\n")
  fmt.Fprint(w,"|  _  |  __/| |__| |__| (_) | | |_| / __/| | |_| |\n")
  fmt.Fprint(w,"|_| |_|\\___/ \\___|\\___|\\___/  |____/_____|_|\\__\\_\\\n")
  fmt.Fprint(w,"\n")
  dt := time.Now()
  fmt.Fprint(w,"[", dt.Format("2006-01-02 15:04:05.00"), "] ", "Container hostname: ", Hostname, "\n")
}
func docker(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"\n")
  fmt.Fprint(w,"                              ##\n")
  fmt.Fprint(w,"                        ## ## ##        ==\n")
  fmt.Fprint(w,"                     ## ## ## ## ##    ===\n")
  fmt.Fprint(w,"                 /`````````````````\\___/ ===\n")
  fmt.Fprint(w,"            ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~~/~ === ~~~\n")
  fmt.Fprint(w,"                 \\______ o           __/\n")
  fmt.Fprint(w,"                   \\    \\         __/\n")
  fmt.Fprint(w,"                    \\____\\_______/\n")
  fmt.Fprint(w," _           _    _                _            _\n")
  fmt.Fprint(w,"| |     ___ | |  | |    ___     __| | ___   ___| | _____ _ __\n")
  fmt.Fprint(w,"| |___ / _ \\| |  | |   / _ \\   / _  |/ _ \\ / __| |/ / _ \\ '__|\n")
  fmt.Fprint(w,"|  _  |  __/| |__| |__| (_) | | (_| | (_) | (__|   <  __/ |\n")
  fmt.Fprint(w,"|_| |_|\\___/ \\___|\\___|\\___/   \\____|\\___/ \\___|_|\\_\\___|_|\n")
  fmt.Fprint(w,"\n")
  dt := time.Now()
  fmt.Fprint(w,"[", dt.Format("2006-01-02 15:04:05.00"), "] ", "Container hostname: ", Hostname, "\n")
}

// [END all]
