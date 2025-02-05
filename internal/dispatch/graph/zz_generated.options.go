// Code generated by github.com/ecordell/optgen. DO NOT EDIT.
package graph

import (
	defaults "github.com/creasty/defaults"
	helpers "github.com/ecordell/optgen/helpers"
)

type ConcurrencyLimitsOption func(c *ConcurrencyLimits)

// NewConcurrencyLimitsWithOptions creates a new ConcurrencyLimits with the passed in options set
func NewConcurrencyLimitsWithOptions(opts ...ConcurrencyLimitsOption) *ConcurrencyLimits {
	c := &ConcurrencyLimits{}
	for _, o := range opts {
		o(c)
	}
	return c
}

// NewConcurrencyLimitsWithOptionsAndDefaults creates a new ConcurrencyLimits with the passed in options set starting from the defaults
func NewConcurrencyLimitsWithOptionsAndDefaults(opts ...ConcurrencyLimitsOption) *ConcurrencyLimits {
	c := &ConcurrencyLimits{}
	defaults.MustSet(c)
	for _, o := range opts {
		o(c)
	}
	return c
}

// ToOption returns a new ConcurrencyLimitsOption that sets the values from the passed in ConcurrencyLimits
func (c *ConcurrencyLimits) ToOption() ConcurrencyLimitsOption {
	return func(to *ConcurrencyLimits) {
		to.Check = c.Check
		to.LookupResources = c.LookupResources
		to.ReachableResources = c.ReachableResources
		to.LookupSubjects = c.LookupSubjects
	}
}

// DebugMap returns a map form of ConcurrencyLimits for debugging
func (c ConcurrencyLimits) DebugMap() map[string]any {
	debugMap := map[string]any{}
	debugMap["Check"] = helpers.DebugValue(c.Check, false)
	debugMap["LookupResources"] = helpers.DebugValue(c.LookupResources, false)
	debugMap["ReachableResources"] = helpers.DebugValue(c.ReachableResources, false)
	debugMap["LookupSubjects"] = helpers.DebugValue(c.LookupSubjects, false)
	return debugMap
}

// ConcurrencyLimitsWithOptions configures an existing ConcurrencyLimits with the passed in options set
func ConcurrencyLimitsWithOptions(c *ConcurrencyLimits, opts ...ConcurrencyLimitsOption) *ConcurrencyLimits {
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithOptions configures the receiver ConcurrencyLimits with the passed in options set
func (c *ConcurrencyLimits) WithOptions(opts ...ConcurrencyLimitsOption) *ConcurrencyLimits {
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithCheck returns an option that can set Check on a ConcurrencyLimits
func WithCheck(check uint16) ConcurrencyLimitsOption {
	return func(c *ConcurrencyLimits) {
		c.Check = check
	}
}

// WithLookupResources returns an option that can set LookupResources on a ConcurrencyLimits
func WithLookupResources(lookupResources uint16) ConcurrencyLimitsOption {
	return func(c *ConcurrencyLimits) {
		c.LookupResources = lookupResources
	}
}

// WithReachableResources returns an option that can set ReachableResources on a ConcurrencyLimits
func WithReachableResources(reachableResources uint16) ConcurrencyLimitsOption {
	return func(c *ConcurrencyLimits) {
		c.ReachableResources = reachableResources
	}
}

// WithLookupSubjects returns an option that can set LookupSubjects on a ConcurrencyLimits
func WithLookupSubjects(lookupSubjects uint16) ConcurrencyLimitsOption {
	return func(c *ConcurrencyLimits) {
		c.LookupSubjects = lookupSubjects
	}
}
