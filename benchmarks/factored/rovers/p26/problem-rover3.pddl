(define (problem roverprob3965) (:domain rover)
(:objects
	rover7store - store
	rover3store - store
	rover4store - store
	rover9store - store
	high_res - mode
	rover2store - store
	low_res - mode
	rover5store - store
	objective1 - objective
	objective0 - objective
	objective3 - objective
	objective2 - objective
	objective5 - objective
	objective4 - objective
	camera0 - camera
	rover8store - store
	camera1 - camera
	waypoint12 - waypoint
	camera3 - camera
	camera2 - camera
	rover0store - store
	general - lander
	camera6 - camera
	rover6store - store
	waypoint25 - waypoint
	waypoint24 - waypoint
	waypoint27 - waypoint
	waypoint26 - waypoint
	waypoint21 - waypoint
	waypoint20 - waypoint
	waypoint23 - waypoint
	waypoint22 - waypoint
	waypoint29 - waypoint
	waypoint28 - waypoint
	camera4 - camera
	waypoint8 - waypoint
	waypoint9 - waypoint
	waypoint6 - waypoint
	waypoint7 - waypoint
	waypoint4 - waypoint
	waypoint5 - waypoint
	waypoint2 - waypoint
	waypoint3 - waypoint
	waypoint0 - waypoint
	waypoint1 - waypoint
	waypoint18 - waypoint
	waypoint19 - waypoint
	camera5 - camera
	waypoint10 - waypoint
	waypoint11 - waypoint
	colour - mode
	waypoint13 - waypoint
	waypoint14 - waypoint
	waypoint15 - waypoint
	waypoint16 - waypoint
	waypoint17 - waypoint
	rover1store - store

	(:private
		rover3 - rover
	)
)
(:init
	(visible waypoint0 waypoint11)
	(visible waypoint11 waypoint0)
	(visible waypoint0 waypoint29)
	(visible waypoint29 waypoint0)
	(visible waypoint1 waypoint0)
	(visible waypoint0 waypoint1)
	(visible waypoint1 waypoint4)
	(visible waypoint4 waypoint1)
	(visible waypoint1 waypoint6)
	(visible waypoint6 waypoint1)
	(visible waypoint1 waypoint15)
	(visible waypoint15 waypoint1)
	(visible waypoint1 waypoint25)
	(visible waypoint25 waypoint1)
	(visible waypoint2 waypoint9)
	(visible waypoint9 waypoint2)
	(visible waypoint2 waypoint16)
	(visible waypoint16 waypoint2)
	(visible waypoint2 waypoint24)
	(visible waypoint24 waypoint2)
	(visible waypoint3 waypoint10)
	(visible waypoint10 waypoint3)
	(visible waypoint3 waypoint13)
	(visible waypoint13 waypoint3)
	(visible waypoint3 waypoint26)
	(visible waypoint26 waypoint3)
	(visible waypoint3 waypoint27)
	(visible waypoint27 waypoint3)
	(visible waypoint4 waypoint6)
	(visible waypoint6 waypoint4)
	(visible waypoint4 waypoint9)
	(visible waypoint9 waypoint4)
	(visible waypoint4 waypoint10)
	(visible waypoint10 waypoint4)
	(visible waypoint5 waypoint1)
	(visible waypoint1 waypoint5)
	(visible waypoint5 waypoint6)
	(visible waypoint6 waypoint5)
	(visible waypoint5 waypoint11)
	(visible waypoint11 waypoint5)
	(visible waypoint5 waypoint14)
	(visible waypoint14 waypoint5)
	(visible waypoint5 waypoint19)
	(visible waypoint19 waypoint5)
	(visible waypoint5 waypoint23)
	(visible waypoint23 waypoint5)
	(visible waypoint6 waypoint11)
	(visible waypoint11 waypoint6)
	(visible waypoint6 waypoint19)
	(visible waypoint19 waypoint6)
	(visible waypoint6 waypoint22)
	(visible waypoint22 waypoint6)
	(visible waypoint7 waypoint5)
	(visible waypoint5 waypoint7)
	(visible waypoint7 waypoint6)
	(visible waypoint6 waypoint7)
	(visible waypoint7 waypoint14)
	(visible waypoint14 waypoint7)
	(visible waypoint7 waypoint20)
	(visible waypoint20 waypoint7)
	(visible waypoint7 waypoint21)
	(visible waypoint21 waypoint7)
	(visible waypoint7 waypoint23)
	(visible waypoint23 waypoint7)
	(visible waypoint7 waypoint27)
	(visible waypoint27 waypoint7)
	(visible waypoint8 waypoint0)
	(visible waypoint0 waypoint8)
	(visible waypoint8 waypoint2)
	(visible waypoint2 waypoint8)
	(visible waypoint8 waypoint14)
	(visible waypoint14 waypoint8)
	(visible waypoint8 waypoint15)
	(visible waypoint15 waypoint8)
	(visible waypoint8 waypoint25)
	(visible waypoint25 waypoint8)
	(visible waypoint8 waypoint27)
	(visible waypoint27 waypoint8)
	(visible waypoint9 waypoint10)
	(visible waypoint10 waypoint9)
	(visible waypoint9 waypoint24)
	(visible waypoint24 waypoint9)
	(visible waypoint10 waypoint8)
	(visible waypoint8 waypoint10)
	(visible waypoint10 waypoint15)
	(visible waypoint15 waypoint10)
	(visible waypoint10 waypoint20)
	(visible waypoint20 waypoint10)
	(visible waypoint10 waypoint21)
	(visible waypoint21 waypoint10)
	(visible waypoint10 waypoint23)
	(visible waypoint23 waypoint10)
	(visible waypoint11 waypoint16)
	(visible waypoint16 waypoint11)
	(visible waypoint11 waypoint18)
	(visible waypoint18 waypoint11)
	(visible waypoint12 waypoint1)
	(visible waypoint1 waypoint12)
	(visible waypoint12 waypoint4)
	(visible waypoint4 waypoint12)
	(visible waypoint12 waypoint6)
	(visible waypoint6 waypoint12)
	(visible waypoint12 waypoint9)
	(visible waypoint9 waypoint12)
	(visible waypoint12 waypoint16)
	(visible waypoint16 waypoint12)
	(visible waypoint12 waypoint23)
	(visible waypoint23 waypoint12)
	(visible waypoint13 waypoint5)
	(visible waypoint5 waypoint13)
	(visible waypoint13 waypoint20)
	(visible waypoint20 waypoint13)
	(visible waypoint14 waypoint6)
	(visible waypoint6 waypoint14)
	(visible waypoint14 waypoint10)
	(visible waypoint10 waypoint14)
	(visible waypoint14 waypoint21)
	(visible waypoint21 waypoint14)
	(visible waypoint15 waypoint22)
	(visible waypoint22 waypoint15)
	(visible waypoint15 waypoint23)
	(visible waypoint23 waypoint15)
	(visible waypoint15 waypoint27)
	(visible waypoint27 waypoint15)
	(visible waypoint15 waypoint28)
	(visible waypoint28 waypoint15)
	(visible waypoint16 waypoint9)
	(visible waypoint9 waypoint16)
	(visible waypoint16 waypoint14)
	(visible waypoint14 waypoint16)
	(visible waypoint17 waypoint10)
	(visible waypoint10 waypoint17)
	(visible waypoint17 waypoint24)
	(visible waypoint24 waypoint17)
	(visible waypoint17 waypoint25)
	(visible waypoint25 waypoint17)
	(visible waypoint18 waypoint10)
	(visible waypoint10 waypoint18)
	(visible waypoint18 waypoint14)
	(visible waypoint14 waypoint18)
	(visible waypoint18 waypoint16)
	(visible waypoint16 waypoint18)
	(visible waypoint18 waypoint27)
	(visible waypoint27 waypoint18)
	(visible waypoint19 waypoint7)
	(visible waypoint7 waypoint19)
	(visible waypoint19 waypoint10)
	(visible waypoint10 waypoint19)
	(visible waypoint19 waypoint12)
	(visible waypoint12 waypoint19)
	(visible waypoint19 waypoint15)
	(visible waypoint15 waypoint19)
	(visible waypoint19 waypoint21)
	(visible waypoint21 waypoint19)
	(visible waypoint20 waypoint8)
	(visible waypoint8 waypoint20)
	(visible waypoint20 waypoint11)
	(visible waypoint11 waypoint20)
	(visible waypoint20 waypoint15)
	(visible waypoint15 waypoint20)
	(visible waypoint20 waypoint16)
	(visible waypoint16 waypoint20)
	(visible waypoint20 waypoint21)
	(visible waypoint21 waypoint20)
	(visible waypoint20 waypoint28)
	(visible waypoint28 waypoint20)
	(visible waypoint21 waypoint2)
	(visible waypoint2 waypoint21)
	(visible waypoint21 waypoint3)
	(visible waypoint3 waypoint21)
	(visible waypoint21 waypoint8)
	(visible waypoint8 waypoint21)
	(visible waypoint21 waypoint13)
	(visible waypoint13 waypoint21)
	(visible waypoint21 waypoint22)
	(visible waypoint22 waypoint21)
	(visible waypoint21 waypoint23)
	(visible waypoint23 waypoint21)
	(visible waypoint21 waypoint26)
	(visible waypoint26 waypoint21)
	(visible waypoint22 waypoint18)
	(visible waypoint18 waypoint22)
	(visible waypoint22 waypoint28)
	(visible waypoint28 waypoint22)
	(visible waypoint22 waypoint29)
	(visible waypoint29 waypoint22)
	(visible waypoint23 waypoint1)
	(visible waypoint1 waypoint23)
	(visible waypoint23 waypoint17)
	(visible waypoint17 waypoint23)
	(visible waypoint23 waypoint28)
	(visible waypoint28 waypoint23)
	(visible waypoint24 waypoint8)
	(visible waypoint8 waypoint24)
	(visible waypoint24 waypoint10)
	(visible waypoint10 waypoint24)
	(visible waypoint24 waypoint13)
	(visible waypoint13 waypoint24)
	(visible waypoint24 waypoint15)
	(visible waypoint15 waypoint24)
	(visible waypoint24 waypoint21)
	(visible waypoint21 waypoint24)
	(visible waypoint24 waypoint22)
	(visible waypoint22 waypoint24)
	(visible waypoint25 waypoint2)
	(visible waypoint2 waypoint25)
	(visible waypoint25 waypoint14)
	(visible waypoint14 waypoint25)
	(visible waypoint26 waypoint9)
	(visible waypoint9 waypoint26)
	(visible waypoint26 waypoint10)
	(visible waypoint10 waypoint26)
	(visible waypoint26 waypoint22)
	(visible waypoint22 waypoint26)
	(visible waypoint27 waypoint22)
	(visible waypoint22 waypoint27)
	(visible waypoint27 waypoint24)
	(visible waypoint24 waypoint27)
	(visible waypoint27 waypoint28)
	(visible waypoint28 waypoint27)
	(visible waypoint28 waypoint6)
	(visible waypoint6 waypoint28)
	(visible waypoint28 waypoint17)
	(visible waypoint17 waypoint28)
	(visible waypoint28 waypoint24)
	(visible waypoint24 waypoint28)
	(visible waypoint28 waypoint25)
	(visible waypoint25 waypoint28)
	(visible waypoint28 waypoint29)
	(visible waypoint29 waypoint28)
	(visible waypoint29 waypoint7)
	(visible waypoint7 waypoint29)
	(visible waypoint29 waypoint10)
	(visible waypoint10 waypoint29)
	(visible waypoint29 waypoint19)
	(visible waypoint19 waypoint29)
	(at_soil_sample waypoint0)
	(at_soil_sample waypoint1)
	(at_soil_sample waypoint2)
	(at_soil_sample waypoint3)
	(at_rock_sample waypoint3)
	(at_rock_sample waypoint4)
	(at_rock_sample waypoint6)
	(at_rock_sample waypoint8)
	(at_rock_sample waypoint9)
	(at_rock_sample waypoint10)
	(at_soil_sample waypoint11)
	(at_soil_sample waypoint12)
	(at_rock_sample waypoint12)
	(at_rock_sample waypoint14)
	(at_rock_sample waypoint15)
	(at_rock_sample waypoint19)
	(at_rock_sample waypoint20)
	(at_soil_sample waypoint21)
	(at_soil_sample waypoint23)
	(at_soil_sample waypoint24)
	(at_rock_sample waypoint24)
	(at_soil_sample waypoint25)
	(at_rock_sample waypoint25)
	(at_rock_sample waypoint26)
	(at_rock_sample waypoint28)
	(at_rock_sample waypoint29)
	(at_lander general waypoint27)
	(channel_free general)
	(empty rover0store)
	(empty rover1store)
	(empty rover2store)
	(at rover3 waypoint3)
	(available rover3)
	(store_of rover3store rover3)
	(empty rover3store)
	(equipped_for_soil_analysis rover3)
	(equipped_for_rock_analysis rover3)
	(equipped_for_imaging rover3)
	(can_traverse rover3 waypoint3 waypoint10)
	(can_traverse rover3 waypoint10 waypoint3)
	(can_traverse rover3 waypoint3 waypoint13)
	(can_traverse rover3 waypoint13 waypoint3)
	(can_traverse rover3 waypoint3 waypoint27)
	(can_traverse rover3 waypoint27 waypoint3)
	(can_traverse rover3 waypoint10 waypoint4)
	(can_traverse rover3 waypoint4 waypoint10)
	(can_traverse rover3 waypoint10 waypoint9)
	(can_traverse rover3 waypoint9 waypoint10)
	(can_traverse rover3 waypoint10 waypoint14)
	(can_traverse rover3 waypoint14 waypoint10)
	(can_traverse rover3 waypoint10 waypoint17)
	(can_traverse rover3 waypoint17 waypoint10)
	(can_traverse rover3 waypoint10 waypoint18)
	(can_traverse rover3 waypoint18 waypoint10)
	(can_traverse rover3 waypoint10 waypoint19)
	(can_traverse rover3 waypoint19 waypoint10)
	(can_traverse rover3 waypoint10 waypoint20)
	(can_traverse rover3 waypoint20 waypoint10)
	(can_traverse rover3 waypoint10 waypoint21)
	(can_traverse rover3 waypoint21 waypoint10)
	(can_traverse rover3 waypoint10 waypoint26)
	(can_traverse rover3 waypoint26 waypoint10)
	(can_traverse rover3 waypoint10 waypoint29)
	(can_traverse rover3 waypoint29 waypoint10)
	(can_traverse rover3 waypoint27 waypoint7)
	(can_traverse rover3 waypoint7 waypoint27)
	(can_traverse rover3 waypoint27 waypoint8)
	(can_traverse rover3 waypoint8 waypoint27)
	(can_traverse rover3 waypoint27 waypoint15)
	(can_traverse rover3 waypoint15 waypoint27)
	(can_traverse rover3 waypoint27 waypoint24)
	(can_traverse rover3 waypoint24 waypoint27)
	(can_traverse rover3 waypoint4 waypoint6)
	(can_traverse rover3 waypoint6 waypoint4)
	(can_traverse rover3 waypoint4 waypoint12)
	(can_traverse rover3 waypoint12 waypoint4)
	(can_traverse rover3 waypoint9 waypoint16)
	(can_traverse rover3 waypoint16 waypoint9)
	(can_traverse rover3 waypoint14 waypoint5)
	(can_traverse rover3 waypoint5 waypoint14)
	(can_traverse rover3 waypoint17 waypoint23)
	(can_traverse rover3 waypoint23 waypoint17)
	(can_traverse rover3 waypoint17 waypoint28)
	(can_traverse rover3 waypoint28 waypoint17)
	(can_traverse rover3 waypoint18 waypoint22)
	(can_traverse rover3 waypoint22 waypoint18)
	(can_traverse rover3 waypoint21 waypoint2)
	(can_traverse rover3 waypoint2 waypoint21)
	(can_traverse rover3 waypoint29 waypoint0)
	(can_traverse rover3 waypoint0 waypoint29)
	(can_traverse rover3 waypoint8 waypoint25)
	(can_traverse rover3 waypoint25 waypoint8)
	(can_traverse rover3 waypoint6 waypoint11)
	(can_traverse rover3 waypoint11 waypoint6)
	(can_traverse rover3 waypoint23 waypoint1)
	(can_traverse rover3 waypoint1 waypoint23)
	(empty rover4store)
	(empty rover5store)
	(empty rover6store)
	(empty rover7store)
	(empty rover8store)
	(empty rover9store)
	(calibration_target camera0 objective3)
	(supports camera0 high_res)
	(supports camera0 low_res)
	(on_board camera1 rover3)
	(calibration_target camera1 objective3)
	(supports camera1 high_res)
	(calibration_target camera2 objective0)
	(supports camera2 colour)
	(supports camera2 high_res)
	(supports camera2 low_res)
	(calibration_target camera3 objective0)
	(supports camera3 colour)
	(supports camera3 high_res)
	(calibration_target camera4 objective4)
	(supports camera4 high_res)
	(calibration_target camera5 objective4)
	(supports camera5 high_res)
	(supports camera5 low_res)
	(calibration_target camera6 objective3)
	(supports camera6 high_res)
	(supports camera6 low_res)
	(visible_from objective0 waypoint0)
	(visible_from objective0 waypoint1)
	(visible_from objective0 waypoint2)
	(visible_from objective0 waypoint3)
	(visible_from objective0 waypoint4)
	(visible_from objective0 waypoint5)
	(visible_from objective0 waypoint6)
	(visible_from objective0 waypoint7)
	(visible_from objective0 waypoint8)
	(visible_from objective0 waypoint9)
	(visible_from objective0 waypoint10)
	(visible_from objective0 waypoint11)
	(visible_from objective0 waypoint12)
	(visible_from objective0 waypoint13)
	(visible_from objective0 waypoint14)
	(visible_from objective0 waypoint15)
	(visible_from objective0 waypoint16)
	(visible_from objective1 waypoint0)
	(visible_from objective1 waypoint1)
	(visible_from objective1 waypoint2)
	(visible_from objective1 waypoint3)
	(visible_from objective1 waypoint4)
	(visible_from objective1 waypoint5)
	(visible_from objective1 waypoint6)
	(visible_from objective1 waypoint7)
	(visible_from objective1 waypoint8)
	(visible_from objective1 waypoint9)
	(visible_from objective1 waypoint10)
	(visible_from objective1 waypoint11)
	(visible_from objective1 waypoint12)
	(visible_from objective1 waypoint13)
	(visible_from objective1 waypoint14)
	(visible_from objective1 waypoint15)
	(visible_from objective1 waypoint16)
	(visible_from objective1 waypoint17)
	(visible_from objective1 waypoint18)
	(visible_from objective1 waypoint19)
	(visible_from objective1 waypoint20)
	(visible_from objective1 waypoint21)
	(visible_from objective1 waypoint22)
	(visible_from objective1 waypoint23)
	(visible_from objective1 waypoint24)
	(visible_from objective1 waypoint25)
	(visible_from objective1 waypoint26)
	(visible_from objective2 waypoint0)
	(visible_from objective2 waypoint1)
	(visible_from objective2 waypoint2)
	(visible_from objective2 waypoint3)
	(visible_from objective2 waypoint4)
	(visible_from objective2 waypoint5)
	(visible_from objective2 waypoint6)
	(visible_from objective3 waypoint0)
	(visible_from objective3 waypoint1)
	(visible_from objective3 waypoint2)
	(visible_from objective3 waypoint3)
	(visible_from objective4 waypoint0)
	(visible_from objective4 waypoint1)
	(visible_from objective4 waypoint2)
	(visible_from objective4 waypoint3)
	(visible_from objective4 waypoint4)
	(visible_from objective5 waypoint0)
	(visible_from objective5 waypoint1)
	(visible_from objective5 waypoint2)
	(visible_from objective5 waypoint3)
	(visible_from objective5 waypoint4)
	(visible_from objective5 waypoint5)
	(visible_from objective5 waypoint6)
	(visible_from objective5 waypoint7)
	(visible_from objective5 waypoint8)
	(visible_from objective5 waypoint9)
	(visible_from objective5 waypoint10)
	(visible_from objective5 waypoint11)
	(visible_from objective5 waypoint12)
	(visible_from objective5 waypoint13)
	(visible_from objective5 waypoint14)
	(visible_from objective5 waypoint15)
	(visible_from objective5 waypoint16)
	(visible_from objective5 waypoint17)
	(visible_from objective5 waypoint18)
	(visible_from objective5 waypoint19)
	(visible_from objective5 waypoint20)
	(visible_from objective5 waypoint21)
	(visible_from objective5 waypoint22)
	(visible_from objective5 waypoint23)
	(visible_from objective5 waypoint24)
	(visible_from objective5 waypoint25)
	(visible_from objective5 waypoint26)
	(visible_from objective5 waypoint27)
	(visible_from objective5 waypoint28)
	(visible_from objective5 waypoint29)
)
(:goal
	(and
		(communicated_soil_data waypoint24)
		(communicated_soil_data waypoint2)
		(communicated_soil_data waypoint1)
		(communicated_soil_data waypoint12)
		(communicated_soil_data waypoint11)
		(communicated_soil_data waypoint21)
		(communicated_soil_data waypoint3)
		(communicated_soil_data waypoint25)
		(communicated_rock_data waypoint6)
		(communicated_rock_data waypoint28)
		(communicated_rock_data waypoint20)
		(communicated_image_data objective1 low_res)
		(communicated_image_data objective5 high_res)
		(communicated_image_data objective0 high_res)
	)
)
)