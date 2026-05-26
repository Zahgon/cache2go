/*
 * Simple caching library with expiration capabilities
 *     Copyright (c) 2013-2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE.txt
 */

package cache2go

import (
	"sync"
	"time"
)

// CacheItem is an individual cache item
// Parameter data contains the user-set value in the cache.
type CacheItem struct {
	sync.RWMutex

	// The item's key.
	key interface{}
	// The item's data.
	data interface{}
	// How long will the item live in the cache when not being accessed/kept alive.
	lifeSpan time.Duration

	// Creation timestamp.
	createdOn time.Time
	// Last access timestamp.
	accessedOn time.Time
	// How often the item was accessed.
	accessCount int64

	// Callback method triggered right before removing the item from the cache
	aboutToExpire []func(key interface{})
}

// NewCacheItem returns a newly created CacheItem.
// Parameter key is the item's cache-key.
// Parameter lifeSpan determines after which time period without an access the item
// will get removed from the cache.
// Parameter data is the item's value.
func NewCacheItem(key interface{}, lifeSpan time.Duration, data interface{}) *CacheItem {
	_ = "STUB: not implemented"
	return nil
}

// KeepAlive marks an item to be kept for another expireDuration period.
func (item *CacheItem) KeepAlive() { _ = "STUB: not implemented"; return }

// LifeSpan returns this item's expiration duration.
func (item *CacheItem) LifeSpan() time.Duration {
	_ = "STUB: not implemented"
	// immutable
	return *new(time.Duration)
}

// AccessedOn returns when this item was last accessed.
func (item *CacheItem) AccessedOn() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// CreatedOn returns when this item was added to the cache.
func (item *CacheItem) CreatedOn() time.Time {
	_ = "STUB: not implemented"
	// immutable
	return *new(time.Time)
}

// AccessCount returns how often this item has been accessed.
func (item *CacheItem) AccessCount() int64 { _ = "STUB: not implemented"; return 0 }

// Key returns the key of this cached item.
func (item *CacheItem) Key() interface{} {
	_ = "STUB: not implemented"
	// immutable
	return nil
}

// Data returns the value of this cached item.
func (item *CacheItem) Data() interface{} {
	_ = "STUB: not implemented"
	// immutable
	return nil
}

// SetAboutToExpireCallback configures a callback, which will be called right
// before the item is about to be removed from the cache.
func (item *CacheItem) SetAboutToExpireCallback(f func(interface{})) {
	_ = "STUB: not implemented"
	return
}

// AddAboutToExpireCallback appends a new callback to the AboutToExpire queue
func (item *CacheItem) AddAboutToExpireCallback(f func(interface{})) {
	_ = "STUB: not implemented"
	return
}

// RemoveAboutToExpireCallback empties the about to expire callback queue
func (item *CacheItem) RemoveAboutToExpireCallback() { _ = "STUB: not implemented"; return }
