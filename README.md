# Configo

Configo is configuration parser for Go. It uses [TOML](https://github.com/toml-lang/toml) as config file format. There are some Go parser for TOML, but I prefer TOML parser library from [Naoya Inada](https://github.com/naoina/toml)
 
# Usage

Basically, you just create a TOML file and put it inside any location you like, example:

`conf/myproject.toml`

~~~
title="My project"
port=":8123"
~~~

Here's the Go source code to get the "port" configuration:

~~~go
package main

import (
	"fmt"
	"github.com/bpdp/configo"
)

type Config struct {
	Title string
	Port  string
}

func main() {

	var cnf Config

	if err := configo.ReadConfig("conf/lapmachine.toml", &cnf); err != nil {
		fmt.Println("Config Load Error: %g", err)
	}

	fmt.Println("Title: " + cnf.Title + ", Port: " + cnf.Port)

}

~~~

# License

~~~
Copyright 2015, Bambang Purnomosidi D. P.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
~~~

# Author

[Bambang Purnomosidi D. P.](http://bpdp.name)