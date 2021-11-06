// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ozonmp/com_message_api/v1/com_message_api.proto

package com_message_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Message) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for From

	// no validation rules for To

	// no validation rules for Text

	if v, ok := interface{}(m.GetDatetime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "Datetime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}

// Validate checks the field values on CreateMessageV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMessageV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetFrom()) < 1 {
		return CreateMessageV1RequestValidationError{
			field:  "From",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetTo()) < 1 {
		return CreateMessageV1RequestValidationError{
			field:  "To",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Text

	if m.GetDatetime() == nil {
		return CreateMessageV1RequestValidationError{
			field:  "Datetime",
			reason: "value is required",
		}
	}

	return nil
}

// CreateMessageV1RequestValidationError is the validation error returned by
// CreateMessageV1Request.Validate if the designated constraints aren't met.
type CreateMessageV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMessageV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMessageV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMessageV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMessageV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMessageV1RequestValidationError) ErrorName() string {
	return "CreateMessageV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMessageV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMessageV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMessageV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMessageV1RequestValidationError{}

// Validate checks the field values on CreateMessageV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMessageV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMessageV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateMessageV1ResponseValidationError is the validation error returned by
// CreateMessageV1Response.Validate if the designated constraints aren't met.
type CreateMessageV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMessageV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMessageV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMessageV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMessageV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMessageV1ResponseValidationError) ErrorName() string {
	return "CreateMessageV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMessageV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMessageV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMessageV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMessageV1ResponseValidationError{}

// Validate checks the field values on DescribeMessageV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeMessageV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMessageId() <= 0 {
		return DescribeMessageV1RequestValidationError{
			field:  "MessageId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeMessageV1RequestValidationError is the validation error returned by
// DescribeMessageV1Request.Validate if the designated constraints aren't met.
type DescribeMessageV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeMessageV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeMessageV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeMessageV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeMessageV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeMessageV1RequestValidationError) ErrorName() string {
	return "DescribeMessageV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeMessageV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeMessageV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeMessageV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeMessageV1RequestValidationError{}

// Validate checks the field values on DescribeMessageV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeMessageV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeMessageV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeMessageV1ResponseValidationError is the validation error returned by
// DescribeMessageV1Response.Validate if the designated constraints aren't met.
type DescribeMessageV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeMessageV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeMessageV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeMessageV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeMessageV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeMessageV1ResponseValidationError) ErrorName() string {
	return "DescribeMessageV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeMessageV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeMessageV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeMessageV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeMessageV1ResponseValidationError{}

// Validate checks the field values on ListMessageV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListMessageV1Request) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListMessageV1RequestValidationError is the validation error returned by
// ListMessageV1Request.Validate if the designated constraints aren't met.
type ListMessageV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMessageV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMessageV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMessageV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMessageV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMessageV1RequestValidationError) ErrorName() string {
	return "ListMessageV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListMessageV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMessageV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMessageV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMessageV1RequestValidationError{}

// Validate checks the field values on ListMessageV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListMessageV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetValue() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMessageV1ResponseValidationError{
					field:  fmt.Sprintf("Value[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListMessageV1ResponseValidationError is the validation error returned by
// ListMessageV1Response.Validate if the designated constraints aren't met.
type ListMessageV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMessageV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMessageV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMessageV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMessageV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMessageV1ResponseValidationError) ErrorName() string {
	return "ListMessageV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListMessageV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMessageV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMessageV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMessageV1ResponseValidationError{}

// Validate checks the field values on RemoveMessageV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveMessageV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMessageId() <= 0 {
		return RemoveMessageV1RequestValidationError{
			field:  "MessageId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveMessageV1RequestValidationError is the validation error returned by
// RemoveMessageV1Request.Validate if the designated constraints aren't met.
type RemoveMessageV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveMessageV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveMessageV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveMessageV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveMessageV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveMessageV1RequestValidationError) ErrorName() string {
	return "RemoveMessageV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveMessageV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveMessageV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveMessageV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveMessageV1RequestValidationError{}

// Validate checks the field values on RemoveMessageV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveMessageV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	return nil
}

// RemoveMessageV1ResponseValidationError is the validation error returned by
// RemoveMessageV1Response.Validate if the designated constraints aren't met.
type RemoveMessageV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveMessageV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveMessageV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveMessageV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveMessageV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveMessageV1ResponseValidationError) ErrorName() string {
	return "RemoveMessageV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveMessageV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveMessageV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveMessageV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveMessageV1ResponseValidationError{}