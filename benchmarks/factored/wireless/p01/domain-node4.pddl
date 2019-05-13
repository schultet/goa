(define (domain wireless)
	(:requirements :strips :factored-privacy :typing)
(:types
	 node level message - object 
 	 base sensor - node 
 )
(:constants
	base  - base
	node1 node2 node3 node4 node5  - sensor
	Zero Low Normal High  - level
)
(:predicates
	(not-message-at ?m - message ?n - node)
	(not-message-data ?m - message ?s - sensor)
	(not-sending ?from - sensor ?to - node ?m - message)
	(neighbor ?n1 - node ?n2 - node)
	(has-data ?n - node ?s - sensor)
	(higher ?l1 - level ?l2 - level)
	(next ?l1 - level ?l2 - level)
	(is-message-at ?m - message ?n - node)
	(message-data ?m - message ?s - sensor)
	(sending ?from - sensor ?to - node ?m - message)

	(:private
		(energy ?s - sensor ?lv - level)
	)
)

(:action generate-data_node4
	:parameters (?e0 - level ?e1 - level)
	:precondition (and
		(energy node4 ?e0)
		(higher ?e0 Zero)
		(next ?e0 ?e1)
	)
	:effect (and
		(not (energy node4 ?e0))
		(energy node4 ?e1)
		(has-data node4 node4)
	)
)


(:action add-to-message_node4
	:parameters (?s2 - sensor ?m - message)
	:precondition (and
		(has-data node4 ?s2)
		(is-message-at ?m node4)
		(not-message-data ?m ?s2)
	)
	:effect (and
		(not (has-data node4 ?s2))
		(not (not-message-data ?m ?s2))
		(message-data ?m ?s2)
	)
)


(:action send-message_node4
	:parameters (?receiver - node ?m - message ?e0 - level ?e1 - level)
	:precondition (and
		(energy node4 ?e0)
		(higher ?e0 Zero)
		(next ?e0 ?e1)
		(neighbor node4 ?receiver)
		(is-message-at ?m node4)
		(not-message-at ?m ?receiver)
		(not-sending node4 ?receiver ?m)
	)
	:effect (and
		(not (energy node4 ?e0))
		(not-message-at ?m node4)
		(not (is-message-at ?m node4))
		(energy node4 ?e1)
		(not (not-sending node4 ?receiver ?m))
		(sending node4 ?receiver ?m)
	)
)


(:action receive-message_node4
	:parameters (?sender - sensor ?m - message)
	:precondition (and
		(not-message-at ?m node4)
		(sending ?sender node4 ?m)
	)
	:effect (and
		(not-sending ?sender node4 ?m)
		(not (sending ?sender node4 ?m))
		(not (not-message-at ?m node4))
		(is-message-at ?m node4)
	)
)


(:action get-data-from-message_node4
	:parameters (?s - sensor ?m - message)
	:precondition (and
		(is-message-at ?m node4)
		(message-data ?m ?s)
	)
	:effect (and
		(not-message-data ?m ?s)
		(not (message-data ?m ?s))
		(has-data node4 ?s)
	)
)

)