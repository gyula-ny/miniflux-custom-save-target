# miniflux-custom-save-target

This is just a fake wallabag service (sorry, wallabag, it's me, not you) - to be able to save articles from the miniflux
rss reader to my custom backend (a [pocketbase](https://pocketbase.io/) instance) - with the miniflux "save" button. At a particular point of spacetime
there was one person (me) to whom this was useful and seemed the simplest. _**Strictly for internal, self-hosted, unauthenticated scenario**_

Expects environment variables - provided by, e.g. in a systemd unit file:

```[Unit]
Description=fake wallabag api

[Service]
Type=simple
Restart=always
RestartSec=5s
Environment=PORT=9999
Environment=TARGET_URL=YOUR CUSTOM TARGET URL THAT EXPECTS POST REQUEST WITH THE ARTICLE ITEM
Environment=GIN_MODE=release
ExecStart=PATH TO THE BINARY

[Install]
WantedBy=multi-user.target
