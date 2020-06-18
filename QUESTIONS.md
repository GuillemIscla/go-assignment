# Questions

For this assignment you also have to answer a couple of questions.
There is no correct answer and none is mandatory, if you don't know just skip it.

 - **What do you think of the initial project structure ?**
 <br /> Being a small project does not really matter but if we were to expand there should be defined some architectural structure. If we were to choose DDD (Domain Driven Design), the Knight or any fighter would be consider entities, so Id would be mandatory. Also, if the purpose of the arena keeps similar, I would consider it an application, not an entitiy (since does things but it has not really a life-cycle). Of course the structure should change to adapt to the typical DDD structure. Among many things, I would go to some refactor too when it comes to mutability of the entities, in general terms entities can perfectly change (be updated) but that only really makes sense before saving it to the repository. 


 - **What you will improve from your solution ?**
 <br /> - Add SQL database, please see notes
 <br /> - Centralize configuration
 <br /> - Clear separation between definition and implementation in some places
 <br /> - Consider a parsing module
 <br /> - Structure the error codes to have a more consistent API

 - **For you, what are the boundaries of a service inside a "micro-service" architecture ?**
 <br /> One should be able to describe the purpose of the service in terms of the business logic of the application. After that it should be fairly independent of the other services since they are loosy coupled and therefore the communication between them should be minimal. As for the coding, it should inherit common code with other services when they are working with the same concepts, when two micro-services share enough conceptual similarities one could consider to code them in the same project but deploy them separately. 


 - **For you, what are the most relevant usage for SQL, NoSQL, key-value and document store databases ?**
 <br /> **SQL:** Storage for structured data which is highly related.
 <br /> **NoSQL:** Storage for data with a free structure that where you want to run customized queries. There term is quite ambiguous, there are many.
 <br /> **key-value:** Storage for data you just want to query by its id, non-related. Considered NoSql
 <br /> **document store:** Storage for documents. Considered key-value, so is also NoSql. The metadata of the documents could be used for optimization purposes.
   

