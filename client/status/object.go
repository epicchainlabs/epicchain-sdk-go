package apistatus

import (
	"errors"

	"github.com/epicchainlabs/epicchain-api-go/v2/object"
	"github.com/epicchainlabs/epicchain-api-go/v2/status"
)

var (
	// ErrObjectLocked is an instance of ObjectLocked error status. It's expected to be used for [errors.Is]
	// and MUST NOT be changed.
	ErrObjectLocked ObjectLocked
	// ErrObjectAlreadyRemoved is an instance of ObjectAlreadyRemoved error status. It's expected to be used for [errors.Is]
	// and MUST NOT be changed.
	ErrObjectAlreadyRemoved ObjectAlreadyRemoved
	// ErrLockNonRegularObject is an instance of LockNonRegularObject error status. It's expected to be used for [errors.Is]
	// and MUST NOT be changed.
	ErrLockNonRegularObject LockNonRegularObject
	// ErrObjectAccessDenied is an instance of ObjectAccessDenied error status. It's expected to be used for [errors.Is]
	// and MUST NOT be changed.
	ErrObjectAccessDenied ObjectAccessDenied
	// ErrObjectNotFound is an instance of ObjectNotFound error status. It's expected to be used for [errors.Is]
	// and MUST NOT be changed.
	ErrObjectNotFound ObjectNotFound
	// ErrObjectOutOfRange is an instance of ObjectOutOfRange error status. It's expected to be used for [errors.Is]
	// and MUST NOT be changed.
	ErrObjectOutOfRange ObjectOutOfRange
)

// ObjectLocked describes status of the failure because of the locked object.
// Instances provide [StatusV2] and error interfaces.
type ObjectLocked struct {
	v2 status.Status
}

const defaultObjectLockedMsg = "object is locked"

func (x ObjectLocked) Error() string {
	msg := x.v2.Message()
	if msg == "" {
		msg = defaultObjectLockedMsg
	}

	return errMessageStatusV2(
		globalizeCodeV2(object.StatusLocked, object.GlobalizeFail),
		msg,
	)
}

// Is implements interface for correct checking current error type with [errors.Is].
func (x ObjectLocked) Is(target error) bool {
	switch target.(type) {
	default:
		return errors.Is(Error, target)
	case ObjectLocked, *ObjectLocked:
		return true
	}
}

// implements local interface defined in [ErrorFromV2] func.
func (x *ObjectLocked) fromStatusV2(st *status.Status) {
	x.v2 = *st
}

// ErrorToV2 implements [StatusV2] interface method.
// If the value was returned by [ErrorFromV2], returns the source message.
// Otherwise, returns message with
//   - code: LOCKED;
//   - string message: "object is locked";
//   - details: empty.
func (x ObjectLocked) ErrorToV2() *status.Status {
	x.v2.SetCode(globalizeCodeV2(object.StatusLocked, object.GlobalizeFail))
	x.v2.SetMessage(defaultObjectLockedMsg)
	return &x.v2
}

// LockNonRegularObject describes status returned on locking the non-regular object.
// Instances provide [StatusV2] and error interfaces.
type LockNonRegularObject struct {
	v2 status.Status
}

const defaultLockNonRegularObjectMsg = "locking non-regular object is forbidden"

func (x LockNonRegularObject) Error() string {
	msg := x.v2.Message()
	if msg == "" {
		msg = defaultLockNonRegularObjectMsg
	}

	return errMessageStatusV2(
		globalizeCodeV2(object.StatusLockNonRegularObject, object.GlobalizeFail),
		msg,
	)
}

// Is implements interface for correct checking current error type with [errors.Is].
func (x LockNonRegularObject) Is(target error) bool {
	switch target.(type) {
	default:
		return errors.Is(Error, target)
	case LockNonRegularObject, *LockNonRegularObject:
		return true
	}
}

// implements local interface defined in [ErrorFromV2] func.
func (x *LockNonRegularObject) fromStatusV2(st *status.Status) {
	x.v2 = *st
}

// ErrorToV2 implements [StatusV2] interface method.
// If the value was returned by [ErrorFromV2], returns the source message.
// Otherwise, returns message with
//   - code: LOCK_NON_REGULAR_OBJECT;
//   - string message: "locking non-regular object is forbidden";
//   - details: empty.
func (x LockNonRegularObject) ErrorToV2() *status.Status {
	x.v2.SetCode(globalizeCodeV2(object.StatusLockNonRegularObject, object.GlobalizeFail))
	x.v2.SetMessage(defaultLockNonRegularObjectMsg)
	return &x.v2
}

// ObjectAccessDenied describes status of the failure because of the access control violation.
// Instances provide [StatusV2] and error interfaces.
type ObjectAccessDenied struct {
	v2 status.Status
}

const defaultObjectAccessDeniedMsg = "access to object operation denied"

func (x ObjectAccessDenied) Error() string {
	msg := x.v2.Message()
	if msg == "" {
		msg = defaultObjectAccessDeniedMsg
	}

	return errMessageStatusV2(
		globalizeCodeV2(object.StatusAccessDenied, object.GlobalizeFail),
		msg,
	)
}

// Is implements interface for correct checking current error type with [errors.Is].
func (x ObjectAccessDenied) Is(target error) bool {
	switch target.(type) {
	default:
		return errors.Is(Error, target)
	case ObjectAccessDenied, *ObjectAccessDenied:
		return true
	}
}

// implements local interface defined in [ErrorFromV2] func.
func (x *ObjectAccessDenied) fromStatusV2(st *status.Status) {
	x.v2 = *st
}

