#!/usr/bin/env python3
# -*- coding: utf-8 -*-

domain_tmpl = """(define (domain production-site)
  (:requirements :typing :factored-privacy)
  (:types
    product property - object
  )
  (:constants
    {products} - product
    {properties} - property
  )
  (:predicates
    (:private
      (satisfies ?o - product ?p - property)
    )
    (completed ?o - product)
  )
  (:functions
    (total-cost)
  )
  (:action Deliver
    :parameters (?o - product)
    :precondition (and
      (not (completed ?o))
{requirements}
    )
    :effect (and 
      (completed ?o)
      (increase (total-cost) {C})
    )
  )
  (:action refine
    :parameters (?o - product ?p - property)
    :precondition (and
      (not (satisfies ?o ?p))
    )
    :effect (and
      (satisfies ?o ?p)
    )
  )
)
"""

problem_tmpl = """(define (problem production-site)
  (:domain production-site)
  (:init
    (= (total-cost) 0)
{initials}
  )
  (:goal
    (and
{completions}
    )
  )
  (:metric minimize (total-cost))
)
"""

comptmpl = "      (completed {})"
ncomptmpl = "    (not (completed {}))"


def create_domain(products, properties, nonproperties, agent):
    d = {
        'products': " ".join(products),
        'properties': (" ".join(properties) + '\n    ' + " ".join(nonproperties)),
        #'requirements': '\n'.join(reqtmpl.format(p, p) for p in properties)}
        'requirements': (
            '      (satisfies ?o {})\n' * len(properties) +
            '      (not (satisfies ?o {}))\n' * len(nonproperties)).format(
                *properties, *nonproperties),
        'C': 4 if agent == 0 else 5,
    }
    return domain_tmpl.format(**d)


def create_problem(products):
    d = {
        'initials': "\n".join(ncomptmpl.format(p) for p in products),
        'completions': "\n".join(comptmpl.format(p) for p in products)
        }
    return problem_tmpl.format(**d)


if __name__ == '__main__':
    import sys
    import pathlib
    num_agents, num_products, num_properties, num_nonproperties, folder = sys.argv[1:6]
    pathlib.Path(folder).mkdir(parents=True, exist_ok=True) 

    products = ["p{}".format(x) for x in range(int(num_products))]
    properties = ["prop{}".format(x) for x in range(int(num_properties))]
    nonproperties = ["alt{}".format(y) for y in range(int(num_nonproperties))]

    for agent in range(int(num_agents)):
        domfile = 'domain-{}.pddl'.format(agent)
        with open(pathlib.PurePath(folder, domfile), 'w') as f: 
            f.write(create_domain(products, properties, nonproperties, agent))
        probfile = 'problem-{}.pddl'.format(agent)
        with open(pathlib.PurePath(folder, probfile), 'w') as f: 
            f.write(create_problem(products))
        print("DONE agent {}".format(agent))
