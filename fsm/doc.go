// Package fsm implements a finite state machine as described by:
//
//	(A)<-[0]-(A)-[1]->(B)
//	(C)<-[0]-(B)-[1]->(A)
//	(B)<-[0]-(C)-[1]->(A)
package fsm
