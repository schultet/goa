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
            (unused ?o - foo)
            (used ?o - foo)
        )
        (done) 
  )


  (:action finish
    :parameters ()
    :precondition (and
            (not (done))
            (used k)
            (used l)
            (used m)
            (used n)
            (used o)
            (used p)
            (used q)
            (used r)
            (used s)
            (used t)
            (used u)
            (used v)
            (used w)
            (used x)
            (used y)
            (used z)
    )
    :effect (and 
            (done)
    )
  )
  
  
  (:action use
    :parameters (?o - foo)
    :precondition (and
            (unused ?o)
    )
    :effect (and
            (used ?o)
            (not (unused ?o))
    )
  )
)