// ErrorToV2 implements [StatusV2] interface method.
// If the value was returned by [ErrorFromV2], returns the source message.
// Otherwise, returns message with
//   - code: ACCESS_DENIED;
//   - string message: "access to object operation denied";
//   - details: empty.
func (x ObjectAccessDenied) ErrorToV2() *status.Status {
	x.v2.SetCode(globalizeCodeV2(object.StatusAccessDenied, object.GlobalizeFail))
	x.v2.SetMessage(defaultObjectAccessDeniedMsg)
	return &x.v2
}

// WriteReason writes human-readable access rejection reason.
func (x *ObjectAccessDenied) WriteReason(reason string) {
	object.WriteAccessDeniedDesc(&x.v2, reason)
}

// Reason returns human-readable access rejection reason returned by the server.
// Returns empty value is reason is not presented.
func (x ObjectAccessDenied) Reason() string {
	return object.ReadAccessDeniedDesc(x.v2)
}

// ObjectNotFound describes status of the failure because of the missing object.
// Instances provide [StatusV2] and error interfaces.
type ObjectNotFound struct {
	v2 status.Status
}

const defaultObjectNotFoundMsg = "object not found"

func (x ObjectNotFound) Error() string {
	msg := x.v2.Message()
	if msg == "" {
		msg = defaultObjectNotFoundMsg
	}

	return errMessageStatusV2(
		globalizeCodeV2(object.StatusNotFound, object.GlobalizeFail),
		msg,
	)
}

// Is implements interface for correct checking current error type with [errors.Is].
func (x ObjectNotFound) Is(target error) bool {
	switch target.(type) {
	default:
		return errors.Is(Error, target)
	case ObjectNotFound, *ObjectNotFound:
		return true
	}
}

// implements local interface defined in [ErrorFromV2] func.
func (x *ObjectNotFound) fromStatusV2(st *status.Status) {
	x.v2 = *st
}

// ErrorToV2 implements [StatusV2] interface method.
// If the value was returned by [ErrorFromV2], returns the source message.
// Otherwise, returns message with
//   - code: OBJECT_NOT_FOUND;
//   - string message: "object not found";
//   - details: empty.
func (x ObjectNotFound) ErrorToV2() *status.Status {
	x.v2.SetCode(globalizeCodeV2(object.StatusNotFound, object.GlobalizeFail))
	x.v2.SetMessage(defaultObjectNotFoundMsg)
	return &x.v2
}

// ObjectAlreadyRemoved describes status of the failure because object has been
// already removed. Instances provide Status and StatusV2 interfaces.
type ObjectAlreadyRemoved struct {
	v2 status.Status
}

const defaultObjectAlreadyRemovedMsg = "object already removed"

func (x ObjectAlreadyRemoved) Error() string {
	msg := x.v2.Message()
	if msg == "" {
		msg = defaultObjectAlreadyRemovedMsg
	}

	return errMessageStatusV2(
		globalizeCodeV2(object.StatusAlreadyRemoved, object.GlobalizeFail),
		msg,
	)
}

// Is implements interface for correct checking current error type with [errors.Is].
func (x ObjectAlreadyRemoved) Is(target error) bool {
	switch target.(type) {
	default:
		return errors.Is(Error, target)
	case ObjectAlreadyRemoved, *ObjectAlreadyRemoved:
		return true
	}
}

// implements local interface defined in [ErrorFromV2] func.
func (x *ObjectAlreadyRemoved) fromStatusV2(st *status.Status) {
	x.v2 = *st
}

// ErrorToV2 implements [StatusV2] interface method.
// If the value was returned by [ErrorFromV2], returns the source message.
// Otherwise, returns message with
//   - code: OBJECT_ALREADY_REMOVED;
//   - string message: "object already removed";
//   - details: empty.
func (x ObjectAlreadyRemoved) ErrorToV2() *status.Status {
	x.v2.SetCode(globalizeCodeV2(object.StatusAlreadyRemoved, object.GlobalizeFail))
	x.v2.SetMessage(defaultObjectAlreadyRemovedMsg)
	return &x.v2
}

// ObjectOutOfRange describes status of the failure because of the incorrect
// provided object ranges.
// Instances provide [StatusV2] and error interfaces.
type ObjectOutOfRange struct {
	v2 status.Status
}

const defaultObjectOutOfRangeMsg = "out of range"

func (x ObjectOutOfRange) Error() string {
	msg := x.v2.Message()
	if msg == "" {
		msg = defaultObjectOutOfRangeMsg
	}

	return errMessageStatusV2(
		globalizeCodeV2(object.StatusOutOfRange, object.GlobalizeFail),
		msg,
	)
}

// Is implements interface for correct checking current error type with [errors.Is].
func (x ObjectOutOfRange) Is(target error) bool {
	switch target.(type) {
	default:
		return errors.Is(Error, target)
	case ObjectOutOfRange, *ObjectOutOfRange:
		return true
	}
}

// implements local interface defined in [ErrorFromV2] func.
func (x *ObjectOutOfRange) fromStatusV2(st *status.Status) {
	x.v2 = *st
}

// ErrorToV2 implements [StatusV2] interface method.
// If the value was returned by [ErrorFromV2], returns the source message.
// Otherwise, returns message with
//   - code: OUT_OF_RANGE;
//   - string message: "out of range";
//   - details: empty.
func (x ObjectOutOfRange) ErrorToV2() *status.Status {
	x.v2.SetCode(globalizeCodeV2(object.StatusOutOfRange, object.GlobalizeFail))
	x.v2.SetMessage(defaultObjectOutOfRangeMsg)
	return &x.v2
}
