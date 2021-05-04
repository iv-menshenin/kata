/*
  This is a test of pointers to complex data structures,
  it may seem to someone that in this case it will not matter whether we pass the value by reference or by value.
  However, tests still show that when using a reference to a value, the runtime needs 1 more memory allocation.
*/
package main

import (
	"github.com/streadway/amqp"
	"time"
)

type (
	some1 struct {
		RequestID int64
	}
	some2 struct {
		RequestID int64
	}
)

func (v *some1) MarshalJSON() ([]byte, error) {
	return []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nil
}

// additional memory allocation is required here
func (v *some1) MarshalAMQP() (*amqp.Publishing, error) {
	body, err := v.MarshalJSON()
	if err != nil {
		return nil, err
	}
	pub := &amqp.Publishing{
		DeliveryMode:    amqp.Persistent,
		AppId:           "someunit",
		Type:            "someunit",
		ContentType:     "application/json",
		ContentEncoding: "UTF-8",
		Timestamp:       time.Now(),
		Body:            body,
		Headers: amqp.Table{
			"Request-ID": v.RequestID,
		},
	}
	return pub, nil
}

func (v *some2) MarshalJSON() ([]byte, error) {
	return []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nil
}

// no additional memory allocation
func (v *some2) MarshalAMQP() (pub amqp.Publishing, err error) {
	var body []byte
	if body, err = v.MarshalJSON(); err != nil {
		return
	}
	return amqp.Publishing{
		DeliveryMode:    amqp.Persistent,
		AppId:           "someunit",
		Type:            "someunit",
		ContentType:     "application/json",
		ContentEncoding: "UTF-8",
		Timestamp:       time.Now(),
		Body:            body,
		Headers: amqp.Table{
			"Request-ID": v.RequestID,
		},
	}, nil
}
