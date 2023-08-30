package gooseneck

import (
	"github.com/patrickmn/go-cache"
)

type Cache = cache.Cache

var NewCache = cache.New
