(define (domain taxi)
	(:requirements :strips :factored-privacy :typing)
(:types
	 location agent - object 
 	 taxi passenger - agent 
 )
(:constants
	t1 - taxi
	p1 - passenger
)
(:predicates
	(directly-connected ?l1 - location ?l2 - location)
	(at ?a - agent ?l - location)
	(in ?p - passenger ?t - taxi)
	(empty ?t - taxi)
	(free ?l - location)
)

(:action drive_t1
	:parameters (?from - location ?to - location)
	:precondition (and
		(at t1 ?from)
		(directly-connected ?from ?to)
		(free ?to)
	)
	:effect (and
		(not (at t1 ?from))
		(not (free ?to))
		(at t1 ?to)
		(free ?from)
	)
)

)
