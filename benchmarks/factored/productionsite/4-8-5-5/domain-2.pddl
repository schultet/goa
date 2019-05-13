(define (domain production-site)
  (:requirements :typing :factored-privacy)
  (:types
    product property - object
  )
  (:constants
    p0 p1 p2 p3 p4 p5 p6 p7 - product
    prop0 prop1 prop2 prop3 prop4
    alt0 alt1 alt2 alt3 alt4 - property
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
      (satisfies ?o prop2)
      (satisfies ?o prop3)
      (satisfies ?o prop4)
      (not (satisfies ?o alt0))
      (not (satisfies ?o alt1))
      (not (satisfies ?o alt2))
      (not (satisfies ?o alt3))
      (not (satisfies ?o alt4))

    )
    :effect (and 
      (completed ?o)
      (increase (total-cost) 5)
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
