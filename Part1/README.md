Part1: Thinking about elevators
---------------------------

Not for handing in, just for thinking about. Talk to other groups, assistants, or even people who have taken the course in previous years.

Brainstorm some techniques you could use to prevent a user from being hopelessly stranded, waiting for an elevator that will never arrive. Think about the [worst-case](http://xkcd.com/748/) behaviour of the system.
 - What if the software controlling one of the elevators suddenly crashes?
   - Stop the elevator and restart the software
   - Orders may be lost if they haven't been persisted. This means we need persistance.
   - Orders should be distributed to the other elevators, and this should be done as soon as possible. That way the user doesn't need to wait for the software to come back online before an elevator can serve them.
 - What if it doesn't crash, but hangs?
   - Something should check periodically if the software is responding, and restart it if it's not.
 - What if a message between machines is lost?
   - Lost packages need to be retransmitted. This can be done through the use of handshaking. 
 - What if the network cable is suddenly disconnected? Then re-connected?
   - With proper handshaking we can ensure no information is lost due to packet loss during the disconnect.
   - If this happens while the car is moving. Maybe stop the car. 
   - If the other elevators regularily ping all other elevators in the system, they will know if one has been disconnected and can send a replacement elevator instead
 - What if a user of the system is being a troll?
   - Ban them
 - What if the elevator car never arrives at its destination?
   - This can indicate that something is wrong with the elevator in question. After a certain time, send another elevator.
