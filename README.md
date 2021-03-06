
YouTube Downloader
==================

A small server that serves YouTube videos and audio using `youtube-dl`. Because I am tired of all the fishy websites out there.

**Things to note**: 
* There currently is no nice frontend just a plain HTTP API.
* To run this application `youtube-dl` has to be installed somewhere in `$PATH`

Endpoints
---------

The following endpoints are available:

```
GET /<id>/video?format=<format>&filename=<filename>
```

```
GET /<id>/audio?format=<format>&filename=<filename>
```

If you want to know what formats are available for audio and video you have to look at the `youtube-dl` documentation. 

This application will serve the downloaded file directly and set the `Content-Disposition` header accordingly.

Errors will be returned in the following format

```
{
  "error": {
    "msg": "this is an error message"
  }
}
```

How to use
----------

Start the server by executing the following command

```
$ ripvid -address=<ip> -port=<port>
```

then just open your browser and head to `http://<ip>:<port>/BBJa32lCaaY/video?format=mp4&filename=ohmy`. The download should start as soon as the server has downloaded the video/audio using `youtube-dl`.
