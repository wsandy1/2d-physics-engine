package main

import "github.com/wsandy1/2d-physics-engine/vector2"

type PhysicsEngine struct {
	RigidBodies []RigidBody
	Gravity     vector2.Vec2
	substeps    uint16
}

func (e *PhysicsEngine) Update(tps float32) {
	var dt float32 = 1 / tps
	// sub_dt := dt /float32(e.substeps)
	// for i := 0; i < int(e.substeps); i++ {
	e.ApplyGravity()
	e.N2Solve()
	e.VerletSolve(dt)
	// }
}

// add global gravity acceleration vector to the acceleration of each RigidBody
func (e *PhysicsEngine) ApplyGravity() {
	for i := range e.RigidBodies {
		e.RigidBodies[i].Accelerate(e.Gravity)
	}
}

// integrate using Verlet method to find next position
func (e *PhysicsEngine) VerletSolve(dt float32) {
	for i := range e.RigidBodies {
		vel := vector2.Sub(e.RigidBodies[i].current_position, e.RigidBodies[i].last_position)
		e.RigidBodies[i].last_position = e.RigidBodies[i].current_position
		e.RigidBodies[i].current_position = vector2.Add(vector2.Add(e.RigidBodies[i].current_position, vel), vector2.ConstMul(vector2.ConstMul(e.RigidBodies[i].acceleration, dt), dt))
		e.RigidBodies[i].acceleration = vector2.Vec2{X: 0, Y: 0}
	}
}

// if you need a comment to understand what this does you should go back to school...
func (e *PhysicsEngine) N2Solve() {
	for i := range e.RigidBodies {
		var resultant vector2.Vec2
		for j := range e.RigidBodies[i].point_forces {
			resultant = vector2.Add(resultant, e.RigidBodies[i].point_forces[j].Vector)
		}
		e.RigidBodies[i].Accelerate(vector2.ConstDiv(resultant, float32(e.RigidBodies[i].mass)))
	}
}

func (e *PhysicsEngine) MomentSolve() {
	for i := range e.RigidBodies {
		var total_torque float32
		for j := range e.RigidBodies[i].point_forces {
			// because the CoM is at {0,0}, the origin vector of the force is the same as the radius vector CoM->Force Origin
			// rotate 90 to give perpendicular radius vector
			r_perp := e.RigidBodies[i].point_forces[j].Origin.Rotate(90)
			// torque exerted by the force on the CoM is the dot product of the perpendicular radius vector from CoM->ForceOrigin and the force vector.
			torque := vector2.Dot(r_perp, e.RigidBodies[i].point_forces[j].Vector)
			total_torque = total_torque + torque
		}
		// angular acceleration = torque / moment of inertia
		a := total_torque / e.RigidBodies[i].moi
		e.RigidBodies[i].AngularAccelerate(a)
	}
}
