# Govecamole

A [Go](https://golang.org) wrapper around [vecamole](https://github.com/zerebubuth/vecamole), which is a C wrapper around a small part of [Mapnik](http://mapnik.org).

## Motivation

Just playing around with Go, but hoping to move in the direction of a nice, easy package for vector tile rendering and raster fallback.

## Installation

First, install [vecamole](https://github.com/zerebubuth/vecamole). That's the only hard bit. From then, it should be as simple as:

```
go get github.com/zerebubut/govecamole
```

## Functionality

Currently outputs a single vector tile to `tile.pbf` containing a single node. Wow, much capability.
