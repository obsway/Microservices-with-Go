package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Registry interface {

	// Register create a service instance record in the registry
	Register(ctx context.Context, instanceID string,
		serviceName string, hostPort string) error

	// Deregister remove a service instance from the registry
	Deregister(ctx context.Context, instanceID string, serviceName string) error

	// ServiceAddresses returns a list of the address of active instances of
	// the given service.
	ServiceAddresses(ctx context.Context, serviceID string) ([]string, error)

	// ReportHealthyState is a push mechanism for reporting healthy state to
	// the registry
	ReportHealthyState(instanceID string, serviceName string) error
}

// ErrNotFound is returned when no service addresses are found.
var ErrNotFound = errors.New("no_service addresses found")

func GenerateInstanceID(serviceName string) string {
	// time.Now().UnixNano(): return the current time with Nanosecond.
	// rand.Source: represent a source of randomness
	// rand.Rand: Wraps a Source to provide a richer set of operation
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
