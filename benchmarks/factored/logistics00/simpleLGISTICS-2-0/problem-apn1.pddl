(define (problem logistics-0) (:domain logistics)
(:objects
	p - package
	y - location
    x - location
    z - location

	(:private
		a1 - airplane
	)
)
(:init
	(p-at p z)
	(at a1 x)
)
(:goal
    (p-at p x)
)
)
