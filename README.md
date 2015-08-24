# Govecamole

A [Go](https://golang.org) wrapper library around [vecamole](https://github.com/zerebubuth/vecamole), which is a C wrapper around a small part of [Mapnik](http://mapnik.org).

## Motivation

Just playing around with Go, but hoping to move in the direction of a nice, easy library for vector tile rendering and raster fallback.

## Installation

First, install [vecamole](https://github.com/zerebubuth/vecamole). That's the only hard bit. From then, it should be as simple as:

```
go get github.com/zerebubut/govecamole
```

## Functionality

Small wrapper API, which can currently load a Mapnik `Map` object, configure it and render vector tiles using it.
