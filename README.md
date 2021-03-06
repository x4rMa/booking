booking
============

Webapp booking with Golang

Usage
-----

Install Go with MacPorts

    sudo port install go

Setup Go workspace

    mkdir -p $HOME/go
    export GOPATH=$HOME/go
    export GO=$HOME/go

Get booking and CompileDaemon

    go get github.com/githubnemo/CompileDaemon
    go get github.com/bborbe/go/dirof
    go get github.com/bborbe/go/unique
    go get github.com/bborbe/booking/bin/booking_server

Start Dev-Mode with autorecompile and restart

    sh $GO/src/github.com/bborbe/booking/scripts/booking.sh devstart

Open webapp in browser

    open http://localhost:48567/

Make your own changes

    sublime $GO/src/github.com/bborbe/booking

Install Database
----------------
```
aptitude install postgresql
```

```
sudo su postgres -c 'psql'
```

```
CREATE USER booking WITH PASSWORD 'myPassword';
CREATE DATABASE booking WITH ENCODING 'utf8';
GRANT SELECT ON ALL TABLES IN SCHEMA public TO booking;
GRANT ALL PRIVILEGES ON DATABASE booking to booking;
```

Documentation
-------------

http://godoc.org/github.com/bborbe/booking

Copyright and license
---------------------

    Copyright (c) 2015, Benjamin Borbe <bborbe@rocketnews.de>
    All rights reserved.

    Redistribution and use in source and binary forms, with or without
    modification, are permitted provided that the following conditions are
    met:

       * Redistributions of source code must retain the above copyright
         notice, this list of conditions and the following disclaimer.
       * Redistributions in binary form must reproduce the above
         copyright notice, this list of conditions and the following
         disclaimer in the documentation and/or other materials provided
         with the distribution.

    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
    "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
    LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
    A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
    OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
    SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
    LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
    DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
    THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
    (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
    OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
