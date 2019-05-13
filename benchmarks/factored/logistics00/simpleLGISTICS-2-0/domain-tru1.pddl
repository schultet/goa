(define (domain logistics)
	(:requirements :factored-privacy :typing)
(:types
	 location truck airplane package - object 
 )
(:predicates
	(p-at ?p - package ?l - location)

	(:private
        (at ?t - truck ?l - location)
        (in ?p - package ?t - truck)
	)
)

(:action load-truck
	:parameters (?t - truck ?p - package ?l - location)
	:precondition (and
		(at ?t ?l)
		(p-at ?p ?l)
	)
	:effect (and
		(not (p-at ?p ?l))
		(in ?p ?t)
	)
)


(:action unload-truck
	:parameters (?t - truck ?p - package ?l - location)
	:precondition (and
		(at ?t ?l)
		(in ?p ?t)
	)
	:effect (and
		(not (in ?p ?t))
		(p-at ?p ?l)
	)
)


(:action drive-truck
	:parameters (?t - truck ?from - location ?to - location)
	:precondition (and
		(at ?t ?from)
	)
	:effect (and
		(not (at ?t ?from))
		(at ?t ?to)
	)
)

)
