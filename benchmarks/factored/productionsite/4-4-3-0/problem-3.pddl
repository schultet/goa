(define (problem production-site)
  (:domain production-site)
  (:init
    (= (total-cost) 0)
    (not (completed p0))
    (not (completed p1))
    (not (completed p2))
    (not (completed p3))
  )
  (:goal
    (and
      (completed p0)
      (completed p1)
      (completed p2)
      (completed p3)
    )
  )
  (:metric minimize (total-cost))
)
