package loginCounter

import "time"

type (
	/*
			we have a login mechanism that provides user_id every time a user logs in
		    you need to implement the function of counting the number of logins of the specified user for the last 5 minutes
		    you also need to implement a function that returns the user_id who has logged in the most times in the last 5 minutes
		    as it is already clear, the system stores data only for the last 5 minutes
		    the system must be responsive to load: more than 100 thousand operations per second
	*/
	Counter interface {
		login(id uint64)
		count(id uint64) (count int)
		maxLogged() uint64
	} //
)

const (
	startBandwidth = 100000
	actualTime     = time.Minute * 5
	bucketCount    = 100
	bucketTime     = actualTime / bucketCount
	bucketSize     = startBandwidth * (actualTime / time.Second) / bucketCount
)
