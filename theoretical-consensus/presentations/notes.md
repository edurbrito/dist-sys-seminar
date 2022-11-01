Today
I will present to you the progress 
that I have been doing so far
regarding my research on
the recent advancements on
permissionless consensus
with the final intended focus
on resource constrained environments
for example, mobile and edge compution


Just to remember everyone
I first started to dive into the roots
of the consensus problem
in the realms of distributed systems
and consensus can be pretty much defined as
reaching an agreement between multiple parties
in the presence of potentially faulty individuals.
This definition was pioneered by Lamport some time ago
and along with it there came 
the first very basic syncronous solutions for the consensus problem


They noticed that some interesting challenges were not being considered
and so they then defined one of the most famous problems in distributed systems
The byzantine generals problem
Along with it, they also proposed the first protocols that were already considering
some types of faulty or byzantine behaviour
After that, some practical implementations and asynchronous mechanisms started to appear
But at this point, there were some very strong assumtions in the table 
that would not model interesting challenges of todays internet
These old mechanisms were mostly based on the knowledge or the role definition
of the participating peers, and so only able to run in permissioned systems


In 2008 finally, the first practical consensus mechanism 
that could run in a permissionless environment
was proposed, with the creation of Bitcoin
It is a proof-of-work consensus protocol 
that resembles a ”replicated state machine”
where the independent participants reach agreement not only
about transactional values, but also about their order - naturally
forming the underlying structure of what is now known as a
blockchain.
And despite the existance of very well defined 
Byzantine fault tolerant consensus protocols 
that run very fast in permissioned environments
The existance of Proof-of-Work is justified
by the necessity of dealing with today’s
sparse and unpredictable networks of 
anonymously and dynamically participating devices.


But what are the building blocks of permissionless consensus?
Most of the well known protocols follow some specific architectural decisions
Some authors divide them into phases and concisely 
We may have a first phase of information, or block, creation known as mining
A second phase for gossiping and spreading the information through the network
And a third phase where the listening nodes validate that info and agree on it
And this is when the consensus is reached
This is the simplistic view, of course
There may be some other components, like incentive mechanisms
Or other block proposal schemes
That I also addressed in the report


Bitcoin's proof-of-work follows more or less this arrangement
But there are some particularities, for example, 
the block generation interval, or the block size
That when combined allow for the definition of a trilemma 
between decentralization, security and scalability
Summarizing, relaxing the security requirements 
may allow for more scalability, both of
which, consequently, have hands tied with decentralization.


Taking proof-of-work as a reference
I followed with my research
And ended up reviewing other block finalization schemas
Like Proof-of-Stake, Proof-of-Space, Proof-of-Elapsed-Time, or Proof-of-Location
And the conclusion is that
All of them present some trade-offs and not all of them may attempt to entirely
solve the permissionless consensus problem
But may be adapted to specific situations and goals


I also came accross different approaches 
Specifically different data structures
That try to model the consensus problem
a bit different from the typical blockchain approach
For example
The XRP Ledge consensus adopts an overlapping clique configuration 
and voting mechanism
And the IOTA network uses a Directed Acyclic graph 
for fast and decentralized transaction validation


What is missing is to research about 
the state of the art
and last trials on applying 
these permissionless consensus mechanisms
In resource constrained networks
