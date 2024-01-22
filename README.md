# scale-maker
scale-maker is a multi-propose internal built scale, perf, load generator which contains the binaries and tools needed to perform scale, perf or load tests. scale-maker was built on top of the open source framework and libraries such as go-fiber, client-go and apimachinery. In addition, it integrates with common load test tools like stress-ng and fio. scale-maker acts as an controller which the users can use to spin up measurable load containers via the built-in swagger UI or by calling the APIs directly.

The scale-maker container can be used to scale up the number of nodes and amount of loads within the K8S clusters. Here are some of the examples how scale-maker can be used to perform various debugging and troubleshooting.

1. Troubleshoot noisy neighbours - stress-ng pods and their CPU/memory resources can be configured to troubleshoot and debug the noisy neighbour issues. Here’s an example Mayank Kumar forwarded.
2. Provide perf/stress test binaries - fio binary was needed for GP3 root volume testing. The fio binary was added as a result. 
3. Verify zombie processes/pid limits - stress-ng has a number of parameters to spin up large number of zombie processes or normal load processes.
4. Debug CPU, memory and IO related issues - stress-ng can spin up separate or combine CPU/memory/IO work loads. Scale-maker can even randomly execute stress test from the 300+ stress-ng supported stress test cases. 

## Documentations
### Run the program locally
```
make run
```
### API Doc
If running the program locally, visit the local URL

[http://localhost:3000/docs](http://localhost:3000/docs)

### Workflow Diagram
![scale-maker](https://github.com/lliu8080/scale_maker/blob/main/docs/scale-maker.png)

### Command to Update Swagger Docs

See the swag [document](https://github.com/swaggo/swag/blob/master/README.md) for details.
```
swag fmt
swag init
```
