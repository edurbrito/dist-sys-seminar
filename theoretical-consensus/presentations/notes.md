Today
I will present to you the final stage 
of my research about the recent advancements on
permissionless consensus
with a special focus on the transition to 
resource-constrained environments
for example, mobile and edge computing.


Just to recall from the past presentations,
I first started to dive into the roots
of the consensus problem in the realms of distributed systems.
And consensus can be pretty much defined as
reaching an agreement between multiple parties
in the presence of potentially faulty individuals.
This definition was pioneered by Lamport some time ago
and along with it there came the first very basic 
synchronous solutions for the consensus problem.


These initial attempts failed at considering interesting challenges.
And allegorically, one of the most famous problems in distributed systems
was finally born: The Byzantine general's problem.
Along with it, the first protocols that were already considering
some types of faulty or Byzantine behavior were finally proposed.
But at this point, these mechanisms were mostly based on 
the knowledge or the role definition
of the participating peers, and so only able to run in permissioned systems.


In 2008, finally, the first practical consensus mechanism 
that could run in a permissionless environment
was proposed, with the creation of Bitcoin.
It is a proof-of-work consensus protocol 
that resembles a ”replicated state machine”
where the independent participants reach agreement not only
about transactional values, but also about their order - naturally
forming the underlying structure of what is now known as a
blockchain.
And despite the existence of very well-defined 
Byzantine Fault-Tolerant consensus protocols 
that run very fast and optimally in permissioned environments,
The appearance of their permissionless counterparts is justified
by the necessity of dealing with today’s
sparse and unpredictable networks of 
anonymously and dynamically participating devices.


But what are the building blocks of permissionless consensus?
Most of the well known protocols follow some specific architectural decisions
Some authors ended up identifying distinct phases.
And so, very concisely, the first is the phase of information, or block, creation, known as mining.
A second phase is then needed for gossiping and spreading that novel piece of information through the network.
And then, there is a third phase where the listening nodes validate that information and agree on it.
And this is when finally consensus is reached
This is the simplistic view, of course.
But it captures the essence of most of the permissionless consensus protocols out there. There may be some other components, like incentive mechanisms, or other block proposal schemes, that I reviewed in the report.


Bitcoin's proof-of-work follows, more or less, this arrangement.
But there are some particularities that should be mentioned, for example, 
the block generation interval, or the block size,
That when combined allow for the definition of a trilemma 
between decentralization, security and scalability.
Summarizing, relaxing the security requirements 
may allow for more scalability, both of
which, consequently, have hands tied with decentralization.


Taking proof-of-work as a reference
I followed with my research
And ended up reviewing other block finalization schemas.
    

Like Proof-of-Stake, Proof-of-Space, Proof-of-Elapsed-Time, or Proof-of-Location.
And the conclusion is that all of them present some trade-offs 
and not all of them may attempt to entirely
solve the permissionless consensus problem,
But may be adapted to specific use cases and network goals.


I also came across different approaches, 
Specifically different data structures
That try to model the consensus problem
a bit differently from the typical blockchain approach.
For example, the XRP Ledger consensus adopts 
an overlapping clique configuration and voting mechanism.
And the IOTA network uses a Directed Acyclic graph 
for fast and scalable transaction validation.


Finally, I also reviewed some of the most recent
attempts to adapt the consensus problem
to resource-constrained environments.


These environments are often characterized by
the limited computational resources, heterogeneity, 
mobility, or dynamism.
Having in mind these characteristics, 
and the drawbacks of the centralized architectures of today, it
may be postulated the need for a decentralized or permissionless shift that addresses some the most important challenges
of these deployments: scalability, security, and privacy
as enablers of a decentralized future.


The creation of protocols that could address these challenges
should aim to provide a solution for a key set of aspects 
that may have not been totally considered before,
like fairer resource consumption, the heterogeneity of the devices, or
overall fault-tolerance, openness, and security against a multitude of new attacks, without excluding, of course, 
the other general requirements of a permissionless system as seen before.
Some attempts to address all these challenges
have been seen out there in the wild, for example with the IOTA Tangle protocol, or some projects that aim at running the new Ethereum protocol on edgy devices. However, these solutions not only have some undesired centralization aspects, but they also fail at taking advantage of the unique characteristics of these computing environments. An interesting path that I may take as future work is, for example, Proof-of-Location, that may benefit from the easiness of movement or targeted location deployment of a lot of these devices, in order to reach consensus.


And so, to conclude,
I would like to quickly show the practical component of my research.
A simple prototype that was developed to demonstrate 
the effectiveness of Proof-of-Work as a permissionless consensus protocol.
This prototype was made to be run on any browser, in any device, 
and is served as a static webpage.
It runs in a seemingly decentralized fashion, where the devices
are connected in a peer-to-peer way, and all of them
compete to mine blocks and reach consensus.
The rest of the implementation details are in the report.
And so, if we access the web page
We can see it running and the different tabs or devices 
showing the same view while competing to mine the next blocks.