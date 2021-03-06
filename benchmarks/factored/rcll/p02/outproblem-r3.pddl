(define (problem rcll-production-steps_task)
(:domain rcll-production-steps)
(:objects
	r1 r2 r3 - robot
	p10 - product
	silver_base_p10 grey_cap_p10 gate2_delivery_p10 - step
	p70 - product
	silver_base_p70 blue_ring_p70 orange_ring_p70 yellow_ring_p70 black_cap_p70 gate2_delivery_p70 - step
)
(:init
	(add-one zero one)
	(add-one one two)
	(add-one two three)
	(subtract three zero three)
	(subtract two zero two)
	(subtract one zero one)
	(subtract zero zero zero)
	(subtract three one two)
	(subtract two one one)
	(subtract one one zero)
	(subtract three two one)
	(subtract two two zero)
	(has-step p10 silver_base_p10)
	(has-step p10 grey_cap_p10)
	(has-step p10 gate2_delivery_p10)
	(initial-step silver_base_p10)
	(has-step p70 silver_base_p70)
	(has-step p70 blue_ring_p70)
	(has-step p70 orange_ring_p70)
	(has-step p70 yellow_ring_p70)
	(has-step p70 black_cap_p70)
	(has-step p70 gate2_delivery_p70)
	(initial-step silver_base_p70)
	(input-location cs1_in cs1)
	(input-location cs2_in cs2)
	(input-location rs1_in rs1)
	(input-location rs2_in rs2)
	(input-location ds_in ds)
	(output-location cs1_out cs1)
	(output-location cs2_out cs2)
	(output-location rs1_out rs1)
	(output-location rs2_out rs2)
	(output-location bs_out bs)
    (robot1-precedes r2)
    (robot1-precedes r3)
    (robot2-precedes r3)
	;(robot1-assigned-machine cs1)
	;(robot1-assigned-machine rs1)
	;(robot1-assigned-machine bs)
	;(robot2-assigned-machine cs2)
	;(robot2-assigned-machine rs2)
	;(robot2-assigned-machine bs)
	(robot3-at-init)
	(robot3-assigned-machine bs)
	(step-at-machine silver_base_p10 bs)
	(step-at-machine grey_cap_p10 cs1)
	(step-at-machine gate2_delivery_p10 ds)
	(step-precedes silver_base_p10 grey_cap_p10)
	(step-precedes grey_cap_p10 gate2_delivery_p10)
	(step-at-machine silver_base_p70 bs)
	(step-at-machine blue_ring_p70 rs2)
	(step-at-machine orange_ring_p70 rs2)
	(step-at-machine yellow_ring_p70 rs1)
	(step-at-machine black_cap_p70 cs2)
	(step-at-machine gate2_delivery_p70 ds)
	;(step-completed silver_base_p70)
	;(step-completed orange_ring_p70)
	(step-precedes silver_base_p70 blue_ring_p70)
	(step-precedes blue_ring_p70 orange_ring_p70)
	(step-precedes orange_ring_p70 yellow_ring_p70)
	(step-precedes yellow_ring_p70 black_cap_p70)
	(step-precedes black_cap_p70 gate2_delivery_p70)
	(material-required silver_base_p10 zero)
	(material-required grey_cap_p10 zero)
	(material-required gate2_delivery_p10 zero)
	(material-required blue_ring_p70 one)
	(material-required orange_ring_p70 zero)
	(material-required yellow_ring_p70 two)
	(material-required silver_base_p70 zero)
	(material-required black_cap_p70 zero)
	(material-required gate2_delivery_p70 zero)
	(material-stored rs2 zero)
	(material-stored rs1 zero)
	(material-stored cs2 zero)
	(material-stored cs1 zero)
	(material-stored ds zero)
	(material-stored bs zero)
	(= (path-length bs_in bs_out) 98)
	(= (path-length bs_in cs1_in) 562)
	(= (path-length bs_in cs1_out) 638)
	(= (path-length bs_in cs2_in) 500)
	(= (path-length bs_in cs2_out) 393)
	(= (path-length bs_in ds_in) 391)
	(= (path-length bs_in rs1_in) 191)
	(= (path-length bs_in rs1_out) 300)
	(= (path-length bs_in rs2_in) 212)
	(= (path-length bs_in rs2_out) 115)
	(= (path-length bs_out bs_in) 98)
	(= (path-length bs_out cs1_in) 615)
	(= (path-length bs_out cs1_out) 691)
	(= (path-length bs_out cs2_in) 553)
	(= (path-length bs_out cs2_out) 446)
	(= (path-length bs_out ds_in) 445)
	(= (path-length bs_out rs1_in) 244)
	(= (path-length bs_out rs1_out) 353)
	(= (path-length bs_out rs2_in) 286)
	(= (path-length bs_out rs2_out) 155)
	(= (path-length cs1_in bs_in) 562)
	(= (path-length cs1_in bs_out) 615)
	(= (path-length cs1_in cs1_out) 129)
	(= (path-length cs1_in cs2_in) 214)
	(= (path-length cs1_in cs2_out) 209)
	(= (path-length cs1_in ds_in) 269)
	(= (path-length cs1_in rs1_in) 418)
	(= (path-length cs1_in rs1_out) 312)
	(= (path-length cs1_in rs2_in) 470)
	(= (path-length cs1_in rs2_out) 584)
	(= (path-length cs1_out bs_in) 638)
	(= (path-length cs1_out bs_out) 691)
	(= (path-length cs1_out cs1_in) 129)
	(= (path-length cs1_out cs2_in) 210)
	(= (path-length cs1_out cs2_out) 290)
	(= (path-length cs1_out ds_in) 362)
	(= (path-length cs1_out rs1_in) 489)
	(= (path-length cs1_out rs1_out) 393)
	(= (path-length cs1_out rs2_in) 552)
	(= (path-length cs1_out rs2_out) 666)
	(= (path-length cs2_in bs_in) 500)
	(= (path-length cs2_in bs_out) 553)
	(= (path-length cs2_in cs1_in) 214)
	(= (path-length cs2_in cs1_out) 210)
	(= (path-length cs2_in cs2_out) 238)
	(= (path-length cs2_in ds_in) 329)
	(= (path-length cs2_in rs1_in) 352)
	(= (path-length cs2_in rs1_out) 353)
	(= (path-length cs2_in rs2_in) 511)
	(= (path-length cs2_in rs2_out) 596)
	(= (path-length cs2_out bs_in) 393)
	(= (path-length cs2_out bs_out) 446)
	(= (path-length cs2_out cs1_in) 209)
	(= (path-length cs2_out cs1_out) 290)
	(= (path-length cs2_out cs2_in) 238)
	(= (path-length cs2_out ds_in) 252)
	(= (path-length cs2_out rs1_in) 245)
	(= (path-length cs2_out rs1_out) 258)
	(= (path-length cs2_out rs2_in) 417)
	(= (path-length cs2_out rs2_out) 489)
	(= (path-length ds_in bs_in) 391)
	(= (path-length ds_in bs_out) 445)
	(= (path-length ds_in cs1_in) 269)
	(= (path-length ds_in cs1_out) 362)
	(= (path-length ds_in cs2_in) 329)
	(= (path-length ds_in cs2_out) 252)
	(= (path-length ds_in rs1_in) 254)
	(= (path-length ds_in rs1_out) 141)
	(= (path-length ds_in rs2_in) 292)
	(= (path-length ds_in rs2_out) 406)
	(= (path-length rs1_in bs_in) 191)
	(= (path-length rs1_in bs_out) 244)
	(= (path-length rs1_in cs1_in) 418)
	(= (path-length rs1_in cs1_out) 489)
	(= (path-length rs1_in cs2_in) 352)
	(= (path-length rs1_in cs2_out) 245)
	(= (path-length rs1_in ds_in) 254)
	(= (path-length rs1_in rs1_out) 260)
	(= (path-length rs1_in rs2_in) 321)
	(= (path-length rs1_in rs2_out) 287)
	(= (path-length rs1_out bs_in) 300)
	(= (path-length rs1_out bs_out) 353)
	(= (path-length rs1_out cs1_in) 312)
	(= (path-length rs1_out cs1_out) 393)
	(= (path-length rs1_out cs2_in) 353)
	(= (path-length rs1_out cs2_out) 258)
	(= (path-length rs1_out ds_in) 141)
	(= (path-length rs1_out rs1_in) 260)
	(= (path-length rs1_out rs2_in) 208)
	(= (path-length rs1_out rs2_out) 322)
	(= (path-length rs2_in bs_in) 212)
	(= (path-length rs2_in bs_out) 286)
	(= (path-length rs2_in cs1_in) 470)
	(= (path-length rs2_in cs1_out) 552)
	(= (path-length rs2_in cs2_in) 511)
	(= (path-length rs2_in cs2_out) 417)
	(= (path-length rs2_in ds_in) 292)
	(= (path-length rs2_in rs1_in) 321)
	(= (path-length rs2_in rs1_out) 208)
	(= (path-length rs2_in rs2_out) 142)
	(= (path-length rs2_out bs_in) 115)
	(= (path-length rs2_out bs_out) 155)
	(= (path-length rs2_out cs1_in) 584)
	(= (path-length rs2_out cs1_out) 666)
	(= (path-length rs2_out cs2_in) 596)
	(= (path-length rs2_out cs2_out) 489)
	(= (path-length rs2_out ds_in) 406)
	(= (path-length rs2_out rs1_in) 287)
	(= (path-length rs2_out rs1_out) 322)
	(= (path-length rs2_out rs2_in) 142)
	(= (path-length start bs_in) 73)
	(= (path-length start bs_out) 122)
	(= (path-length start cs1_in) 536)
	(= (path-length start cs1_out) 611)
	(= (path-length start cs2_in) 474)
	(= (path-length start cs2_out) 367)
	(= (path-length start ds_in) 365)
	(= (path-length start rs1_in) 165)
	(= (path-length start rs1_out) 274)
	(= (path-length start rs2_in) 266)
	(= (path-length start rs2_out) 169)
	(= (total-cost) 0)

	(= (path-length bs_in bs_in) 0)
	(= (path-length bs_out bs_out) 0)
	(= (path-length cs1_in cs1_in) 0)
	(= (path-length cs1_out cs1_out) 0)
	(= (path-length cs2_in cs2_in) 0)
	(= (path-length cs2_out cs2_out) 0)
	(= (path-length ds_in ds_in) 0)
	(= (path-length rs1_in rs1_in) 0)
	(= (path-length rs1_out rs1_out) 0)
	(= (path-length rs2_in rs2_in) 0)
	(= (path-length rs2_out rs2_out) 0)
	(= (path-length start start) 0)

	(= (path-length bs_in start) 0)
	(= (path-length bs_out start) 0)
	(= (path-length cs1_in start) 0)
	(= (path-length cs1_out start) 0)
	(= (path-length cs2_in start) 0)
	(= (path-length cs2_out start) 0)
	(= (path-length ds_in start) 0)
	(= (path-length rs1_in start) 0)
	(= (path-length rs1_out start) 0)
	(= (path-length rs2_in start) 0)
	(= (path-length rs2_out start) 0)
)
(:goal (and
	;(step-completed silver_base_p10)
	(step-completed grey_cap_p10)
	;(step-completed gate2_delivery_p10)
	;(step-completed silver_base_p70)
	(step-completed blue_ring_p70)
	;(step-completed orange_ring_p70)
	;(step-completed yellow_ring_p70)
	;(step-completed black_cap_p70)
	;(step-completed gate2_delivery_p70)
	;(cap-buffered cs1)
	;(cap-buffered cs2)
	;(robot-at r1 bs_in)
	;(robot-holding-product r1 p70)
	;(robot-holding-material r1)
	;(not (material-at cs2_out))
	;(not (material-at cs1_out))
)))
