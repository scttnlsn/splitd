splitd
======

Splits newline-separated HTTP response body into multiple HTTP requests.

Makes a long-held HTTP request to specified URL reading newline separated messages from the response body.  For each message recieved, a separate HTTP request is made to a designated URL (passing along the message).  Intended to be used in conjunction with with [broadcastd](http://github.com/scttnlsn/broadcastd) and [queued](http://github.com/scttnlsn/queued).

Getting Started
---------------

**Install:**

Ensure Go is installed and then run:

    $ go get
    $ make
    $ sudo make install

**Run:**

    $ splitd -source-url="http://localhost:5454" -dest-url="http://localhost:5353/messages" [options]

## CLI Options

* **-source-url="..."** - URL from which to read messages (required)
* **-source-auth=""** - HTTP basic auth password required for source requests
* **-source-method="GET"** - HTTP request method for source requests
* **-dest-url="..."** - URL to which messages are sent (required)
* **-dest-auth=""** - HTTP basic auth password required for destination requests
* **-dest-method="POST"** - HTTP request method for destination requests