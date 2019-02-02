package ingress

import "errors"

// ErrUnknownPod is returned when an unknown pod is mapped.
var ErrUnknownPod = errors.New("unknown pod id")

// ErrUnsupportedEpochDepth is returned when an unsupported epoch depth is
// received in an OrderFragmentMapping.
var ErrUnsupportedEpochDepth = errors.New("unsupported epoch depth")

// ErrInvalidNumberOfPods is returned when an insufficient number of pods are
// mapped.
var ErrInvalidNumberOfPods = errors.New("invalid number of pods")

// ErrInvalidNumberOfOrderFragments is returned when a pod is mapped to an
// insufficient number of order fragments, or too many order fragments.
var ErrInvalidNumberOfOrderFragments = errors.New("invalid number of order fragments")

// ErrInvalidOrderFragmentMapping is returned when an order fragment mapping is
// of an invalid length.
var ErrInvalidOrderFragmentMapping = errors.New("invalid order fragment mappings")

// ErrInvalidEpochDepth is returned when an invalid epoch depth is provided
// upon verification.
var ErrInvalidEpochDepth = errors.New("invalid epoch depth")

// ErrCannotOpenOrderFragments is returned when none of the pods were available
// to receive order fragments
var ErrCannotOpenOrderFragments = errors.New("cannot open order fragments: no pod received an order fragment")
