# headerquality

Headers media range relative quality indicator.

[![Build Status](https://travis-ci.org/linkosmos/headerquality.svg?branch=master)](https://travis-ci.org/linkosmos/headerquality)
[![GoDoc](http://godoc.org/github.com/linkosmos/headerquality?status.svg)](http://godoc.org/github.com/linkosmos/headerquality)
[![BSD License](http://img.shields.io/badge/license-BSD-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)

The values in the slice are sorted descending by qualification (the most qualified parameters will have the lowest indexes in the slice).
Unqualified parameters are interpreted as having the maximum qualification of `1`, as defined in the HTTP/1.1 specification.

### Usage

```go
func homeHandler(rw http.ResponseWriter, req *http.Request) {

  if languages := headerquality.Factors("Accept-Language", req); len(languages) > 0 {
    for _, language := range languages {
      if IsLocaleSupported(language.Factor) {
        // Set language for usage in views
      }
    }
  }
}
```

### RFC info:

[HTTP/1.1 Header Field Definitions](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html)

### Inspired by

> https://github.com/martini-contrib/acceptlang

### License

Copyright (c) 2015, linkosmos
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of headerquality nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
