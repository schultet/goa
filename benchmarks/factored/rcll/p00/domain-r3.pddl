(define (domain rcll-production-steps)
(:requirements 
    :strips
	:typing
	:factored-privacy
)

(:types
	robot - object
	location - object
	input output - location
	ds_input rs_input cs_input - input
	bs_output rs_output cs_output - output
	s_location - location
	machine - object
	base_station ring_station cap_station delivery_station - machine
	product - object
	step - object
	material_counter - object
)

(:constants
	; stations
	bs - base_station
	rs1 - ring_station
	rs2 - ring_station
	cs1 - cap_station
	cs2 - cap_station
	ds - delivery_station

	; locations
	start - s_location
	bs_in - bs_output
	bs_out - bs_output
	rs1_in - rs_input
	rs1_out - rs_output
	rs2_in - rs_input
	rs2_out - rs_output
	cs1_in - cs_input
	cs1_out - cs_output
	cs2_in - cs_input
	cs2_out - cs_output
	ds_in - ds_input

	; materials
	zero one two three - material_counter
)

(:predicates
	; stations
	(input-location ?il - input ?m - machine)
	(output-location ?ol - output ?m - machine)
	(conveyor-full ?m - machine)
	
	; ring stations
	(material-required ?s - step ?r - material_counter)
	(material-stored ?m - machine ?r - material_counter)
	(subtract ?minuend ?subtrahend ?difference - material_counter)
	(add-one ?summand ?sum - material_counter)

	; cap_station stations
	(cap-buffered ?m - cap_station)
	
	; steps
	(has-step ?p - product ?s - step)
	(step-completed ?s - step)
	(initial-step ?s - step)
	(step-precedes ?s1 ?s2 - step)
	(step-at-machine ?s - step ?m - machine)
	
	; products
	(product-at ?p - product ?l - location)
	
	; materials
	(material-at ?l - location)
	
	; locations
	(location-occupied ?l - location)

	(robot3-precedes ?r2 - robot)
    (:private
	    ; robots
	    (robot3-at ?l - location)
	    (robot3-at-init)
	    (robot3-holding-material)
	    (robot3-holding-product ?p - product)
	    (robot3-holding-something)
	    (robot3-recently-moved)
	    (robot3-assigned-machine ?m - machine)
    )
)

(:functions
	;(material-stored ?m - ring_station) - number
	;(material-required ?s - step) - number
	; paths
	(path-length ?l1 ?l2 - location) - number
	; cost
	(total-cost) - number
)

(:action dispense-material
	:parameters (?m - base_station ?o - bs_output)
	:precondition (and
		(not (conveyor-full ?m))
	)
	:effect (and
		(conveyor-full ?m)
		(material-at ?o)
		(increase (total-cost) 1)
	)
)

(:action dispense-product
	:parameters (?p - product ?s - step ?m - base_station ?o - bs_output)
	:precondition (and
		(has-step ?p ?s)
		(step-at-machine ?s ?m)
		(initial-step ?s)
		(not (step-completed ?s))
		(not (conveyor-full ?m))
	)
	:effect (and
		(conveyor-full ?m)
		(product-at ?p ?o)
		(step-completed ?s)
		(increase (total-cost) 1)
	)
)

(:action mount-ring
	:parameters (?m - ring_station ?p - product ?s1 ?s - step ?i - rs_input ?o - rs_output ?mi ?mr ?mf - material_counter)
	:precondition (and
		(product-at ?p ?i)
		(has-step ?p ?s)
		(step-at-machine ?s ?m)
		(not (step-completed ?s))
		(step-completed ?s1)
		(step-precedes ?s1 ?s)
		(input-location ?i ?m)
		(output-location ?o ?m)
		;(>= (material-stored ?m) (material-required ?s))
		(material-required ?s ?mr)
		(material-stored ?m ?mi)
		(subtract ?mi ?mr ?mf)
	)
	:effect (and
		(not (product-at ?p ?i))
		(product-at ?p ?o)
		(step-completed ?s)
		;(decrease (material-stored ?m) (material-required ?s))
		(not (material-stored ?m ?mi))
		(material-stored ?m ?mf)
		(increase (total-cost) 1)
	)
)

(:action buffer-cap
	:parameters (?m - cap_station ?i - cs_input ?o - cs_output)
	:precondition (and
		(input-location ?i ?m)
		(output-location ?o ?m)
		(material-at ?i)
		(not (cap-buffered ?m))
	)
	:effect (and
		(not (material-at ?i))
		(material-at ?o)
		(cap-buffered ?m)
		(increase (total-cost) 1)
	)
)

(:action mount-cap
	:parameters (?m - cap_station ?p - product ?s - step ?i - cs_input ?o - cs_output)
	:precondition (and
		(product-at ?p ?i)
		(has-step ?p ?s)
		(step-at-machine ?s ?m)
		(not (step-completed ?s))
		(input-location ?i ?m)
		(output-location ?o ?m)
		(cap-buffered ?m)
	)
	:effect (and
		(not (product-at ?p ?i))
		(product-at ?p ?o)
		(not (cap-buffered ?m))
		(step-completed ?s)
		(increase (total-cost) 1)
	)
)

(:action deliver
	:parameters (?p - product ?s - step ?m - delivery_station ?i - ds_input)
	:precondition (and
		(product-at ?p ?i)
		(has-step ?p ?s)
		(step-at-machine ?s ?m)
		(not (step-completed ?s))
	)
	:effect (and
		(not (product-at ?p ?i))
		(not (conveyor-full ?m))
		(step-completed ?s)
		(increase (total-cost) 1)
	)
)

(:action discard-material
	:parameters (?m - delivery_station ?i - ds_input)
	:precondition (and
		(material-at ?i)
	)
	:effect (and
		(not (material-at ?i))
		(not (conveyor-full ?m))
		(increase (total-cost) 1)
	)
)

(:action r3-insert-cap
	:parameters (?m - cap_station ?i - cs_input)
	:precondition (and
		(robot3-assigned-machine ?m)
		(robot3-at ?i)
		(not (conveyor-full ?m))
		(not (robot3-holding-something))
		(input-location ?i ?m)
		(not (cap-buffered ?m))
	)
	:effect (and
		(conveyor-full ?m)
		(material-at ?i)
		(not (robot3-recently-moved))
		(increase (total-cost) 30)
	)
)

(:action r3-pickup-material
	:parameters (?o - output ?m - machine)
	:precondition (and
		(robot3-assigned-machine ?m)
		(robot3-at ?o)
		(not (robot3-holding-something))
		(output-location ?o ?m)
		(material-at ?o)
	)
	:effect (and
		(robot3-holding-material)
		(robot3-holding-something)
		(not (material-at ?o))
		(not (conveyor-full ?m))
		(not (robot3-recently-moved))
		(increase (total-cost) 15)
	)
)

(:action r3-pickup-product
	:parameters (?o - output ?p - product ?m - machine)
	:precondition (and
		(robot3-assigned-machine ?m)
		(robot3-at ?o)
		(output-location ?o ?m)
		(not (robot3-holding-something))
		(product-at ?p ?o)
	)
	:effect (and
		(robot3-holding-product ?p)
		(robot3-holding-something)
		(not (product-at ?p ?o))
		(not (conveyor-full ?m))
		(not (robot3-recently-moved))
		(increase (total-cost) 15)
	)
)

(:action r3-insert-product
	:parameters (?i - input ?p - product ?m - machine)
	:precondition (and
		(robot3-at ?i)
		(input-location ?i ?m)
		(robot3-holding-product ?p)
		(not (conveyor-full ?m))
	)
	:effect (and
		(not (robot3-holding-product ?p))
		(product-at ?p ?i)
		(conveyor-full ?m)
		(not (robot3-recently-moved))
		(not (robot3-holding-something))
		(increase (total-cost) 15)
	)
)

(:action r3-insert-material
	:parameters (?i - rs_input ?m - ring_station ?mi ?mf - material_counter)
	:precondition (and
		(robot3-at ?i)
		(input-location ?i ?m)
		(robot3-holding-material)
		;(< (material-stored ?m) 3)
		(material-stored ?m ?mi)
		(add-one ?mi ?mf)
	)
	:effect (and
		(not (robot3-holding-material))
		;(increase (material-stored ?m) 1)
		(not (material-stored ?m ?mi))
		(material-stored ?m ?mf)
		(not (robot3-recently-moved))
		(not (robot3-holding-something))
		(increase (total-cost) 15)
	)
)

(:action r3-drop-material
	:parameters ()
	:precondition (and
		(robot3-holding-material)
	)
	:effect (and
		(not (robot3-holding-material))
		(not (robot3-holding-something))
		(not (robot3-recently-moved))
		;(increase (total-cost) 1)
	)
)

(:action r3-transport-material
	:parameters (?o - output ?i - rs_input ?m - ring_station)
	:precondition (and
		(input-location ?i ?m)
		(robot3-at ?o)
		(not (robot3-recently-moved))
		(robot3-holding-material)
		(not (location-occupied ?i))
	)
	:effect (and
		(not (robot3-at ?o))
		(not (material-at ?o))
		(not (location-occupied ?o))
		(location-occupied ?i)
		(robot3-at ?i)
		(robot3-recently-moved)
		(increase (total-cost) (path-length ?o ?i))
		;(increase (total-cost) 1)
	)
)

(:action r3-transport-product
	:parameters (?p - product ?o - output ?i - input ?m - machine ?s1 ?s2 - step)
	:precondition (and
		(robot3-holding-product ?p)
		(has-step ?p ?s1)
		(has-step ?p ?s2)
		(step-at-machine ?s2 ?m)
		(step-precedes ?s1 ?s2)
		(step-completed ?s1)
		(input-location ?i ?m)
		(not (step-completed ?s2))
		(robot3-at ?o)
		(not (robot3-recently-moved))
		(not (location-occupied ?i))
	)
	:effect (and
		(not (robot3-at ?o))
		(not (location-occupied ?o))
		(location-occupied ?i)
		(robot3-at ?i)
		(robot3-recently-moved)
		(increase (total-cost) (path-length ?o ?i))
		;(increase (total-cost) 1)
	)
)

(:action r3-move
	:parameters (?l1 ?l2 - location)
	:precondition (and
		(robot3-at ?l1)
		(not (robot3-holding-something))
		(not (robot3-recently-moved))
		(not (location-occupied ?l2))
	)
	:effect (and
		(not (robot3-at ?l1))
		(not (location-occupied ?l1))
		(location-occupied ?l2)
		(robot3-at ?l2)
		(robot3-recently-moved)
		(increase (total-cost) (path-length ?l1 ?l2))
		;(increase (total-cost) 15)
	)
)

(:action r3-move-in
	:parameters (?l - s_location)
	:precondition (and
		(robot3-at-init)
		;(not (exists (?_r - robot) (and (not (robot3-precedes ?_r)) (robot3-at-init))))
		(not (location-occupied ?l))
	)
	:effect (and
		(not (robot3-at-init))
		(location-occupied ?l)
		(robot3-at ?l)
		(increase (total-cost) 10)
	)
)
)

