(define (problem depotprob5646) (:domain depot)
(:objects
	distributor1 - distributor
	distributor0 - distributor
	distributor2 - distributor
	depot0 - depot
	depot1 - depot
	depot2 - depot
	pallet5 - pallet
	pallet4 - pallet
	pallet7 - pallet
	pallet6 - pallet
	pallet1 - pallet
	pallet0 - pallet
	pallet3 - pallet
	pallet2 - pallet
	pallet9 - pallet
	pallet8 - pallet
	truck1 - truck
	truck0 - truck
	crate5 - crate
	crate4 - crate
	crate1 - crate
	crate0 - crate
	crate3 - crate
	crate2 - crate

	(:private
		hoist0 - hoist
	)
)
(:init
	(at pallet0 depot0)
	(clear crate2)
	(at pallet1 depot1)
	(clear pallet1)
	(at pallet2 depot2)
	(clear crate5)
	(at pallet3 distributor0)
	(clear crate4)
	(at pallet4 distributor1)
	(clear pallet4)
	(at pallet5 distributor2)
	(clear pallet5)
	(at pallet6 distributor1)
	(clear pallet6)
	(at pallet7 depot0)
	(clear pallet7)
	(at pallet8 depot0)
	(clear crate3)
	(at pallet9 distributor0)
	(clear pallet9)
	(at truck0 distributor1)
	(at truck1 depot0)
	(at hoist0 depot0)
	(available depot0 hoist0)
	(at crate0 depot2)
	(on crate0 pallet2)
	(at crate1 depot2)
	(on crate1 crate0)
	(at crate2 depot0)
	(on crate2 pallet0)
	(at crate3 depot0)
	(on crate3 pallet8)
	(at crate4 distributor0)
	(on crate4 pallet3)
	(at crate5 depot2)
	(on crate5 crate1)
)
(:goal
	(and
		(on crate0 pallet0)
		(on crate1 pallet5)
		(on crate2 pallet4)
		(on crate3 pallet7)
		(on crate4 pallet9)
		(on crate5 pallet1)
	)
)
)