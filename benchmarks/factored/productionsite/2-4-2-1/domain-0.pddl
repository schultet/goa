(define (domain production-site)
  (:requirements :typing :factored-privacy)
  (:types
    product property - object
  )
  (:constants
    p0 p1 p2 - product
    prop0 prop1
    alt0 - property
  )
  (:predicates
    (:private
      (satisfies ?o - product ?p - property)
    )
    (completed ?o - product)
  )
  (:functions
    (total-cost)
  )
  (:action Deliver
    :parameters (?o - product)
    :precondition (and
      (not (completed ?o))
      (satisfies ?o prop0)
      (satisfies ?o prop1)
      (not (satisfies ?o alt0))

    )
    :effect (and 
      (completed ?o)
      (increase (total-cost) 4)
    )
  )
  (:action refine
    :parameters (?o - product ?p - property)
    :precondition (and
      (not (satisfies ?o ?p))
    )
    :effect (and
      (satisfies ?o ?p)
    )
  )
)
