# Notes

## Database

For the test I did not add a sql database due to lack of time, requires some setup that I could not really copy & paste from the deployments I have done since I use to make use of Kubernetes rather than docker compose. Let me note several hints on deploying SQL databases.

**If deploying to the cloud**
 <br /> If you are to use GCP to hold your deployment, you are recommended to actually not deploy a SQL docker image but use the sidecar container patter to deploy a cloudsql-proxy with an image such as "gcr.io/cloudsql-docker/gce-proxy:1.16". One of the advantages of choosing this deployment is a better and more direct communication fdor the database than if you need to communicate over containers. This sidecar pattern is used in logging too.
 
**Regarding the database docker image**
 <br /> For local deployments you can use the official image of mysql or mariadb also works. To use this image you are required to add as a variable the root password. Using kubernetes you are encouraged to use secrets to provide that. Also you can mention the scripts that the container should run when being created, you can use that to create the schema that will hold your deployment. You can also add the tables but for a more robust deployment is recommended to use database evolutions. For an exercise such as this would be fine to use the root password for the application but initial scripts such as the schema creation should be used to create additional users for the applications and developers. Again advisable to use some obfuscating mechanism to add the passwords to the scripts. The scripts will be found in /docker-entrypoint-initdb.d as explained in https://hub.docker.com/_/mysql, to place the scripts there you can create your own mysql docker image.
  
**Docker compose**
 <br /> The docker compose file (if not using kubernetes) should contain the information for instantiating both containers. Regarding the go application, nothing special as in env variables is required (if not passing the db password through variables). I choosed port 8080 for the server to listen (in /adapters/http/adapter.go. This could be improved by centralizing such configuration). Database should use a port that would be specified in the go configuration (see next line) as well as the root password (or password of the user that the app will use).
 
**Connecting from the application**
 <br /> To create the database connection you can use this syntax in go: 	
 <pre><code>db, err := sql.Open("mysql","user:password@tcp(127.0.0.1:3306)/hello")</code></pre>

 <br /> To query the database: 
 <pre><code>rows, err := db.Query("select id, name from users where id = ?", 1) </code></pre>
 but to avoid sql injection attacks is recommended db.Prepare
 
 <br /> 
 <br /> To parse the results 
 <pre><code> err := rows.Scan(&id, &name) // inside a loop</code></pre>

## A note about immutability
Coming from a functional language I changed some signatures since I found more appropiate. Of course this is my coding style, I could be talked into changing my mind and not a fan of refactoring just for the sake of it. Nevertheless I thought it worth making a point and it was enforced in the instructions of the exercise. This is from the git diff:
 <pre><code>
-	ListKnights() []*domain.Knight
+	ListKnights() []domain.Knight
</code></pre>

The knight is an entity stored in the database, it makes not sense for the code to change it if is it not to be stored to the database again, otherwise we would have two different versions of the same entity (in code and in database). That is why I enforced that by changing this signature among others.