(define (domain logistics)
    (:requirements :factored-privacy :typing)
    (:types location truck airplane package - object)
    (:predicates
    	  (p-at ?p - package ?l - location)
    
        (:private
            (at ?a - airplane ?l - location)
    	      (in ?p - package ?a - airplane)
        )
    )
    
    (:action load-airplane
    	:parameters (?a - airplane ?p - package ?l - location)
    	:precondition (and
    		(p-at ?p ?l)
    		(at ?a ?l)
    	)
    	:effect (and
    		(not (p-at ?p ?l))
    		(in ?p ?a)
    	)
    )
    
    (:action unload-airplane
    	:parameters (?a - airplane ?p - package ?l - location)
    	:precondition (and
    		(in ?p ?a)
    		(at ?a ?l)
    	)
    	:effect (and
    		(not (in ?p ?a))
    		(p-at ?p ?l)
    	)
    )
    
    (:action fly-airplane
    	:parameters (?a - airplane ?from - location ?to - location)
    	:precondition (and
    		(at ?a ?from)
        )
    	:effect (and
    		(not (at ?a ?from))
    		(at ?a ?to)
    	)
    )
)
