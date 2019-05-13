(define (domain production-site)
  (:requirements :typing :factored-privacy)
  (:types
    product property - object
  )
  (:constants
    p0 p1 p2 p3 p4 - product
    prop0 prop1 prop2 prop3 prop4
     - property
  )
  (:predicates
    (:private
      (satisfies ?o - product ?p - property)
    )
    (completed ?o - product)
  )
  (:action Finish
    :parameters (?o - product)
    :precondition (and
      (not (completed ?o))
      (satisfies ?o prop0)
      (satisfies ?o prop1)
      (satisfies ?o prop2)
      (satisfies ?o prop3)
      (satisfies ?o prop4)

    )
    :effect (completed ?o)
  )
  (:action process
    :parameters (?o - product ?p - property)
    :precondition (and
      (not (satisfies ?o ?p))
    )
    :effect (and
      (satisfies ?o ?p)
    )
  )
)
