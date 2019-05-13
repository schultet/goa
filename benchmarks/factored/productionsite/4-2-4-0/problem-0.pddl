(define (problem production-site)
  (:domain production-site)
  (:init
    (= (total-cost) 0)
    (not (completed p0))
    (not (completed p1))
  )
  (:goal
    (and
      (completed p0)
      (completed p1)
    )
  )
  (:metric minimize (total-cost))
)
