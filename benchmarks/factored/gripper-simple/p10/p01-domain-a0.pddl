;; Whatever
;;

(define (domain whatever)
  (:requirements :typing :factored-privacy)
  (:types
      foo - object
  )
  
  (:constants
    a b c d e f g h i j k l m n o p q r s t u v w x y z - foo
  )

  (:predicates
        (:private
            (a0-unused ?o - foo)
            (a0-used ?o - foo)
        )
        (done)
  )


  (:action a0-finish
    :parameters ()
    :precondition (and
            (not (done))
            (a0-used k)
            (a0-used l)
            (a0-used m)
            (a0-used n)
            (a0-used o)
            (a0-used p)
            (a0-used q)
            (a0-used r)
            (a0-used s)
            (a0-used t)
            (a0-used u)
            (a0-used v)
            (a0-used w)
            (a0-used x)
            (a0-used y)
            (a0-used z)
    )
    :effect (and 
            (done)
    )
  )
  
  
  (:action a0-use
    :parameters (?o - foo)
    :precondition (and
            (a0-unused ?o)
    )
    :effect (and
            (a0-used ?o)
            (not (a0-unused ?o))
    )
  )
)

